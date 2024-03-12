package controller

import (
	"clean-API/internal/model"
	"clean-API/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ConferenceController interface {
	getConferences(*gin.Context)
	getConference(*gin.Context)
	createConference(*gin.Context)
	updateConference(*gin.Context)
	deleteConference(*gin.Context)
}

type conferenceController struct {
	conferenceService service.ConferenceService
}

func newConferenceController(conferenceService service.ConferenceService) ConferenceController {
	return &conferenceController{
		conferenceService: conferenceService,
	}
}

func (c conferenceController) getConferences(context *gin.Context) {

	conferences, err := c.conferenceService.GetAllConferences()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conferences."})
		return
	}
	context.JSON(http.StatusOK, conferences)
}

func (c conferenceController) getConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}

	conference, err := c.conferenceService.GetConferenceById(conferenceId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conference."})
		return
	}

	context.JSON(http.StatusOK, conference)
}

func (c conferenceController) createConference(context *gin.Context) {

	var conference model.Conference
	err := context.ShouldBindJSON(&conference)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	userId := context.GetInt64("userId")
	conference.UserID = userId

	conference, err = c.conferenceService.Save(conference)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create conference. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Conference created!", "conference": conference})
}

func (c conferenceController) updateConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}

	userId := context.GetInt64("userId")
	conference, err := c.conferenceService.GetConferenceById(conferenceId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch conference."})
		return
	}

	if conference.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update conference."})
		return
	}

	var updatedConference model.Conference

	err = context.ShouldBindJSON(&updatedConference)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	updatedConference.ID = conferenceId
	err = c.conferenceService.Update(updatedConference)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update conference."})
		return
	}
	updatedConference.UserID = userId
	context.JSON(http.StatusOK, gin.H{"message": "Conference updated successfully!", "conference": updatedConference})
}

func (c conferenceController) deleteConference(context *gin.Context) {
	conferenceId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse conference id."})
		return
	}
	userId := context.GetInt64("userId")
	conference, err := c.conferenceService.GetConferenceById(conferenceId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the conference."})
		return
	}

	if conference.UserID != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete conference."})
		return
	}
	err = c.conferenceService.Delete(conference)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the conference."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Conference deleted successfully!"})
}
