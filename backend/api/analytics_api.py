from flask import Blueprint, jsonify

analytics_api = Blueprint("analytics_api", __name__)

# Пример: Получение данных для графика
@analytics_api.route("/get-chart-data", methods=["GET"])
def get_chart_data():
    # Здесь должна быть логика получения данных с биржи или базы данных
    data = {
        "labels": ["2024-12-01", "2024-12-02", "2024-12-03"],
        "prices": [40000, 40500, 41000],
        "volumes": [120, 150, 130]
    }
    return jsonify(data)
