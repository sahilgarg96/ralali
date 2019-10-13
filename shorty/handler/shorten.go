package handler

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"math/rand"
	"ralali/shorty/render"
	"ralali/shorty/resource"
)

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

// Generate a random slug. The final length of this function is twice size
func slugify(url string, size int) string {
	b, err := generateRandomBytes(size)
	if err != nil {
		panic(err)
	}
	return hex.EncodeToString(b)
}

func shortenURL(c *gin.Context, url string, prefferedCode string) (string, error) {

	err := errors.New("Invalid URL")
	if resource.ValidateURL(url) != true {
		c.AbortWithError(422, err)
		return "", err
	}

	generatedCode := false
	shortenedURL := ""

	alreadyExists, errExists := GetURL(c, prefferedCode)
	err1 := errors.New("Already Exists Preferntial Shortcode")
	if errExists == nil {
		c.AbortWithError(409, err1)
	} else {
		generatedCode = true
		shortenedURL = prefferedCode
	}

	for !generatedCode {
		shortenedURL = slugify(url, 5)

		alreadyExists, errExists := GetURL(c, shortenedURL)

		if errExists != nil {
			generatedCode = true
		}
	}

	redisClient := c.MustGet("redis").(redis.Client)

	err = redisClient.Set(shortenedURL, url, 0).Err()

	if err != nil {
		panic(err)
	}
	return shortenedURL, nil
}

func GetURL(c *gin.Context, slug string) (string, error) {

	redisClient := c.MustGet("redis").(redis.Client)

	shortenedURL, err := redisClient.Get(slug).Result()
	return shortenedURL, err
}

func ShortenURLHandler(c *gin.Context) {

	type IncomingData struct {
		URL       string `json:"url" form:"url"`
		ShortCode string `json:"shortcode" form:"shortcode"`
	}

	var json IncomingData

	if err := c.ShouldBind(&json); err != nil {
		fmt.Println("err", err)
	}

	shortUrl, _ := shortenURL(c, json.URL, json.ShortCode)

	type Response struct {
		ShortCode string `json:"shortcode" form:"shortcode"`
	}

	var rsp Response

	rsp.ShortCode = shortUrl

	render.Render(c, gin.H{"payload": rsp})
}

func GetURLHandler(c *gin.Context) {

	type IncomingData struct {
		ShortCode string `json:"shortcode" form:"shortcode"`
	}

	var json IncomingData

	if err := c.ShouldBind(&json); err != nil {
		fmt.Println("err", err)
	}

	getUrl, err := GetURL(c, json.ShortCode)

	err1 := errors.New("Shortcode not found")

	if err1 != nil {
		c.AbortWithError(404, err1)
	}

	type Response struct {
		Location string `json:"location" form:"location"`
	}

	c.Status(302)

	var rsp Response

	rsp.Location = getUrl

	render.Render(c, gin.H{"payload": rsp})
}

func GetURLStats(c *gin.Context, slug string) (string, error) {

	redisClient := c.MustGet("redis").(redis.Client)

	json, err := redisClient.Get(slug + "stats").Result()

	return json, err
}

func GetShortcodeStatsHandler(c *gin.Context) {

	type IncomingData struct {
		ShortCode string `json:"shortcode" form:"shortcode"`
	}

	var json IncomingData

	if err := c.ShouldBind(&json); err != nil {
		fmt.Println("err", err)
	}

	statsJson, err := GetURLStats(c, json.ShortCode)

	err = errors.New("Shortcode not found")

	if err != nil {
		c.AbortWithError(404, err)
	}

	type Response struct {
		StartDate     string `json:"startDate" form:"startDate"`
		LastSeenDate  string `json:"lastSeenDate" form:"lastSeenDate"`
		RedirectCount int32  `json:"redirectCount" form:"redirectCount"`
	}

	b := []byte(string(statsJson))
	var rsp Response

	err := json.Unmarshal(b, &rsp)

	if err != nil {
		c.AbortWithError(404, err)
	}

	c.Status(200)

	render.Render(c, gin.H{"payload": rsp})
}
