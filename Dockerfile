FROM golang:1.16-alpine as groot
WORKDIR /groot

COPY go.mod ./ 
RUN go mod download
COPY . .

RUN go build -o groot ./cmd/groot/main.go

EXPOSE 3000
CMD ["./groot"]