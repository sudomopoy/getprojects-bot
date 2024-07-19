FROM golang:1.20

WORKDIR /home

COPY . .


RUN go mod tidy
RUN go build github.com/Mohsenpoureiny/getprojects-bot

CMD ./getprojects-bot