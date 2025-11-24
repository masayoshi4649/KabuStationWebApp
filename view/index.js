const ORDERBOOK_ENDPOINT = "/book";

/*
### 機能
- `/book` から取得した板データをHTMLテーブルへ描画する。

### 引数およびその型
- `data` Array<Object> - 取得した板データ配列。

### 返り値およびその型
- なし
*/
function renderOrderBook(data) {
    const tbody = document.getElementById("orderbook-body");
    tbody.innerHTML = "";

    data.forEach(item => {
        const tr = document.createElement("tr");
        tr.classList.add("orderbook-row");
        tr.dataset.price = item.Price;

        // 売気配
        const askTd = document.createElement("td");
        askTd.classList.add(
            "orderbook-ask",
            "has-text-link",
            "has-text-right"
        );
        askTd.textContent = item.SellQty > 0 ? item.SellQty.toLocaleString() : "";

        // 価格（中央寄せ）
        const priceTd = document.createElement("td");
        priceTd.classList.add(
            "orderbook-price",
            "has-text-centered"   // ★ 追加
        );
        priceTd.textContent = item.Price.toFixed(2);

        if (item.Current) {
            priceTd.classList.add(
                "has-background-primary-dark",
                "has-text-white",
                "has-text-weight-bold"
            );
        }

        // 買気配
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
/*
### 機能
- テーブル行クリック時に価格をコンソールへ出力する。

### 引数およびその型
- なし

### 返り値およびその型
- なし
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
/*
### 機能
- 板データをAPIから取得し、描画する。

### 引数およびその型
- なし

### 返り値およびその型
- Promise<void>
*/
async function fetchOrderBook() {
    try {
        const response = await axios.get(ORDERBOOK_ENDPOINT);
        renderOrderBook(Array.isArray(response.data) ? response.data : []);
    } catch (error) {
        console.error("取得に失敗しました", error);
    }
}

// ----------------------------------------
/*
### 機能
- 初期描画を行い、1秒間隔で板データ取得・更新を続ける。

### 引数およびその型
- なし

### 返り値およびその型
- なし
*/
function startOrderBookPolling() {
    fetchOrderBook();
    setInterval(fetchOrderBook, 1000);
}

// ----------------------------------------
/*
### 機能
- キャンセルエリアのチェック状態に応じて背景色を切り替える。

### 引数およびその型
- なし

### 返り値およびその型
- なし
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


