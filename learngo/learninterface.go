package learngo

import (
	"bytes"
	"fmt"
	"io"
)

/*
* Best Pratices:
 * - 精简的 interface 更加灵活, io.Writer, io.Reader, interface{}
 * - 开放实现类而不是 interface, 例如 golang 的 DB
 * - 使用者为使用的类提供 interface
 * - 设计方法时参数尽量是 interface 的

 * 如上可见, golang 的 interface 并不像其他面向对象语言的接口, 更像是为了精简设计提供的片段(Fragmentation)
 * 简言之: 在 Go 中 "accepts interfaces, return structs"
 * 由于Go 的灵活 interface 特性, 你可以用 on-the-fly interface 来实现 IOC, 也可以用 interface 对引入的
 * 包进行 mock;  see: https://rakyll.org/interface-pollution/
*/
// RunCaseInterface
func RunCaseInterface() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	intIncr := IntCounter(0)
	var incr Incrementer = &intIncr
	for i := 0; i < 10; i++ {
		fmt.Println("incr: ", incr.Increment())
	}

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write([]byte("1234567890_1234\n1"))
	wc.Close()

	// NOTICE: type conversion (如果转类型失败会 panic)
	// bwc := wc.(*BufferedWriterCloser)
	// fmt.Printf("%v\n", bwc)

	// NOTICE: 有处理的转类型方式
	r, ok := wc.(io.Reader)
	if ok {
		fmt.Println(r)
	} else {
		fmt.Println("Conversion failed")
	}

	// NOTICE: 空 interface 体, 和匿名函数作用差不多, 都是为了 on-the-fly-define
	var myObj interface{} = NewBufferedWriterCloser()
	// 这里做了一个 WriterCloser 接口的类型判断
	if wc, ok = myObj.(WriterCloser); ok {
		wc.Write([]byte("ABCDEFGHIJ_KLMNO"))
		wc.Close()
	}

	// NOTICE: 通过 interface{} 我们可以模糊地声明类型，类似于 java 的 Object 类一样
	var i interface{} = true
	switch i.(type) {
	case int:
		fmt.Println("i interface{} is an integer")
	case bool:
		fmt.Println("i interface{} is a boolean type")
	case string:
		fmt.Println("i interface{} is a string")
	default:
		fmt.Printf("i interface{} type is %t", i)
	}
}

type Incrementer interface {
	Increment() int
}

// IntCounter 这个实现类是一个 int 类型, 可以视为在 int 本身上增加方法, 类似于 extends
type IntCounter int

// Increment 使用实现类本身来存值
func (ic *IntCounter) Increment() int {
	*ic++
	return int(*ic)
}

/*
* 注意: 官方并不鼓励开放的接口用 interface 的方式
* link: https://github.com/golang/go/wiki/CodeReviewComments#interfaces
 */
// Writer 接口命名规则: Action + er
type Writer interface {
	Write([]byte) (int, error)
}

type Closer interface {
	Close() error
}

type WriterCloser interface {
	Writer
	Closer
}

// ConsoleWriter 可以看到 go 的实现类并不强制实现接口
type ConsoleWriter struct{}

func (cw ConsoleWriter) Write(data []byte) (int, error) {

	n, err := fmt.Println(string(data))
	return n, err
}

type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8) // 申请 8 bytes 的缓冲区来读取
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}

	return n, nil
}

func (bwc *BufferedWriterCloser) Close() error {

	// buffer 里还有数据就继续打印
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// Constructor
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
