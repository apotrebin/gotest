package main

// MAP
import "fmt"

func main() {

	// ---------------  Case # 1
	//var colors map[string]string

	// ---------------  Case # 2
	colors := map[string]string{
		"red":   "#ff00000",
		"green": "#ff333333",
		"white": "dsasgdfgdfg",
	}

	// ---------------  Case # 3
	//colors := make(map[int]string)

	//colors[111] = "#gg545454"
	//colors[222] = "#gg545454"
	//
	delete(colors, "sdfsdf")

	//fmt.Println(colors)

	printMap(colors)
}

func printMap(m map[string]string) {
	for k, v := range m {
		fmt.Println("Hex is:", v, "and color is", k)
	}
}
