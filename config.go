package main

import (
	"fmt"
	"strings"
)

// validateConfig は、設定値の整合性チェックを行います。
//
// 機能:
//   - [SYSTEM] の SYSID/SYSPW が片方のみ指定されていないか検証する
//
// 引数およびその型:
//   - cfg Config: 検証対象の設定値
//
// 返り値およびその型:
//   - error: 設定が不正な場合はエラー、それ以外は nil
func validateConfig(cfg Config) error {
	sysID := strings.TrimSpace(cfg.System.SysID)
	sysPW := strings.TrimSpace(cfg.System.SysPW)

	if (sysID == "") != (sysPW == "") {
		return fmt.Errorf("[SYSTEM] は SYSID と SYSPW を両方指定してください")
	}
	return nil
}
