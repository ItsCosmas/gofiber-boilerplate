package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// Validate validates the input struct
func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)

	if err != nil {
		var errors []string
		// TODO
		// Better Error Handling Here
		fmt.Println("------ List of tag fields with error ---------")
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(
				errors,
				fmt.Sprintf("`%v` with value `%v` doesn't satisfy the `%v` constraint", err.Field(), err.Value(), err.Tag()),
			)
			fmt.Println(err.StructField())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
		}

		fmt.Println("---------------")

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: strings.Join(errors, ","),
		}
	}

	return nil
}

// ParseBody is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBody(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := c.BodyParser(body); err != nil {
		fmt.Println("=======================")
		fmt.Println(err)
		fmt.Println("=======================")
		return fiber.ErrBadRequest
	}

	return nil
}

// ParseBodyAndValidate is helper function for parsing the body.
// Is any error occurs it will panic.
// Its just a helper function to avoid writing if condition again n again.
func ParseBodyAndValidate(c *fiber.Ctx, body interface{}) *fiber.Error {
	if err := ParseBody(c, body); err != nil {
		return err
	}

	return Validate(body)
}

// CUSTOM VALIDATION RULES =============================================

// Password validation rule: required,min=6,max=100
var _ = validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
	l := len(fl.Field().String())

	return l >= 6 && l < 100
})
