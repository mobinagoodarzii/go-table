package main

import (
	"fmt"
	"github.com/goodarzi/models"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Let's make a table!!")

	fmt.Println("Enter the number of columns:")
	var X int
	fmt.Scanln(&X)

	fmt.Println("Enter the number of rows:")
	var Y int
	fmt.Scanln(&Y)

	time.Sleep(time.Second * 1)
	models.CallClear()

	rand.Seed(time.Now().Unix())
	tableIntValues := rand.Perm(X * Y)
	strNum := fmt.Sprintf("%d", X*Y)
	sqrWidth := len(strNum) + 2
	tableStrValues := make([]string, len(tableIntValues))
	for i, v := range tableIntValues {
		strV := fmt.Sprintf("%d", v)
		diffStr := len(strNum) - len(strV)
		for j := 0; j < diffStr; j++ {
			strV = strV + " "
		}
		tableStrValues[i] = strV
	}
	fmt.Print("┌")
	for i := 0; i < X-1; i++ {
		for j := 0; j < sqrWidth; j++ {
			fmt.Print("-")
		}
		fmt.Print("┬")
	}
	for j := 0; j < sqrWidth; j++ {
		fmt.Print("-")
	}
	fmt.Print("┐\n")
	for i := 0; i < Y; i++ {
		for j := 0; j < X; j++ {
			fmt.Print("|")
			fmt.Print(" ")
			fmt.Print(tableStrValues[i*X+j])
			fmt.Print(" ")
		}
		fmt.Print("|\n")
		fmt.Print("└")
		for j := 0; j < X-1; j++ {
			for k := 0; k < sqrWidth; k++ {
				fmt.Print("-")
			}
			fmt.Print("┴")
		}
		for k := 0; k < sqrWidth; k++ {
			fmt.Print("-")
		}
		fmt.Print("┘")
		fmt.Println()
	}
}
