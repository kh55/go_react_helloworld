# 開発環境
FROM golang:1.21

WORKDIR /api
COPY . .
#ENV GO111MODULE=auto

#RUN go get github.com/pilu/fresh
#CMD ["fresh"]

RUN go build main.go