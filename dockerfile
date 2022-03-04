FROM golang:latest
RUN mkdir /app
COPY main /app/main
WORKDIR /app
# RUN go build -o main . 
CMD ["/app/main"]