FROM golang:1.10

RUN apt-get update && apt-get upgrade -y

RUN apt-get install docker -y

RUN apt-get install docker-compose -y

RUN docker-compose --version

WORKDIR /go/src/app
COPY . .

# RUN test -n "$TRAVIS_TAG" || bash tests/prepare.sh
