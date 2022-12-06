package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"nextbasis-service-v-0.1/db/repository/models"
	"nextbasis-service-v-0.1/usecase"
)

// FireBaseUIDHandler ...
type FireBaseUIDHandler struct {
	Handler
}

// SelectAll ...
func (h *FireBaseUIDHandler) GetFirebaseUIDData(ctx *fiber.Ctx) error {

	type firebaseUID struct {
		ListUID []string `json:"agentUidList"`
	}

	ObjectData := new(firebaseUID)

	c := ctx.Locals("ctx").(context.Context)
	uc := usecase.FireBaseUIDUC{ContractUC: h.ContractUC}
	res, err := uc.SelectAll(c, models.FireStoreUserParameter{By: "def.created_date"})
	if err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}

	for _, data := range res {
		ObjectData.ListUID = append(ObjectData.ListUID, data.UID)
	}

	return h.SendResponse(ctx, ObjectData, nil, nil, 0)
}

func (h *FireBaseUIDHandler) SyncData(ctx *fiber.Ctx) error {
	c := ctx.Locals("ctx").(context.Context)
	uc := usecase.FireBaseUIDUC{ContractUC: h.ContractUC}
	res, err := uc.SyncData(c)
	if err != nil {

	}
	return h.SendResponse(ctx, res, nil, nil, 0)
}
