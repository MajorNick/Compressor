package main

import (
	"fmt"
	"github.com/icza/bitio"
	"io"
	"math"
	"os"
)

const bitSize int = 12

func Check(err error, s string) {
	if err != nil {
		fmt.Println(s)
		os.Exit(1)
	}
}

func Compress(path, filename string) {
	sequences := makeSequenceMap()
	data, err := os.ReadFile(path)
	Check(err, "Error In Reading File")
	file, err := os.Create(filename)
	Check(err, "Error In Creating File")
	var wr io.Writer
	wr = file

	w := bitio.NewWriter(wr)
	err = w.WriteBool(true)
	Check(err, "error in writing")
	writeBitSize(w)
	it := 257
	if len(data) == 0 {
		return
	}
	tmp := string(data[0])
	for i := 1; i < len(data); i++ {
		if sequences[tmp+string(data[i])] > 0 {
			tmp += string(data[i])
		} else {

			bits := intToBinary(sequences[tmp])
			for i := 0; i < len(bits); i++ {
				var b bool
				if int(bits[i]-'0') == 1 {
					b = true
				} else {
					b = false
				}
				err = w.WriteBool(b)
				Check(err, "Error in Writing in File")
			}

			if it < int(math.Pow(2, float64(bitSize))) {
				sequences[tmp+string(data[i])] = it
				it++
			}
			tmp = string(data[i])
		}
	}
	w.Close()

}
func makeSequenceMap() map[string]int {
	res := make(map[string]int)
	for i := 0; i < 256; i++ {
		res[string(i)] = i
	}
	return res
}
func intToBinary(n int) string {
	var s, res string
	for n > 0 {
		s += string(int('0' + n%2))
		n /= 2
	}
	for i := 0; i < bitSize-len(s); i++ {
		res += string(int('0'))
	}
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
func writeBitSize(w *bitio.Writer) {
	str := intToBinary(bitSize)
	str = str[4:]
	for i := 0; i < 8; i++ {
		var b bool
		if int(str[i]-'0') == 1 {
			b = true
		} else {
			b = false
		}
		err := w.WriteBool(b)
		if err != nil {
			fmt.Println("error in writing file")
		}
	}

}
