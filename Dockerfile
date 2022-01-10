FROM golang:1.17.5

COPY . .

RUN go get

CMD go run getprojects