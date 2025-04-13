package routes

import (
	"net/http"
	"strconv"

	"example.com/restapp/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not convert to event id"})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch event"})
	}

	event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event"})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user registered!"})

}

func cancelRegistration(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not convert to event id"})
		return
	}

	var event models.Event

	event.ID = eventId

	err = event.CancelRegistration(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not cancel registration"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "registeration cancelled!"})

}
