FROM ubuntu:22.04

WORKDIR /MuxiFresh-Be-2.0/form

COPY rpc/form form-rpc
COPY rpc/etc/form.yaml etc/form-rpc.yaml

COPY api/form form-api
COPY api/etc/form-api.yaml etc/form-api.yaml

COPY start.sh ./

RUN chmod +x ./start.sh

CMD ["./start.sh"]