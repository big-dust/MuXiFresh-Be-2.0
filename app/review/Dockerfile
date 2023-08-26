FROM centos:8.1.1911

WORKDIR /MuxiFresh-Be-2.0/review

COPY cmd/api/review .
COPY cmd/api/etc etc

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]