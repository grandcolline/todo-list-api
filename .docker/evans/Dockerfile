FROM debian:stretch

ARG EVANS_VERSION=0.8.5

RUN apt-get update && \
	apt-get install -y wget && \
	rm -rf /var/lib/apt/lists/*

RUN apt-get install -y bash

# Install evans
RUN wget https://github.com/ktr0731/evans/releases/download/${EVANS_VERSION}/evans_linux_amd64.tar.gz && \
	tar xvf evans_linux_amd64.tar.gz && \
	mv evans /usr/local/bin/ && \
	rm evans_linux_amd64.tar.gz

WORKDIR /proto

# Install grpc middleware
RUN mkdir -p github.com/mwitkow/go-proto-validators && \
	wget https://raw.githubusercontent.com/mwitkow/go-proto-validators/master/validator.proto \
		-O github.com/mwitkow/go-proto-validators/validator.proto

COPY cli /usr/local/bin/cli

ENTRYPOINT ["bash"]
