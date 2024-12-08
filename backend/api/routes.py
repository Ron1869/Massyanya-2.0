from flask import Blueprint, jsonify

api = Blueprint('api', __name__)

@api.route('/start-trading', methods=['POST'])
def start_trading():
    # TODO: добавить логику старта торговли
    return jsonify({"status": "Trading started"})

@api.route('/stop-trading', methods=['POST'])
def stop_trading():
    # TODO: добавить логику остановки торговли
    return jsonify({"status": "Trading stopped"})

@api.route('/forecast', methods=['POST'])
def forecast():
    # TODO: добавить логику прогноза
    return jsonify({"status": "Forecast initiated"})
