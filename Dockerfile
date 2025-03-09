FROM golang:1.24.1-alpine AS builder
WORKDIR /build
COPY go.mod ./
COPY  . .
RUN ls
RUN go build -o /app/b612lpp.hellow


FROM golang:1.24.1-alpine
COPY --from=builder /app/b612lpp.hellow /app/b612lpp.hellow
CMD [ "/app/b612lpp.hellow" ]
