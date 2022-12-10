package day9

type Direction struct {
	x int
	y int
}
type Point struct {
	x int
	y int
}

func (p *Point) Move(direction Direction) {
	p.x += direction.x
	p.y += direction.y
}

func (p *Point) touching(anotherPoint Point) bool {
	return abs(anotherPoint.x-p.x) <= 1 && abs(anotherPoint.y-p.y) <= 1
}

func (p *Point) Follow(anotherPoint Point) {
	if !p.touching(anotherPoint) {

		if anotherPoint.x != p.x {
			p.x += (anotherPoint.x - p.x) / abs(anotherPoint.x-p.x)
		}
		if anotherPoint.y != p.y {
			p.y += (anotherPoint.y - p.y) / abs(anotherPoint.y-p.y)
		}
	}

}
