const ORDERBOOK_ENDPOINT = "/book";

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
        const isLongChecked = cancelLong.checked;
        const isShortChecked = cancelShort.checked;

        let backgroundColor = "";
        if (isLongChecked && isShortChecked) {
            backgroundColor = "#A15FC2";
        } else if (isLongChecked) {
            backgroundColor = "#FF6685";
        } else if (isShortChecked) {
            backgroundColor = "#4258FF";
        }

        cancelContainer.style.backgroundColor = backgroundColor;
    };

    cancelLong.addEventListener("change", updateBackground);
    cancelShort.addEventListener("change", updateBackground);
    updateBackground();
}

setupRowClick();
startOrderBookPolling();
setupCancelState();
