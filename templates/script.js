// Получаем элементы чата и кнопки
const chatBox = document.getElementById('chatBox');
const sendButton = document.getElementById('sendButton');

// Функция отправки команды
function sendMessage() {
    const message = chatBox.innerText.trim(); // Получаем текст из чата
    if (message !== "") {
        // Здесь обработка команды (например, отправка на сервер или добавление в историю)
        console.log(`Команда отправлена: ${message}`);

        // Добавляем сообщение в чат
        const chatHistory = document.createElement('p');
        chatHistory.textContent = `Вы: ${message}`;
        chatBox.appendChild(chatHistory);

        // Очищаем поле ввода
        chatBox.innerText = "";
    }
}

// Обработка кнопки "Отправить"
sendButton.addEventListener('click', sendMessage);

// Обработка клавиши Enter в чате
chatBox.addEventListener('keydown', (event) => {
    if (event.key === 'Enter') {
        event.preventDefault(); // Предотвращаем перенос строки
        sendMessage();
    }
});

// Пример работы с настройками
document.getElementById('settingsForm').addEventListener('submit', (event) => {
    event.preventDefault();
    const market = document.getElementById('marketSelect').value;
    const addCoin = document.getElementById('addCoinInput').value;
    const removeCoin = document.getElementById('removeCoinInput').value;
    const timeframe = document.getElementById('timeframeSelect').value;
    const forecast = document.getElementById('forecastSettings').value;

    console.log('Настройки сохранены:', { market, addCoin, removeCoin, timeframe, forecast });

    // Логика сохранения данных...
});
