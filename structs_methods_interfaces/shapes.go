package main

type Rectangle struct {
	Width  float64
	Height float64
}

// using a struct in a function (in this case from a test) is defined by creating method name "rectangle" and then struct itself "Rectangle"
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}