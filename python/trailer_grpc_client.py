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


def run():
    channel = grpc.insecure_channel("localhost:50051")
    stub = trailer_pb2_grpc.trailerStub(channel)
    response = stub.Init(trailer_pb2.Config())
    print("trailer_pb2 Init received: ", response.code)

    response = stub.Start(trailer_pb2.Request())
    print("trailer_pb2 Start received: ", response.code)

    response = stub.Status(trailer_pb2.Request())
    print("trailer_pb2 Status received: ", response)

    response = stub.Service(trailer_pb2.ServiceRequest())
    print("trailer_pb2 Service received: ", response.code)

    response = stub.Query(trailer_pb2.DataRowsRequest())
    print("trailer_pb2 Query received: ", response.row)

    response = stub.Schema(trailer_pb2.SchemaRequest())
    print("trailer_pb2 Schema received: ", response)

    response = stub.Stop(trailer_pb2.Request())
    print("trailer_pb2 Stop received: ", response)


if __name__ == "__main__":
    run()
