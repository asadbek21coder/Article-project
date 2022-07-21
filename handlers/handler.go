package handlers

import (
	"strconv"

	"project6/config"
	"project6/storage"

	"github.com/gin-gonic/gin"
)

type HandlerImpl struct {
	strg storage.StorageI
	cfg  config.Config
}

func NewHandler(s storage.StorageI, cfg config.Config) HandlerImpl {
	return HandlerImpl{
		strg: s,
		cfg:  cfg,
	}
}

func (h *HandlerImpl) parseOffsetQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("offset", h.cfg.DefaultOffset))
}

func (h *HandlerImpl) parseLimitQueryParam(c *gin.Context) (int, error) {
	return strconv.Atoi(c.DefaultQuery("limit", h.cfg.DefaultLimit))
}
