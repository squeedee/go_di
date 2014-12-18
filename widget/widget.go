package widget

import (
	"fmt"
)

var Builder WidgetBuilderer = &WidgetBuilder{}

type WidgetBuilderer interface {
	BuildWidget(property string) Runner
}

type WidgetBuilder struct {
}

func (builder *WidgetBuilder) BuildWidget(property string) Runner {
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
