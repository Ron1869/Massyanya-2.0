from flask import Blueprint, jsonify, request

data_api = Blueprint("data_api", __name__)

saved_coins = ["BTC", "ETH", "XRP"]  # Пример данных

# Получение сохранённых монет
@data_api.route("/get-saved-coins", methods=["GET"])
def get_saved_coins():
    return jsonify({"coins": saved_coins})

# Добавление новой монеты
@data_api.route("/add-coin", methods=["POST"])
def add_coin():
    coin = request.json.get("coin", "").upper()
    if coin and coin not in saved_coins:
        saved_coins.append(coin)
        return jsonify({"message": f"Монета {coin} успешно добавлена!"}), 200
    return jsonify({"error": "Монета уже существует или некорректное название!"}), 400

# Удаление монеты
@data_api.route("/remove-coin", methods=["POST"])
def remove_coin():
    coin = request.json.get("coin", "").upper()
    if coin in saved_coins:
        saved_coins.remove(coin)
        return jsonify({"message": f"Монета {coin} успешно удалена!"}), 200
    return jsonify({"error": "Монета не найдена!"}), 404
