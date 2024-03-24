FROM golang:1.22-alpine

WORKDIR /app

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN go build -ldflags='-s -w -extldflags "-static"' -o /discord-bot

EXPOSE 8080

CMD [ "/discord-bot" ]