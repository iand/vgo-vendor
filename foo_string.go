// Code generated by "stringer -type=Foo"; DO NOT EDIT.

package main

import "strconv"

const _Foo_name = "FooOneFooTwoFooThree"

var _Foo_index = [...]uint8{0, 6, 12, 20}

func (i Foo) String() string {
	if i < 0 || i >= Foo(len(_Foo_index)-1) {
		return "Foo(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Foo_name[_Foo_index[i]:_Foo_index[i+1]]
}