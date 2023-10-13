FROM golang:latest

WORKDIR $GOPATH/src/go-boilerplate

# Install package
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

RUN git config --global --add safe.directory $GOPATH/src/go-boilerplate

# Install golangci-lint
# binary will be $GOPATH/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.54.2

# Install Swag
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Install Staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

# Install Govulncheck
RUN go install golang.org/x/vuln/cmd/govulncheck@latest

# Install Wire
RUN go install github.com/google/wire/cmd/wire@latest

# Install Protocol Buffer Compiler
RUN curl -L https://github.com/protocolbuffers/protobuf/releases/download/v24.3/protoc-24.3-linux-x86_64.zip -o /tmp/protoc.zip && \
  unzip -o /tmp/protoc.zip -d /usr/local bin/protoc && \
  unzip -o /tmp/protoc.zip -d /usr/local include/*

# Install the Protocol Buffer Compiler plugins for Go
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

# Install the mockgen tool
RUN go install go.uber.org/mock/mockgen@latest

# Clean up
RUN apt-get clean \
  && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*
