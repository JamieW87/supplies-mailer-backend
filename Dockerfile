FROM golang:1.20

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/server/

EXPOSE 8080

CMD [ "./main" ]

# docker run --env-file=env.docker [image name]