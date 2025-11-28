const ORDERBOOK_ENDPOINT = "/book";

// ----------------------------------------
/**
 * チェック状態に応じた基調色を算出する。
 *
 * ロング・ショートの選択有無を判定し、強調用の 16 進カラーコードを返す。
 * どちらも未選択の場合は背景リセット用に空文字を返す。
 *
 * @function determineSelectionColor
 * @param {boolean} isLongChecked - ロングが選択されているかどうか。
 * @param {boolean} isShortChecked - ショートが選択されているかどうか。
 * @returns {string} 選択状態に対応する 16 進カラーコード。未選択時は空文字を返します。
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
 * 指定した要素に市松模様の背景を適用する。
 *
 * ベースカラーの上にグリッド状のオーバーレイを重ね、選択状態を視覚化する。
 *
 * @function applyCheckerboardBackground
 * @param {HTMLElement} target - 背景を適用する対象要素。
 * @param {string} baseColor - ベースとなるカラーコード。空文字の場合は元の背景に戻す。
 * @param {number} angleDeg - チェック枠の回転角度（度）。
 * @param {boolean} shiftPattern - 斜めチェック枠の位置を半マスずらすかどうか。
 * @returns {void} 返り値はありません。
 */
function applyCheckerboardBackground(target, baseColor, angleDeg = 0, shiftPattern = false) {
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
    const overlayPrimary = `linear-gradient(${angleDeg}deg, rgba(255, 255, 255, ${patternOpacity}) 25%, transparent 25%, transparent 75%, rgba(255, 255, 255, ${patternOpacity}) 75%, rgba(255, 255, 255, ${patternOpacity}))`;
    const overlaySecondary = `linear-gradient(${angleDeg + 90}deg, rgba(255, 255, 255, ${patternOpacity}) 25%, transparent 25%, transparent 75%, rgba(255, 255, 255, ${patternOpacity}) 75%, rgba(255, 255, 255, ${patternOpacity}))`;

    target.style.backgroundColor = baseColor;
    const offset = shiftPattern ? halfSize / 2 : 0;
    const firstPos = `${offset}px ${offset}px`;
    const secondPos = `${offset + halfSize}px ${offset + halfSize}px`;
    target.style.backgroundImage = `${overlayPrimary}, ${overlaySecondary}`;
    target.style.backgroundPosition = `${firstPos}, ${secondPos}`;
    target.style.backgroundSize = `${patternSize}px ${patternSize}px`;
}

// ----------------------------------------
/**
 * 指定した要素にストライプ模様の背景を適用する。
 *
 * ベースカラーの上に垂直ストライプを重ね、選択状態を視覚的に分かりやすくする。
 *
 * @function applyStripedBackground
 * @param {HTMLElement} target - 背景を適用する対象要素。
 * @param {string} baseColor - ベースとなるカラーコード。空文字の場合は元の背景に戻す。
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
 * 取得した板情報をテーブルへ描画する。
 *
 * 受け取った配列を 1 行ずつ組み立て、強調用のクラスや data 属性を付与する。
 *
 * @function renderOrderBook
 * @param {Array<Object>} data - API から取得した板データ配列。
 * @returns {void} 返り値はありません。
 */
