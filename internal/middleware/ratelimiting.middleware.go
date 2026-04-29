package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type Client struct {
	limiter     *rate.Limiter
	lastRequest time.Time
}

var (
	mu      sync.Mutex
	clients = make(map[string]*Client)
)

func getClientIP(ctx *gin.Context) string {
	ip := ctx.ClientIP()
	if ip == "::1" {
		ip = "127.0.0.1"
	} else if ip == "" {
		ip = ctx.Request.RemoteAddr
	}

	return ip
}

func getRateLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	client, exists := clients[ip]
	if !exists {
		limiter := rate.NewLimiter(5, 15) // 5 request/sec, brust 10
		newClient := &Client{limiter, time.Now()}
		clients[ip] = newClient
		return limiter
	}

	client.lastRequest = time.Now()
	return client.limiter
}

func CleanupClient() {
	for {
		time.Sleep(time.Minute)
		mu.Lock()
		for ip, client := range clients {
			if time.Since(client.lastRequest) > time.Minute {
				delete(clients, ip)
			}
		}
		mu.Unlock()
	}
}

func RateLimitingMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ip := getClientIP(ctx)
		limiter := getRateLimiter(ip)
		if !limiter.Allow() {
			ctx.JSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
