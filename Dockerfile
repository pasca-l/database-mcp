# syntax=docker/dockerfile:1

FROM golang:1.26-trixie

ENV HOME=/home/local/
WORKDIR /home/local/app/

RUN apt-get update && apt-get upgrade -y

COPY */go.mod */go.sum ./
