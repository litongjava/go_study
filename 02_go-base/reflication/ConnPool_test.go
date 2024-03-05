package reflication

import (
	"fmt"
	"reflect"
	"testing"
)

type ConnPool interface {
	Id() int
}

type MyConnPool struct {
	id int
}

func (mcp *MyConnPool) Id() int {
	return mcp.id
}

func TestReflication(t *testing.T) {
	connPoolA := &MyConnPool{id: 1}
	connPoolB := &MyConnPool{id: 2}

	fmt.Println(connPoolA.Id())
	fmt.Println(connPoolB.Id())

	//reflect.Value
	reflectValue := reflect.ValueOf(connPoolA)
	//uintptr
	pointer := reflectValue.Pointer()
	areSame := pointer == reflect.ValueOf(connPoolB).Pointer()

	if areSame {
		fmt.Println("connPoolA 和 connPoolB 指向同一个对象")
	} else {
		fmt.Println("connPoolA 和 connPoolB 不是指向同一个对象")
	}
}
