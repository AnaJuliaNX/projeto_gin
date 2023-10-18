package validacoes

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

// Vai retornar um booleano, ou seja, se a validação é verdadeira ou falsa
func ValidacaoTitulo(field validator.FieldLevel) bool {
	return strings.Contains(field.Field().String(), "terra-do-nunca")
}
