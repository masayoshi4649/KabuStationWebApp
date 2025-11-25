const ORDERBOOK_ENDPOINT = "/book";

// ----------------------------------------
/**
 * チェック状態に応じてキャンセル/決済エリアで使う基調色を決定する。
 *
 * 機能:
 *   - ロングとショートの選択組み合わせに応じて 3 パターンの色を返す
 *   - 未選択時は空文字を返し背景を初期化しやすくする
 *
 * @function determineSelectionColor
 * @param {boolean} isLongChecked - ロングが選択されているかどうか。
 * @param {boolean} isShortChecked - ショートが選択されているかどうか。
 * @returns {string} 選択状態に応じた 16 進カラーコード。未選択の場合は空文字を返します。
 */
function determineSelectionColor(isLongChecked, isShortChecked) {
    if (isLongChecked && isShortChecked) {
        return "#A15FC2";
    }
    if (isLongChecked) {
        return "#FF6685";
    }
    if (isShortChecked) {
        return "#4258FF";
    }
    return "";
}

// ----------------------------------------
/**
 * 指定した要素に粗めの斜めチェッカーボード背景を適用する。
 *
 * 機能:
 *   - 基調色を背景色として設定し、半透明の白でパターンを重ねる
 *   - 粗めのマス目 (約 28px 角) で背景の変化を視認しやすくする
 *   - 基調色が空の場合は背景設定をリセットする
 *
 * @function applyCheckerboardBackground
 * @param {HTMLElement} target - 背景を適用する要素。
 * @param {string} baseColor - 基調となるカラーコード。空文字の場合はリセット。
 * @returns {void} 返り値はありません。
 */
function applyCheckerboardBackground(target, baseColor) {
    if (!target) {
        return;
    }

    if (!baseColor) {
        target.style.backgroundColor = "";
        target.style.backgroundImage = "";
        target.style.backgroundSize = "";
        target.style.backgroundPosition = "";
        return;
    }

    const patternOpacity = 0.28;
    const patternSize = 28;
    const halfSize = patternSize / 2;
    const overlay = `linear-gradient(45deg, rgba(255, 255, 255, ${patternOpacity}) 25%, transparent 25%, transparent 75%, rgba(255, 255, 255, ${patternOpacity}) 75%, rgba(255, 255, 255, ${patternOpacity}))`;

    target.style.backgroundColor = baseColor;
    target.style.backgroundImage = `${overlay}, ${overlay}`;
    target.style.backgroundPosition = `0 0, ${halfSize}px ${halfSize}px`;
    target.style.backgroundSize = `${patternSize}px ${patternSize}px`;
}

// ----------------------------------------
/**
 * 指定した要素に横線のストライプ背景を適用する。
 *
 * 機能:
 *   - 基調色を背景に設定し、半透明の白い横線を重ねる
 *   - 粗めのピッチで状態差分を視認しやすくする
 *   - 基調色が空の場合は背景設定をリセットする
 *
 * @function applyStripedBackground
 * @param {HTMLElement} target - 背景を適用する要素。
 * @param {string} baseColor - 基調となるカラーコード。空文字の場合はリセット。
 * @returns {void} 返り値はありません。
 */
function applyStripedBackground(target, baseColor) {
    if (!target) {
        return;
    }

    if (!baseColor) {
        target.style.backgroundColor = "";
        target.style.backgroundImage = "";
        target.style.backgroundSize = "";
        target.style.backgroundPosition = "";
        return;
    }

    const stripeOpacity = 0.26;
    const stripeSize = 18;
    const overlay = `repeating-linear-gradient(180deg, rgba(255, 255, 255, ${stripeOpacity}) 0, rgba(255, 255, 255, ${stripeOpacity}) ${stripeSize / 2}px, transparent ${stripeSize / 2}px, transparent ${stripeSize}px)`;

    target.style.backgroundColor = baseColor;
    target.style.backgroundImage = overlay;
    target.style.backgroundPosition = "0 0";
    target.style.backgroundSize = `${stripeSize}px ${stripeSize}px`;
}

// ----------------------------------------
/**
 * 板データを HTML テーブルへ描画する。
 *
 * 受け取った配列を行ごとに生成し、現在値行や数量に応じてクラスを付与して更新する。
 *
 * @function renderOrderBook
 * @param {Array<Object>} data - API から取得した板行データの配列。
 * @returns {void} 返り値はありません。
 */
function renderOrderBook(data) {
    const tbody = document.getElementById("orderbook-body");
    tbody.innerHTML = "";

    data.forEach(item => {
        const tr = document.createElement("tr");
        tr.classList.add("orderbook-row");
        tr.dataset.price = item.Price;

        // 売り数量セル
        const askTd = document.createElement("td");
        askTd.classList.add(
            "orderbook-ask",
            "has-text-link",
            "has-text-right"
        );
        askTd.textContent = item.SellQty > 0 ? item.SellQty.toLocaleString() : "";

        // 中央価格セル
        const priceTd = document.createElement("td");
        priceTd.classList.add(
            "orderbook-price",
            "has-text-centered"
        );
        priceTd.textContent = item.Price.toFixed(2);

        if (item.Current) {
            priceTd.classList.add(
                "has-background-primary-dark",
                "has-text-white",
                "has-text-weight-bold"
            );
        }

        // 買い数量セル
        const bidTd = document.createElement("td");
        bidTd.classList.add(
            "orderbook-bid",
            "has-text-danger",
            "has-text-left"
        );
        bidTd.textContent = item.BuyQty > 0 ? item.BuyQty.toLocaleString() : "";

        tr.appendChild(askTd);
        tr.appendChild(priceTd);
        tr.appendChild(bidTd);

        tbody.appendChild(tr);
    });
}

