package jsonop

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Ops int

const (
	OpAdd Ops = iota + 1
	OpDelete
	OpPrint
)

func getJsonUnmarsh(jsb []byte) map[string]interface{} {
	//fmt.Println("jsb:", string(jsb))
	if jsb == nil {
		return nil
	}

	res := make(map[string]interface{})
	if err := json.Unmarshal(jsb, &res); err != nil {
		fmt.Println("failed to Unmarshal")
	}
	return res
}

func compareTwo(one, two map[string]interface{}) bool {
	//fmt.Println(one, two)
	if one == nil || two == nil {
		return false
	}
	return reflect.DeepEqual(one, two)
}

func IsJsonEqual(jsA, jsB []byte) bool {
	return compareTwo(getJsonUnmarsh(jsA), getJsonUnmarsh(jsB))
}

func actionByOp(objA, objB interface{}, op int) interface{} {
	if reflect.TypeOf(objA) != reflect.TypeOf(objB) {
		//type different abort op, return objA
		return objA
	}

	switch v := objA.(type) {
	case string:
		fmt.Println(v)
	case int32, int64:
		fmt.Println(v)
	//case SomeCustomType:
	//	fmt.Println(v)
	default:
		fmt.Println("unknown")
	}
	//return interface{}
	return objA
}

func jsonOps(obA, obB map[string]interface{}, op Ops) interface{} {

	for key, val := range objA {
		if valB, ok := obB[val]; !ok {
			//A element not in B map, skip it.
			continue
		}
		if val.(type) != valB.(type) {
			//type different, skip it.
			continue
		}

		switch v := val.(type) {

		case []interface{}:
			listObj := val.([]interface{})
			fmt.Printf("key:%s val:[", key)
			for _, strV := range listObj {
				fmt.Printf("%v,", strV)
			}
			fmt.Printf("]\n")
			break
		case bool:
			fmt.Println("type:", v, " key:", key, " val:", obj)
			break
		case string:
			fmt.Println("type:", v, " key:", key, " val:", obj)
			break
		case int32, int64, float64:
			fmt.Println("type:", v, " key:", key, " val:", obj)
			break
		case map[string]interface{}:
			mapObj := obj.(map[string]interface{})
			fmt.Println("key:", key, "{")
			for keyB, valB := range mapObj {
				parseInterface(keyB, valB)
			}
			fmt.Println("}")
			break
		default:
			fmt.Println("unknown key:", key, " type:", reflect.TypeOf(obj))
		}
	}
}

//func JsonDiff(jsA, jsB []byte) (bool,
func JsonAdd(jsA, jsB []byte) []byte {
	//TODO

	return nil
}

func parseInterface(key string, obj interface{}) {
	//fmt.Println(" obj:", obj)
	switch v := obj.(type) {

	case []interface{}:
		listObj := obj.([]interface{})
		fmt.Printf("key:%s val:[", key)
		for _, strV := range listObj {
			fmt.Printf("%v,", strV)
		}
		fmt.Printf("]\n")

		break
	case bool:
		fmt.Println("type:", v, " key:", key, " val:", obj)
		break
	case string:
		fmt.Println("type:", v, " key:", key, " val:", obj)
		break
	case int32, int64, float64:
		fmt.Println("type:", v, " key:", key, " val:", obj)
		break
	case map[string]interface{}:
		mapObj := obj.(map[string]interface{})
		fmt.Println("key:", key, "{")
		for keyB, valB := range mapObj {
			parseInterface(keyB, valB)
		}
		fmt.Println("}")
		break
	default:
		fmt.Println("unknown key:", key, " type:", reflect.TypeOf(obj))
	}
}

func TraversalJson(jsn []byte) {
	for key, valA := range getJsonUnmarsh(jsn) {
		parseInterface(key, valA)
	}
}
