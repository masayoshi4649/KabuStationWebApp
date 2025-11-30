const ORDERBOOK_ENDPOINT = "/book";

// ----------------------------------------
/**
 * チェック状態に応じた強調色を返します。
 *
 * ロング / ショートの選択状況から、表示に使う 16 進カラーコードを返します。
 * どちらも未選択の場合は背景をリセットするために空文字を返します。
 *
 * @function determineSelectionColor
 * @param {boolean} isLongChecked - 買建が選択されているか。
 * @param {boolean} isShortChecked - 売建が選択されているか。
 * @returns {string} 選択状態に応じた 16 進カラーコード。未選択時は空文字。
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
 * 指定要素にチェック柄の背景を適用します。
 *
 * ベースカラーの上に格子状のオーバーレイを重ね、選択状態を視覚的に示します。
 *
 * @function applyCheckerboardBackground
 * @param {HTMLElement} target - 背景を適用する対象要素。
 * @param {string} baseColor - ベースとなるカラーコード。未指定の場合は背景をリセット。
 * @param {number} angleDeg - パターンの角度（度数法）。
 * @param {boolean} shiftPattern - 模様を半マスずらすかどうか。
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
 * 指定要素にストライプ柄の背景を適用します。
 *
 * ベースカラーの上に平行なストライプを重ね、選択状態を視覚的に示します。
 *
 * @function applyStripedBackground
 * @param {HTMLElement} target - 背景を適用する対象要素。
 * @param {string} baseColor - ベースとなるカラーコード。未指定の場合は背景をリセット。
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
 * 取得した板情報をテーブルへ描画します。
 *
 * 1 行を 1 tick として生成し、価格セルには data-price を付与します。
 *
 * @function renderOrderBook
 * @param {Array<Object>} data - API から受け取った板データ配列。
 * @returns {void} 返り値はありません。
 */
