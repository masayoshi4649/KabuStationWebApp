package main

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	sessionCookieName = "kswa_session"
	sessionTTL        = 24 * time.Hour
)

type sessionEntry struct {
	ExpiresAt time.Time
}

var (
	sessionMu    sync.RWMutex
	sessionStore = map[string]sessionEntry{}
)

type ReqLoginPOST struct {
	ID       string `form:"id" json:"id"`
	PW       string `form:"pw" json:"pw"`
	Username string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// isAuthEnabled は、ログイン機能を有効化する設定かどうかを判定します。
//
// 機能:
//   - 設定ファイルの [SYSTEM] SYSID/SYSPW が両方指定されているかを判定する
//
// 引数およびその型:
//   - なし
//
// 返り値およびその型:
//   - bool: 有効な場合 true、無効な場合 false
func isAuthEnabled() bool {
	return strings.TrimSpace(cfg.System.SysID) != "" && strings.TrimSpace(cfg.System.SysPW) != ""
}

// issueSessionToken は、セッション識別子として利用するランダムトークンを生成します。
//
// 機能:
//   - 暗号学的に安全な乱数から 32 バイトのトークンを生成する
//   - URL セーフな base64 文字列へエンコードして返す
//
// 引数およびその型:
//   - なし
//
// 返り値およびその型:
//   - string: セッショントークン
//   - error: 乱数生成に失敗した場合のエラー
func issueSessionToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// storeSession は、セッション情報をメモリ上に保存します。
//
// 機能:
//   - トークンと有効期限を対応付けて保存する
//
// 引数およびその型:
//   - token string: セッション識別子
//   - ttl time.Duration: セッション有効期間
//
// 返り値およびその型:
//   - time.Time: 失効時刻
func storeSession(token string, ttl time.Duration) time.Time {
	expiresAt := time.Now().Add(ttl)
	sessionMu.Lock()
	sessionStore[token] = sessionEntry{ExpiresAt: expiresAt}
	sessionMu.Unlock()
	return expiresAt
}

// deleteSession は、セッション情報を削除します。
//
// 機能:
//   - 指定されたトークンのセッションを破棄する
//
// 引数およびその型:
//   - token string: セッション識別子
//
// 返り値およびその型:
//   - なし
func deleteSession(token string) {
	if token == "" {
		return
	}
	sessionMu.Lock()
	delete(sessionStore, token)
	sessionMu.Unlock()
}

// getValidSessionToken は、リクエストに紐づく有効なセッションを取得します。
//
// 機能:
//   - Cookie からセッショントークンを読み取る
//   - メモリ上のセッションストアと照合し、期限内の場合のみ有効とする
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - string: 有効なセッショントークン（無効時は空文字）
func getValidSessionToken(c *gin.Context) string {
	token, err := c.Cookie(sessionCookieName)
	if err != nil || token == "" {
		return ""
	}

	now := time.Now()
	sessionMu.RLock()
	entry, ok := sessionStore[token]
	sessionMu.RUnlock()
	if !ok {
		return ""
	}
	if now.After(entry.ExpiresAt) {
		deleteSession(token)
		return ""
	}
	return token
}

// setSessionCookie は、セッション Cookie をレスポンスへ設定します。
//
// 機能:
//   - トークンを HttpOnly Cookie としてクライアントへ保存する
//   - TLS の有無に応じて Secure 属性を切り替える
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//   - token string: セッショントークン
//   - expiresAt time.Time: 失効時刻
//
// 返り値およびその型:
//   - なし
func setSessionCookie(c *gin.Context, token string, expiresAt time.Time) {
	secure := c.Request.TLS != nil
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     sessionCookieName,
		Value:    token,
		Path:     "/",
		Expires:  expiresAt,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

// clearSessionCookie は、セッション Cookie を削除するための Set-Cookie を返します。
//
// 機能:
//   - Cookie の有効期限を過去に設定してクライアント側の保持を解除する
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - なし
func clearSessionCookie(c *gin.Context) {
	secure := c.Request.TLS != nil
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     sessionCookieName,
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   secure,
		SameSite: http.SameSiteLaxMode,
	})
}

