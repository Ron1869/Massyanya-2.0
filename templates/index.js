// Элементы страницы
const addCoinButton = document.getElementById('addCoinButton');
const removeCoinButton = document.getElementById('removeCoinButton');
const addCoinInput = document.getElementById('addCoinInput');
const removeCoinInput = document.getElementById('removeCoinInput');
const savedCoinsList = document.getElementById('savedCoinsList');

// Пример сохранённых монет
let savedCoins = ['BTC', 'ETH', 'XRP'];

// Функция для рендера списка сохранённых монет
function renderSavedCoins() {
    savedCoinsList.innerHTML = '';

    if (savedCoins.length === 0) {
        savedCoinsList.innerHTML = '<li>Нет сохранённых монет</li>';
        return;
    }

    savedCoins.forEach((coin, index) => {
        const listItem = document.createElement('li');
        listItem.textContent = coin;

        const deleteButton = document.createElement('button');
        deleteButton.textContent = 'Удалить';
        deleteButton.style.marginLeft = '10px';
        deleteButton.addEventListener('click', () => {
            savedCoins.splice(index, 1);
            renderSavedCoins();
        });

        listItem.appendChild(deleteButton);
        savedCoinsList.appendChild(listItem);
    });
}

// Обработчик кнопки "Добавить монету"
addCoinButton.addEventListener('click', () => {
    const newCoin = addCoinInput.value.trim().toUpperCase();
    if (newCoin && !savedCoins.includes(newCoin)) {
        savedCoins.push(newCoin);
        addCoinInput.value = '';
        renderSavedCoins();
    }
});

// Обработчик кнопки "Удалить монету"
removeCoinButton.addEventListener('click', () => {
    const coinToRemove = removeCoinInput.value.trim().toUpperCase();
    if (coinToRemove && savedCoins.includes(coinToRemove)) {
        savedCoins = savedCoins.filter((coin) => coin !== coinToRemove);
        removeCoinInput.value = '';
        renderSavedCoins();
    }
});

// static/js/index.js
document.getElementById('startButton').addEventListener('click', async () => {
    const response = await fetch('/api/trading/start-trading', { method: 'POST' });
    const result = await response.json();
    alert(result.message);
});

document.getElementById('stopButton').addEventListener('click', async () => {
    const response = await fetch('/api/trading/stop-trading', { method: 'POST' });
    const result = await response.json();
    alert(result.message);
});

document.getElementById("startButton").addEventListener("click", () => {
    fetch("/start-trading", { method: "POST" })
        .then(response => response.json())
        .then(data => alert(data.status))
        .catch(err => console.error(err));
});

document.getElementById("stopButton").addEventListener("click", () => {
    fetch("/stop-trading", { method: "POST" })
        .then(response => response.json())
        .then(data => alert(data.status))
        .catch(err => console.error(err));
});


// Инициализация
document.addEventListener('DOMContentLoaded', () => {
    renderSavedCoins();
});
