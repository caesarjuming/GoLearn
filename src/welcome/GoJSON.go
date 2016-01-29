package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Response1 struct {
	Page   int
	Fruits []string
}

type Response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("str")
	fmt.Println(string(strB))

	listB := []string{"aaa", "bbb", "ccc"}
	listBA, _ := json.Marshal(listB)
	fmt.Println(string(listBA))

	mapB := map[string]int{"aa": 1, "bb": 2, "cc": 3}
	mapBA, _ := json.Marshal(mapB)
	fmt.Println(string(mapBA))

	structRes1B := &Response1{
		Page:   100,
		Fruits: []string{"QQ", "WW", "EE"},
	}
	structRes1BA, _ := json.Marshal(structRes1B)
	fmt.Println(string(structRes1BA))

	structRes2B := &Response2{
		Page:   200,
		Fruits: []string{"ccc", "ddd"},
	}
	structRes2BA, _ := json.Marshal(structRes2B)
	fmt.Println(string(structRes2BA))

	bytes := []byte(`{"a":6.13,"strs":["aa","bb","cc"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(bytes, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["a"].(float64)
	fmt.Println(num)

	strss := dat["strs"].([]interface{})
	fmt.Println(strss)
	fmt.Println(strss[0].(string))

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	os.Exit(1)

}
