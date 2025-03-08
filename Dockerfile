FROM golang:1.24.1
WORKDIR /app
COPY * ./
RUN go build -o /b612lpp.hellow
CMD [ "/b612lpp.hellow" ]
