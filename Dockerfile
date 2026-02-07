FROM golang:1.25-trixie

ENV HOME=/home/local/
WORKDIR /home/local/app/

RUN apt-get update && apt-get upgrade -y

COPY */go.mod */go.sum ./
