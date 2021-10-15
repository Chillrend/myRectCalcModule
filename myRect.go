package myRect

type Rect struct {
	len, wid int
}

func (rec Rect) calAndReturn() (int, int, int) {
	return rec.len * rec.wid, rec.len, rec.wid
}

// func main() {
// 	r := Rect{10, 12}
// 	fmt.Println("Panjang dan Lebar kotak: ", r)
// 	area, len, wid := r.calAndReturn()
// 	fmt.Println("Luas kotak, Panjang, dan Lebar kotak: ", area, len, wid)
// }
