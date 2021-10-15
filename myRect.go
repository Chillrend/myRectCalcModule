package myRect

type Rect struct {
	Len, Wid int
}

func (rec Rect) CalAndReturn() (int, int, int) {
	return rec.Len * rec.Wid, rec.Len, rec.Wid
}

// func main() {
// 	r := Rect{10, 12}
// 	fmt.Println("Panjang dan Lebar kotak: ", r)
// 	area, len, wid := r.calAndReturn()
// 	fmt.Println("Luas kotak, Panjang, dan Lebar kotak: ", area, len, wid)
// }
