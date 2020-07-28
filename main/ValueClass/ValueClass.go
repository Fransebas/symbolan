package ValueClass

const (
	NUMERIC_CONSTANT            = "NUMERIC_CONSTANT"
	CONSTANT                    = "CONSTANT"
	VARIABLE                    = "VARIABLE"
	CONSTANT_EXPRESSION         = "CONSTANT_EXPRESSION"
	VARIABLE_EXPRESSION         = "VARIABLE_EXPRESSION"
	NUMERIC_CONSTANT_EXPRESSION = "NUMERIC_CONSTANT_EXPRESSION"
	SYSTEM_FUNCTION             = "SYSTEM_FUNCTION"
)

func IsNumeric(classByValues string) bool {
	return classByValues == NUMERIC_CONSTANT || classByValues == NUMERIC_CONSTANT_EXPRESSION
}

func IsVariable(classByValues string) bool {
	return classByValues == VARIABLE || classByValues == VARIABLE_EXPRESSION
}

func IsConstant(classByValues string) bool {
	return classByValues == CONSTANT || classByValues == CONSTANT_EXPRESSION
}
