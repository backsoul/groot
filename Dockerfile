FROM golang:1.18-alpine as groot
WORKDIR /groot

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o groot ./cmd/groot/main.go

EXPOSE 8000
CMD ["./groot"]