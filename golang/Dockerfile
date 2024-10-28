FROM golang:1.23-alpine

RUN apk add --no-cache bash ffmpeg

WORKDIR /app

COPY . .

CMD ["bash"]

#COPY go.mod go.sum ./
#RUN go mod download