from flask import Flask
from minio import Minio
from PIL import Image
import matplotlib.pyplot as plt
import lxml
import numpy as np
import pandas as pd
import investpy

from minio import Minio

# Пример создания клиента MinIO
client = Minio(
    "play.min.io",
    access_key="minioadmin",
    secret_key="minioadmin",
    secure=False,
)
print("Minio успешно импортирован!")


print("Все импорты успешно выполнены!")
