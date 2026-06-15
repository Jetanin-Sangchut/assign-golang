// =============================================================
// A04: Two Sum — Logic & Map (O(n) solution)
// =============================================================
// หาสองตำแหน่งใน slice ที่ค่ารวมกันได้เท่ากับ target
// วิธีที่ดี: ใช้ map เก็บค่าที่ผ่านมา → ค้นหาได้ O(1) → รวมทั้งหมด O(n)
// ห้ามใช้ nested loop (O(n²)) ซึ่งช้ากว่า 10,000 เท่าสำหรับข้อมูลใหญ่
// =============================================================

package main

import "fmt"

// TwoSum หาสองตำแหน่งใน nums ที่บวกกันได้เท่ากับ target
//
// กลยุทธ์ "complement pattern":
//   สำหรับแต่ละ num ที่เจอ → complement = target - num
//   ถ้า complement อยู่ใน map แล้ว → เจอคู่คำตอบแล้ว
//   ถ้ายังไม่มี → เก็บ num ลง map รอให้คนอื่นมาจับคู่
//
// คืนค่า (-1, -1) ถ้าไม่มีคู่ที่ตรงเงื่อนไข
func TwoSum(nums []int, target int) (int, int) {
	// seen เก็บ value → index ของตัวเลขที่ผ่านมาแล้ว
	seen := make(map[int]int)

	for i, num := range nums {
		complement := target - num // ค่าที่เราต้องการจับคู่

		// ค้นหา complement ใน map — O(1) average
		if j, ok := seen[complement]; ok {
			return j, i // nums[j] + nums[i] == target
		}

		// ยังไม่เจอคู่ → เก็บ num ไว้ให้รอบถัดไปใช้
		seen[num] = i
	}

	return -1, -1
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	fmt.Printf("nums: %v, target: %d\n", nums, target)

	i, j := TwoSum(nums, target)

	if i == -1 {
		fmt.Println("คำตอบ: ไม่มีคู่ที่ตรงเงื่อนไข")
		return
	}

	// คาดหวัง: indices [0, 1] เพราะ nums[0]+nums[1] = 2+7 = 9
	fmt.Printf("คำตอบ: ตำแหน่ง [%d, %d] → %d + %d = %d\n",
		i, j, nums[i], nums[j], nums[i]+nums[j])
}
