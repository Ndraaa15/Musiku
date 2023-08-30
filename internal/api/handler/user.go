package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Ndraaa15/musiku/global/errors"
	e "github.com/Ndraaa15/musiku/global/errors"
	"github.com/Ndraaa15/musiku/global/response"
	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func (h *Handler) Register(ctx *gin.Context) {
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

	req := entity.UserRegister{}

	if err = ctx.ShouldBindJSON(&req); err != nil {
		err = errors.ErrInvalidRequest
		message = "Invalid request body"
		code = http.StatusBadRequest
		return
	}

	user, err := h.User.Register(&req, c)

	if err != nil {
		err = errors.ErrInternalServer
		message = "Failed to register user"
		code = http.StatusInternalServerError
		return
	}

	select {
	case <-c.Done():
		err = e.ErrRequestTimeout
		code = http.StatusRequestTimeout
		message = "Request timeout"
	default:
		message = "Success to register user and please verify your email"
		data = user
	}
}

func (h *Handler) VerifyAccount(ctx *gin.Context) {
	id := ctx.Param("id")

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
		http.Redirect(ctx.Writer, ctx.Request, "http://localhost:3000/login", http.StatusFound)
	}()

	uuid, err := uuid.FromString(id)
	if err != nil {
		err = errors.ErrInvalidRequest
		message = "Invalid request body"
		code = http.StatusBadRequest
		return
	}

	_, err = h.User.VerifyAccount(uuid, c)

	select {
	case <-c.Done():
		err = errors.ErrRequestTimeout
		code = http.StatusRequestTimeout
		message = "Request timeout"
	}

}

func (h *Handler) Login(ctx *gin.Context) {
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
		err = e.ErrRequestTimeout
		code = http.StatusRequestTimeout
	default:

	}

}
