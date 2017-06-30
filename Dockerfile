FROM prom/busybox

MAINTAINER Aurelien Chaudagne

COPY cmd/sys-status-srv/sys-status-srv /home/

WORKDIR /home

EXPOSE 8080

CMD ["/home/sys-status-srv"]