package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination"`
}

type Pagination struct {
	Page       int  `json:"page"`
	Limit      int  `json:"limit"`
	TotalPages int  `json:"total_pages"`
	HasNext    bool `json:"has_next"`
	HasPrev    bool `json:"has_prev"`
}

func NewPagination(page int, limit int, total int) *Pagination {
	totalPages := total / limit
	hasNext := page < totalPages
	hasPrev := page > 1

	return &Pagination{
		Page:       page,
		Limit:      limit,
		TotalPages: totalPages,
		HasNext:    hasNext,
		HasPrev:    hasPrev,
	}
}

func SuccessReponse(c *gin.Context, code int, data interface{}, pagination *Pagination) {
	c.JSON(http.StatusOK, ResponseData{
		Code:       code,
		Message:    msg[code],
		Data:       data,
		Pagination: pagination,
	})
}

func ErrorResponse(c *gin.Context, code int, message string) {
	if message == "" {
		message = msg[code]
	}

	c.JSON(http.StatusOK, ResponseData{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
