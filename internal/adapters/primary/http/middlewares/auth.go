package middlewares

import (
	"fmt"
	"net/http"

	shared_err "github.com/axel-andrade/finance_planner_api/internal/core/domain/errors"
	"github.com/axel-andrade/finance_planner_api/internal/infra"
	"github.com/gin-gonic/gin"
)

func Authorize(dependencies *infra.Dependencies) gin.HandlerFunc {
	tokenManagerHandler := dependencies.TokenManagerHandler
	sessionRepo := dependencies.SessionRepository

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
		c.Set("user_id", userId)

		c.Next()
	}
}
