# -*- coding: utf-8 -*-

import os
import sys
import grpc

def main():
    import hello_pb2
    import hello_pb2_grpc
    print('hello jjgame.')
    channel = grpc.insecure_channel('192.168.70.128:20000')
    stub = hello_pb2_grpc.TesterStub(channel)
    req = hello_pb2.HelloReq()
    req.msg = 'jjgame client.'
    msg = stub.SayHello(req)
    print(msg.msg)

if __name__ == '__main__':
    cur_dir = os.getcwd()
    sys.path.append(os.path.join(cur_dir, '/scripts/protoc'))
    main()