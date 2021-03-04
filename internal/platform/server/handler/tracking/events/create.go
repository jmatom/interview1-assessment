package tracking_events

import (
	"net/http"

	tracking_event "interview1-assessment/internal"

	"github.com/gin-gonic/gin"
)

type createRequest struct {
	Uid string `json:"uid" binding:"required"`
	Url string `json:"url" binding:"required"`
}

// createEventHandler returns an HTTP handler to increment a counter
func CreateEventHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		/*
			course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			}
		*/
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		_, err := tracking_event.NewTrackingEvent(req.Uid, req.Url)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		// Event is accepted because counter will be increased eventually
		ctx.String(http.StatusAccepted, "Event registered")
	}
}
