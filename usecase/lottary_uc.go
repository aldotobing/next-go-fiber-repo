package usecase

import (
	"context"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"nextbasis-service-v-0.1/db/repository"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/pkg/functioncaller"
	"nextbasis-service-v-0.1/pkg/logruslogger"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/server/requests"
	"nextbasis-service-v-0.1/usecase/viewmodel"
)

// LottaryUC ...
type LottaryUC struct {
	*ContractUC
}

// BuildBody ...
func (uc LottaryUC) BuildBody(data *models.Lottary, res *viewmodel.LottaryVM) {
	res.ID = data.ID
	res.SerialNo = data.SerailNo
	res.Status = data.Status
	res.CustomerCode = data.CustomerCode
	res.CustomerName = data.CustomerName.String
	res.CreatedAt = data.CreatedAt
	res.UpdatedAt = data.UpdatedAt.String
	res.DeletedAt = data.DeletedAt.String

	res.Year = data.Year.String
	res.Quartal = data.Quartal.String
	res.Sequence = data.Sequence.String
	res.BranchName = data.BranchName.String
	res.RegionCode = data.RegionCode.String
	res.RegionName = data.RegionName.String
	res.RegionGroup = data.RegionGroup.String
	res.CustomerCpName = data.CustomerCpName.String
	res.CustomerLevel = data.CustomerLevel.String
	res.CustomerType = data.CustomerType.String
}

// FindAll ...
func (uc LottaryUC) FindAll(c context.Context, parameter models.LottaryParameter) (out []viewmodel.LottaryVM, p viewmodel.PaginationVM, err error) {
	parameter.Offset, parameter.Limit, parameter.Page, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.LottaryOrderBy, models.LottaryOrderByrByString)

	repo := repository.NewLottaryRepository(uc.DB)
	data, count, err := repo.FindAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	p = uc.setPaginationResponse(parameter.Page, parameter.Limit, count)

	for _, datum := range data {
		var temp viewmodel.LottaryVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.LottaryVM, 0)
	}

	return
}

// SelectAll ...
func (uc LottaryUC) SelectAll(c context.Context, parameter models.LottaryParameter) (out []viewmodel.LottaryVM, err error) {
	_, _, _, parameter.By, parameter.Sort = uc.setPaginationParameter(parameter.Page, parameter.Limit, parameter.By, parameter.Sort, models.LottaryOrderBy, models.LottaryOrderByrByString)

	repo := repository.NewLottaryRepository(uc.DB)
	data, err := repo.SelectAll(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	for _, datum := range data {
		var temp viewmodel.LottaryVM
		uc.BuildBody(&datum, &temp)

		out = append(out, temp)
	}

	if out == nil {
		out = make([]viewmodel.LottaryVM, 0)
	}

	return
}

// FindByID ...
func (uc LottaryUC) FindByID(c context.Context, parameter models.LottaryParameter) (out viewmodel.LottaryVM, err error) {
	repo := repository.NewLottaryRepository(uc.DB)
	data, err := repo.FindByID(c, parameter)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// FindByCustomerCode ...
func (uc LottaryUC) FindByCustomerCode(c context.Context, customerCode string) (out viewmodel.LottaryVM, err error) {
	repo := repository.NewLottaryRepository(uc.DB)
	data, err := repo.FindByCustomerCode(c, customerCode)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	uc.BuildBody(&data, &out)

	return
}

// Add ...
func (uc LottaryUC) Import(c context.Context, in requests.LottaryRequestHeader) (out viewmodel.LottaryVM, err error) {
	for _, datum := range in.Detail {

		cusrepo := repository.NewWebCustomerRepository(uc.DB)
		cus, errcus := cusrepo.FindByCodes(c, models.WebCustomerParameter{Code: `'` + datum.CustomerCode + `'`})

		if errcus == nil {
			if len(cus) > 0 {

				customer := cus[0]
				for i := 0; i < datum.Jumlah; i++ {

					SerailNo := uc.GenerateRandNo(c, models.LottaryParameter{CustomerID: customer.ID.String, Quartal: in.Quartal, Year: in.Year})
					sequence := i + 1
					repo := repository.NewLottaryRepository(uc.DB)
					datamodel := new(viewmodel.LottaryVM)
					datamodel.CustomerID = customer.ID.String
					datamodel.SerialNo = SerailNo
					datamodel.Quartal = in.Quartal
					datamodel.Year = in.Year
					datamodel.Sequence = strconv.Itoa(sequence)
					err = repo.Add(c, *datamodel)
					if err != nil {
						logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
						return
					}
				}
			} else {
				cacheKeyerr := "err_import_lottary: " + *&datum.CustomerCode + " jumlah : " + strconv.Itoa(datum.Jumlah)
				datamodel := new(viewmodel.LottaryVM)
				datamodel.CustomerCode = datum.CustomerCode
				datamodel.Year = in.Year
				datamodel.Quartal = in.Quartal

				errjsonData, errmars := json.Marshal(datamodel)
				if errmars == nil {
					uc.RedisClient.Client.Set(cacheKeyerr, errjsonData, time.Hour*168).Err()
				}

			}
		}
	}

	return
}

// Update ...
func (uc LottaryUC) Update(c context.Context, id string, in requests.LottaryRequest) (out viewmodel.LottaryVM, err error) {
	out = viewmodel.LottaryVM{
		ID:           id,
		CustomerCode: in.CustomerCode,
	}

	repo := repository.NewLottaryRepository(uc.DB)
	out.ID, err = repo.Update(c, out)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc LottaryUC) Delete(c context.Context, in string) (err error) {
	repo := repository.NewLottaryRepository(uc.DB)
	_, err = repo.Delete(c, in)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

// Delete ...
func (uc LottaryUC) DeleteAll(c context.Context) (err error) {
	repo := repository.NewLottaryRepository(uc.DB)
	err = repo.DeleteAll(c)
	if err != nil {
		logruslogger.Log(logruslogger.WarnLevel, err.Error(), functioncaller.PrintFuncName(), "query", c.Value("requestid"))
		return
	}

	return
}

func (uc LottaryUC) GenerateRandNo(c context.Context, parameter models.LottaryParameter) string {
	var code string
	repo := repository.NewLottaryRepository(uc.DB)
	for {
		code = str.RandAllAlphanumericString(8)
		parameter.SerialNo = code
		_, err := repo.FindExsistingLottaryCupon(c, parameter)
		if strings.Contains(err.Error(), "no rows in result set") {
			break
		}
	}
	return code
}
