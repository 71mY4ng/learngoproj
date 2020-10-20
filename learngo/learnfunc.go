package learngo

import (
	"fmt"
)

const bigStructSize = 100

type bigStruct struct {
	a [bigStructSize]int
}

func newBigStruct() bigStruct {
	var b bigStruct

	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return b
}

func newBigStructPtr() *bigStruct {
	var b bigStruct

	for i := 0; i < bigStructSize; i++ {
		b.a[i] = i
	}
	return &b
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// 传入引用可以在方法中操作被引用的结构体对象
func (g *greeter) greetAndManipParent() {
	fmt.Println(g.greeting, g.name)
	g.name = "Went"
}

func RunCaseFunction() {

	for i := 0; i < 5; i++ {
		// 匿名函数
		// 以下这种为带参数形式的，这样传递 i 是为了线程安全性
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
	fmt.Println("message is ", g.name)
	g.greetAndManipParent()
	fmt.Println("message is ", g.name)

	d, err := divide(5.0, 3.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

	// 和 divide 一样但是用的是 on-the-fly-define 的形式
	var divideInline func(float64, float64) (float64, error)
	divideInline = func(dividend, divisor float64) (float64, error) {
		if divisor == 0.0 {
			return 0.0, fmt.Errorf("Cannot divide by zero")
		}
		return dividend / divisor, nil
	}

	d, err = divideInline(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

// divide Comma-Error 的风格, 比起直接 panic 这样方便给调用者处理
func divide(dividend, divisor float64) (float64, error) {
	if divisor == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return dividend / divisor, nil
}
