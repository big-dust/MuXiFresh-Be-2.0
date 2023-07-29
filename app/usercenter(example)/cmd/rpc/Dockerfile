FROM ubuntu:20.04
RUN [ "mkdir", "/usr/local/usercenter-rpc" ]
RUN [ "mkdir", "/usr/local/usercenter-rpc/etc" ]
WORKDIR /usr/local/usercenter-rpc
COPY ./usercenter-rpc /usr/local/usercenter-rpc
COPY ./etc /usr/local/usercenter-api/etc
CMD ["./usercenter-rpc", "-f", "etc/usercenter.yaml"]