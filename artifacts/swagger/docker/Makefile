build-image:
	docker build -t gen-tf-model-swagger .

build:
	docker run -v ${PWD}/out:/gen/out gen-tf-model-swagger --proto_path=/gen/serving --proto_path=/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:/gen/out  tf-serving.proto