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
	fmt.Println("SHA256 MAPPER V0.2")
	//sum := sha256.Sum256([]byte("a"))
	//tryPoint([]byte("a"))

	//fmt.Printf("%x", sum)
	//execSeries(100)
	twoPoints(1000000, "1", "11111")
}

func tryPoint(r []byte) {
	var a, b, h = swu.HashToPoint(r)
	fmt.Println("Hash Value")
	str1 := fmt.Sprintf("0x%x", h)
	fmt.Println("%x", str1)
	printPoint(a, b)
	return
}

func printHash(h_ [32]byte) {
	fmt.Println("Hash Value")
	str1 := fmt.Sprintf("0x%x", h_)
	fmt.Println("%x", str1)
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
func twoPoints(maxDist int, str1 string, str2 string) {

	var dist int64
	x1, y1, h1 := swu.HashToPoint([]byte(str1))
	x2, y2, h2 := swu.HashToPoint([]byte(str2))

	// Print Initial Data
	fmt.Println("Preamble")
	fmt.Println("Point 1")
	printHash(h1)
	printPoint(x1, y1)
	fmt.Println("Point 2")
	printHash(h2)
	printPoint(x2, y2)

	for i := 1; i < maxDist; i++ {
		dist += 1
		mult := big.NewInt(dist)
		xt, yt := c.ScalarMult(x1, y1, mult.Bytes())
		if (x2.Cmp(xt) == 0) && (y2.Cmp(yt) == 0) {
			fmt.Println("Multiple Found")
			fmt.Println("%v", dist)
			printPoint(xt, yt)
			break
		}

	}
	fmt.Println("NO MATCHES")
	fmt.Println("%v", dist)
}
