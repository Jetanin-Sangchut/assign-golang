# assign-golang

Go concurrency, interfaces, and HTTP exercises -- worker pool, mutex, shapes, two-sum, JSON API.

แบบฝึกหัด Go ครอบคลุม concurrency, interface, data structure, และ HTTP server
พร้อม comment ภาษาไทยในโค้ดทุกไฟล์ เหมาะสำหรับผู้เริ่มต้นเรียนรู้ Go


## สารบัญ / Table of Contents

- [Prerequisites](#prerequisites)
- [Repository Structure](#repository-structure)
- [Exercise Overview](#exercise-overview)
- [How to Run](#how-to-run)
- [Modern Go Features Used](#modern-go-features-used)
- [Notes for Windows Users](#notes-for-windows-users)
- [License](#license)


## Prerequisites

| Requirement | Version |
|-------------|---------|
| Go          | 1.26+   |

ตรวจสอบเวอร์ชัน Go ที่ติดตั้งอยู่ด้วยคำสั่ง:

```bash
go version
```

โปรเจกต์นี้ใช้ฟีเจอร์ของ Go 1.22 ถึง 1.26 หากใช้เวอร์ชันต่ำกว่า 1.26 จะ compile ไม่ผ่าน


## Repository Structure

```
assign-golang/
├── A01_worker_pool/       # Concurrency: goroutines + channels + WaitGroup
│   ├── go.mod
│   └── main.go
├── A02_safe_counter/      # Thread-safety: sync.Mutex
│   ├── go.mod
│   └── main.go
├── A03_shape_interface/   # Polymorphism: interface + implicit satisfaction
│   ├── go.mod
│   └── main.go
├── A04_two_sum/           # Algorithm: map-based O(n) solution
│   ├── go.mod
│   └── main.go
├── A05_http_json_api/     # Web: net/http + JSON encode/decode
│   ├── go.mod
│   └── main.go
├── .gitignore
├── LICENSE
└── README.md
```

แต่ละ exercise เป็น Go module อิสระ มี `go.mod` แยกของตัวเอง สามารถ run ได้แยกกัน


## Exercise Overview

| No. | Folder | Title (TH) | Title (EN) | Concepts |
|-----|--------|-------------|------------|----------|
| A01 | `A01_worker_pool` | Concurrency พื้นฐาน | Worker Pool | goroutines, buffered channels, sync.WaitGroup, `wg.Go()` |
| A02 | `A02_safe_counter` | Thread-Safety ด้วย sync.Mutex | Safe Counter | sync.Mutex, Lock/Unlock, race condition prevention |
| A03 | `A03_shape_interface` | Polymorphism ผ่าน Interface | Shape Interface | interface, implicit satisfaction, polymorphism |
| A04 | `A04_two_sum` | Logic & Map (O(n)) | Two Sum | map[int]int, complement pattern, O(n) time complexity |
| A05 | `A05_http_json_api` | Web Server ด้วย net/http | HTTP JSON API | net/http, encoding/json, HTTP status codes, method routing |

### A01: Worker Pool -- Concurrency พื้นฐาน

เขียน `RunWorkers(numWorkers, numJobs int)` สร้าง worker จำนวนคงที่รับงานจาก channel พร้อมกัน
แต่ละ worker พิมพ์ "Worker [ID] processing job [Job ID]" จำลองงานด้วย `time.Sleep`
main รอจนทุก worker เสร็จด้วย `sync.WaitGroup`

### A02: Safe Counter -- Thread-Safety ด้วย sync.Mutex

สร้าง `SafeCounter` struct มี `Inc()` เพิ่มค่า 1 และ `Value()` คืนค่าปัจจุบัน
รองรับ 1,000 goroutines เรียก `Inc()` พร้อมกัน ค่าสุดท้ายต้องเท่ากับ 1,000 พอดี ห้ามเกิด race condition

### A03: Shape Interface -- Polymorphism ผ่าน Interface

ประกาศ `Shape` interface มี `Area() float64`
สร้าง `Rectangle` (Width, Height) และ `Circle` (Radius) implement `Area()`
เขียน `PrintArea(s Shape)` รับ shape ชนิดใดก็ได้

### A04: Two Sum -- Logic & Map (O(n))

กำหนด `nums := []int{2, 7, 11, 15}`, `target := 9`
หา 2 index ที่ค่ารวมกันเท่ากับ target โดยใช้ complement pattern กับ `map[int]int` ให้ได้ O(n)

### A05: HTTP JSON API -- Web Server ด้วย net/http

สร้าง server ที่ port 8080 endpoint `POST /hello`
รับ JSON `{"name":"Somchai"}` ตอบ `{"message":"Hello, Somchai!"}`
method ไม่ใช่ POST ตอบ 405, JSON ผิดรูปแบบ ตอบ 400, name ว่าง ตอบ 400


## How to Run

แต่ละ exercise อยู่ใน folder แยก เข้า folder แล้วรันด้วย `go run main.go`

```bash
# A01: Worker Pool
cd A01_worker_pool
go run main.go

# A02: Safe Counter
cd A02_safe_counter
go run main.go

# A03: Shape Interface
cd A03_shape_interface
go run main.go

# A04: Two Sum
cd A04_two_sum
go run main.go

# A05: HTTP JSON API
cd A05_http_json_api
go run main.go
```

สำหรับ A05 หลัง server เริ่มทำงานแล้ว ทดสอบด้วย curl ใน terminal อีกหน้าต่าง:

```bash
# Success case (200)
curl -X POST -H "Content-Type: application/json" \
     -d '{"name":"Somchai"}' http://localhost:8080/hello

# Empty name (400)
curl -X POST -H "Content-Type: application/json" \
     -d '{"name":""}' http://localhost:8080/hello

# Invalid JSON (400)
curl -X POST -H "Content-Type: application/json" \
     -d 'not json' http://localhost:8080/hello

# Wrong method (405)
curl http://localhost:8080/hello
```

ตรวจสอบ race condition ใน A02 ด้วย Go race detector:

```bash
cd A02_safe_counter
go run -race main.go
```


## Modern Go Features Used

โปรเจกต์นี้ใช้ฟีเจอร์สมัยใหม่ของ Go ตั้งแต่เวอร์ชัน 1.22 ถึง 1.26

| Feature | Since | Used In | Description |
|---------|-------|---------|-------------|
| `for i := range n` | Go 1.22 | A01, A02 | Integer range loop แทน `for i := 0; i < n; i++` |
| `"POST /hello"` method routing | Go 1.22 | A05 | ระบุ HTTP method ใน ServeMux pattern ได้โดยตรง |
| `wg.Go(func() { ... })` | Go 1.25 | A01, A02 | เปิด goroutine พร้อมจัดการ WaitGroup อัตโนมัติ แทน `wg.Add(1)` + `defer wg.Done()` |
| `errors.AsType[T](err)` | Go 1.26 | A05 | Generic version ของ `errors.As` ไม่ต้องประกาศตัวแปรชั่วคราว |


## Notes for Windows Users

### GOTMPDIR Configuration

หากใช้ Windows ที่มี Application Control policy (เช่น Windows ในองค์กร) อาจพบ error เมื่อรัน `go run` เนื่องจาก Go สร้าง temp executable ใน `%TMP%` ซึ่งอาจถูก block

แก้ไขด้วยการตั้ง `GOTMPDIR` ไปยัง folder ที่อนุญาตให้รันโปรแกรม:

```powershell
# PowerShell — ตั้งค่าชั่วคราวในเซสชันนี้
$env:GOTMPDIR = "C:\GoTmp"
New-Item -ItemType Directory -Force -Path $env:GOTMPDIR

# หรือตั้งค่าถาวร
[System.Environment]::SetEnvironmentVariable("GOTMPDIR", "C:\GoTmp", "User")
```

```bash
# Git Bash / WSL
export GOTMPDIR="C:/GoTmp"
mkdir -p "$GOTMPDIR"
```


## License

MIT License -- see [LICENSE](LICENSE) for details.

Copyright (c) 2026 @Jay-Jetanin