function renderOrderBook(data) {
    const tbody = document.getElementById("orderbook-body");
    tbody.innerHTML = "";

    data.forEach((item) => {
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

        const askTd = document.createElement("td");
        askTd.classList.add(
            "orderbook-ask",
            "px-3",
            "py-1.5",
            "text-right",
            "text-emerald-200",
            "font-semibold",
            "bg-emerald-500/5",
            "first:rounded-l-2xl",
            "leading-tight"
        );
        askTd.textContent = item.SellQty > 0 ? item.SellQty.toLocaleString() : "";

        const priceTd = document.createElement("td");
        priceTd.classList.add(
            "orderbook-price",
            "px-3",
            "py-1.5",
            "text-center",
            "font-semibold",
            "text-slate-100",
            "tracking-tight",
            "bg-slate-800/60",
            "leading-tight"
        );
        priceTd.textContent = Number.isFinite(item.Price) ? item.Price.toFixed(TICK_CONFIG.decimals) : "--";

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

        const bidTd = document.createElement("td");
        bidTd.classList.add(
            "orderbook-bid",
            "px-3",
            "py-1.5",
            "text-left",
            "text-rose-200",
            "font-semibold",
            "bg-rose-500/5",
            "last:rounded-r-2xl",
            "leading-tight"
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
 * テーブル行のクリックイベントを設定し、開始価格へ反映します。
 *
 * @function setupRowClick
 * @returns {void} 返り値はありません。
 */
function setupRowClick() {
    const tbody = document.getElementById("orderbook-body");
    const startInput = document.getElementById("immediate__start_price");
    tbody.addEventListener("click", (event) => {
        const tr = event.target.closest("tr.orderbook-row");
        if (!tr) {
            return;
        }
        const price = parseFloat(tr.dataset.price);
        if (!Number.isFinite(price)) {
            return;
        }
        if (startInput) {
            startInput.value = price.toFixed(TICK_CONFIG.decimals);
            startInput.dispatchEvent(new Event("input"));
        }
    });
}

// ----------------------------------------
/**
 * 板情報を API から取得し、描画します。
 *
 * @function fetchOrderBook
 * @returns {Promise<void>} 描画完了までの Promise。
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
 * 板情報の定期取得を開始します。
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
 * 注文取消カードの背景更新を設定します。
 *
 * @function setupCancelState
 * @returns {void} 返り値はありません。
 */
function setupCancelState() {
    const cancelSurface = document.getElementById("cancel__selection_surface");
    const cancelLong = document.getElementById("cancel__long");
    const cancelShort = document.getElementById("cancel__short");

    if (!cancelSurface || !cancelLong || !cancelShort) {
        return;
    }

    const updateBackground = () => {
        const backgroundColor = determineSelectionColor(cancelLong.checked, cancelShort.checked);
        applyStripedBackground(cancelSurface, backgroundColor);
    };

    cancelLong.addEventListener("change", updateBackground);
    cancelShort.addEventListener("change", updateBackground);
    updateBackground();
}

// ----------------------------------------
/**
 * 注文取消ボタンの送信処理を設定します。
 *
 * - 現在のチェック状態を payload として /order/cancel へ送信します。
 * - 送信後はチェックをクリアし、背景も更新します。
 * - 結果に応じてトースト通知を表示します。
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
            short: cancelShort.checked,
        };

        cancelLong.checked = false;
        cancelShort.checked = false;
        cancelLong.dispatchEvent(new Event("change"));
        cancelShort.dispatchEvent(new Event("change"));

        try {
            const response = await axios.post("/order/cancel", payload);
            if (response.status === 200) {
                iziToast.success({
                    title: "成功",
                    message: "取消リクエストを送信しました",
                });
            } else {
                iziToast.error({
                    title: "失敗",
                    message: "取消リクエストを送信できませんでした",
                });
            }
        } catch (error) {
            console.error("注文取消の送信に失敗しました", error);
            iziToast.error({
                title: "失敗",
                message: "取消リクエストを送信できませんでした",
            });
        }
    });
}

// ----------------------------------------
/**
 * 建玉返済カードの背景更新を設定します。
 *
 * @function setupCloseState
 * @returns {void} 返り値はありません。
 */
function setupCloseState() {
    const closeSurface = document.getElementById("close__selection_surface");
    const closeLong = document.getElementById("close__long");
    const closeShort = document.getElementById("close__short");
    const closeOnlyProfit = document.getElementById("close__only_profit");

    if (!closeSurface || !closeLong || !closeShort || !closeOnlyProfit) {
        return;
    }

    const updateBackground = () => {
        const hasProfitOnly = closeOnlyProfit.checked;
        const baseColor = determineSelectionColor(closeLong.checked, closeShort.checked);
        const angleDeg = hasProfitOnly ? 45 : 0;
        const shiftPattern = false;

        applyCheckerboardBackground(closeSurface, baseColor, angleDeg, shiftPattern);
    };

    closeOnlyProfit.addEventListener("change", updateBackground);
    closeLong.addEventListener("change", updateBackground);
    closeShort.addEventListener("change", updateBackground);
    updateBackground();
}

// ----------------------------------------
/**
 * 建玉返済ボタンの送信処理を設定します。
 *
 * - long/short/only_profit を payload にまとめ、/order/close へ送信します。
 * - 送信後はサイドをリセットし、含み益のみをオンに戻します。
 * - トーストで結果を通知します。
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
            only_profit: closeOnlyProfit.checked,
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
                    message: "返済リクエストを送信しました",
                });
            } else {
                iziToast.error({
                    title: "失敗",
                    message: "返済リクエストを送信できませんでした",
                });
            }
        } catch (error) {
            console.error("建玉返済の送信に失敗しました", error);
            iziToast.error({
                title: "失敗",
                message: "返済リクエストを送信できませんでした",
            });
        }
    });
}

// ----------------------------------------
/**
 * 1 tick あたりの値幅と小数点桁数を取得します。
 *
 * @function getTickConfig
 * @returns {{tickValue: number, decimals: number}} 1 tick の値幅と小数点桁数。
 */
function getTickConfig() {
    const tickRaw = document.body?.dataset?.onetick;
    const parsed = parseFloat(tickRaw);
    const tickValue = Number.isFinite(parsed) && parsed > 0 ? parsed : 0.25;
    const decimals = tickValue < 1 ? Math.max(2, `${tickValue}`.split(".")[1]?.length || 2) : 2;
    return { tickValue, decimals };
}

const TICK_CONFIG = getTickConfig();

// ----------------------------------------
/**
 * 入力値を検証し、1 tick 単位で正規化します。
 *
 * @function normalizeInterval
 * @param {HTMLInputElement} input - 間隔入力フィールド。
 * @returns {number} コマンドで扱う正規化済みの数値。
 */
function normalizeInterval(input) {
    const parsed = parseFloat(input?.value);
    if (Number.isFinite(parsed) && parsed > 0) {
        return parsed;
    }
    if (input) {
        input.value = TICK_CONFIG.tickValue.toFixed(TICK_CONFIG.decimals);
    }
    return TICK_CONFIG.tickValue;
}

// ----------------------------------------
/**
 * 指定条件に基づき、プレビュー表示用の文字列を作成します。
 *
 * @function buildPricePreview
 * @param {number} startPrice - 開始価格。
 * @param {number} step - 間隔。
 * @param {number} size - 件数。
 * @param {boolean} ascending - true の場合は昇順、false の場合は降順。
 * @returns {string} プレビューに表示するテキスト。
 */
function buildPricePreview(startPrice, step, size, ascending) {
    if (!Number.isFinite(startPrice) || !Number.isFinite(step)) {
        return "--";
    }
    const safeSize = Math.max(1, Math.floor(size));
    const arrow = ascending ? "↑" : "↓";
    const direction = ascending ? 1 : -1;
    const maxPreview = 2;
    const values = [];

    for (let i = 0; i < Math.min(safeSize, maxPreview); i += 1) {
        const value = startPrice + direction * step * i;
        values.push(value.toFixed(TICK_CONFIG.decimals));
    }

    if (safeSize > maxPreview) {
        const last = startPrice + direction * step * (safeSize - 1);
        values.push("…", last.toFixed(TICK_CONFIG.decimals));
    }

    return values.join(` ${arrow} `);
}

// ----------------------------------------
/**
 * 即時注文の計算・プレビュー更新・送信補助を設定します。
 *
 * @function setupImmediateCalculator
 * @returns {void}
 */
function setupImmediateCalculator() {
    const startInput = document.getElementById("immediate__start_price");
    const intervalInput = document.getElementById("immediate__interval");
    const sizeInput = document.getElementById("immediate__size");
    const previewLabel = document.getElementById("immediate__preview_label");
    const previewText = document.getElementById("immediate__preview_text");
    const previewBox = document.getElementById("immediate__preview_box");
    const intervalUp = document.getElementById("immediate__interval_up");
    const intervalDown = document.getElementById("immediate__interval_down");
    const directionInputs = Array.from(document.querySelectorAll("input[name='immediate__direction']"));
    const typeInputs = Array.from(document.querySelectorAll("input[name='immediate__type']"));
    const sizeButtons = Array.from(document.querySelectorAll("[data-immediate-size-delta]"));
    const startButtons = Array.from(document.querySelectorAll("[data-immediate-start-step]"));

    if (!startInput || !intervalInput || !sizeInput || !previewLabel || !previewText || !previewBox) {
        return;
    }

    const ensureStartPrice = () => {
        const parsed = parseFloat(startInput.value);
        if (Number.isFinite(parsed)) {
            return parsed;
        }
        const fallback = parseFloat(document.body?.dataset?.current);
        const validFallback = Number.isFinite(fallback) ? fallback : TICK_CONFIG.tickValue;
        startInput.value = validFallback.toFixed(TICK_CONFIG.decimals);
        return validFallback;
    };

    const formatTickDelta = (tickSteps) => {
        const value = tickSteps * TICK_CONFIG.tickValue;
        const fixed = value.toFixed(TICK_CONFIG.decimals);
        const numeric = Number.isFinite(Number(fixed)) ? Number(fixed) : 0;
        const sign = numeric >= 0 ? "+" : "";
        return `${sign}${numeric.toString()}`;
    };

    const ensureSize = () => {
        const parsed = parseInt(sizeInput.value, 10);
        if (Number.isFinite(parsed) && parsed > 0) {
            return parsed;
        }
        sizeInput.value = "1";
        return 1;
    };

    const applyPreviewTheme = (dir, type) => {
        const themes = {
            buy_limit: ["bg-gradient-to-b", "from-transparent", "to-rose-500/40", "text-white"],
            buy_stop: ["bg-gradient-to-t", "from-transparent", "to-rose-500/40", "text-white"],
            sell_limit: ["bg-gradient-to-t", "from-transparent", "to-blue-500/40", "text-white"],
            sell_stop: ["bg-gradient-to-b", "from-transparent", "to-blue-500/40", "text-white"],
        };
        previewBox.className = "mt-3 rounded-xl bg-slate-950/70 p-3 text-slate-100 shadow";
        const key = `${dir}_${type}`;
        if (themes[key]) {
            previewBox.classList.add(...themes[key]);
        }
    };

    const update = () => {
        const baseInterval = normalizeInterval(intervalInput);
        const size = ensureSize();
        const directionInput = directionInputs.find((input) => input.checked);
        const typeInput = typeInputs.find((input) => input.checked);
        const dir = directionInput ? directionInput.value : "buy";
        const type = typeInput ? typeInput.value : "limit";

        const ascending = dir === "buy" ? type === "stop" : type === "limit";
        const startPrice = ensureStartPrice();

        const preview = buildPricePreview(startPrice, baseInterval, size, ascending);
        const label = `${dir === "buy" ? "買い" : "売り"}${type === "limit" ? "指値" : "逆指値"}`;

        previewLabel.textContent = label;
        previewText.textContent = preview;
        applyPreviewTheme(dir, type);
    };

    const adjustInterval = (delta) => {
        const current = normalizeInterval(intervalInput);
        const next = Math.max(current + delta, TICK_CONFIG.tickValue);
        intervalInput.value = next.toFixed(TICK_CONFIG.decimals);
        update();
    };

    const adjustSize = (delta) => {
        const current = ensureSize();
        const next = Math.max(1, current + delta);
        sizeInput.value = next.toString();
        update();
    };

    const adjustStartPrice = (tickSteps) => {
        const base = ensureStartPrice();
        const next = base + tickSteps * TICK_CONFIG.tickValue;
        startInput.value = next.toFixed(TICK_CONFIG.decimals);
        update();
    };

    const applyStartButtonLabels = () => {
        startButtons.forEach((btn) => {
            const tickSteps = parseInt(btn.dataset.immediateStartStep, 10);
            if (!Number.isFinite(tickSteps)) {
                return;
            }
            btn.textContent = formatTickDelta(tickSteps);
        });
    };

    [startInput, intervalInput, sizeInput, ...directionInputs, ...typeInputs].forEach((el) => {
        el?.addEventListener("input", update);
        el?.addEventListener("change", update);
    });

    intervalUp?.addEventListener("click", () => adjustInterval(TICK_CONFIG.tickValue));
    intervalDown?.addEventListener("click", () => adjustInterval(-TICK_CONFIG.tickValue));
    startButtons.forEach((btn) => {
        const tickSteps = parseInt(btn.dataset.immediateStartStep, 10);
        if (!Number.isFinite(tickSteps)) {
            return;
        }
        btn.addEventListener("click", () => adjustStartPrice(tickSteps));
    });
    sizeButtons.forEach((btn) => {
        btn?.addEventListener("click", () => {
            const delta = parseInt(btn.dataset.immediateSizeDelta, 10);
            if (Number.isFinite(delta)) {
                adjustSize(delta);
            }
        });
    });

    applyStartButtonLabels();
    update();
}

setupRowClick();
startOrderBookPolling();
setupCancelState();
setupCloseState();
setupCancelButton();
setupCloseButton();
setupImmediateCalculator();
