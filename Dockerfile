FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /web-app
EXPOSE 8080
CMD [ "/web-app" ]