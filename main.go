package main

import (
	"interview/models"
	"interview/service"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/api/covid/summary", func(c *gin.Context) {
		c.JSON(200, logic())
	})
	r.Run()

}
func logic() models.Summary {
	res, err := service.GetWongnaiData()
	if err == nil {
		provinceCount := make(map[string]int)
		ageGroups := make(map[string]int)

		for _, detail := range res.Data {
			if detail.Province == nil {
				defaultProvince := "N/A"
				detail.Province = &defaultProvince
			}
			provinceCount[*detail.Province]++
			if detail.Age != nil {
				switch {
				case *detail.Age >= 0 && *detail.Age <= 30:
					ageGroups["0-30"]++
				case *detail.Age >= 31 && *detail.Age <= 60:
					ageGroups["31-60"]++
				case *detail.Age > 60:
					ageGroups["61+"]++
				}
			} else {
				ageGroups["N/A"]++
			}

		}
		return models.Summary{provinceCount, ageGroups}
	}
	return models.Summary{}
}
