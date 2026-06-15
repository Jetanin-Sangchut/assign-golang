// =============================================================
// A01: Worker Pool — Concurrency พื้นฐาน
// =============================================================
// แทนที่จะเปิด goroutine ใหม่ทุกงาน เราใช้ worker จำนวนคงที่
// แต่ละ worker รับงานจาก channel ร่วมกัน แล้ว loop ต่อไปเรื่อยๆ
// พอ channel ถูก close() และหมดแล้ว worker จะออกเองอัตโนมัติ
// ใช้ sync.WaitGroup เพื่อรอจนกว่า worker ทุกตัวจะเสร็จ
// =============================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

// RunWorkers เปิด worker จำนวน numWorkers ตัว เพื่อประมวลผลงาน numJobs ชิ้น
func RunWorkers(numWorkers int, numJobs int) {
	fmt.Printf("เริ่ม worker pool: %d workers, %d jobs\n\n", numWorkers, numJobs)

	// 1. สร้าง channel สำหรับส่งงาน (buffered = numJobs)
	// buffered ทำให้ sender ใส่งานได้ทันทีโดยไม่ block
	// ถ้าไม่ buffer → sender block → worker ยังไม่ได้เริ่ม → deadlock
	jobs := make(chan int, numJobs)

	// 2. สร้าง channel สำหรับรับผลลัพธ์ (buffered = numJobs)
	// worker ส่งผลได้ทันทีโดยไม่ต้องรอ main มารับ
	results := make(chan int, numJobs)

	// 3. เตรียม WaitGroup ติดตาม worker ที่ยังทำงานอยู่
	var wg sync.WaitGroup

	// 4. เปิด worker (Go 1.22+: for i := range n แทน for i := 0; i < n; i++)
	for i := range numWorkers {
		workerID := i + 1 // นับจาก 1 เพื่อ output ที่อ่านง่าย

		// wg.Go (Go 1.25+) แทนรูปแบบเก่า wg.Add(1) + defer wg.Done()
		wg.Go(func() {
			// range channel: รับค่าจนกว่า channel จะถูก close() และว่างเปล่า
			for jobID := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, jobID)
				time.Sleep(1 * time.Second) // จำลองการทำงานที่ใช้เวลา
				results <- jobID            // ส่งผลลัพธ์กลับ
			}
			fmt.Printf("Worker %d เสร็จสิ้น\n", workerID)
		})
	}

	// 5. ใส่งานเข้า channel (buffered ทำให้ทำได้ทันทีทั้งหมด)
	fmt.Println("กำลังส่งงานเข้า queue...")
	for j := range numJobs {
		jobID := j + 1
		jobs <- jobID
		fmt.Printf("  ส่ง job %d เข้าคิว\n", jobID)
	}

	// 6. close(jobs) บอก worker ว่า "ไม่มีงานใหม่แล้ว"
	// ถ้าไม่ close → worker รอตลอด → deadlock
	close(jobs)
	fmt.Println("\nใส่งานครบแล้ว รอ worker ทั้งหมด...")

	// 7. รอ worker ทุกตัวเสร็จก่อน แล้วค่อยปิด results
	// ปลอดภัยเพราะ results buffered = numJobs → worker ไม่มีทาง block ตอน send
	// ถ้าเปลี่ยน buffer ให้เล็กลง ต้องใช้ pattern: go func(){ wg.Wait(); close(results) }()
	wg.Wait()
	close(results)

	// 8. พิมพ์ผลลัพธ์ (worker เสร็จหมดแล้ว output ไม่ปนกัน)
	fmt.Println("\nงานที่เสร็จแล้ว:")
	for result := range results {
		fmt.Printf("  Job %d เสร็จ\n", result)
	}
}

func main() {
	RunWorkers(3, 5)
}
