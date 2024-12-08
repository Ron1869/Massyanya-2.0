from flask import Blueprint, jsonify

trading_api = Blueprint("trading_api", __name__)

@trading_api.route("/start-trading", methods=["POST"])
def start_trading():
    # Здесь вызывается логика для старта торговли
    return jsonify({"message": "Торговля успешно запущена!"})

@trading_api.route("/stop-trading", methods=["POST"])
def stop_trading():
    # Здесь вызывается логика для остановки торговли
    return jsonify({"message": "Торговля успешно остановлена!"})
