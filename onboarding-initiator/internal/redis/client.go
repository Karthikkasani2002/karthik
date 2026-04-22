package redis

import (
    "context"
    "crypto/tls"
    "crypto/x509"
    "fmt"
    "os"

    "github.com/redis/go-redis/v9"

    "onboarding/internal/config"
)

type Client struct {
    rdb *redis.Client
}

func New(cfg config.Config) *Client {

    addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

    // ✅ Load CA certificate
    caCert, err := os.ReadFile(cfg.RedisCA)
    if err != nil {
        panic(err)
    }

    // ✅ Create CA pool
    caCertPool := x509.NewCertPool()
    if ok := caCertPool.AppendCertsFromPEM(caCert); !ok {
        panic("failed to append CA cert")
    }

    r := redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: cfg.RedisPass,

        TLSConfig: &tls.Config{
            MinVersion: tls.VersionTLS12,
            RootCAs:    caCertPool, // ✅ VERY IMPORTANT
        },
    })

    ctx := context.Background()

    _, err = r.Ping(ctx).Result()
    if err != nil {
        panic(err)
    }

    return &Client{rdb: r}
}

func (c *Client) Close() {
    c.rdb.Close()
}
