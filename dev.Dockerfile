FROM golang:1.22.2-bookworm

RUN apt update \
    && apt install -y --no-install-recommends sudo locales \
    && sed -i 's/^# *\(ja_JP.UTF-8\)/\1/' /etc/locale.gen \
    && locale-gen \
    && update-locale LANG=ja_JP.UTF-8 \
    && rm -rf /var/lib/apt/lists/*

ENV SHELL=/usr/bin/bash
ENV LANGUAGE=ja_JP.UTF-8
ENV LC_ALL=ja_JP.UTF-8
ENV LANG=ja_JP.UTF-8
ENV TZ=Asia/Tokyo

WORKDIR /usr/src/app

COPY go.mod ./
RUN go mod download

COPY . .