function renderOrderBook(data) {
    const tbody = document.getElementById("orderbook-body");
    tbody.innerHTML = "";

    data.forEach(item => {
        const tr = document.createElement("tr");
        tr.classList.add(
            "orderbook-row",
            "rounded-2xl",
            "border",
            "border-slate-800/60",
            "bg-slate-900/60",
            "text-slate-100",
            "shadow",
            "shadow-slate-950/40",
            "transition",
            "hover:-translate-y-[1px]",
            "hover:bg-slate-800/80",
            "hover:shadow-lg",
            "hover:shadow-slate-900/50"
        );
        tr.dataset.price = item.Price;

        // 売数量セル
        const askTd = document.createElement("td");
        askTd.classList.add(
            "orderbook-ask",
            "px-3",
            "py-2",
            "text-right",
            "text-emerald-300",
            "font-semibold",
            "bg-emerald-500/5",
            "first:rounded-l-2xl"
        );
        askTd.textContent = item.SellQty > 0 ? item.SellQty.toLocaleString() : "";

        // 価格セル
        const priceTd = document.createElement("td");
        priceTd.classList.add(
            "orderbook-price",
            "px-3",
            "py-2",
            "text-center",
            "font-semibold",
            "text-slate-100",
            "tracking-tight",
            "bg-slate-800/60"
        );
        priceTd.textContent = item.Price.toFixed(2);

        if (item.Current) {
            priceTd.classList.add(
                "bg-gradient-to-r",
                "from-sky-600/70",
                "to-fuchsia-600/60",
                "text-white",
                "font-bold",
                "ring-1",
                "ring-sky-300/60",
                "shadow-lg",
                "shadow-slate-900/60"
            );
        }

        // 買数量セル
        const bidTd = document.createElement("td");
        bidTd.classList.add(
            "orderbook-bid",
            "px-3",
            "py-2",
            "text-left",
            "text-rose-300",
            "font-semibold",
            "bg-rose-500/5",
            "last:rounded-r-2xl"
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
 * テーブル行のクリックイベントを設定し、選択価格をログへ出力する。
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
 * 板情報を API から取得し、描画する。
 *
 * @function fetchOrderBook
 * @returns {Promise<void>} 描画完了までの Promise。通信に失敗した場合はエラーを出力する。
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
 * 板情報の取得を開始し、1 秒間隔でポーリングする。
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
 * 注文取消カードの背景をチェック状態に応じて更新する。
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
 * 注文取消ボタンの送信処理をセットアップする。
 *
 * - 現在のチェック状態をペイロードとして /order/cancel へ送信する。
 * - 送信後にチェック状態と背景をリセットする。
 * - 結果に応じてトーストを表示する。
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
                    title: "注文取消",
                    message: "送信が完了しました",
                });
            } else {
                iziToast.error({
                    title: "注文取消",
                    message: "送信が失敗しました",
                });
            }
        } catch (error) {
            console.error("注文取消の送信に失敗しました", error);
            iziToast.error({
                title: "注文取消",
                message: "送信が失敗しました",
            });
        }
    });
}

// ----------------------------------------
/**
 * 建玉返済カードの背景をチェック状態に応じて更新する。
 *
 * @function setupCloseState
 * @returns {void} 返り値はありません。
 */
function setupCloseState() {
    const closeContainer = document.getElementById("close");
    const closeLong = document.getElementById("close__long");
    const closeShort = document.getElementById("close__short");
    const closeOnlyProfit = document.getElementById("close__only_profit");

    if (!closeContainer || !closeLong || !closeShort || !closeOnlyProfit) {
        return;
    }

    const updateBackground = () => {
        const hasProfitOnly = closeOnlyProfit.checked;
        const baseColor = determineSelectionColor(closeLong.checked, closeShort.checked);
        const angleDeg = hasProfitOnly ? 45 : 0;
        const shiftPattern = false;

        applyCheckerboardBackground(closeContainer, baseColor, angleDeg, shiftPattern);
    };

    closeOnlyProfit.addEventListener("change", updateBackground);
    closeLong.addEventListener("change", updateBackground);
    closeShort.addEventListener("change", updateBackground);
    updateBackground();
}

// ----------------------------------------
/**
 * 取引終了の送信処理を設定します。
 *
 * - close ブロックの long/short/only_profit を /order/close へ送信します。
 * - 送信後にチェック状態を初期化し、背景更新用の change イベントを発火させます。
 * - 成否に応じてトーストで結果を知らせます。
 *
 * @function setupCloseButton
 * @returns {void} 返り値はありません。
 */
function setupCloseButton() {
    const closeButton = document.getElementById("close__btn");
    const closeLong = document.getElementById("close__long");
    const closeShort = document.getElementById("close__short");
    const closeOnlyProfit = document.getElementById("close__only_profit");

    if (!closeButton || !closeLong || !closeShort || !closeOnlyProfit) {
        return;
    }

    closeButton.addEventListener("click", async () => {
        const payload = {
            long: closeLong.checked,
            short: closeShort.checked,
            only_profit: closeOnlyProfit.checked
        };

        closeLong.checked = false;
        closeShort.checked = false;
        closeOnlyProfit.checked = true;
        closeLong.dispatchEvent(new Event("change"));
        closeShort.dispatchEvent(new Event("change"));
        closeOnlyProfit.dispatchEvent(new Event("change"));

        try {
            const response = await axios.post("/order/close", payload);
            if (response.status === 200) {
                iziToast.success({
                    title: "成功",
                    message: "送信に成功しました",
                });
            } else {
                iziToast.error({
                    title: "失敗",
                    message: "送信に失敗しました",
                });
            }
        } catch (error) {
            console.error("建玉決済ボタンの送信に失敗しました", error);
            iziToast.error({
                title: "失敗",
                message: "送信に失敗しました",
            });
        }
    });
}

setupRowClick();
startOrderBookPolling();
setupCancelState();
setupCloseState();
setupCancelButton();
setupCloseButton();
