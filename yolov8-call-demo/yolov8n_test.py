# Copyright (C) 2023 wwhai
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU Affero General Public License for more details.
#
# You should have received a copy of the GNU Affero General Public License
# along with this program.  If not, see <http://www.gnu.org/licenses/>.

from ultralytics import YOLO
import numpy as np

import cv2
from PIL import Image


def Train():
    model = YOLO("yolov8n.pt")
    model.info()
    model.train(data="coco8.yml", epochs=100, imgsz=640)


import base64


def RunModel():
    model = YOLO("yolov8n.pt")
    # 1. 读取图像文件
    image = cv2.imread("./1.png")

    # 2. 将图像转换为Base64字符串
    # image_base64 就是来自RPC客户端的字节
    image_base64 = base64.b64encode(cv2.imencode(".png", image)[1]).decode()

    # 3. 将Base64字符串解码为NumPy数组
    decoded_image = np.frombuffer(base64.b64decode(image_base64), np.uint8)
    decoded_image = cv2.imdecode(decoded_image, cv2.IMREAD_COLOR)
    results = model.predict(source=decoded_image, save=True, save_txt=True)
    for result in results:
        boxes = result.boxes  # Boxes object for bbox outputs
        masks = result.masks  # Masks object for segmentation masks outputs
        keypoints = result.keypoints  # Keypoints object for pose outputs
        probs = result.probs  # Probs object for classification outputs
        print(boxes)
        print(masks)
        print(keypoints)
        print(probs)


RunModel()
