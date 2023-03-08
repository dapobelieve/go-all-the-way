FROM --platform=linux/amd64 golang:1.19-alpine as builder

WORKDIR /recipe-api

COPY go.mod /recipe-api/go.mod
COPY go.sum /recipe-api/go.sum
COPY *.go /recipe-api/

RUN go mod download
RUN go build -o /bin/entrypoint .

FROM --platform=linux/amd64 alpine:latest 

COPY --from=builder /bin/entrypoint /bin/entrypoint

EXPOSE 8080
CMD ["/bin/entrypoint"]
