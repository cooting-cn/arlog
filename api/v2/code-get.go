package v2

import (
	"arlog/utils/res"
	"github.com/gin-gonic/gin"
)

func GetCode(c *gin.Context) {

	response := Response{
		Languages: []Language{
			{"Vue", "#41b883", 46.8},
			{"Go", "#00ADD8", 34.5},
			{"JavaScript", "#f1e05a", 16.3},
			{"CSS", "#563d7c", 0.9},
			{"Dockerfile", "#384d54", 0.7},
			{"Other", "#EDEDED", 0.8},
		},
	}

	res.Ask(c, 200, response)
	return
}

type Language struct {
	Language string  `json:"language"`
	Color    string  `json:"color"`
	Percent  float64 `json:"percent"`
}

type Response struct {
	Languages []Language `json:"languages"`
}
