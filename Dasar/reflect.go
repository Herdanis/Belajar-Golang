package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true"`
}

type Contohlagi struct {
	Name string
	Nick string `required:"true"`
}

func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"tes"}
	sampleType := reflect.TypeOf(sample)

	fmt.Println(sampleType.NumField())
	fmt.Println(sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))

	fmt.Println(IsValid(sample))

	contoh := Contohlagi{"tes", ""}
	fmt.Println(IsValid(contoh))

}
