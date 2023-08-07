FROM ubuntu:22.04

WORKDIR /MuxiFresh-Be-2.0/test

COPY rpc/test test-rpc
COPY rpc/etc/test.yaml etc/test-rpc.yaml

COPY api/test test-api
COPY api/etc/test-api.yaml etc/test-api.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]