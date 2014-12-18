package main

import (
	"fmt"
	. "github.com/squeedee/go_di/widget"
	. "github.com/squeedee/go_di/widget/fakes")

func main() {
	fmt.Println("Welcome to Rash")

	Builder.BuildWidget("foo").Run()

	Builder = &FakeWidgetBuilder{}
	fakeWidget := Builder.BuildWidget("foo")
	fakeWidget.Run()

	fmt.Printf("Fake widget what happened: %s", fakeWidget.(FakeRunner).GetWhatHappened())
}
