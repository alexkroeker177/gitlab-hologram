FROM ubuntu

WORKDIR /usr/src/app
ENV goversion=1.20.4

RUN apt-get update
RUN apt-get install curl -y

RUN curl -OL https://golang.org/dl/go${goversion}.linux-amd64.tar.gz
RUN tar -C /usr/local -xvf go${goversion}.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -o bitnode ./cmd/docker/

ENV BITNODE_LOCAL_PORT=9060
ENV BITNODE_LOCAL_ADDRESS=0.0.0.0:${BITNODE_LOCAL_PORT}

EXPOSE ${BITNODE_LOCAL_PORT}

CMD ["/usr/src/app/bitnode"]
