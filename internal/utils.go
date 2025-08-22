// Package internal
package internal

import (
	"fmt"
	"reflect"
)

func PrintLines(s any) {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.Type().Kind() == reflect.String {
			fmt.Println(f.String())
		}
	}
}
