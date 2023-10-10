package testx

import (
	"github.com/agiledragon/gomonkey/v2"
	"reflect"
)

func generateReplacement(v reflect.Value, results ...interface{}) interface{} {
	var values []reflect.Value

	if v.Type().NumOut() < len(results) {
		panic("too many results for target")
	}

	for i := 0; i < v.Type().NumOut(); i++ {
		tOut := v.Type().Out(i)
		if i < len(results) {
			vRes := reflect.ValueOf(results[i])
			if tOut != vRes.Type() {
				panic("results type not match")
			}
			values = append(values, vRes)
		} else {
			values = append(values, reflect.Zero(tOut))
		}
	}

	replacementValue := reflect.MakeFunc(v.Type(), func(args []reflect.Value) []reflect.Value {
		return values
	})

	return replacementValue.Interface()
}

// PatchResult 替换一个函数的实现为直接返回结果，target是目标函数，results是要替换的结果
func PatchResult(target interface{}, results ...interface{}) *gomonkey.Patches {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Func {
		panic("target has to be a Func")
	}

	replacement := generateReplacement(v, results...)

	return gomonkey.ApplyFuncReturn(target, replacement)
}

// Patch ... 替换一个函数的实现，target是目标函数，replacement是替换的实现
func Patch(target, replacement interface{}) *gomonkey.Patches {
	return gomonkey.ApplyFunc(target, replacement)
}

// PatchMethodResult ...替换一个结构体的方法实现为直接返回结果，target是目标方法函数，results是要替换的结果
func PatchMethodResult(target interface{}, results ...interface{}) *gomonkey.Patches {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Func {
		panic("target has to be a Func")
	}

	replacement := generateReplacement(v, results...)
	return PatchMethod(target, replacement)
}

// PatchMethod ...替换一个结构体的方法实现，target是目标方法函数，replacement是替换的实现（replacement第一个参数是receiver）
func PatchMethod(target, replacement interface{}) *gomonkey.Patches {
	v := reflect.ValueOf(target)
	if v.Kind() != reflect.Func {
		panic("target has to be a Func")
	}

	if v.Type().NumIn() < 1 {
		panic("target must be a method")
	}

	t := v.Type().In(0)
	methodName := ""

	println(t.NumMethod())

	//Trick 枚举所有method获取当前method名字
	for i := 0; i < t.NumMethod(); i++ {
		if t.Method(i).Func.Pointer() == v.Pointer() {
			methodName = t.Method(i).Name
			break
		}
	}

	return gomonkey.ApplyMethod(t, methodName, replacement)
}
