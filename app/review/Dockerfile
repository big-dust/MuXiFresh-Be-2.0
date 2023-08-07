FROM ubuntu:22.04

WORKDIR /MuxiFresh-Be-2.0/review

COPY cmd/api/review .
COPY cmd/api/etc etc

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]