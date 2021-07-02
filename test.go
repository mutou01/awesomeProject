package main

import "errors"

type T1 struct {
	error
}

func (t *T1)Hello()(*T1){
	if 1>0 {
		t.error = errors.New("aaaaa")
	}
	return  t
}

func (t *T1)World()(string,error){
	if t.error!=nil {
		return  "",t.error
	}
	return "",nil
}


func test(){

	t:= &T1{}


	t.Hello()

	t.World()


	t.Hello().World()

	if t.error==nil {

	}
}