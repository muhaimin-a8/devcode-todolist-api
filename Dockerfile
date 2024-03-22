FROM golang:1.21.2 as builder

ARG CGO_ENABLED=0
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN make build

FROM scratch
COPY --from=builder /app/.env .env
COPY --from=builder /app/bin/devcode_todolist_api bin/devcode_todolist_api
ENV APP_PORT=3030

EXPOSE 3030
ENTRYPOINT ["bin/devcode_todolist_api"]