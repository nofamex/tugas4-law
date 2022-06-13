FROM golang:1.18 AS Production
WORKDIR /app
COPY go.mod .env ./
RUN go mod tidy
COPY . .
RUN go build -o simakara
EXPOSE 5000
CMD /app/simakara