# backend/api/trading_api.py
from flask import Blueprint, jsonify
from backend.services.trading_service import start_trading, stop_trading

trading_api = Blueprint('trading_api', __name__)

@trading_api.route('/start-trading', methods=['POST'])
def start():
    result = start_trading()
    return jsonify(result)

@trading_api.route('/stop-trading', methods=['POST'])
def stop():
    result = stop_trading()
    return jsonify(result)
