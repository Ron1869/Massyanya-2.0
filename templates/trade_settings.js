// Получаем элементы формы
const tradeSettingsForm = document.getElementById('tradeSettingsForm');
const settingsMessage = document.getElementById('settingsMessage');

// Обработчик отправки формы
tradeSettingsForm.addEventListener('submit', (event) => {
    event.preventDefault(); // Предотвращаем перезагрузку страницы

    // Собираем данные из формы
    const settings = {
        entryAmount: parseFloat(document.getElementById('entry-amount').value),
        leverage: parseFloat(document.getElementById('leverage').value),
        stopLoss: parseFloat(document.getElementById('stop-loss').value),
        takeProfit: parseFloat(document.getElementById('take-profit').value),
        breakevenBuffer: parseFloat(document.getElementById('breakeven-buffer').value),
        volumeThreshold: parseFloat(document.getElementById('volume-threshold').value),
    };

    // Валидация данных (пример)
    if (settings.entryAmount < 1 || settings.entryAmount > 100) {
        showMessage('Ошибка: Сумма входа должна быть от 1 до 100%', 'error');
        return;
    }

    if (settings.stopLoss >= settings.takeProfit) {
        showMessage('Ошибка: Тейк-профит должен быть больше стоп-лосса!', 'error');
        return;
    }

    // Здесь можно отправить данные на сервер (заменить на реальный запрос)
    console.log('Настройки сохранены:', settings);

    // Показываем сообщение об успешном сохранении
    showMessage('Настройки успешно сохранены!', 'success');
});

// Функция для отображения сообщения
function showMessage(message, type) {
    settingsMessage.textContent = message;
    settingsMessage.style.display = 'block';
    settingsMessage.style.color = type === 'success' ? 'green' : 'red';

    // Скрываем сообщение через 3 секунды
    setTimeout(() => {
        settingsMessage.style.display = 'none';
    }, 3000);
}
