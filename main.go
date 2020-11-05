package main

import (
	"fmt"

	"github.com/xueliangliang/uuidgeneratormock/pkg/uuidgeneratormock"
)

func main() {
	uuidgeneratormock.Init([]string{"1", "2", "3", "4", "5"})
	fmt.Println(uuidgeneratormock.UUIDGeneratorMock())
	fmt.Println(uuidgeneratormock.UUIDGeneratorMock())
	fmt.Println(uuidgeneratormock.UUIDGeneratorMock())
	uuidgeneratormock.Reset()
	fmt.Println(uuidgeneratormock.UUIDGeneratorMock())
}
