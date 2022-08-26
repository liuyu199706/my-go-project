package main

import (
	"fmt"
	"hash/crc32"
)

func CRC32(str string) uint32 {
	return crc32.ChecksumIEEE([]byte(str))
}

func main() {
	fmt.Println(CRC32("123456"))
}
