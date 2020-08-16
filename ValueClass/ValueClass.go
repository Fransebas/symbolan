package ValueClass

type ValueClass string

const (
	DERIVATIVE         ValueClass = "DERIVATIVE"
	MATH_SYMBOL        ValueClass = "MATH_SYMBOL"
	NUMERIC_CONSTANT   ValueClass = "NUMERIC_CONSTANT"
	IMAGINARY_CONSTANT ValueClass = "IMAGINARY_CONSTANT"
	CONSTANT           ValueClass = "CONSTANT"
	VARIABLE           ValueClass = "VARIABLE"

	CONSTANT_EXPRESSION         ValueClass = "CONSTANT_EXPRESSION"
	VARIABLE_EXPRESSION         ValueClass = "VARIABLE_EXPRESSION"
	NUMERIC_CONSTANT_EXPRESSION ValueClass = "NUMERIC_CONSTANT_EXPRESSION"
	SYSTEM_FUNCTION             ValueClass = "SYSTEM_FUNCTION"
	SIGN                        ValueClass = "SIGN"
)

const (
	DERIVATIVE_RULE ValueClass = "DERIVATIVE_RULE"
	NUMERIC_RULE    ValueClass = "NUMERIC_RULE"
	CONST_RULE      ValueClass = "CONST_RULE"
	VAR_RULE        ValueClass = "VAR_RULE"
	GENERIC_RULE    ValueClass = "GENERIC_RULE"

	EXPR_GENERIC_RULE  ValueClass = "EXPR_GENERIC_RULE"
	EXPR_CONSTANT_RULE ValueClass = "EXPR_CONSTANT_RULE"
	EXPR_VARIABLE_RULE ValueClass = "EXPR_VARIABLE_RULE"

	NUMERIC_TREE_RULE  ValueClass = "NUMERIC_TREE_RULE"
	CONST_TREE_RULE    ValueClass = "CONST_TREE_RULE"
	VAR_TREE_RULE      ValueClass = "VAR_TREE_RULE"
	FUNCTION_TREE_RULE ValueClass = "FUNCTION_TREE_RULE"

	EXPR_NUMERIC_TREE_RULE  ValueClass = "EXPR_NUMERIC_TREE_RULE"
	EXPR_CONST_TREE_RULE    ValueClass = "EXPR_CONST_TREE_RULE"
	EXPR_VAR_TREE_RULE      ValueClass = "EXPR_VAR_TREE_RULE"
	EXPR_FUNCTION_TREE_RULE ValueClass = "EXPR_FUNCTION_TREE_RULE"

	CUSTOM_FUNCTION_RULE ValueClass = "CUSTOM_FUNCTION_RULE"
)

func IsConstantRule(classByValues ValueClass) bool {
	return classByValues == CONST_RULE || classByValues == EXPR_CONSTANT_RULE || classByValues == NUMERIC_RULE
}

func IsGenericRule(classByValues ValueClass) bool {
	return classByValues == GENERIC_RULE || classByValues == EXPR_GENERIC_RULE
}

func IsVariableRule(classByValues ValueClass) bool {
	return classByValues == VAR_RULE || classByValues == EXPR_VARIABLE_RULE || classByValues == DERIVATIVE_RULE
}

func IsNumeric(classByValues ValueClass) bool {
	return classByValues == NUMERIC_CONSTANT || classByValues == NUMERIC_CONSTANT_EXPRESSION
}

func IsTreeNumeric(classByValues ValueClass) bool {
	return classByValues == NUMERIC_TREE_RULE || classByValues == EXPR_NUMERIC_TREE_RULE
}

func IsTreeConstant(classByValues ValueClass) bool {
	return classByValues == CONST_TREE_RULE || classByValues == EXPR_CONST_TREE_RULE
}

func IsTreeVar(classByValues ValueClass) bool {
	return classByValues == VAR_TREE_RULE || classByValues == EXPR_VAR_TREE_RULE
}

func IsTreeFunction(classByValues ValueClass) bool {
	return classByValues == FUNCTION_TREE_RULE || classByValues == EXPR_FUNCTION_TREE_RULE
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
