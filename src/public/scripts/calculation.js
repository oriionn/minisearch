let element = document.getElementById("result");
let result = eval(document.querySelector(`input[name="q"]`).value);

element.textContent = result;
