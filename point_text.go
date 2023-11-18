package sprint

import "fmt"

type Point struct {
	X, Y float32
	Text string
}

func PointText(p Point) Point {

resultPoint := Point{}

resultPoint.X = p.X
resultPoint.Y =p.Y
resultPoint.Text = fmt.Sprintf("Text at (%f, %f)", resultPoint.X , resultPoint.Y)

return resultPoint
}
