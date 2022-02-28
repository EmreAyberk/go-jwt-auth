FROM golang:1.16-alpine as stage1

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./cmd/api/main.go


FROM golang:1.16-alpine as stage2


COPY --from=stage1 /app/main ./main

EXPOSE 3000

CMD ["./main"]