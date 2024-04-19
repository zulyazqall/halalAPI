package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *App) startService() error {
	// userRepo := userRepository.NewRepository(app.db)
	// userUC := userUseCase.NewUseCase(userRepo, app.cfg)
	// userCTRL := userV1.NewHandlers(userUC)

	domain := app.gin.Group("/api/v1/")
	// domain.GET("/ping", func(c *gin.Context) error {
	// 	return c.String(http.StatusOK, "Hello Word 👋")
	// })

	domain.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	// userCTRL.UserRoutes(domain, app.cfg)

	return nil
}
