package usecase

import "time"

// VAGeneratorUC ...
type VAGeneratorUC struct {
	*ContractUC
}

func (uc VAGeneratorUC) Generate(VaPartnerCode, BranchCode, currentVaCount *string) string {
	formater := "0601"
	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	strnow := now.Format(formater)
	return strnow
}
