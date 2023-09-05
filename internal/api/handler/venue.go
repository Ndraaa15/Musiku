package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/global/response"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllVenue(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 6000*time.Millisecond)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		data = "Success to get all venue"
	}
}

func (h *Handler) RentVenue(ctx *gin.Context) {
	c, cancel := context.WithTimeout(ctx.Request.Context(), 6000*time.Millisecond)
	defer cancel()

	var (
		err     error
		message string
		code    = http.StatusOK
		data    interface{}
	)

	defer func() {
		if err != nil {
			response.Error(ctx, code, err, message, data)
			return
		}
		response.Success(ctx, code, message, data)
	}()

	select {
	case <-c.Done():
		code = http.StatusRequestTimeout
		message = errors.ErrRequestTimeout.Error()
		return
	default:
		data = "Success to rent venue"
	}
}
