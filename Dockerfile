FROM golang:1.23.3-alpine

RUN apk add --no-cache ffmpeg

RUN apk add --no-cache python3 py3-pip
RUN pip3 install openai-whisper

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

EXPOSE 8080

CMD ["go", "run", "bot/main.go"]
