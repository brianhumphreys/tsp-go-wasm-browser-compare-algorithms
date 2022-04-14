package main

import (
	"fmt"
	"math"
	"syscall/js"
)

// func prettyJson(input string) (string, error) {
// 	var raw interface{}
// 	if err := json.Unmarshal([]byte(input), &raw); err != nil {
// 		return "", err
// 	}
// 	pretty, err := json.MarshalIndent(raw, "", "  ")
// 	if err != nil {
// 		return "", err
// 	}
// 	return string(pretty), nil
// }

// func jsonWrapper() js.Func {
// 	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		if len(args) != 1 {
// 			return "Invalid no of arguments passed"
// 		}
// 		inputJSON := args[0].Get("json").String()
// 		fmt.Printf("input %s\n", inputJSON)
// 		pretty, err := prettyJson(inputJSON)
// 		if err != nil {
// 			fmt.Printf("unable to convert to json %s\n", err)
// 			return err.Error()
// 		}
// 		return pretty
// 	})
// 	return jsonFunc
// }

func cost(vertices []Vertex) float64 {
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
	x float64
	y float64
}

func jsValueToVertexArray(args []js.Value) []Vertex {
	length := args[0].Get("length").Int()

	resultArray := make([]Vertex, length)

	for i := 0; i < length; i++ {
		index := fmt.Sprintf("%d", i)
		x := float64(args[0].Get(index).Get("x").Int())
		y := float64(args[0].Get(index).Get("y").Int())
		resultArray[i] = Vertex{x: x, y: y}
	}

	return resultArray
}

func vertexArrayToInterfaceMap(vertices []Vertex) map[string]interface{} {
	resultArray := map[string]interface{}{
		"length": len(vertices),
	}
	for i := 0; i < len(vertices); i++ {
		resultArray[fmt.Sprintf("%d", i)] = map[string]interface{}{"x": vertices[i].x, "y": vertices[i].y}
	}
	return resultArray
}

func costWrapper() js.Func {
	costFunction := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 1 {
			return "Invalid number of arguments passed.  Expecting 2."
		}

		vertexArray := jsValueToVertexArray(args)

		totalCost := cost(vertexArray)

		return totalCost
	})

	return costFunction
}

func distance(vertex1 Vertex, vertex2 Vertex) float64 {

	return math.Pow(math.Pow(vertex1.x-vertex2.x, 2)+math.Pow(vertex1.y-vertex2.y, 2), 0.5)
}

func distanceWrapper() js.Func {
	distanceFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(args) != 2 {
			return "Invalid number of arguments passed.  Expecting 2."
		}

		vertex1 := Vertex{x: float64(args[0].Index(0).Int()), y: float64(args[0].Index(1).Int())}
		vertex2 := Vertex{x: float64(args[1].Index(0).Int()), y: float64(args[1].Index(1).Int())}

		return distance(vertex1, vertex2)
	})
	return distanceFunc
}

func main() {
	fmt.Println("Go Web Assembly")
	js.Global().Set("cost", costWrapper())
	js.Global().Set("distance", distanceWrapper())

	<-make(chan bool)
}
