package jsonop

func boolOps(boolA, boolB bool, op Ops) bool {

	switch op {
	case OpAdd:
		return boolA || boolB
	case OpDelete:
		return boolA
	}
	return boolA
}

func intOps(intA, intB int64, op Ops) int64 {
	switch op {
	case OpAdd:
		return intA + intB
	case OpDelete:
		return intA - intB
	}
	return intA
}

func float64Ops(floatA, floatB float64, op Ops) float64 {

	switch op {
	case OpAdd:
		return floatA + floatB
	case OpDelete:
		return floatA - floatB
	}
	return floatA
}

func stringOps(strA, strB string, op Ops) string {
	switch op {
	case OpAdd:
		return strA + strB
	case OpDelete:
		return strA
	}
	return strA
}

func sliceOps(sliceA, sliceB []interface{}, op Ops) []interface{} {
	switch op {
	case OpAdd:

		for _, elmt := range sliceB {
			sliceA = append(sliceA, elmt)
		}
		return sliceA
	}
	return sliceA

}
