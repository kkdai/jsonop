package jsonop

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"
)

func TestNil(t *testing.T) {
	if IsJsonEqual(nil, nil) {
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

	if !IsJsonEqual(byt1, byt1) {
		t.Errorf("Equal fail on json simple struct")
	}

	if !IsJsonEqual(byt1, byt2) {
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

	if !IsJsonEqual(byt3, byt3) {
		t.Errorf("Equal fail on json array.")
	}
	if IsJsonEqual(byt3, byt4) {
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

	if !IsJsonEqual(byt5, byt5) {
		t.Errorf("Equal fail on nest struct")
	}
	if IsJsonEqual(byt5, byt6) {
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
	if IsJsonEqual(jsnA, jsnB) == false {
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
			"num":6,
			"strs":["a","b"],
			"stra": {
				"num2": 7,
				"strA": "c",
				"num3": 8
				} 
			}`)
	TraversalJson(byt6)
}
