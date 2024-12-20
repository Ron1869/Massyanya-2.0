# config/settings.py
import os
from pathlib import Path
import logging
from logging.handlers import RotatingFileHandler

from dotenv import load_dotenv

# Выбираем абсолютный путь - берём за основу этот файл и поднимаемся на две директории выше
BASE_DIR = Path(__file__).resolve().parent.parent
IMG_DIR = os.path.join(BASE_DIR, "predict_stock_crypto", "img")#поменял с IMG_DIR = os.path.join(BASE_DIR, "img")
STOCKS_CRYPTO_DIR = os.path.join(BASE_DIR, "stocks_list")
DATABASE_DIR = os.path.join(BASE_DIR, "database")
DATABASE = os.path.join(DATABASE_DIR, "users.db")
CONFIG_DIR = os.path.join(BASE_DIR, "config")
LOGGING_DIR = os.path.join(CONFIG_DIR, "logging")
LOGGING = os.path.join(LOGGING_DIR, "logfile.log")

# Здесь задана глобальная конфигурация для всех логгеров
# Если требуется более детальный сбор логов, измените level=logging. на INFO
logging.basicConfig(
    handlers=[RotatingFileHandler(LOGGING, maxBytes=10**5, backupCount=5)],
    level=logging.ERROR,
    format="[%(asctime)s] %(levelname)s [%(name)s.%(funcName)s:%(lineno)d] %(message)s",
    datefmt="%Y-%m-%d | T%H:%M:%S",
)

# Количество избранных валют у каждого пользователя
WATCHLIST_POSITION = 3

COUNTRY = "United States"

# специальное отображение графиков для pyplot fivethirtyeight
PYPLOT_SET_GRAPH = "fivethirtyeight"

# Загрузим переменные окружения
load_dotenv('C:/Project_Massyanya_New/config/.env')

# При работе с переменными окружения, будем использовать getenv
# В отличии от environ, getenv не вызывает исключения,
# а возвращает None
try:
    TOKEN = os.getenv("TOKEN")
except Exception as error:
    logging.exception(
        f'({error}) - отсутствует переменная окружения "TOKEN"',
        stack_info=True,
    )
try:
    CHAT_ID = int(os.getenv("CHATID"))
except Exception as error:
    logging.exception(
        f'({error}) - отсутствует переменная окружения "CHAT_ID"',
        stack_info=True,
    )
