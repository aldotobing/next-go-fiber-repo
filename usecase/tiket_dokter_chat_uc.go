package usecase

import (
	"context"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// TicketDokterChatUC ...
type TicketDokterChatUC struct {
	*ContractUC
}

// BuildBody ...
func (uc TicketDokterChatUC) BuildBody(data *models.TicketDokterChat, res *viewmodel.TicketDocterChatVM) {
	res.ID = data.ID
	res.TicketDokterID = data.TicketDokterID
	res.ChatBy = data.ChatBy
	res.Description = data.Description
	res.CreatedDate = data.CreatedDate
}

// SelectAll ...
func (uc TicketDokterChatUC) FindByTicketDocterID(c context.Context, parameter models.TicketDokterChatParameter, ticketDokterID string) (res []viewmodel.TicketDocterChatVM, err error) {
	_, _, _, _, parameter.Sort = uc.setPaginationParameter(0, 0, parameter.By, parameter.Sort, models.TicketDokterChatOrderBy, models.TicketDokterChatOrderByrByString)

	repo := repository.NewTicketDokterChatRepository(uc.DB)
	data, err := repo.FindByTicketDocterID(c, parameter, ticketDokterID)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	for i := range data {
		var temp viewmodel.TicketDocterChatVM
		uc.BuildBody(&data[i], &temp)

		res = append(res, temp)
	}

	if res == nil {
		res = make([]viewmodel.TicketDocterChatVM, 0)
	}

	return res, err
}

// Add ...
func (uc TicketDokterChatUC) Add(c context.Context, in *requests.TicketDokterChatRequest) (res viewmodel.TicketDocterChatVM, err error) {
	repo := repository.NewTicketDokterChatRepository(uc.DB)
	data := models.TicketDokterChat{
		TicketDokterID: &in.TicketDocterID,
		ChatBy:         &in.ChatBy,
		Description:    &in.Description,
		CreatedDate:    &in.CreatedDate,
	}
	data.ID, data.CreatedDate, err = repo.Add(c, &data)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return res, err
	}

	res = viewmodel.TicketDocterChatVM{
		ID:             data.ID,
		TicketDokterID: data.TicketDokterID,
		Description:    data.Description,
		ChatBy:         data.ChatBy,
		CreatedDate:    data.CreatedDate,
	}

	return res, err
}

// Delete ...
func (uc TicketDokterChatUC) Delete(c context.Context, id string) (res models.TicketDokterChat, err error) {
	repo := repository.NewTicketDokterChatRepository(uc.DB)
	err = repo.Delete(c, id)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}
