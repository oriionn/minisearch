let element = document.getElementById("result");
let query = document.querySelector(`input[name="q"]`)
let result = eval(query.value);

element.textContent = result;
