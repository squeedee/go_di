package widget

import (
	"fmt"
)

var NewWidget = func(property string) Runner {
	return &Widget{Property: property}
}

type Runner interface {
	Run()
}

type Widget struct {
	Property string
}

func (widget *Widget) Run() {
	fmt.Printf("Run Was Called with: %s\n", widget.Property)
}
