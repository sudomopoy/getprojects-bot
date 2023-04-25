FROM golang:1.17.5

WORKDIR /home

COPY . .

RUN go get

CMD go run github.com/Mohsenpoureiny/getprojects-bot