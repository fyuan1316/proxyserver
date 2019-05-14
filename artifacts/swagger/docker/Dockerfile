FROM golang

WORKDIR /
RUN apt update && \
    apt install unzip && \
    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.7.0/protoc-3.7.0-linux-x86_64.zip && \
    unzip protoc-3.7.0-linux-x86_64.zip -d protoc && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway && \
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

ENV PATH=$PATH:/protoc/bin

WORKDIR /gen

RUN git clone https://github.com/tensorflow/serving.git && \
    git clone https://github.com/tensorflow/tensorflow.git && \
    ln -s /gen/tensorflow/tensorflow /gen/serving/tensorflow
ADD tf-serving.proto /gen/serving/
ENTRYPOINT [ "protoc" ]
