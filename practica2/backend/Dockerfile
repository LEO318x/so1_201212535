FROM golang:1.20-alpine3.16
WORKDIR /app/src
ENV GOPATH=/app
COPY go.mod /app/src
COPY go.sum /app/src
RUN go mod download
COPY . /app/src
EXPOSE 8000
CMD ["go", "run", "main.go"]