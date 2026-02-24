package main
import ("fmt"; "math") // math.Sqrt

type Circle struct{
	x, y, r float64
}

type Rectangle struct {
	x1, x2, y1, y2 float64
}

func main(){
	c := new(Circle)
	c = &Circle{2, 2, 4}
	fmt.Println(*c, math.Sqrt(4))
	fmt.Println("Area of Circle", areaCircle(c))
	fmt.Println("Area of Circle using method", c.area())
	rect := new(Rectangle)
	rect = &Rectangle{1, 2, 3, 4}
	fmt.Println("Area of Rectange:", rectangleArea(rect))
	fmt.Println("Area of Rectange using method:", rect.area())
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2-x1
	b := y2-y2
	fmt.Println(a, b, a*a+b*b)
	distance := math.Sqrt(a*a + b*b)
	return distance
}

func areaCircle(c *Circle) float64 {
	fmt.Println(c.r)
	area := math.Pi * (c.r) * (c.r) 
	return area
}

func rectangleArea(rect *Rectangle) float64 {
	// 1, 3, 1, 4
	// 1, 3, 2, 3
	l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	w := distance(rect.x1, rect.y1, rect.x2, rect.y1)
	return l*w
}

func (c *Circle) area() float64{
	fmt.Println(c.r)
	area := math.Pi * (c.r) * (c.r) 
	return area
}
// receiver
func (rect *Rectangle) area() float64{
	l := distance(rect.x1, rect.y1, rect.x1, rect.y2)
	w := distance(rect.x1, rect.y1, rect.x2, rect.y1)
	return l*w
}