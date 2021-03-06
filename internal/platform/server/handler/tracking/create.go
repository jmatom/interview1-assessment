package handler_tracking

import (
	"net/http"

	"interview1-assessment/internal/tracking"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Uid string `json:"uid" binding:"required"`
	Url string `json:"url" binding:"required"`
}

// createEventHandler returns an HTTP handler to increment a counter
func CreateEventHandler(counter tracking.CounterRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		trackingEvent, err := tracking.NewTrackingEvent(req.Uid, req.Url)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		counter.AddVisit(trackingEvent)

		ctx.String(http.StatusCreated, "")
	}
}
