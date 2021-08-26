FROM golang:1.16-alpine

WORKDIR /go/src/app

COPY . .

RUN go mod download && go mod verify

RUN go build -o /app

ARG ADDR

ENV ADDR=${ADDR}

EXPOSE ${ADDR}

CMD [ "/app" ]
