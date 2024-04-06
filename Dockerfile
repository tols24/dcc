FROM golang:1.22.1

WORKDIR /bin

COPY . ./

RUN go get
RUN go build -o main .

ENTRYPOINT [ "/bin/main" ]