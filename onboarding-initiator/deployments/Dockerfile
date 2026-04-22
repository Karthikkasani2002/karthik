FROM golang:1.22 as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=1 go build \
-o onboarding-service \
./cmd/service


FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/onboarding-service .

USER 10001

EXPOSE 8080

ENTRYPOINT ["/app/onboarding-service"]
