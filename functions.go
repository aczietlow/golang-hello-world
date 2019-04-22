package functions

func functionss() {
	// Can use the blank identifier if we only want 1 of a functions return values
	area, _ := rectProps2(15.2, 15.6)
	println(area)
}

//func functionname(parametername type) returntype {
func addNum(a int, b int) int {
	var sum = a + b
	return sum
}

// multiple return types
// func functionname(parametername type) (float32, float64)
func rectProps(length, width float64)(float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// return values can be named
func rectProps2(length, width float64)(area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

// Blank identifier is _
