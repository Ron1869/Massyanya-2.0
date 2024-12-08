from flask import Blueprint, jsonify

telegram_api = Blueprint("telegram_api", __name__)

@telegram_api.route("/start", methods=["POST"])
def start_telegram_bot():
    # Логика для запуска Telegram-бота
    return jsonify({"message": "Telegram-бот успешно запущен!"}), 200

@telegram_api.route("/stop", methods=["POST"])
def stop_telegram_bot():
    # Логика для остановки Telegram-бота
    return jsonify({"message": "Telegram-бот успешно остановлен!"}), 200
