package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)


type Input struct {
	Msisdn string `json:"msisdn"`
}

type Parsed struct {
	MnoIdentifier     string `json:"mno_identifier"`
	CountryCode       int32  `json:"country_code"`
	SubscriberNumber  string `json:"subscriber_number"`
	CountryIdentifier string `json:"country_identifier"`
}

func ParseMsisdn(msisdn string) (*Parsed, error) {
	return &Parsed{
		MnoIdentifier:     "",
		CountryCode:       0,
		SubscriberNumber:  "",
		CountryIdentifier: "",
	}, nil
}

func PostMsisdn(c *gin.Context) {
	var input Input
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log.Println(input)

	parsed, err := ParseMsisdn(input.Msisdn)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, parsed)
}

func main() {
	router := gin.Default()

	router.POST("/msisdn", PostMsisdn)

	router.Run(":8080")
}
