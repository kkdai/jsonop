package jsonop

import "testing"

func TestEqualNil(t *testing.T) {
	if !JsonEqual(nil, nil) {
		t.Errorf("Equal cannot handle nil\n")
	}

	byt1 := []byte(`{}`)
	if !JsonEqual(byt1, byt1) {
		t.Errorf("Equal cannot handle empty json\n")
	}
}

func TestEqual(t *testing.T) {
	bytEmpty := []byte(`{}`)

	if !JsonEqual(bytEmpty, bytEmpty) {
		t.Errorf("Equal failed on empty map.")
	}

	byt1 := []byte(`{
			"num":6,
			"strs":"a",
			"num2":7 }`)

	byt2 := []byte(`{
			"strs":"a",
			"num":6,
			"num2":7 }`)

	if !JsonEqual(byt1, byt1) {
		t.Errorf("Equal fail on json simple struct\n")
	}

	if !JsonEqual(byt1, byt2) {
		t.Errorf("Equal fail on json address change\n")
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
		t.Errorf("Equal fail on json array.\n")
	}
	if JsonEqual(byt3, byt4) {
		t.Errorf("Equal fail on json array, address change should treat as different.\n")
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

func TestPrint(t *testing.T) {
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
	PrintJson(byt6)
}

func TestJsonOpAdd(t *testing.T) {

	if !JsonEqual(JsonAdd(nil, nil), nil) {
		t.Errorf("Json Add failed on nil \n")
	}

	bytEmpty := []byte(`{}`)

	byt1 := []byte(`{
		"num":6,
		"strs":"a",
		"num2":7 }`)

	if !JsonEqual(JsonAdd(byt1, nil), nil) {
		t.Errorf("Json add failed on jsonB nil\n")
	}

	if !JsonEqual(JsonAdd(nil, byt1), nil) {
		t.Errorf("Json add failed on jsonA nil\n")
	}

	if !JsonEqual(JsonAdd(bytEmpty, byt1), byt1) {
		t.Errorf("Json add failed on empty map A\n")
	}

	if !JsonEqual(JsonAdd(byt1, bytEmpty), byt1) {
		t.Errorf("Json add failed on empty map A\n")
	}

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
