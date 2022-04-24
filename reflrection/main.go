package main

import (
	"fmt"
	"reflect"
)

func main() {
	var data interface{} = map[string]string{
		"text": "hello, world!",
	}

	// switch type check
	switch t := data.(type) {
	case string:
		fmt.Println("data is string")
		fmt.Println(t)
	case map[string]string:
		fmt.Println("data is map[string]string")
		for k, v := range t {
			fmt.Println(k, v)
		}
	}

	fmt.Println()

	// reflect type check
	switch rv := reflect.ValueOf(data); rv.Kind() {
	case reflect.String:
		fmt.Println(rv.Type().String())
		fmt.Println(rv.String())
	case reflect.Map:
		fmt.Println(rv.Type().String())
		fmt.Println(rv.String())
		iter := rv.MapRange()
		for iter.Next() {
			k := iter.Key()
			v := iter.Value()
			fmt.Println(k, v)
		}
	}
}
