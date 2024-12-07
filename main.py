import sys
from flask import Flask, request, jsonify, render_template
from dotenv import load_dotenv
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


# --- Главная страница ---
@app.route("/")
def index():
    return render_template("index.html")


# --- График ---
@app.route("/analytics")
def analytics():
    return render_template("analytics.html")


# --- Настройки торговли ---
@app.route("/trade-settings")
def trade_settings():
    return render_template("trade_settings.html")


# --- База данных ---
@app.route("/data")
def data():
    return render_template("data.html")


# --- IP и Ключи ---
@app.route("/keys")
def keys():
    return render_template("keys.html")


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
        # Симулируем данные для теста
        data = {
            "2024-12-01": 1.200,
            "2024-12-02": 1.205,
            "2024-12-03": 1.210
        }
        return jsonify({"message": "Предсказания успешно выполнены!", "data": data})
    except Exception as e:
        return jsonify({"error": str(e)}), 500


# --- Точка входа ---
if __name__ == "__main__":
    print("Инициализация проекта Massyanya-2.0...")
    print("Запуск веб-сервера...")
    app.run(host="0.0.0.0", port=5000)
