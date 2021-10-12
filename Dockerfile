FROM golang

RUN mkdir /aastra-monitor

ADD . /aastra-monitor

WORKDIR /aastra-monitor

RUN go build -o main .

CMD ["/aastra-monitor/main"]