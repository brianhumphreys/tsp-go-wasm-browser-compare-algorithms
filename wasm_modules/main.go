package main

import (
	"encoding/json"
	"fmt"
	"math"
	"syscall/js"
)

func prettyJson(input string) (string, error) {
	var raw interface{}
	if err := json.Unmarshal([]byte(input), &raw); err != nil {
		return "", err
	}
	pretty, err := json.MarshalIndent(raw, "", "  ")
	if err != nil {
		return "", err
	}
	return string(pretty), nil
}

func jsonWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputJSON := args[0].Get("json").String()
		fmt.Printf("input %s\n", inputJSON)
		pretty, err := prettyJson(inputJSON)
		if err != nil {
			fmt.Printf("unable to convert to json %s\n", err)
			return err.Error()
		}
		return pretty
	})
	return jsonFunc
}

func cost(vertices [][]float64) float64 {
	total := 0.0
	for i := 1; i < len(vertices); i++ {
		total += distance(vertices[i-1], vertices[i])
	}
	total += distance(vertices[len(vertices)-1], vertices[0])
	return total
}

// func jsValueToInt2DArray(vals js.Value) [][]int {
// 	array := [][]int{}
// 	for i := 0; ; i++ {

// 		defer func() {
// 			if err := recover(); err != nil {
// 				return
// 			}
// 		}()
// 		x := vals.Index(i).Index(0).Int()
// 		fmt.Println("errrr: ", err)
// 		y := vals.Index(i).Index(1).Int()
// 		array = append(array, []int{x, y})
// 		fmt.Println(x)
// 		fmt.Println(y)
// 		// fmt.Printf("%T\n", vari)
// 		// fmt.Println(args[0].Index(i))
// 	}
// 	return array
// }

type Vertex struct {
	array int
}

func costWrapper() js.Func {
	costFunction := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed.  Expecting 2."
		}

		fmt.Println("poop")
		var collection []interface{}

		length := args[1].Int()
		// array := [][]int{}

		for i := 0; i < length; i++ {
			x := args[0].Index(i).Index(0).Int()
			y := args[0].Index(i).Index(1).Int()
			collection = append(collection, []int{x, y})
		}

		fmt.Println(collection)

		// v1 := Vertex{x: 1, y: 2}
		return map[string]interface{}{
			"result": "restsss",
		}
	})

	return costFunction
}

func distance(vertex1 []float64, vertex2 []float64) float64 {

	return math.Pow(math.Pow(vertex1[0]-vertex2[0], 2)+math.Pow(vertex1[1]-vertex2[1], 2), 0.5)
}

func distanceWrapper() js.Func {
	distanceFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed.  Expecting 2."
		}

		vertex1 := []float64{float64(args[0].Index(0).Int()), float64(args[0].Index(0).Int())}
		vertex2 := []float64{float64(args[1].Index(0).Int()), float64(args[1].Index(1).Int())}

		return distance(vertex1, vertex2)
	})
	return distanceFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("cost", costWrapper())
	js.Global().Set("distance", distanceWrapper())
	js.Global().Set("formatJSON", jsonWrapper())

	<-make(chan bool)
}
