FROM centos:centos7
WORKDIR /root
EXPOSE 9090

COPY  nctr  /usr/bin/
COPY nctr.yaml /root
ENTRYPOINT ["nctr","start","--config","/root/nctr.yaml"]