package main

import (
	"fmt"
	"os"
)

func main() {
	// open the files
	sda, _ := os.Open("./sda")
	sdb, _ := os.Open("./sdb")
	sdc, _ := os.Open("./sdc")
	sdd, _ := os.Open("./sdd")

	files := []*os.File{sda, sdb, sdc, sdd}

	// loop to take each byte in our order and concat together pair bytes
	// 0110 << 1    1100 & 0111 => 0100. 0100 << 1 => 1000
	//	sdaBits := 0b1000 // 0x8
	//	sdbBits := 0b0100 // 0x4
	//	sdcBits := 0b0010 // 0x2
	//	sddBits := 0b0001 // 0x1

	// drives := []int{sdaBits, sdbBits, sdcBits, sddBits}

	// currentLoopBits := 0b0110 // 0110 -> 0100 or 0010
	var result string
	var concat string
	for _ = range 35 {
		var fdrive *os.File
		var sdrive *os.File
		var pdrive *os.File
		var qdrive *os.File

		for j := 3; j >= 0; j-- {
			//var oR int
			var f, s, p, q int

			q = j % 4
			p = (j + 1) % 4
			f = (q + 2) % 4
			s = (p + 2) % 4

			fdrive = files[f]
			sdrive = files[s]

			pdrive = files[p]
			qdrive = files[q]

			fBuf := make([]byte, 16)
			_, _ = fdrive.Read(fBuf)

			sBuf := make([]byte, 16)
			_, _ = sdrive.Read(sBuf)

			_, _ = pdrive.Read(make([]byte, 16))
			_, _ = qdrive.Read(make([]byte, 16))

			//	fmt.Printf("first cat: %s\n second cat: %s\n", fBuf, sBuf)
			concat = string(fBuf) + string(sBuf)
			result += concat

			concat = ""
		}
		// fmt.Printf("fdrive = %s\n", )

	}

	fmt.Printf(result)
}
