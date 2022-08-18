package handlers

import (
	"database/sql"
	"fmt"
	"go/types"
	"net/http"
	"strings"

	"nextbasis-service-v-0.1/pkg/jwe"
	"nextbasis-service-v-0.1/pkg/responsedto"
	"nextbasis-service-v-0.1/pkg/str"
	"nextbasis-service-v-0.1/usecase"
	"nextbasis-service-v-0.1/usecase/viewmodel"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	jwtFiber "github.com/gofiber/jwt/v2"
)

type Handler struct {
	FiberApp   *fiber.App
	ContractUC *usecase.ContractUC
	Jwe        jwe.Credential
	Db         *sql.DB
	Validator  *validator.Validate
	Translator ut.Translator
	JwtConfig  jwtFiber.Config
}

//base send response
func (h Handler) SendResponse(ctx *fiber.Ctx, data interface{}, meta interface{}, err interface{}, code int) error {
	if code == 0 && err != nil {
		code = http.StatusUnprocessableEntity
		err = err.(error).Error()
	}

	if code != http.StatusOK && err != nil {
		return h.SendErrorResponseWithCode(ctx, err, code)
	}

	return h.SendSuccessResponse(ctx, data, meta)
}

//send response if status code 200
func (h Handler) SendSuccessResponse(ctx *fiber.Ctx, data interface{}, meta interface{}) error {
	response := responsedto.SuccessResponse(data, meta)

	return ctx.Status(http.StatusOK).JSON(response)
}

//send response if status code != 200
func (h Handler) SendErrorResponse(ctx *fiber.Ctx, err interface{}, code int) error {
	response := responsedto.ErrorResponse(err)

	return ctx.Status(code).JSON(response)
}

func (h Handler) SendErrorResponseWithCode(ctx *fiber.Ctx, err interface{}, code int) error {
	response := responsedto.ErrorResponseWithCode(code, err)

	return ctx.Status(code).JSON(response)
}

//extract error message from validator
func (h Handler) ExtractErrorValidationMessages(error validator.ValidationErrors) map[string][]string {
	errorMessage := map[string][]string{}
	errorTranslation := error.Translate(h.Translator)

	for _, err := range error {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			strings.Replace(errorTranslation[err.Namespace()], err.StructField(), err.StructField(), -1),
		)
	}

	return errorMessage
}

//handling error
func (h Handler) RequestBodyHandling(ctx *fiber.Ctx, input *types.Type) error {
	if err := ctx.BodyParser(input); err != nil {
		return h.SendResponse(ctx, nil, nil, err, http.StatusBadRequest)
	}
	if err := h.Validator.Struct(input); err != nil {
		errMessage := h.ExtractErrorValidationMessages(err.(validator.ValidationErrors))
		return h.SendResponse(ctx, nil, nil, errMessage, http.StatusBadRequest)
	}

	return nil
}

// GetUser
func (h Handler) GetUser(ctx *fiber.Ctx) (res viewmodel.UserVM) {
	res.ID = fmt.Sprintf("%v", ctx.Locals("user_id"))

	return res
}
