package learngo

import (
	"fmt"
	"log"
	"net/http"
)

// RunCaseBasicDefer defer 将声明代码的执行延后到了方法退出
// 方法里可以有多个 defer, 遵循 LIFO (后进先出) 的顺序执行, 类似于将声明代码放入栈中
func RunCaseBasicDefer() {

	a := "start"
	defer fmt.Println(a)
	a = "end"
}

// RunCasePanicDeferRecover defer关键字执行于 panic 之前, recover() 在panic 后拥有了err 的返回值
func RunCasePanicDeferRecover() {

	fmt.Println("about to panic")

	defer func() {
		// 此处利用了 if condition block 中的求值表达式
		if err := recover(); err != nil {
			log.Println("Error: ", err)
		}
	}()

	panic("Smth went wrong")

	fmt.Println("meh")
}

func RunCaseUrlGet() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}
