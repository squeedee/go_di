package main

import (
	"fmt"
	"github.com/squeedee/go_di/widget"
	"github.com/squeedee/go_di/widget/fakes"
)

func main() {
	fmt.Println("Welcome to Rash")

	// Normal use
	widgetInstance := widget.NewWidget("foo")
	widgetInstance.Run()

	// stubbing
	var fakeWidget fakes.FakeRunner = &fakes.FakeWidget{}

	widget.NewWidget = func(property string) widget.Runner {
		return fakeWidget
	}

	// 'normal' use with stub
	widgetInstance2 := widget.NewWidget("foo")
	widgetInstance2.Run()

	// 'assertion'
	fmt.Printf("Fake widget what happened: %s", fakeWidget.GetWhatHappened())
}
