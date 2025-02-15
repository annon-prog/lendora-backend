package route

import (
	"notification-service/controllers"

	"github.com/gin-gonic/gin"
)

/**
 * @Function Name: Routes
 *
 * @Description:
 * Function used to handle route set up
 *
 * @Params:
 * router *gin.Engine
 *
 * @Returns:
 */
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
