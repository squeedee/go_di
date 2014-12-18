package main

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

// ***** NO Touch above

type FakeWidgetBuilder struct {
}

func (fakeBuilder *FakeWidgetBuilder) BuildWidget(property string) Runner {
	return &FakeWidget{}
}

type FakeRunner interface {
	Runner
	GetWhatHappened() string
}

type FakeWidget struct {
	WhatHappened string
}

func (widget *FakeWidget) Run() {
	widget.WhatHappened = "Ran Run"
}

func (widget *FakeWidget) GetWhatHappened() string {
	return widget.WhatHappened
}

func main() {
	fmt.Println("Welcome to Rash")

	Builder.BuildWidget("foo").Run()

	Builder = &FakeWidgetBuilder{}
	fakeWidget := Builder.BuildWidget("foo")
	fakeWidget.Run()

	fmt.Printf("Fake widget what happened: %s", fakeWidget.(FakeRunner).GetWhatHappened())
}
