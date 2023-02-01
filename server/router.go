package server

import (
	"epc.com/api/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	v1 := router.Group("/hk/api/v1")
	{
		memberGroup := v1.Group("/member")
		{
			member := new(controllers.MemberController)
			memberGroup.POST("/login", member.Login)
			memberGroup.POST("/register-member", member.RegisterMember)
			memberGroup.POST("/register-pm", member.RegisterPM)
			memberGroup.POST("/forgot-password", member.ForgotPassword)
			memberGroup.POST("/change-password", member.ChangePassword)
			memberGroup.GET("/profile", member.Profile)
		}

		taskMonitoringGroup := v1.Group("/task-monitoring")
		{
			dailyTaskGroup := taskMonitoringGroup.Group("/daily")
			{
				dailyTask := new(controllers.DailyTaskController)
				dailyTaskGroup.GET("/", dailyTask.GetAllTask)
			}
		}
	}

	return router
}
