package service

import (
	"encoding/json"
	"fmt"
	"interview/models"
	"io"
	"net/http"
)

func GetWongnaiData() (models.CovidData, error) {
	response, err := http.Get("https://static.wongnai.com/devinterview/covid-cases.json")
	if err != nil {
		return models.CovidData{}, fmt.Errorf("เกิดข้อผิดพลาดในการดึงข้อมูล: %v\n", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return models.CovidData{}, fmt.Errorf("เกิดข้อผิดพลาดในการอ่านข้อมูล: %v\n", err)
	}

	var res models.CovidData
	err = json.Unmarshal(body, &res)
	if err != nil {
		return models.CovidData{}, fmt.Errorf("เกิดข้อผิดพลาดในการแปลง JSON: %v\n", err)
	}
	return res, nil
}
