package jsonop

import (
	"encoding/json"
	"fmt"
	"reflect"
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

func addOp(objA, objB interface{}) interface{} {
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
	case string:
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

	default:
		fmt.Println("unknown key:", key, " type:", reflect.TypeOf(obj))
	}
}

func TraversalJson(jsn []byte) {
	for key, valA := range getJsonUnmarsh(jsn) {
		parseInterface(key, valA)
	}
}
