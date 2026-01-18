package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"restapi-go/models"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	eventIdParam := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	userID := context.GetInt64("userID")
	err = models.RegisterUserForEvent(userID, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event"})
}

func unregisterFromEvent(context *gin.Context) {
	eventIdParam := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdParam, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event ID"})
		return
	}
	userID := context.GetInt64("userID")
	err = models.UnregisterUserFromEvent(userID, eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not unregister from event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Successfully unregistered from event"})
}
