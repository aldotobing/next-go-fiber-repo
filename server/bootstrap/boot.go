package bootstrap

import (
	"nextbasis-service-v-0.1/usecase"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type Bootstrap struct {
	App        *fiber.App
	ContractUC usecase.ContractUC
	Validator  *validator.Validate
	Translator ut.Translator
}
