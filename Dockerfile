FROM golang:1.20

WORKDIR /home

COPY . .

RUN go mod tidy

CMD go run github.com/Mohsenpoureiny/getprojects-bot