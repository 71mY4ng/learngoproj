package learngo

import (
	"fmt"
)

// RunCasePointers 从此例子中得知: go 的赋值操作都是进行拷贝的，数据结构中是否存在内含的指针，影响拷贝后变量的操作结果
func RunCasePointers() {
	var a int = 42
	// referencing
	var b *int = &a
	fmt.Println(a, b)
	a = 24
	// dereferencing (by adding a '*')
	fmt.Println(a, *b)
	*b = 15
	fmt.Println(a, *b)

	arr := [3]int{1, 2, 3}
	arra := &arr[0]
	arrb := &arr[1]
	fmt.Printf("array arr values: %v, arra pointer: %p, arrb pointer: %p\n", arr, arra, arrb)
	// 下面的例子可以看出:
	// go 里面基本类型, array 和 struct 的赋值是拷贝的
	barr := arr
	fmt.Println("array barr: ", barr)
	arr[2] = 99
	fmt.Printf("array arr: %v, array barr: %v\n", arr, barr)

	var ms *myStruct
	ms = new(myStruct)
	(*ms).foo = 42
	fmt.Println("`(*ms).foo`: ", (*ms).foo)
	// 这么做是等价的，因为前者是后者去语法糖化后的结果
	ms.foo = 35
	fmt.Println("`ms.foo`: ", ms.foo)

	// 下面的例子可以看出:
	// go 里 slice 和 map 的赋值是拷贝了数据结构中的引用
	// slice 和 map 内含了指针
	amap := map[string]string{"f": "b", "baz": "buz"}
	bmap := amap
	fmt.Printf("map amap: %v, map bmap: %v\n", amap, bmap)
	amap["f"] = "q"
	fmt.Printf("map amap: %v, map bmap: %v\n", amap, bmap)

	aslice := []int{1, 2, 3}
	bslice := aslice
	fmt.Printf("slice aslice: %v, slice bslice: %v\n", aslice, bslice)
	fmt.Printf("slice aslice addr: %p, slice bslice addr: %p\n", aslice, bslice)
	aslice[2] = 66
	fmt.Printf("slice aslice: %v, slice bslice: %v\n", aslice, bslice)
}

type myStruct struct {
	foo int
}
