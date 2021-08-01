# sets the base image
FROM golang:1.16 AS base

ENV GOSUMDB="off"
ENV HOME=/usr/home
ENV GOROOT="/usr/local/go"
ENV GOPATH=$HOME/go
ENV PATH="${PATH}:${GOROOT}/bin:${GOPATH}/bin:usr/local/bin"


RUN mkdir -p $HOME\
 && mkdir -p $GOPATH

WORKDIR /src

# ==================
# Dev Container
FROM base as dev

# RUN echo "$PWD"
# Cache dependencies
COPY go.mod go.sum /src/
RUN go mod download -x

RUN chmod -R a+rwX $HOME

# ==================
# Build Container

From dev AS build

ENV CGO_ENABLED=0
ENV GOOS=linux

# Build executables
COPY . /src/
RUN go install -v ./...

RUN chmod -R a+rwX $HOME