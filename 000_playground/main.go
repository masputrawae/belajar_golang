package main

import (
	"fmt"
)

func main() {
	var buah [3]string
	buah[0] = "Apel"
	buah[1] = "Jeruk"
	buah[2] = "Mangga"

    siswa := [5]string{"Andi", "Jono", "Budi", "Upin", "Ipin"}

	for i, b := range buah {
		fmt.Printf("[%d]: %s\n", i+1, b)
	}

    fmt.Println("=========================")

    for i, s := range siswa {
        fmt.Printf("[%d]: %s\n", i+1, s)
    }
}
