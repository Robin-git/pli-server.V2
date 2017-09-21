FROM golang:1.9
EXPOSE 8000
WORKDIR /go/src/gloo-server
COPY . .
RUN go-wrapper download
RUN go-wrapper install
CMD ["go-wrapper", "run"]
