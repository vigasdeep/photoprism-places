package api

import (
	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism-places/internal/config"
	"github.com/photoprism/photoprism/pkg/txt"
)

var (
	ErrUnauthorized = gin.H{"code": 401, "error": txt.UcFirst(config.ErrUnauthorized.Error())}
)
