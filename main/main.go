package main

import (
	//"crypto/sha256"
	"crypto/elliptic"
	"fmt"
	"math/big"
	"strings"

	"github.com/tyr-coder/mapper/swu"
)

var c = elliptic.P256()

func main() {
	fmt.Println("SHA256 MAPPER V0.1")
	//sum := sha256.Sum256([]byte("a"))
	//tryPoint([]byte("a"))

	//fmt.Printf("%x", sum)
	execSeries(100)

}

func tryPoint(r []byte) {
	var a, b, h = swu.HashToPoint(r)
	fmt.Println("Hash Value")
	str1 := fmt.Sprintf("0x%x", h)
	fmt.Println("%x", str1)
	printPoint(a, b)
	return
}

func printPoint(x, y *big.Int) {
	str1 := fmt.Sprintf("0x%x", x)
	str2 := fmt.Sprintf("0x%x", y)
	fmt.Println("x= ")
	fmt.Println("%x", str1)
	fmt.Println("y= ")
	fmt.Println("%x", str2)
	if c.IsOnCurve(x, y) {
		fmt.Println("on curve")
	}
}

func execSeries(maxtimes int) {
	var buf strings.Builder
	var str = "nsa"

	for i := 1; i < maxtimes; i++ {
		buf.WriteString(str)
		var arg1 = []byte(buf.String())
		a, b, h := swu.HashToPoint(arg1)
		fmt.Println("Input:")
		fmt.Println(buf.String())
		fmt.Println("Hash =")
		str1 := fmt.Sprintf("0x%x", h)
		fmt.Println("%x", str1)
		printPoint(a, b)
	}

}
