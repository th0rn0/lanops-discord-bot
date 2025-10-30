FROM golang:1.24-alpine
LABEL maintainer="Thornton Phillis (dev@th0rn0.co.uk)"

WORKDIR /app

ENV CGO_ENABLED=1
RUN apk add --no-cache gcc musl-dev

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN go build -ldflags='-s -w -extldflags "-static"' -o /discord-bot

EXPOSE 9888

CMD [ "/discord-bot" ]