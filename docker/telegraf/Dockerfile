FROM telegraf:latest

LABEL maintainer="Ricky Lu <fantasy9830@gmail.com>"

COPY telegraf.conf /etc/telegraf/telegraf.conf

COPY ca.pem /etc/telegraf/ca.pem
COPY cert.pem /etc/telegraf/cert.pem
COPY key.pem /etc/telegraf/key.pem
