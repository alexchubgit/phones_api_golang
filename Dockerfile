FROM golang:1.16-alpine

ENV APP_HOME /app
RUN mkdir -p "$APP_HOME"

WORKDIR "$APP_HOME"

COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

RUN go build -o server

EXPOSE 8000

CMD ["/app/server"]
