package handlers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", "0"))
}

func parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", "10"))
}
