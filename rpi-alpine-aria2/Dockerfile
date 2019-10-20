FROM arm32v6/alpine:3.8
MAINTAINER Yangxuan <yangxuan8282@gmail.com>

ENV ARIA2_VERSION=1.35.0
ENV ARIA2_URL=https://github.com/q3aql/aria2-static-builds/releases/download/v${ARIA2_VERSION}/aria2-${ARIA2_VERSION}-linux-gnu-arm-rbpi-build1.tar.bz2

RUN apk add bash curl tar
RUN curl -sSL $ARIA2_URL | tar xj --strip 1 -C /usr/local/bin/ aria2-${ARIA2_VERSION}-linux-gnu-arm-rbpi-build1/aria2c 
RUN adduser -D aria2 
RUN mkdir -p /etc/aria2 
RUN touch /etc/aria2/session.lock 

COPY aria2.conf /etc/aria2/aria2.conf 
COPY ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
VOLUME /home/aria2 /etc/aria2

EXPOSE 6800

CMD set -xe && \
 chown -R aria2:aria2 /home/aria2 && \
 aria2c --conf-path=/etc/aria2/aria2.conf
