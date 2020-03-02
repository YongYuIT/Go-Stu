package excel

import "strconv"

type Point struct {
	X string
	Y string
}

func (thiz *Point) GetNextX() *Point {
	chars := []byte(thiz.X)
	lastChar := chars[len(chars)-1]
	if lastChar != 'Z' {
		chars[len(chars)-1] = (lastChar + 1)
	} else {
		if len(chars) == 1 {
			return &Point{"AA", thiz.Y}
		} else {
			chars[0] += 1
			chars[1] = 'A'
		}
	}
	return &Point{string(chars), thiz.Y}
}

func (thiz *Point) GetNextY() *Point {
	y, _ := strconv.Atoi(thiz.Y)
	y += 1
	return &Point{thiz.X, strconv.Itoa(y)}
}

func (thiz *Point) GetPosition() string {
	return thiz.X + thiz.Y
}

func (thiz *Point) GetDate() *Point {
	new_x := thiz.X
	new_y := "2"
	return &Point{new_x, new_y}
}

func (thiz *Point) GetDataItemName() *Point {
	new_x := thiz.X
	new_y := "3"
	return &Point{new_x, new_y}
}

func (thiz *Point) GetQu() *Point {
	new_x := "A"
	new_y := thiz.Y
	return &Point{new_x, new_y}
}

func (thiz *Point) GetOrgName() *Point {
	new_x := "C"
	new_y := thiz.Y
	return &Point{new_x, new_y}
}

func (thiz *Point) GetOrgType() *Point {
	new_x := "D"
	new_y := thiz.Y
	return &Point{new_x, new_y}
}
