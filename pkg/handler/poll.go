package handler

import (
	"RestApiProj"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createPoll(c *gin.Context) {
	var input RestApiProj.INCreate
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreatePoll(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) Poll(c *gin.Context) {
	var input RestApiProj.INChoice
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.service.Polling(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, []byte("ok"))
}

type getPollingResponse struct {
	Data []RestApiProj.Poll `json:"data"`
}

func (h *Handler) getResult(c *gin.Context) {
	var input RestApiProj.INGet
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	polls, err := h.service.GetPoll(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getPollingResponse{
		Data: polls,
	})
}
