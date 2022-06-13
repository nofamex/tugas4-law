FROM golang:1.18 AS Production
WORKDIR /app
COPY go.mod .env ./
RUN go mod tidy
COPY . .
RUN go build -o read
EXPOSE 5002
CMD /app/read