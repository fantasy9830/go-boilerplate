FROM eclipse-mosquitto:latest

LABEL maintainer="Ricky Lu <fantasy9830@gmail.com>"

COPY mosquitto.conf /mosquitto/config/mosquitto.conf

COPY ca.crt /mosquitto/certs/ca.crt
COPY broker.key /mosquitto/certs/broker.key
COPY broker.crt /mosquitto/certs/broker.crt

EXPOSE 8883
