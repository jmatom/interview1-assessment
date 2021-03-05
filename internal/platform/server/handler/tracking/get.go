package handler_tracking

import (
	"fmt"
	"interview1-assessment/internal/tracking"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createEventHandler returns an HTTP handler to increment a counter
func CreateEventHandlerGetVisits(counter tracking.CounterRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		url, err := tracking.NewUrl(ctx.Query("url"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		visits := counter.GetVisits(url)

		ctx.String(200, "%d", visits)

		fmt.Printf("url hash %s\n", url.Hash())
		fmt.Printf("url string %s\n", url.String())
		fmt.Printf("url visits %d\n", visits)
	}
}
