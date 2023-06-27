from concurrent import futures

import grpc

import paddle_pb2
import paddle_pb2_grpc


class PaddleService(paddle_pb2_grpc.TrainServiceServicer):

    def TrainAndPredict(self, request, context):
        # 把之前的训练代码块放到这里

        # img_byte_arr = BytesIO()
        # im = Image.fromarray(np.uint8(img[0] * 255))
        # im.save(img_byte_arr, format='PNG')
        # img_byte_data = img_byte_arr.getvalue()

        print("Paddle gRPC Server received image data")

        return paddle_pb2.TrainResult(message="Prediction Successful",
                                      true_label=int(1),
                                      predicted_label=int(2),
                                      image_data=bytes("img_byte_data", encoding='utf-8'))


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    paddle_pb2_grpc.add_TrainServiceServicer_to_server(PaddleService(), server)
    server.add_insecure_port('[::]:50051')
    server.start()
    print("Paddle gRPC Server started")

    server.wait_for_termination()


if __name__ == '__main__':
    serve()
