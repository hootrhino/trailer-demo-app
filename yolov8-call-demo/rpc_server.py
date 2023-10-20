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

import grpc
import trailer_pb2
import trailer_pb2_grpc
from concurrent import futures
from ultralytics import YOLO
import numpy as np
import cv2
import base64


class trailerServicer(trailer_pb2_grpc.trailerServicer):
    def Init(self, request, context):
        self.model = YOLO("yolov8n.pt")
        return trailer_pb2.Response(code=0, message="Init ok")

    def Start(self, request, context):
        return trailer_pb2.Response(code=0, message="Start ok")

    def Status(self, request, context):
        return trailer_pb2.StatusResponse(status=0, message="Status Ok")

    def Service(self, request, context):
        print("Service:", request.cmd, request.args)
        # request.args: 是客户端编码过的二进制流 ,在这里可以还原出来
        image_base64 = request.args.decode()
        decoded_image = np.frombuffer(base64.b64decode(image_base64), np.uint8)
        decoded_image = cv2.imdecode(decoded_image, cv2.IMREAD_COLOR)
        results = self.model.predict(source=decoded_image, save=True, save_txt=True)
        return trailer_pb2.ServiceResponse(code=0, data=b"Service ok")

    def Query(self, request, context):
        return trailer_pb2.DataRowsResponse(row=[])

    def Schema(self, request, context):
        return trailer_pb2.SchemaResponse(message="Status Ok", columns=[{}])

    def Stop(self, request, context):
        return trailer_pb2.Response(code=0, message="Stop ok")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    trailer_pb2_grpc.add_trailerServicer_to_server(trailerServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
