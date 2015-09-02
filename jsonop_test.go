package jsonop

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestNil(t *testing.T) {
	if JsonEqual(nil, nil) {
		t.Errorf("Equal cannot handle nil")
	}
}

func TestEqual(t *testing.T) {
	byt1 := []byte(`{
			"num":6,
			"strs":"a",
			"num2":7 }`)

	byt2 := []byte(`{
			"strs":"a",
			"num":6,
			"num2":7 }`)

	if !JsonEqual(byt1, byt1) {
		t.Errorf("Equal fail on json simple struct")
	}

	if !JsonEqual(byt1, byt2) {
		t.Errorf("Equal fail on json address change")
	}

	byt3 := []byte(`{
			"num":6,
			"strs":["a","b"],
			"num2":7 }`)

	byt4 := []byte(`{
			"num":6,
			"strs":["b","a"],
			"num2":7 }`)

	if !JsonEqual(byt3, byt3) {
		t.Errorf("Equal fail on json array.")
	}
	if JsonEqual(byt3, byt4) {
		t.Errorf("Equal fail on json array, address change should treat as different.")
	}

	byt5 := []byte(`{
			"num":6,
			"strs":["a","b"],
			"stra": {
				"num2": 7,
				"strA": "c"
				} 
			}`)

	byt6 := []byte(`{
			"num":6,
			"strs":["a","b"],
			"stra": {
				"num2": 7,
				"strA": "c",
				"num3": 8
				} 
			}`)

	if !JsonEqual(byt5, byt5) {
		t.Errorf("Equal fail on nest struct")
	}
	if JsonEqual(byt5, byt6) {
		t.Errorf("Equal fail on nest struct checking diff.")
	}
}

func TestMarshal(t *testing.T) {
	type testBaseA struct {
		name string `json:"name"`
		age  int    `json:"age"`
	}

	type testBaseB struct {
		name string `json:"name"`
		age  int64  `json:"age"`
	}

	dataA := &testBaseA{name: "test_name", age: 5}
	dataB := &testBaseB{name: "test_name", age: 5}

	jsnA, _ := json.Marshal(dataA)
	jsnB, _ := json.Marshal(&dataB)

	fmt.Println("jsA:", jsnA, " jsB:", jsnB)
	if JsonEqual(jsnA, jsnB) == false {
		t.Errorf("json marshal fail")
	}

	byt := []byte(`{"num":6,"strs":"a"}`)
	defJsn := make(map[string]interface{})
	json.Unmarshal(byt, &defJsn)
	fmt.Println(" Got:", defJsn)
	for key, val := range defJsn {
		fmt.Println("key:", key, " val:", val, " type:", reflect.TypeOf(val))
	}
}

func TestTraversal(t *testing.T) {
	byt6 := []byte(`{
			"bool_val": true,
			"num":6,
			"num_list": [3, 5, 7],
			"strs":["a","b"],
			"stra": {
				"num2": 7,
				"strA": "c",
				"num3": 8
				} 
			}`)
	TraversalJson(byt6)
}

func TestJsonOpAdd(t *testing.T) {

	byt1 := []byte(`{
		"num":6,
		"strs":"a",
		"num2":7 }`)

	byt2 := []byte(`{
		"num":1,
		"strs":"b",
		"num2":3 }`)
	byte12 := []byte(`{"num":7,"num2":10,"strs":"ab"}`)

	if !JsonEqual(JsonAdd(byt1, byt2), byte12) {
		t.Errorf("Json op add failed on basic item\n")
	}

	byt3 := []byte(`{
		"num":6,
		"strs":"a",
		"array_a": [ 1 ,2 ,3,4],
		"num2":7 }`)

	byt4 := []byte(`{
		"num":1,
		"array_a":[5,6],
		"strs":"b",
		"num2":3 }`)

	byte34 := []byte(`{
		"array_a":[1,2,3,4,5,6],
		"num":7,
		"num2":10,
		"strs":"ab"}`)
	if !JsonEqual(JsonAdd(byt3, byt4), byte34) {
		t.Errorf("Json op add failed on list\n")
	}

	byt5 := []byte(`{
		"num":6,
		"strs":"a",
		"matrix_a": {
			"num_a":1,
			"string_a": "a string"
		},
		"array_a": [ 1 ,2 ,3,4],
		"num2":7 }`)

	byt6 := []byte(`{
		"num":1,
		"matrix_a": {
			"num_a":2,
			"string_a": "b string"
		},
		"array_a":[5,6],
		"strs":"b",
		"num2":3 }`)

	byte56 := []byte(`{
		"array_a":[1,2,3,4,5,6],
		"matrix_a":{
			"num_a":3,
			"string_a":"a stringb string"
		},
		"num":7,
		"num2":10,
		"strs":"ab"}`)
	if !JsonEqual(JsonAdd(byt5, byt6), byte56) {
		t.Errorf("Json op add failed on matrix \n")
	}

	byt7 := []byte(`{
		"num":6,
		"strs":"a",
		"matrix_a": {
			"num_a":1,
			"string_a": "a string"
		},
		"array_a": [ 1 ,2 ,3,4],
		"array_b": ["c","d"],
		"num2":7 }`)

	byt8 := []byte(`{
		"num":1,
		"matrix_a": {
			"num_a":2,
			"num_bb":7,
			"string_a": " b string"
		},
		"array_a":[5,6],
		"array_b":["a","b"],
		"array_c":[0,9,8],
		"strB": "B string only",
		"strs":"b",
		"num2":3 }`)

	byte78 := []byte(`{
		"array_a":[1,2,3,4,5,6],
		"array_b":["c","d","a","b"],
		"array_c":[0,9,8],
		"matrix_a":{
			"num_a":3,
			"num_bb":7,
			"string_a":"a string b string"
		},
		"num":7,
		"num2":10,
		"strB":"B string only",
		"strs":"ab"}`)

	if !JsonEqual(JsonAdd(byt7, byt8), byte78) {
		t.Errorf("Json op add failed on B set bigger than A set \n")
	}
}
