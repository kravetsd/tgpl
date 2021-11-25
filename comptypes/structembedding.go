package comptypes

import (
	"fmt"
)

func Embstructs() {
	fmt.Println("Hello structs embedding")
	type Point struct {
		X, Y int
	}
	type Circle struct {
		Radius int
		Point
	}

	type Wheel struct {
		Circle
		Spokes int
	}

	var w Wheel
	w.X = 3
	w.Y = 1
	w.Radius = 15
	w.Spokes = 10

	//w = Wheel{8, 8, 5, 20}                       // compile error: unknown fields
	//w = Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields

	w = Wheel{Spokes: 10, Circle: Circle{Radius: 15, Point: Point{X: 3, Y: 1}}}
	fmt.Printf("%#v\n", w)

	//more preety way to define:

	w = Wheel{
		Spokes: 10,
		Circle: Circle{
			Radius: 15,
			Point: Point{
				X: 3,
				Y: 1,
			},
		},
	}

	fmt.Printf("%#v\n", w)

	w.X = 41
	fmt.Printf("%#v\n", w)

}
