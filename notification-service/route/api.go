package route

import (
	"notification-service/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {

	//starter url
	router.GET("/", controllers.StarterPage)

	//url group used for notifications
	notify := router.Group("/notify")
	{
		//url used to send out email messages
		notify.POST("/email", controllers.SendEmails)
	}

}
