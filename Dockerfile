FROM golang:1.21
WORKDIR /src
COPY . .
RUN go mod download
RUN go build -o /bin/hello ./main.go

CMD ["/bin/hello"]
