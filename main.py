import sys
from flask import Flask, request, jsonify
from extensions.investpy import investpy
from extensions.news_parsing_movies import parse_news
from extensions.minio import MinioClient
from extensions.telegram_bot import TelegramBot  # Telegram бот (добавь свой модуль)
import threading  # Для запуска Telegram-бота параллельно с веб-сервером

# --- Инициализация Flask-приложения ---
app = Flask(__name__)

# --- Настройка MinIO ---
minio_client = MinioClient()
minio_client.connect()  # Проверка подключения к MinIO

# --- Инициализация Telegram-бота ---
bot = TelegramBot()

# --- Главная страница веб-интерфейса ---
@app.route("/")
def index():
    return """
    <h1>Massyanya-2.0</h1>
    <p>Добро пожаловать в веб-интерфейс!</p>
    <ul>
        <li><a href="/start-bot">Запустить Telegram-бота</a></li>
        <li><a href="/run-predictions">Запустить предсказания</a></li>
    </ul>
    """

# --- Запуск Telegram-бота ---
@app.route("/start-bot")
def start_bot():
    try:
        threading.Thread(target=bot.run).start()
        return "Telegram Bot успешно запущен!"
    except Exception as e:
        return f"Ошибка при запуске Telegram Bot: {str(e)}"

# --- Запуск предсказаний ---
@app.route("/run-predictions", methods=["GET"])
def run_predictions():
    try:
        currency_cross = request.args.get("currency_cross", "EUR/USD")
        data = investpy.get_currency_cross_recent_data(currency_cross=currency_cross)
        return jsonify({"message": "Предсказания успешно выполнены!", "data": data.head().to_dict()})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# --- Парсинг новостей ---
@app.route("/parse-news", methods=["GET"])
def parse_news_route():
    try:
        url = request.args.get("url", "https://example.com")
        news = parse_news(url)
        return jsonify({"message": "Новости успешно собраны!", "news": news})
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# --- Точка входа ---
if __name__ == "__main__":
    try:
        print("Инициализация проекта Massyanya-2.0...")
        print("Проверка MinIO подключения...")
        minio_client.connect()
        print("MinIO успешно подключен!")
    except Exception as e:
        print(f"Ошибка подключения MinIO: {str(e)}")
        sys.exit(1)

    # Запуск веб-сервера
    app.run(host="0.0.0.0", port=5000)
