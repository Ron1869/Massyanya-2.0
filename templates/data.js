// Получаем элементы страницы
const refreshDataButton = document.getElementById('refreshDataButton');
const dataTable = document.getElementById('dataTable').getElementsByTagName('tbody')[0];

// Пример данных для обновления
const mockData = [
    { name: 'BTC/USD', date: '02.12.2024', value: 92000.50 },
    { name: 'ETH/USD', date: '02.12.2024', value: 4800.00 },
    { name: 'XRP/USD', date: '02.12.2024', value: 1.20 },
];

// Функция для обновления данных таблицы
function updateTable(data) {
    // Очищаем таблицу
    dataTable.innerHTML = '';

    // Добавляем новые строки
    data.forEach(item => {
        const row = dataTable.insertRow();
        const cellName = row.insertCell(0);
        const cellDate = row.insertCell(1);
        const cellValue = row.insertCell(2);

        cellName.textContent = item.name;
        cellDate.textContent = item.date;
        cellValue.textContent = item.value;
    });
}

// Обработчик кнопки "Обновить данные"
refreshDataButton.addEventListener('click', () => {
    console.log('Обновление данных...');
    updateTable(mockData); // Здесь можно заменить на реальный запрос данных
});
