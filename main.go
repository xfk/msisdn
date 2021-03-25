package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nyaruka/phonenumbers"
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
	trimPlus := strings.TrimLeft(msisdn, "+")
	trimZeroes := strings.TrimLeft(trimPlus, "0")
	prefixed := "+" + trimZeroes

	num, err := phonenumbers.Parse(prefixed, "")
	if err != nil {
		return nil, err
	}

	regionCode := phonenumbers.GetRegionCodeForNumber(num)
	carrier, err := phonenumbers.GetCarrierForNumber(num, regionCode)
	if err != nil {
		return nil, err
	}

	return &Parsed{
		MnoIdentifier:     carrier,
		CountryCode:       *num.CountryCode,
		SubscriberNumber:  phonenumbers.Format(num, phonenumbers.NATIONAL),
		CountryIdentifier: regionCode,
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
