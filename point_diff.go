

package sprint

type Point struct {
	X, Y float32
	Text string
	
}

func PointDiff(p1, p2 Point) Point {

distanceP2 := p2.X + p2.Y
distanceP1 := p1.X + p1.Y 
if (distanceP1) > (distanceP2) {

return p1
}
return p2

}