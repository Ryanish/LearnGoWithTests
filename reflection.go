package main

import "reflect"

func walk(x interface{}, fn func(string)) {
	val := reflect.ValueOf(x)
	field := val.Field(0)
	fn(field.String())
}

/// Skipped the rest of the relfection part as it advised to avoid using it in any case...
