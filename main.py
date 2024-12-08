import sys
from flask import Flask, request, jsonify, render_template
from dotenv import load_dotenv
from minio import Minio
from predict_stock_crypto.messaging.telegram_bot import TelegramBot
import os
import threading
from backend.services.trading_service import start_trading, stop_trading
from backend.services.forecast_service import make_forecast
from backend.api.trading_api import trading_api
from backend.api.analytics_api import analytics_api
from backend.api.data_api import data_api
from backend.api.telegram_api import telegram_api

import logging
logging.basicConfig(level=logging.DEBUG)


# --- Инициализация Flask-приложения ---
app = Flask(__name__)

# --- Регистрация API-маршрутов ---
app.register_blueprint(analytics_api, url_prefix="/api/analytics")
app.register_blueprint(data_api, url_prefix="/api/data")
app.register_blueprint(telegram_api, url_prefix="/api/telegram")
app.register_blueprint(trading_api, url_prefix="/api/trading")

# --- Загрузка переменных из .env ---
load_dotenv(dotenv_path="config/.env")
TOKEN = os.getenv("TOKEN")
CHAT_ID = os.getenv("CHATID")

# --- Настройка MinIO ---
minio_client = Minio(
    "127.0.0.1:9000",
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

# --- Запуск торговли ---
@app.route("/start-trading", methods=["POST"])
def start_trading_route():
    try:
        start_trading()
        return jsonify({"status": "Торговля успешно запущена!"}), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# --- Остановка торговли ---
@app.route("/stop-trading", methods=["POST"])
def stop_trading_route():
    try:
        stop_trading()
        return jsonify({"status": "Торговля успешно остановлена!"}), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# --- Запуск прогнозов ---
@app.route("/forecast", methods=["POST"])
def forecast_route():
    try:
        result = make_forecast()
        return jsonify({"status": "Прогноз успешно выполнен!", "result": result}), 200
    except Exception as e:
        return jsonify({"error": str(e)}), 500

# --- Точка входа ---
if __name__ == "__main__":
    print("Инициализация проекта Massyanya-2.0...")
    print("Запуск веб-сервера...")
    app.run(host="0.0.0.0", port=5000)
