package fakes

import (
	"github.com/squeedee/go_di/widget"
)

type FakeWidgetBuilder struct {
}

func (fakeBuilder *FakeWidgetBuilder) BuildWidget(property string) widget.Runner {
	return &FakeWidget{}
}

type FakeRunner interface {
	widget.Runner
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
