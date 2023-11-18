package sprint

const (
	PI = 3.14
)

type Circle struct {
	Radius  float32
	Diameter  float32
	Area  float32
	Perimeter  float32
}

func GetCircle(r float32) Circle {
newCircle := Circle{
	Radius:    r,
	Diameter:  2*r,
	Area:      r*r*PI,
	Perimeter: PI*2*r,
}
return newCircle
}
