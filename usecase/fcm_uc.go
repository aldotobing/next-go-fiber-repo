package usecase

import (
	"context"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
)

type FCMUC struct {
	*ContractUC
}

// BuildBody ...
func (uc FCMUC) BuildBody(res *models.FireBaseCloudMessage) {
}

func (uc FCMUC) SendMessage(c context.Context, data *models.FireBaseCloudMessage) (res string, err error) {

	// res = data

	dataInterface := map[string]interface{}{
		"body":  data.Body,
		"title": data.Title,
		"type":  data.Type,
	}

	// fmt.Println("Token ", data.Token)

	// token := `f687FAOTQnapQ6BjnKNFEA:APA91bF1rlfor8mKYID9MSESvbkyaGi4bt4KDkwCP39XNMD6WyKIx79qEuvrQxIHEO6P5_neRoVj6M44pEQn4S6FicSQCkP9GLD4tRkJwtqclmWgIkVDs5Z1ieaaW_sDkh15en1GSnr4`
	res, err = uc.ContractUC.Fcm.SendAndroid([]string{data.Token}, data.Title, data.Body, dataInterface)

	if err != nil {
		fmt.Println(err.Error())
	}

	// fmt.Println(uc.ContractUC.Fcm.APIKey)

	return res, err
}
