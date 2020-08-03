package Node

type SimplifyEnum string

// Each level covers the previous
const (
	NONE                SimplifyEnum = "NONE"
	NUMERIC_BASIC_OPS_1 SimplifyEnum = "NUMERIC_BASIC_OPS"
	NUMERIC_FUNCTIONS_2 SimplifyEnum = "NUMERIC_FUNCTIONS"
)

type TreeParsingConfig struct {
	Simplify SimplifyEnum
}

var operations = map[string]bool{
	"+":  true,
	"-":  true,
	"*":  true,
	"/":  true,
	"^":  true,
	"**": true,
}
