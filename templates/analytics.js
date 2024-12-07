// Элементы управления графиком
const chartTypeSelect = document.getElementById('chartTypeSelect');
const timeframeSelect = document.getElementById('timeframeSelect');
const updateChartButton = document.getElementById('updateChartButton');

// Данные для графика
let chartData = [50000, 50500, 51000, 52000, 53000, 54000]; // Пример данных
const labels = ["1 мин", "5 мин", "15 мин", "1 час", "4 часа", "1 день"];

// Настройки графика
let chart;
function createChart(type = 'line') {
    const ctx = document.getElementById('predictionChart').getContext('2d');

    // Если график уже существует, удаляем его перед созданием нового
    if (chart) chart.destroy();

    chart = new Chart(ctx, {
        type: type,
        data: {
            labels: labels,
            datasets: [{
                label: 'Прогноз цены BTC/USDT',
                data: chartData,
                borderColor: 'rgba(255, 111, 0, 1)',
                backgroundColor: 'rgba(255, 111, 0, 0.2)',
                tension: 0.4,
            }]
        },
        options: {
            responsive: true,
            plugins: {
                legend: {
                    display: true,
                    position: 'top',
                },
            },
            scales: {
                x: {
                    title: {
                        display: true,
                        text: 'Таймфрейм'
                    }
                },
                y: {
                    title: {
                        display: true,
                        text: 'Цена (USDT)'
                    }
                }
            }
        }
    });
}

// Инициализация графика
createChart();

// Обновление графика при изменении параметров
updateChartButton.addEventListener('click', () => {
    const selectedType = chartTypeSelect.value;
    const selectedTimeframe = timeframeSelect.value;

    // Здесь можно добавить логику загрузки новых данных
    console.log(`Обновляем график: Тип - ${selectedType}, Таймфрейм - ${selectedTimeframe}`);

    // Пример изменения данных
    chartData = chartData.map(value => value + Math.random() * 1000 - 500); // Имитация новых данных
    createChart(selectedType); // Перестраиваем график
});
