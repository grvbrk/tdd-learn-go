package structs

import "testing"

func TestStructs(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, testcase := range areaTests {
		t.Run(testcase.name, func(t *testing.T) {
			got := testcase.shape.Area()
			if got != testcase.hasArea {
				t.Errorf("%#v got %g want %g", testcase.shape, got, testcase.hasArea)
			}
		})
	}
	// t.Run("Check rectangle perimeter", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	got := rectangle.Perimeter()
	// 	want := 40.0

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("Check rectangle Area", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	got := rectangle.Area()
	// 	want := 72.0

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("Check rectangle Area", func(t *testing.T) {
	// 	rectangle := Rectangle{12.0, 6.0}
	// 	got := rectangle.Area()
	// 	want := 72.0

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// })

	// t.Run("Check circle Area", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	got := circle.Area()
	// 	want := 314.1592653589793

	// 	if got != want {
	// 		t.Errorf("got %g want %g", got, want)
	// 	}
	// })
}
