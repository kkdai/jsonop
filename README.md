jsonop: A JSON Operation library
==================

[![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/kkdai/jsonop/master/LICENSE)  [![GoDoc](https://godoc.org/github.com/kkdai/jsonop?status.svg)](https://godoc.org/github.com/kkdai/jsonop)  [![Build Status](https://travis-ci.org/kkdai/jsonop.svg?branch=master)](https://travis-ci.org/kkdai/jsonop)


##What is this library

Json Operation Library help you to do some json basic operation directly. It support some basic operations such as `Add`, `Substract`, `Equal`, `PrintJson`.


##Operation constraint

#### "nil" is specific data type not mean empty.

Please note, `nil`(none) is specific data type, it is not equal to empty json with `"{}"`.
So, in this case any json operation with `nil` will become `nil`.

        //Equal rule
        None will not equal to Empty Json

        //Addition rule
        NONE (Add) NONE == NONE
        A (Add) NONE == NONE
        NONE (Add) A == NONE
       
#### Some data type don't support some operation

Currently some data type don't support operation (provide your idea if any :) ) as follow:

- `slice`: Only support "Add" as `append()`
- `string`: Only support "Add" as string connect.

Install
---------------
`go get github.com/kkdai/jsonop`


Usage
---------------

```go

package main

import (
	"fmt"
	
    . "github/kkdai/"
)

func main() {
	//nil will equal to nil as a none type.
	fmt.Println(JsonEqual(nil, nil))
	//true

	//Please note: nil is not equal empty json data with "{}"
	byt0 := []byte(`{}`)
	fmt.Println(JsonEqual(nil, byt0))
	//false

	byt1 := []byte(`{
                "num":6,
                "strs":"a",
                "num2":7 }`)

	byt2 := []byte(`{
                "strs":"a",
                "num":6,
                "num2":7 }`)
	//Will treat it as equal.
	fmt.Println(JsonEqual(byt1, byt2))
	//true

	//Sample for json Add
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

    fmt.Println(string(JsonAdd(byt3, byt4)))

	//{
	//  "array_a":[1,2,3,4,5,6],
	//  "num":7,
	//  "num2":10,
	//  "strs":"ab"
	//}

	//Subtract sample
	byt5 := []byte(`{
        "num":6,
        "strs":"a",
        "num2":7 }`)

	byt6 := []byte(`{
        "num":1,
        "strs":"b",
        "num2":3 }`)

	fmt.Println(string(JsonSubtract(byt5, byt6))) //{"num":5,"num2":4,"strs":"a"}


	byt7 := []byte(`{
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
			
    //Print json struct
	PrintJson(byt7)
	
	//{
    //	type: float64  key: num  val: 6
    //	key:num_list val:[3,5,7,]
    //	key:strs val:[a,b,]
    //	key: stra {
    //		type: float64  key: num3  val: 8
    //		type: float64  key: num2  val: 7
    //		type: string  key: strA  val: c
    //	}
    //	type: bool  key: bool_val  val: true
    //}
}

```

Inspired
---------------

- [https://github.com/d4l3k/messagediff](https://github.com/d4l3k/messagediff)
- [Go Playground sample](http://play.golang.org/p/rGCez-W36T)

Project52
---------------

It is one of my [project 52](https://github.com/kkdai/project52).


License
---------------

This package is licensed under MIT license. See LICENSE for details.

