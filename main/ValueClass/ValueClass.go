package ValueClass

type ValueClass string

const (
	DERIVATIVE                  ValueClass = "DERIVATIVE"
	MATH_SYMBOL                 ValueClass = "MATH_SYMBOL"
	NUMERIC_CONSTANT            ValueClass = "NUMERIC_CONSTANT"
	IMAGINARY_CONSTANT          ValueClass = "IMAGINARY_CONSTANT"
	CONSTANT                    ValueClass = "CONSTANT"
	VARIABLE                    ValueClass = "VARIABLE"
	CONSTANT_EXPRESSION         ValueClass = "CONSTANT_EXPRESSION"
	VARIABLE_EXPRESSION         ValueClass = "VARIABLE_EXPRESSION"
	NUMERIC_CONSTANT_EXPRESSION ValueClass = "NUMERIC_CONSTANT_EXPRESSION"
	SYSTEM_FUNCTION             ValueClass = "SYSTEM_FUNCTION"
)

const (
	DERIVATIVE_RULE    ValueClass = "DERIVATIVE_RULE"
	NUMERIC_RULE       ValueClass = "NUMERIC_RULE"
	CONST_RULE         ValueClass = "CONST_RULE"
	VAR_RULE           ValueClass = "VAR_RULE"
	EXPR_RULE          ValueClass = "EXPR_RULE"
	GENERIC_EXPR_RULE  ValueClass = "GENERIC_EXPR_RULE"
	CONSTANT_EXPR_RULE ValueClass = "CONSTANT_EXPR_RULE"
	VARIABLE_EXPR_RULE ValueClass = "VARIABLE_EXPR_RULE"

	RULE_FUNCTION ValueClass = "RULE_FUNCTION"
)

func IsConstantRule(classByValues ValueClass) bool {
	return classByValues == CONST_RULE || classByValues == CONSTANT_EXPR_RULE || classByValues == NUMERIC_RULE
}

func IsGenericRule(classByValues ValueClass) bool {
	return classByValues == EXPR_RULE || classByValues == GENERIC_EXPR_RULE
}

func IsVariableRule(classByValues ValueClass) bool {
	return classByValues == VAR_RULE || classByValues == VARIABLE_EXPR_RULE || classByValues == DERIVATIVE_RULE
}

func IsNumeric(classByValues ValueClass) bool {
	return classByValues == NUMERIC_CONSTANT || classByValues == NUMERIC_CONSTANT_EXPRESSION
}

func IsVariable(classByValues ValueClass) bool {
	return classByValues == VARIABLE || classByValues == VARIABLE_EXPRESSION || classByValues == DERIVATIVE
}

func IsConstant(classByValues ValueClass) bool {
	return classByValues == CONSTANT ||
		classByValues == CONSTANT_EXPRESSION ||
		classByValues == MATH_SYMBOL ||
		classByValues == IMAGINARY_CONSTANT
}
