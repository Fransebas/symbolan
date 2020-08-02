package OperationClass

type OperationClass string

const (
	ADDITION        OperationClass = "ADDITION"
	SUBSTRACTION    OperationClass = "SUBSTRACTION"
	MULTIPLICATION  OperationClass = "MULTIPLICATION"
	DIVISION        OperationClass = "DIVISION"
	EXPONENTIAL     OperationClass = "EXPONENTIAL"
	LOGARITHM       OperationClass = "LOGARITHM"
	SIN             OperationClass = "SIN" // SIN, COS, TAN, etc... not user defined functions
	COS             OperationClass = "COS"
	TAN             OperationClass = "TAN"
	TRANSFORMATION  OperationClass = "TRANSFORMATION"
	EQUATION        OperationClass = "EQUATION"
	LEAF            OperationClass = "LEAF"
	SYSTEM_FUNCTION OperationClass = "SYSTEM_FUNCTION"
	RULE            OperationClass = "RULE"
)
