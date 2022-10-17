package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/icza/bitio"
)
var k int  = 0

func Decompress(path, filename string) {
	red, err := os.Open(path)
	Check(err, "ERROR IN OPEN")
	readBit := bitio.NewReader(red)
	seqMap := makeMap()
	f, err := os.Create(filename)
	Check(err, "ERROR IN CREATING FILE")
	defer f.Close()
	it := 257

	size := getLength(readBit)
	
	last := ""
	cur := ""
	
	for i := 0; i < size; i++ {
		
		bit, err := readBit.ReadBool()
		k++
		Check(err, "ERROR IN READING")
		cur += returnBit(bit)
	}

	n := binToInt(cur)
	data := string(n)
	seqMap[n] = data
	last = data
	writenData := ""
	writenData += data
	for err == nil {

		byte := ""
		var t bool
		t = false
		for i := 0; i < size; i++ {
			bit, err := readBit.ReadBool()
			k++
			if err != nil {
				t = true
			}
			byte += returnBit(bit)
		}
		if t {
			break
		}

		tmp := binToInt(byte)
		
		byte = ""
		if seqMap[tmp] == "" {
			data = seqMap[n]
			data += last
		} else {
			data = seqMap[tmp]
		}

		writenData += data
		last = ""

		last += string(data[0])

		seqMap[it] = seqMap[n] + last
		it++
		data = ""
		n = tmp

	}
	f.Write([]byte(writenData))
	

}

// helper function
func binToInt(bin string) int {

	res, err := strconv.ParseInt(bin, 2, 64)
	Check(err, "error in converting string to binary")
	fmt.Print()
	return int(res)
}

func getLength(rd *bitio.Reader) int {
	var bin string
	for i := 0; i < 8; i++ {
		bit, err := rd.ReadBool()
		Check(err, "ERROR IN READING")
		k++
		bin += returnBit(bit)
	}
	res := binToInt(bin)
	
	return res
}
func makeMap() map[int]string {
	res := make(map[int]string)
	for i := 0; i < 256; i++ {
		res[i] = string(i)
	}
	return res
}
func returnBit(b bool) string {
	if b {
		return "1"
	} else {
		return "0"
	}
}