// ----------------------------------------
/**
 * テーブル行のクリックイベントを監視し、選択した価格をコンソールへ出力する。
 *
 * @function setupRowClick
 * @returns {void} 返り値はありません。
 */
function setupRowClick() {
    const tbody = document.getElementById("orderbook-body");
    tbody.addEventListener("click", (event) => {
        const tr = event.target.closest("tr.orderbook-row");
        if (!tr) return;
        console.log("clicked price:", tr.dataset.price);
    });
}

// ----------------------------------------
/**
 * 板データを API から取得し、描画処理を行う。
 *
 * @function fetchOrderBook
 * @returns {Promise<void>} 描画完了を示す Promise。通信に失敗した場合はエラーをログ出力する。
 */
async function fetchOrderBook() {
    try {
        const response = await axios.get(ORDERBOOK_ENDPOINT);
        renderOrderBook(Array.isArray(response.data) ? response.data : []);
    } catch (error) {
        console.error("板データの取得に失敗しました", error);
    }
}

// ----------------------------------------
/**
 * 初期描画を行い、1 秒ごとに板データを取得・更新するポーリングを開始する。
 *
 * @function startOrderBookPolling
 * @returns {void} 返り値はありません。
 */
function startOrderBookPolling() {
    fetchOrderBook();
    setInterval(fetchOrderBook, 1000);
}

// ----------------------------------------
/**
 * キャンセル選択エリアのチェック状態に応じて背景色を切り替える。
 *
 * @function setupCancelState
 * @returns {void} 返り値はありません。
 */
function setupCancelState() {
    const cancelContainer = document.getElementById("cancel");
    const cancelLong = document.getElementById("cancel__long");
    const cancelShort = document.getElementById("cancel__short");

    if (!cancelContainer || !cancelLong || !cancelShort) {
        return;
    }

    const updateBackground = () => {
        const backgroundColor = determineSelectionColor(cancelLong.checked, cancelShort.checked);
        applyStripedBackground(cancelContainer, backgroundColor);
    };

    cancelLong.addEventListener("change", updateBackground);
    cancelShort.addEventListener("change", updateBackground);
    updateBackground();
}

// ----------------------------------------
/**
 * キャンセルボタン押下時の送信処理をセットアップする。
 *
 * 機能:
 *   - 現在のチェック状態をペイロードに含めて /order/cancel へ送信する
 *   - 送信後にチェックボックスをオフへ戻し背景色を更新する
 *   - レスポンスに応じてトーストを表示する
 *
 * 引数およびその型: なし
 * 返り値およびその型: {void} 返り値はありません。
 *
 * @function setupCancelButton
 * @returns {void} 返り値はありません。
 */
function setupCancelButton() {
    const cancelButton = document.getElementById("cancel__btn");
    const cancelLong = document.getElementById("cancel__long");
    const cancelShort = document.getElementById("cancel__short");

    if (!cancelButton || !cancelLong || !cancelShort) {
        return;
    }

    cancelButton.addEventListener("click", async () => {
        const payload = {
            long: cancelLong.checked,
            short: cancelShort.checked
        };

        cancelLong.checked = false;
        cancelShort.checked = false;
        cancelLong.dispatchEvent(new Event("change"));
        cancelShort.dispatchEvent(new Event("change"));

        try {
            const response = await axios.post("/order/cancel", payload);
            if (response.status === 200) {
                iziToast.success({
                    title: '注文取消',
                    message: '送信が完了しました',
                });
            } else {
                iziToast.error({
                    title: '注文取消',
                    message: '送信が失敗しました',
                });
            }
        } catch (error) {
            console.error("注文取消の送信に失敗しました", error);
            iziToast.error({
                title: '注文取消',
                message: '送信が失敗しました',
            });
        }
    });
}

// ----------------------------------------
/**
 * 決済注文エリアのチェック状態に応じて、斜めチェッカーボード背景で配色を切り替える。
 *
 * 機能:
 *   - キャンセルと同じ組み合わせ判定で基調色を決定する
 *   - 基調色に粗めの斜めチェッカーボード柄を重ねる
 *   - チェック解除時には背景を初期状態に戻す
 *
 * @function setupCloseState
 * @returns {void} 返り値はありません。
 */
function setupCloseState() {
    const closeContainer = document.getElementById("close");
    const closeLong = document.getElementById("close__long");
    const closeShort = document.getElementById("close__short");

    if (!closeContainer || !closeLong || !closeShort) {
        return;
    }

    const updateBackground = () => {
        const baseColor = determineSelectionColor(closeLong.checked, closeShort.checked);
        applyCheckerboardBackground(closeContainer, baseColor);
    };

    closeLong.addEventListener("change", updateBackground);
    closeShort.addEventListener("change", updateBackground);
    updateBackground();
}

setupRowClick();
startOrderBookPolling();
setupCancelState();
setupCloseState();
setupCancelButton();
