FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache go 

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

COPY go.mod go.sum ./
RUN go mod download

COPY .env /app
COPY . .

EXPOSE 8080

CMD ["go", "run", "./src/cmd/main.go"]