//+build svcw
// Code generated by SvcWrap. DO NOT EDIT.

package cxk

import (
	"context"
	"fmt"
	"runtime"

	"github.com/Just-maple/svcw/example/nba"
	"github.com/Just-maple/svcw/wrapflask"
)

type _w_ChrisWu_ElectricVoice struct{ ElectricVoice }

var _ = wrapflask.Register(_w_ChrisWu_ElectricVoice{}, new(ElectricVoice))

func (this _w_ChrisWu_ElectricVoice) AutoTune(p0 context.Context) {

	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)

	//func(context.Context)
	this.ElectricVoice.AutoTune(p0)
	return
}

type _w_CXK_Singer struct{ Singer }

var _ = wrapflask.Register(_w_CXK_Singer{}, new(Singer))

func (this _w_CXK_Singer) Sing(p0 context.Context) {

	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)

	//func(context.Context)
	this.Singer.Sing(p0)
	return
}

type _w_CXK_Dancer struct{ Dancer }

var _ = wrapflask.Register(_w_CXK_Dancer{}, new(Dancer))

func (this _w_CXK_Dancer) Dance(p0 context.Context) {

	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)

	//func(context.Context)
	this.Dancer.Dance(p0)
	return
}

type _w_CXK_NBAStar struct{ nba.Basketball }

var _ = wrapflask.Register(_w_CXK_NBAStar{}, new(nba.Basketball))

func (this _w_CXK_NBAStar) PlayBall(p0 context.Context) {

	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)

	//func(context.Context)
	this.Basketball.PlayBall(p0)
	return
}

type _w_CXK_Rapper struct{ Rapper }

var _ = wrapflask.Register(_w_CXK_Rapper{}, new(Rapper))

func (this _w_CXK_Rapper) Rap(p0 context.Context) {

	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)

	//func(context.Context)
	this.Rapper.Rap(p0)
	return
}
