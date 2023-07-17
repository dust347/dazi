FROM ccr.ccs.tencentyun.com/library/golang:1.19

ENV GO118MODULE=on \
CGO_ENABLED=0\
GOOS=linux \
GOARCH=amd64 \
GOPROXY="https://goproxy.cn,direct"

RUN mkdir -p /app/dazi
RUN git clone https://github.com/dust347/dazi.git /app/dazi
#COPY ./* /app/dazi

WORKDIR /app/dazi

#RUN go mod tidy
RUN go build -o dazi

CMD ["./dazi"]
