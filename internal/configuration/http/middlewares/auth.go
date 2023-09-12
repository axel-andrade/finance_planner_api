package middlewares

import (
	"fmt"
	"net/http"

	"github.com/axel-andrade/finance_planner_api/internal/configuration/bootstrap"
	redis_repositories "github.com/axel-andrade/finance_planner_api/internal/configuration/database/redis/repositories"
	"github.com/axel-andrade/finance_planner_api/internal/configuration/handlers"
	shared_err "github.com/axel-andrade/finance_planner_api/internal/shared/errors"
	"github.com/gin-gonic/gin"
)

func Authorize(dependencies *bootstrap.Dependencies) gin.HandlerFunc {
	tokenManagerHandler := new(handlers.TokenManagerHandler)
	dependencies.Invoke(func(h *handlers.TokenManagerHandler) {
		tokenManagerHandler = h
	})

	sessionRepo := new(redis_repositories.SessionRepository)
	dependencies.Invoke(func(r *redis_repositories.SessionRepository) {
		sessionRepo = r
	})

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) == 0 {
			fmt.Println("message: authorization not informed")
			c.JSON(http.StatusUnauthorized, gin.H{"error": shared_err.UNAUTHORIZED})
			c.Abort()
			return
		}

		encodedToken := authHeader[len("Bearer "):]

		tokenAuth, err := tokenManagerHandler.ExtractTokenMetadata(encodedToken)
		if err != nil {
			fmt.Println("error: error in extract token metadata: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": shared_err.UNAUTHORIZED})
			c.Abort()
			return
		}

		userId, err := sessionRepo.GetAuth(tokenAuth)
		if err != nil {
			fmt.Println("error in get auth: ", err.Error())
			c.JSON(http.StatusUnauthorized, gin.H{"error": shared_err.UNAUTHORIZED})
			c.Abort()
			return
		}

		// TODO: verificar se o usuario existe no banco de dados
		c.Set("user-id", userId)

		c.Next()
	}
}
