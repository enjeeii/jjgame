import grpc
import hello_pb2
import hello_pb2_grpc

def dotest():
    channel = grpc.insecure_channel('192.168.70.128:20000')
    stub = hello_pb2_grpc.TesterStub(channel)
    req = hello_pb2.HelloReq()

    # Simple RPC
    ## sync
    req.msg = 'simple rpc.'
    msg = stub.SayHello(req)
    print(msg.msg)
    ## async
    rep = stub.SayHello.future(req)
    msg = rep.result()
    print(msg.msg)

    # Server-side RPC
    req.msg = 'server-side rpc.'
    for msg in stub.SayHello2(req):
        print(msg.msg)

    def genIterReq(num):
        while num > 0:
            yield req
            num -= 1

    # Client-side RPC
    req.msg = 'client-side rpc.'
    ## sync
    msg = stub.SayHello3(genIterReq(3))
    print(msg.msg)
    ## async
    rep = stub.SayHello3.future(genIterReq(3))
    msg = rep.result()
    print(msg.msg)

    # Bidirectional RPC
    req.msg = 'bidirectional rpc.'
    for msg in stub.SayHello4(genIterReq(3)):
        print(msg.msg)
