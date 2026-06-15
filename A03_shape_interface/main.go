// =============================================================
// A03: Shape Interface — Polymorphism ผ่าน Interface
// =============================================================
// ใน Go, interface กำหนด "พฤติกรรม" ไม่ใช่ inheritance
// type ใดที่มี method ครบตาม interface → ถือว่า implement แล้วอัตโนมัติ
// ไม่ต้องประกาศ "implements" เหมือนภาษาอื่น — เรียกว่า implicit satisfaction
// =============================================================

package main

import (
	"fmt"
	"math"
)

// -------------------------------------------------------------
// Interface: Shape
// -------------------------------------------------------------
// Shape คือ contract — type ใดที่มี Area() float64 ถือว่าเป็น Shape
// compiler ตรวจสอบตอน compile ไม่ใช่ runtime

type Shape interface {
	Area() float64
}

// -------------------------------------------------------------
// Struct: Rectangle
// -------------------------------------------------------------
// Rectangle ไม่ได้ประกาศว่า implement Shape — Go รู้เองจาก method

type Rectangle struct {
	Width  float64
	Height float64
}

// Area ทำให้ Rectangle เป็น Shape โดยอัตโนมัติ
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// -------------------------------------------------------------
// Struct: Circle
// -------------------------------------------------------------
// Circle และ Rectangle ต่างกันโดยสิ้นเชิง แต่ทั้งคู่เป็น Shape
// เพราะต่างก็มี Area() float64

type Circle struct {
	Radius float64
}

// Area คำนวณพื้นที่วงกลม: π × r²
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// -------------------------------------------------------------
// Polymorphic function: PrintArea
// -------------------------------------------------------------
// PrintArea รับ Shape ใดก็ได้ — ไม่สนว่าเป็น Rectangle หรือ Circle
// Go เลือก Area() ที่ถูกต้องให้เองตอน runtime (dynamic dispatch)
// %T แสดงชื่อ type จริงๆ เช่น main.Rectangle หรือ main.Circle

func PrintArea(s Shape) {
	fmt.Printf("%T พื้นที่: %.2f\n", s, s.Area())
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circle := Circle{Radius: 4}

	// ทั้งคู่ส่งเข้า function เดียวกัน — นี่คือ polymorphism
	PrintArea(rect)
	PrintArea(circle)
}
