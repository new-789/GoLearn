package log

import (
	"fmt"
)

type Console struct {
}

func NewConsole(file string) LogInterface {
	return Console{}
}

func (c Console) LogDebug(msg string) {
	fmt.Println("console Debug:", msg)
}

func (c Console) LogWarn(msg string) {
	fmt.Println("console warn:", msg)
}
