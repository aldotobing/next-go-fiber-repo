package usecase

import (
	"context"
	"fmt"

	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
)

// BroadcastUC ...
type BroadcastUC struct {
	*ContractUC
}

// FindAll ...
func (uc BroadcastUC) Broadcast(c context.Context, input *requests.BroadcastRequest) (err error) {
	data, err := CustomerUC{ContractUC: uc.ContractUC}.FindByID(c, models.CustomerParameter{ID: "40903527"})
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	// var to []string
	// for i := range data {
	// }
	fmt.Println(*data.CustomerFCMToken)
	_, err = FCMUC{ContractUC: uc.ContractUC}.SendFCMMessage(c, input.Title, input.Body, *data.CustomerFCMToken)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "send_message", c.Value("requestid"))
		return
	}

	return err
}
