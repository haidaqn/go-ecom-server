FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o go-ecommerce-backend-api ./cmd/server

FROM scratch
COPY ./configs /configs
COPY --from=builder /build/go-ecommerce-backend-api /

ENTRYPOINT [ "/go-ecommerce-backend-api", "configs/dev.yaml" ]