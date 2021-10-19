FROM golang:latest as base

FROM base as dev

RUN apt-get update
RUN apt-get -y install ffmpeg

RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin


WORKDIR /opt/app/borat_bot

RUN chmod a+x /opt/app/borat_bot

CMD ["air"]