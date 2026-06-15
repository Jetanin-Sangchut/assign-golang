// =============================================================
// A02: Safe Counter — Thread-Safety ด้วย sync.Mutex
// =============================================================
// เมื่อ goroutine หลายตัวแก้ไขตัวแปรเดียวกันพร้อมกัน ผลลัพธ์จะผิดพลาด
// นั่นคือ "race condition" — แก้ด้วย sync.Mutex เพื่อให้ทีละตัวเท่านั้น
// =============================================================

package main

import (
	"fmt"
	"sync"
)

// SafeCounter คือ counter ที่ปลอดภัยสำหรับการใช้งานพร้อมกันหลาย goroutine
type SafeCounter struct {
	mu    sync.RWMutex // RWMutex: อ่านพร้อมกันได้หลายตัว แต่เขียนได้ทีละตัว
	count int          // ห้ามอ่าน/เขียนโดยตรง — ต้องผ่าน mu เสมอ
}

// Inc เพิ่มค่า count ทีละ 1 แบบ thread-safe
func (c *SafeCounter) Inc() {
	c.mu.Lock()         // Lock เต็ม: block ทั้งอ่านและเขียน
	defer c.mu.Unlock() // defer ป้องกันลืม Unlock แม้เกิด panic
	c.count++
}

// Value คืนค่า count ปัจจุบันแบบ thread-safe
// ใช้ RLock เพราะอ่านอย่างเดียว → goroutine หลายตัวอ่านพร้อมกันได้
func (c *SafeCounter) Value() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.count
}

func main() {
	fmt.Println("=== A02: Safe Counter ===")
	fmt.Println()

	counter := SafeCounter{}
	var wg sync.WaitGroup

	fmt.Println("เปิด 1000 goroutines แต่ละตัวเรียก counter.Inc() หนึ่งครั้ง...")
	fmt.Println()

	// for range N (Go 1.22+) — ไม่ต้องการตัวแปร loop
	for range 1000 {
		// wg.Go (Go 1.25+) เปิด goroutine พร้อมจัดการ WaitGroup อัตโนมัติ
		wg.Go(func() {
			counter.Inc()
		})
	}

	// บล็อกรอจนกว่า 1000 goroutines จะเสร็จทั้งหมด
	wg.Wait()

	finalCount := counter.Value()
	fmt.Printf("ผลลัพธ์สุดท้าย: %d\n\n", finalCount)

	if finalCount == 1000 {
		fmt.Println("SUCCESS: Mutex ป้องกัน race condition ได้สำเร็จ")
	} else {
		fmt.Printf("FAIL: คาดหวัง 1000 แต่ได้ %d — มี race condition!\n", finalCount)
	}

	fmt.Println()
	fmt.Println("--- ทำไม mutex ถึงสำคัญ ---")
	fmt.Println("ถ้าไม่ lock: goroutine 2 ตัวอาจอ่าน count=5 พร้อมกัน")
	fmt.Println("ทั้งคู่คำนวณ 5+1=6 และเขียน 6 กลับ → สูญเสีย 1 การนับ")
	fmt.Println("mutex บังคับให้ทีละตัว → ผลลัพธ์ถูกต้องเสมอ")
}
