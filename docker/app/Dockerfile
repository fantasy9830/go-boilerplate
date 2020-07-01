FROM golang:latest

WORKDIR $GOPATH/src/app

# Instal package
RUN set -x \
  && apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y --no-install-recommends \
  unzip

# Set Timezone
ARG TZ=UTC
ENV TZ ${TZ}
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# User Aliases
COPY ./aliases.sh /root/aliases.sh
RUN sed -i 's/\r//g' /root/aliases.sh && \
  echo "" >> ~/.bashrc && \
  echo "# Load Custom Aliases" >> ~/.bashrc && \
  echo "source /root/aliases.sh" >> ~/.bashrc && \
  echo "" >> ~/.bashrc

# Install Golint
RUN go get -u golang.org/x/lint/golint

# Install Wire
RUN go get -u github.com/google/wire/cmd/wire

# Install protoc-gen-go
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go

# Install Protoc
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protoc-3.12.3-linux-x86_64.zip -o /tmp/protoc.zip && \
  unzip -o /tmp/protoc.zip -d /usr/local bin/protoc && \
  unzip -o /tmp/protoc.zip -d /usr/local include/*

# Clean up
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
