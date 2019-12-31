package api

import (
	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism-places/internal/config"
	"github.com/photoprism/photoprism-places/internal/util"
)

var (
	ErrUnauthorized = gin.H{"code": 401, "error": util.UcFirst(config.ErrUnauthorized.Error())}
)
