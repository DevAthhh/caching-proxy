package controllers

import (
	"log"
	"net/http"

	"github.com/DevAthhh/caching-proxy/app/cmd/caching"
	"github.com/DevAthhh/caching-proxy/app/cmd/server/internal/utils"
	"github.com/gin-gonic/gin"
)

func Index(origin string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		url := origin + ctx.Param("path")
		cache, _ := caching.GetCache(url)
		if cache == "" {
			body := utils.GetOriginResponse(url)

			if err := caching.SetCache(url, string(body)); err != nil {
				log.Fatalf("error with setting cache: %v", err)
			}

			ctx.HTML(http.StatusOK, "index.html", gin.H{
				"Body": body,
			})
			return
		}
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"Body": cache,
		})
	}
}
