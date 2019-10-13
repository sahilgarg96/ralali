package main

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"flag"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"ralali/shorty/data/redis"
	"regexp"
	"strings"
)

var router *gin.Engine

func init() {

	redisPool, errR := redis.RedisConnect()

	// close once process exits
	defer redisClient.Close()
}

func main() {
	// Set the router as the default one provided by Gin
	router = gin.Default()

	// setup middleware for redis to pass it across project
	router.Use(setupFunc.RedisMiddleware(*redisClient))

	router.POST("/shorten", handler.ShortenURLHandler)
	router.GET("/:shortcode", handler.GetURLHandler)
	router.GET("/:shortcode/stats", handler.GetShortcodeStatsHandler)

	log.Fatal(http.ListenAndServe(":8080", router))
}
