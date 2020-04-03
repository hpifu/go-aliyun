FROM centos:centos7

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN echo "Asia/Shanghai" >> /etc/timezone

COPY docker/go-aliyun /var/docker/go-aliyun
RUN mkdir -p /var/docker/go-aliyun/log

EXPOSE 7060

WORKDIR /var/docker/go-aliyun
CMD [ "bin/aliyun", "-c", "configs/aliyun.json" ]
