FROM golang:1.16-alpine
ADD . /gitrunner
WORKDIR /gitrunner
RUN go build ./main.go
CMD ["./main"]