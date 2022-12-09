package day8

type Point struct {
	row, column int
}

func (point *Point) WithinRange(rows int, columns int) bool {
	return point.row >= 0 && point.row < rows && point.column >= 0 && point.column < columns
}

func (point *Point) Move(direction Point) {
	point.row += direction.row
	point.column += direction.column
}
