package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/photoprism/photoprism-places/internal/config"
	"github.com/photoprism/photoprism-places/internal/entity"
)

// GET /v1/location/:id
func GetLocation(router *gin.RouterGroup, conf *config.Config) {
	router.GET("/location/:id", func(c *gin.Context) {
		id := c.Param("id")

		if len(id) > 16 || len(id) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid location id"})
			return
		}

		gc := conf.Cache()
		cacheKey := fmt.Sprintf("location:%s", id)

		if hit, ok := gc.Get(cacheKey); ok {
			c.JSON(http.StatusOK, hit)
			return
		}

		db := conf.Db()

		location, err := entity.FindLocation(id, db)

		if err != nil {
			log.Error(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		gc.Set(cacheKey, location, 3*time.Hour)

		c.JSON(http.StatusOK, location)
	})
}
