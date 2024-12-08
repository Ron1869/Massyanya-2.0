# Подключение к базе данных
from sqlalchemy import create_engine

engine = create_engine("sqlite:///../database/db.sqlite3")

def get_db_connection():
    return engine.connect()
