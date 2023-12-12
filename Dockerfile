FROM golang:1.21.0
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go build -o server
ENTRYPOINT [ "./server" ]
