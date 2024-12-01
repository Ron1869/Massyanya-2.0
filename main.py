import sys
from flask import Flask, request, jsonify
from dotenv import load_dotenv

from extensions.investpy import investpy
from minio import Minio
from predict_stock_crypto.messaging.telegram_bot import TelegramBot
import os
import threading

# Загрузка переменных из .env
load_dotenv(dotenv_path="config/.env")
TOKEN = os.getenv("TOKEN")
CHAT_ID = os.getenv("CHATID")

# --- Инициализация Flask-приложения ---
app = Flask(__name__)

# --- Настройка MinIO ---
minio_client = Minio(
    "127.0.0.1:9000",  # Убедись, что здесь нет протокола (http://)
    access_key="minioadmin",
    secret_key="minioadmin",
    secure=False
)


# --- Инициализация Telegram-бота ---
bot = TelegramBot(token=TOKEN)

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

# --- Точка входа ---
if __name__ == "__main__":
    print("Инициализация проекта Massyanya-2.0...")
    print("Запуск веб-сервера...")
    app.run(host="0.0.0.0", port=5000)
