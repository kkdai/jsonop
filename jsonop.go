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
	if one == nil || two == nil {
		return false
	}
	return reflect.DeepEqual(one, two)
}

func JsonEqual(jsA, jsB []byte) bool {
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

func jsonOps(objA, objB map[string]interface{}, op Ops) map[string]interface{} {

	//Search all key in A first
	for key, val := range objA {
		valB, ok := objB[key]
		if !ok {
			//A element not in B map, skip it.
			continue
		}

		if reflect.TypeOf(val) != reflect.TypeOf(valB) {
			//the same key but type different, skip it.
			continue
		}

		switch v := val.(type) {

		case []interface{}:
			objA[key] = sliceOps(objA[key].([]interface{}), objB[key].([]interface{}), op)
		case bool:
			objA[key] = boolOps(objA[key].(bool), objB[key].(bool), op)
		case string:
			objA[key] = stringOps(objA[key].(string), objB[key].(string), op)
		case int64, int32:
			objA[key] = intOps(objA[key].(int64), objB[key].(int64), op)
		case float64:
			objA[key] = float64Ops(objA[key].(float64), objB[key].(float64), op)
		case map[string]interface{}:
			objA[key] = jsonOps(objA[key].(map[string]interface{}), objB[key].(map[string]interface{}), op)
		default:
			fmt.Println("unsupport key:", key, " type:", reflect.TypeOf(objA), v)
		}
	}

	for key, val := range objB {
		_, ok := objA[key]
		if ok {
			//A element in A  map, skip it.
			continue
		}

		//if reflect.TypeOf(val) != reflect.TypeOf(valA) {
		//	//the same key bit type different, skip it.
		//	continue
		//}

		switch v := val.(type) {

		case []interface{}:
			var emptySlice []interface{}
			objA[key] = sliceOps(emptySlice, objB[key].([]interface{}), op)
		case bool:
			objA[key] = boolOps(false, objB[key].(bool), op)
		case string:
			objA[key] = stringOps("", objB[key].(string), op)
		case int64, int32:
			objA[key] = intOps(int64(0), objB[key].(int64), op)
		case float64:
			objA[key] = float64Ops(float64(0.0), objB[key].(float64), op)
		case map[string]interface{}:
			objA[key] = jsonOps(objA[key].(map[string]interface{}), objB[key].(map[string]interface{}), op)
		default:
			fmt.Println("unsupport key:", key, " type:", reflect.TypeOf(objA), v)
		}
	}
	return objA
}

func JsonAdd(jsA, jsB []byte) []byte {
	retMap := jsonOps(getJsonUnmarsh(jsA), getJsonUnmarsh(jsB), OpAdd)
	fmt.Println(" Ret map:", retMap)
	retByte, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(" Unmarshall error:", err)
		return nil
	}
	return retByte
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
