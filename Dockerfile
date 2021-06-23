FROM golang:1.16-alpine as builder
WORKDIR /go/src
COPY . .

ARG _TWITTER_KEY
ENV TWITTER_KEY ${_TWITTER_KEY}
ARG _TWITTER_SECRET
ENV TWITTER_SECRET ${_TWITTER_SECRET}

RUN CGO_ENABLED=0 GOOS=linux go build -v -o popular

FROM gcr.io/cloud-builders/gcloud:latest
WORKDIR /go/src
COPY . .

COPY --from=builder /go/src/popular /popular
ENTRYPOINT ["/popular"]
