// Элементы формы
const keysForm = document.getElementById('keysForm');
const keysMessage = document.getElementById('keysMessage');
const resetButton = document.getElementById('resetButton');

// Валидация IP-адреса
function validateIPAddress(ip) {
    const ipRegex = /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/;
    return ipRegex.test(ip);
}

// Обработчик отправки формы
keysForm.addEventListener('submit', (event) => {
    event.preventDefault(); // Предотвращаем перезагрузку страницы

    // Получаем значения полей
    const brokerIP = document.getElementById('broker-ip').value.trim();
    const apiKey = document.getElementById('api-key').value.trim();
    const apiSecret = document.getElementById('api-secret').value.trim();

    // Валидация данных
    if (!validateIPAddress(brokerIP)) {
        showMessage('Ошибка: Некорректный IP-адрес.', 'error');
        return;
    }

    if (!apiKey || !apiSecret) {
        showMessage('Ошибка: Поля API-ключа и секретного ключа обязательны.', 'error');
        return;
    }

    // Имитация сохранения данных
    console.log('Сохраненные данные:', { brokerIP, apiKey, apiSecret });
    showMessage('Данные успешно сохранены!', 'success');
});

// Обработчик кнопки "Сбросить"
resetButton.addEventListener('click', () => {
    document.getElementById('broker-ip').value = '';
    document.getElementById('api-key').value = '';
    document.getElementById('api-secret').value = '';
    showMessage('Форма успешно сброшена.', 'success');
});

// Функция для отображения сообщения
function showMessage(message, type) {
    keysMessage.textContent = message;
    keysMessage.style.display = 'block';
    keysMessage.style.color = type === 'success' ? 'green' : 'red';

    // Скрываем сообщение через 3 секунды
    setTimeout(() => {
        keysMessage.style.display = 'none';
    }, 3000);
}
