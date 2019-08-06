package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	// import ("os" "strconv")
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
		
	arguments := os.Args
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[2], 64)
	
	fmt.Println(min)
	fmt.Println(max)

	var f *os.File
	f = os.Stdin
	
	defer f.Close()
	
	scanner := bufio.NewScanner(f)
	
	for scanner.Scan() {
		fmt.Println(">", scanner.Text())
	}
}