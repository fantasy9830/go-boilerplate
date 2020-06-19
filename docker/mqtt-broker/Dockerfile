FROM eclipse-mosquitto:latest

COPY mosquitto.conf /mosquitto/config/mosquitto.conf

COPY ca.crt /mosquitto/certs/ca.crt
COPY broker.key /mosquitto/certs/broker.key
COPY broker.crt /mosquitto/certs/broker.crt

EXPOSE 8883
