FROM golang:alpine

COPY ./auth_server /auth_server

# COPY ./publicKey.pem /server/publicKey.pem

# COPY ./privateKey.pem /server/privateKey.pem

WORKDIR /auth_server

# ENV GIN_MODE=${GIN_MODE}

RUN go mod download

RUN go mod tidy

RUN go build -o server

CMD ["/server/server"]
