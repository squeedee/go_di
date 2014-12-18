package widget

import (
	"fmt"
)

var Builder WidgetBuilder = &widgetBuilder{}

type WidgetBuilder interface {
	BuildWidget(property string) Runner
}

type widgetBuilder struct {
}

func (builder *widgetBuilder) BuildWidget(property string) Runner {
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
