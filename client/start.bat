python -m grpc_tools.protoc -I../proto --python_out=./scripts/protoc --pyi_out=./scripts/protoc --grpc_python_out=./scripts/protoc hello.proto

python main.py
@pause