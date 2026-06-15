// =============================================================
// A05: HTTP JSON API — Web Server ด้วย net/http
// =============================================================
// สร้าง HTTP server รับ POST request พร้อม JSON body
// แล้วตอบกลับด้วย JSON greeting response บน port 8080
//
// ทดสอบด้วย:
//   curl -X POST -H "Content-Type: application/json" \
//        -d '{"name":"Somchai"}' http://localhost:8080/hello
// =============================================================

package main

import (
	"encoding/json" // encode/decode JSON ↔ struct
	"fmt"
	"log"
	"net/http"
)

// =============================================================
// Structs สำหรับ Request / Response
// =============================================================
// struct tag `json:"name"` บอก encoding/json ว่า field นี้ตรงกับ JSON key ไหน

// HelloRequest คือ JSON body ที่ client ต้องส่งมา
type HelloRequest struct {
	Name string `json:"name"`
}

// HelloResponse คือ JSON ที่ส่งกลับเมื่อสำเร็จ
type HelloResponse struct {
	Message string `json:"message"`
}

// ErrorResponse คือ JSON ที่ส่งกลับเมื่อเกิดข้อผิดพลาด
type ErrorResponse struct {
	Error string `json:"error"`
}

// =============================================================
// writeJSON — helper ลด code ซ้ำ
// =============================================================
// ทุก response ในระบบนี้เป็น JSON → รวมขั้นตอนไว้ที่เดียว
func writeJSON(w http.ResponseWriter, status int, v any) {
	// Content-Type ต้องตั้งก่อน WriteHeader — ถ้าตั้งหลังจะไม่มีผล
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	// Encode เขียนตรงลง ResponseWriter ไม่ต้องสร้าง buffer กลาง
	if err := json.NewEncoder(w).Encode(v); err != nil {
		fmt.Printf("writeJSON error: %v\n", err) // header ส่งไปแล้ว ทำได้แค่ log
	}
}

// =============================================================
// helloHandler — handler หลัก
// =============================================================
// ลงทะเบียนด้วย "POST /hello" (Go 1.22+) ทำให้ ServeMux
// กรอง method ที่ไม่ใช่ POST ออกให้อัตโนมัติ (ได้ 405)
func helloHandler(w http.ResponseWriter, r *http.Request) {

	// ขั้นที่ 1 — ตรวจ method (safety fallback)
	// ServeMux จัดการให้แล้ว แต่ตรวจซ้ำเผื่อมีการเรียก handler โดยตรง
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, ErrorResponse{Error: "method not allowed"})
		return
	}

	// ขั้นที่ 2 — Decode JSON body
	// จำกัดขนาด body ไม่เกิน 1 MB ป้องกัน client ส่ง payload ใหญ่เกินไป
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)

	var req HelloRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		// JSON ผิด format, body ว่าง, หรือใหญ่เกินไป → 400
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "invalid request body"})
		return
	}

	// ขั้นที่ 3 — ตรวจ business rule
	// JSON ถูก format แต่ name ว่าง → ไม่ยอมรับ
	if req.Name == "" {
		writeJSON(w, http.StatusBadRequest, ErrorResponse{Error: "name is required"})
		return
	}

	// ขั้นที่ 4 — ส่ง response สำเร็จ
	resp := HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", req.Name),
	}
	writeJSON(w, http.StatusOK, resp)
}

// =============================================================
// main — ลงทะเบียน route และเริ่ม server
// =============================================================
func main() {
	// "POST /hello" (Go 1.22+) — ระบุ method + path ใน pattern เดียว
	// method อื่น → 405 อัตโนมัติ, path อื่น → 404 อัตโนมัติ
	http.HandleFunc("POST /hello", helloHandler)

	addr := ":8080"
	fmt.Printf("Server เริ่มต้นที่ %s\n", addr)
	fmt.Println(`POST /hello รับ {"name":"Somchai"} → {"message":"Hello, Somchai!"}`)

	// ListenAndServe บล็อกจนกว่า server จะหยุด
	// ถ้าเกิดข้อผิดพลาด (เช่น port ชนกัน) → log.Fatalf จบโปรแกรมด้วย exit code 1
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server error: %v\n", err)
	}
}
