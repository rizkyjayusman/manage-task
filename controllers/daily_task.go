package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DailyTaskController struct{}

func (dt DailyTaskController) GetAllTask(c *gin.Context) {
	requestId := uuid.New().String()
	log.Printf("INFO : [RequestId : %v] Execute DailyTaskController.countTask()", requestId)
	c.JSON(200, gin.H{
		"id":     requestId,
		"code":   http.StatusOK,
		"status": http.StatusText(http.StatusOK),
		"data":   gin.H{},
	})
}
