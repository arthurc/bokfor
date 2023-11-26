FROM golang:1.21-alpine as builder

WORKDIR /src/bokfor
ADD go.mod .
RUN go mod download

ADD . .
RUN go build -o bin/bokfor ./cmd/bokfor

FROM alpine

WORKDIR /bin/

COPY --from=builder /src/bokfor/bin/bokfor .

CMD /bin/bokfor
