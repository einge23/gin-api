package routes

import (
	"gin-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.Login)
		}
		printers := api.Group("/printers")
		{
			printers.GET("/getPrinters", controllers.GetPrinters)
			printers.PUT("/reservePrinter", controllers.ReservePrinter)
		}
		admin := api.Group(("/admin"))
		{
			users := admin.Group("/users")
			{
				users.POST("/create", controllers.CreateUser)
				userId := users.Group(("/:userID"))
				{
					userId.PUT("/setTrained", controllers.SetUserTrained)
				}
			}
		}
	}
}
