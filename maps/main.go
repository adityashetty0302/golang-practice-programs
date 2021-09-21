package main

import "fmt"

func main() {
	// capitals := map[string]string{
	// 	"Japan": "Tokyo",
	// }
	// capitals := map[string]string{}
	// var capitals map[string]string

	capitals := make(map[string]string)
	capitals["India"] = "New Delhi"
	// capitals["India"] = "ND"

	// delete(capitals, "India")
	// fmt.Println(capitals)

	capitals["Japan"] = "Tokyo"
	printMap(capitals)

}

func printMap(m map[string]string) {
	for country, capital := range m {
		fmt.Printf("Capital of %v is %v \n", country, capital)
	}
}
