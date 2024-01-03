package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"nextbasis-service-v-0.1/db/repository/models"
)

func FetchClientDataTarget(params models.CustomerTargetSemesterParameter) (res interface{}) {
	jsonReq, err := json.Marshal(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/customertarget/data/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	type resutlData struct {
		SemesterTarget string `json:"semester_target"`
		CurrentTarget  string `json:"current_target"`
		QuartalTarget  string `json:"quartal_target"`
		AnualTarget    string `json:"anual_target"`
	}

	objectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &objectData)

	return objectData
}

func FetchVisitDay(params models.CustomerParameter) interface{} {
	jsonReq, err := json.Marshal(params)
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://nextbasis.id:8080/mysmagonsrv/rest/customer/visitday/1", bytes.NewBuffer(jsonReq))
	if err != nil {
		fmt.Println("client err")
		fmt.Print(err.Error())
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer C2A5CE6A2292E7745CE5A3F7E68A9")

	resp, err := client.Do(req)
	if err != nil {

		fmt.Print(err.Error())
	}
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	type resutlData struct {
		VisitDay  string `json:"visit_day"`
		VisitWeek string `json:"visit_week"`
	}

	ObjectData := new(resutlData)

	// var responseObject http.Response
	json.Unmarshal(bodyBytes, &ObjectData)

	return ObjectData
}
