package handlers

// SubDistrictHandler ...
type SubDistrictHandler struct {
	Handler
}

// // SelectAll ...
// func (h *SubDistrictHandler) SelectAll(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.SubDistrictParameter{
// 		DistrictID: ctx.Query("district_id"),
// 		Search:     ctx.Query("search"),
// 		Status:     ctx.Query("status"),
// 		By:         ctx.Query("by"),
// 		Sort:       ctx.Query("sort"),
// 	}
// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.SelectAll(c, parameter)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// // FindAll ...
// func (h *SubDistrictHandler) FindAll(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.SubDistrictParameter{
// 		DistrictID: ctx.Query("district_id"),
// 		Search:     ctx.Query("search"),
// 		Status:     ctx.Query("status"),
// 		Page:       str.StringToInt(ctx.Query("page")),
// 		Limit:      str.StringToInt(ctx.Query("limit")),
// 		By:         ctx.Query("by"),
// 		Sort:       ctx.Query("sort"),
// 	}
// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.SelectAll(c, parameter)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// // FindByID ...
// func (h *SubDistrictHandler) FindByID(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	parameter := models.SubDistrictParameter{
// 		ID: ctx.Params("id"),
// 	}
// 	if parameter.ID == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.FindByID(c, parameter)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// // Add ...
// func (h *SubDistrictHandler) Add(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	input := new(requests.SubDistrictRequest)
// 	if err := ctx.BodyParser(input); err != nil {
// 		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
// 	}
// 	if err := h.Validator.Struct(input); err != nil {
// 		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
// 		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
// 	}

// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.Add(c, input)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// // Edit ...
// func (h *SubDistrictHandler) Edit(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	id := ctx.Params("id")
// 	if id == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	input := new(requests.SubDistrictRequest)
// 	if err := ctx.BodyParser(input); err != nil {
// 		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
// 	}
// 	if err := h.Validator.Struct(input); err != nil {
// 		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
// 		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
// 	}

// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.Edit(c, id, input)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }

// // Delete ...
// func (h *SubDistrictHandler) Delete(ctx *fiber.Ctx) error {
// 	c := ctx.Locals("ctx").(context.Context)

// 	id := ctx.Params("id")
// 	if id == "" {
// 		return h.SendResponse(ctx, nil, nil, helper.InvalidParameter, http.StatusBadRequest)
// 	}

// 	uc := usecase.SubDistrictUC{ContractUC: h.ContractUC}
// 	res, err := uc.Delete(c, id)

// 	return h.SendResponse(ctx, res, nil, err, 0)
// }
