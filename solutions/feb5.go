package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("tiddies all day")

	hda, err := os.Open("/home/trshpuppy/stderr/stderr-challenges/sda")
	if err != nil {
		log.Fatalf("Error opening hda\n")
	}
	defer hda.Close()

	hdb, err := os.Open("/home/trshpuppy/stderr/stderr-challenges/sdb")
	if err != nil {
		log.Fatalf("Error opening hdb\n")
	}
	defer hdb.Close()

	hdc, err := os.Open("/home/trshpuppy/stderr/stderr-challenges/sdc")
	if err != nil {
		log.Fatalf("Error opening hdc\n")
	}
	defer hdc.Close()

	dAta := [][]byte{}
	cFlag := false
	//reconstructed := [][]byte{}
	for {
		aBuf := make([]byte, 1)
		_, err := hda.Read(aBuf)

		bBuf := make([]byte, 1)
		_, err = hdb.Read(bBuf)

		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		var cByte byte
		if !cFlag {
			cBuf := make([]byte, 1)
			_, errC := hdc.Read(cBuf)

			if errC != nil {
				if errC.Error() == "EOF" {
					// Reconstruct C instead:
					cByte = aBuf[0] ^ bBuf[0]
					cFlag = true
				}
			} else {
				cByte = cBuf[0]
			}
		} else {
			cByte = aBuf[0] ^ bBuf[0]
		}

		// Make sure the XOR sum is 0:
		checkSum := aBuf[0] ^ bBuf[0] ^ cByte
		if checkSum != 0 {
			fmt.Printf("ERRORRRRRRRRRRRR!!!!!!!!! \n")
			break
		}

		// Reconstruct this stripe and append to reconstructed:
		var newStripe []byte = []byte{aBuf[0], bBuf[0], cByte}
		dAta = append(dAta, newStripe)

		//		fmt.Printf("Newstripe: %v\n", newStripe)
	}

	// [a, b, c] [0, 1, 2]
	reconstructed := []string{}
	for i, value := range dAta {
		parity := (i + 2) % 3

		d1 := (parity + 1) % 3
		d2 := (parity + 2) % 3

		reconstructed = append(reconstructed, string(value[d1]), string(value[d2]))

		// for j, v := range value {
		// 	if j == parity {
		// 		continue
		// 	} else {
		// 		s := string(v)
		// 		reconstucted = append(reconstucted, s)
		// 	}
		// }
	}
	fmt.Printf("Reconstructed = %v\n", reconstructed)
}
