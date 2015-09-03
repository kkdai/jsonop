package jsonop

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Ops int

const (
	OpAdd Ops = iota + 1
	OpSub
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
	if one == nil && two == nil {
		return true
	} else if one == nil || two == nil {
		return false
	}
	return reflect.DeepEqual(one, two)
}

// To determin if two input json content is equal or not
// Please note: none equal to none, empty map equal to empty map
func JsonEqual(jsA, jsB []byte) bool {
	return compareTwo(getJsonUnmarsh(jsA), getJsonUnmarsh(jsB))
}

func jsonOps(objA, objB map[string]interface{}, op Ops) map[string]interface{} {

	if objA == nil || objB == nil {
		// None + B == None
		// A + None == None
		return nil
	}

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

	//B set bigger than A, only need in AddOp
	if op != OpAdd {
		return objA
	}

	for key, val := range objB {
		_, ok := objA[key]
		if ok {
			//A element in A  map, skip it.
			continue
		}

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

//Substract: Json object A substract Json object B
func JsonSubtract(jsA, jsB []byte) []byte {
	retMap := jsonOps(getJsonUnmarsh(jsA), getJsonUnmarsh(jsB), OpSub)
	retByte, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(" Unmarshall error:", err)
		return nil
	}
	return retByte
}

//Add: To add two json data directly.
//Please note the priciple to add two object as follow:
// stringA + stringB = stringA concat stringB
// sliceA + sliceB = sliceA append sliceB
func JsonAdd(jsA, jsB []byte) []byte {
	retMap := jsonOps(getJsonUnmarsh(jsA), getJsonUnmarsh(jsB), OpAdd)
	retByte, err := json.Marshal(retMap)
	if err != nil {
		fmt.Println(" Unmarshall error:", err)
		return nil
	}
	return retByte
}

func parseInterface(key string, obj interface{}, degree int) {

	var indentString string
	for i := 1; i <= degree; i++ {
		indentString = indentString + "\t"
	}

	switch v := obj.(type) {

	case []interface{}:
		listObj := obj.([]interface{})
		fmt.Printf(indentString+"key:%s val:[", key)
		for _, strV := range listObj {
			fmt.Printf("%v,", strV)
		}
		fmt.Printf("]\n")
		break
	case bool:
		fmt.Println(indentString+"type:", v, " key:", key, " val:", obj)
		break
	case string:
		fmt.Println(indentString+"type:", v, " key:", key, " val:", obj)
		break
	case int32, int64, float64:
		fmt.Println(indentString+"type:", v, " key:", key, " val:", obj)
		break
	case map[string]interface{}:
		mapObj := obj.(map[string]interface{})
		fmt.Println(indentString+"key:", key, "{")
		for keyB, valB := range mapObj {
			parseInterface(keyB, valB, degree+1)
		}
		fmt.Println(indentString + "}")
		break
	default:
		fmt.Println(indentString+"unknown key:", key, " type:", reflect.TypeOf(obj))
	}
}

//To print out whole map structure
func PrintJson(jsn []byte) {
	fmt.Printf("{\n")
	for key, valA := range getJsonUnmarsh(jsn) {
		parseInterface(key, valA, 1)
	}
	fmt.Printf("}\n")
}
