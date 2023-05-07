let orderEl = document.getElementById("json_print")
let order = orderEl.textContent
order = JSON.stringify(JSON.parse(order), undefined, 2)
orderEl.textContent = order