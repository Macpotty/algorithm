/*
* @Author: Michael
* @Date:   2016-09-27 11:31:16
* @Last Modified by:   Michael
* @Last Modified time: 2016-10-19 21:26:57
 */

package halfset

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func run() {
	n := readFromFile()
	a := make([]int, 1024)
	sum := HalfSet(n, a)
	fmt.Printf("%d\n", sum)
	writeToFile(sum)
}

func HalfSet(n int, array []int) int {
	if array[n] > 0 {
		return array[n]
	} else {
		array[n] = 1
		for i := 1; i <= n/2; i++ {
			array[n] += HalfSet(i, array)
		}
		return array[n]
	}
}

func readFromFile() int {
	inputFile := "../testfiles/input_halfSet.txt"
	fin, err := os.Open(inputFile)
	defer fin.Close()
	if err != nil {
		panic(err)
	} else {
		buf := make([]byte, 64)
		list, _ := fin.Read(buf)
		if err != nil {
			panic(err)
		} else {
			strings.Split(string(buf[:list]), " ")
		}
	}
}

func writeToFile(num int) {
	outputFile := "../testfiles/output_halfSet.txt"
	fout, err := os.Create(outputFile)
	defer fout.Close()
	if err != nil {
		panic(err)
	} else {
		fout.WriteString(strconv.Itoa(num))
	}
}