// requireLoginMiddleware は、ログイン必須のルートに適用するミドルウェアを返します。
//
// 機能:
//   - 設定により認証が無効な場合は常に通過させる
//   - セッション Cookie が不正な場合は HTML は /login へリダイレクトし、API は 401 を返す
//
// 引数およびその型:
//   - なし
//
// 返り値およびその型:
//   - gin.HandlerFunc: Gin ミドルウェア
func requireLoginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if !isAuthEnabled() {
			c.Next()
			return
		}

		if getValidSessionToken(c) != "" {
			c.Next()
			return
		}

		if isHTMLRequest(c) {
			c.Redirect(http.StatusFound, "/login")
			c.Abort()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "認証が必要です",
		})
		c.Abort()
	}
}

// isHTMLRequest は、HTML レスポンスを期待するリクエストかどうかを判定します。
//
// 機能:
//   - Accept ヘッダーに text/html が含まれる場合に true を返す
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - bool: HTML リクエストの場合 true、それ以外は false
func isHTMLRequest(c *gin.Context) bool {
	accept := c.GetHeader("Accept")
	return strings.Contains(accept, "text/html")
}

// registerAuthRoutes は、ログイン関連のルートを登録します。
//
// 機能:
//   - /login の表示と認証処理を登録する
//   - /logout のセッション破棄処理を登録する
//
// 引数およびその型:
//   - rt *gin.Engine: ルートを登録する Gin エンジン
//
// 返り値およびその型:
//   - なし
func registerAuthRoutes(rt *gin.Engine) {
	rt.GET("/login", handleLoginGET)
	rt.POST("/login", handleLoginPOST)
	rt.POST("/logout", handleLogoutPOST)
}

// handleLoginGET は、ログイン画面を表示します。
//
// 機能:
//   - 認証が無効な場合はトップへリダイレクトする
//   - 認証が有効な場合は login.html を返す
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - なし
func handleLoginGET(c *gin.Context) {
	if !isAuthEnabled() {
		c.Redirect(http.StatusFound, "/")
		return
	}

	hasError := c.Query("error") == "1"
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "ログイン",
		"error": hasError,
	})
}

// handleLoginPOST は、ID/パスワードを検証し、セッションを発行します。
//
// 機能:
//   - 送信された id/pw を [SYSTEM] の SYSID/SYSPW と照合する
//   - 一致した場合はセッション Cookie を発行し、HTML は / へリダイレクトする
//   - 不一致の場合は JSON は 401、フォーム送信は /login?error=1 へリダイレクトする
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - なし
func handleLoginPOST(c *gin.Context) {
	if !isAuthEnabled() {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "認証が無効です",
		})
		return
	}

	var req ReqLoginPOST
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "リクエストが不正です",
		})
		return
	}

	if req.ID == "" {
		req.ID = req.Username
	}
	if req.PW == "" {
		req.PW = req.Password
	}

	if req.ID != cfg.System.SysID || req.PW != cfg.System.SysPW {
		if isJSONRequest(c) {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "ID またはパスワードが違います",
			})
			return
		}
		c.Redirect(http.StatusFound, "/login?error=1")
		return
	}

	token, err := issueSessionToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "セッションの発行に失敗しました",
		})
		return
	}
	expiresAt := storeSession(token, sessionTTL)
	setSessionCookie(c, token, expiresAt)

	if isJSONRequest(c) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
		return
	}
	c.Redirect(http.StatusFound, "/")
}

// handleLogoutPOST は、セッションを破棄してログアウトします。
//
// 機能:
//   - セッションストアから該当トークンを削除する
//   - Cookie を削除し、/login へリダイレクトする
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - なし
func handleLogoutPOST(c *gin.Context) {
	if token := getValidSessionToken(c); token != "" {
		deleteSession(token)
	}
	clearSessionCookie(c)

	if isJSONRequest(c) {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
		})
		return
	}
	c.Redirect(http.StatusFound, "/login")
}

// isJSONRequest は、JSON を期待するリクエストかどうかを判定します。
//
// 機能:
//   - Accept / Content-Type を参照し、application/json の場合に true を返す
//
// 引数およびその型:
//   - c *gin.Context: リクエストコンテキスト
//
// 返り値およびその型:
//   - bool: JSON リクエストの場合 true、それ以外は false
func isJSONRequest(c *gin.Context) bool {
	accept := c.GetHeader("Accept")
	contentType := c.GetHeader("Content-Type")
	return strings.Contains(accept, "application/json") || strings.Contains(contentType, "application/json")
}
