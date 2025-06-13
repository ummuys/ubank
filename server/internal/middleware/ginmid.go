package middleware

import (
	"net/http"
	"strings"
	"ubank/internal/tools"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AddZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(g *gin.Context) {
		logger := logger.With(
			zap.String("method", g.Request.Method),
			zap.String("path", g.Request.URL.Path),
			zap.String("ip", g.ClientIP()),
		)
		g.Set("logger", logger)
		g.Next()
	}
}

func JWTRequest() gin.HandlerFunc {
	return func(g *gin.Context) {
		header := g.GetHeader("Authorization")
		if !strings.HasPrefix(header, "Bearer ") {
			g.JSON(http.StatusUnauthorized, gin.H{"message": "missing or invalid token"})
			return
		}

		token := strings.TrimPrefix(header, "Bearer ")
		claims, err := tools.ValidateJWT(token, tools.KeyFunc)
		if err != nil {
			g.JSON(http.StatusUnauthorized, gin.H{"message": "invalid or expired token"})
			return
		}

		g.Set("login", claims["login"])
		g.Next()
	}
}
