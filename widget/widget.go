package widget

import (
	"fmt"
)

var NewWidget = func(property string) Runner {
	return &Widget{Property: property}
}

type Runner interface {
	Run() string
}

type Widget struct {
	Property string
}

func (widget *Widget) Run() string {
	return fmt.Sprintf("This widget has the property: %s\n", widget.Property)
}
