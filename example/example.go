package main

import (
	"context"

	svcw "github.com/Just-maple/svc-interruption"
	"github.com/Just-maple/svc-interruption/example/cxk"
	"github.com/Just-maple/svc-interruption/example/nba"
	"github.com/Just-maple/svc-interruption/wrapflask"
)

var ctx = context.Background()

// language=go
const interruptionCode = `
	f, _, _, _ := runtime.Caller(0)
	name := runtime.FuncForPC(f).Name()
	fmt.Printf("call:%v\n", name)
	defer fmt.Printf("done:%v\n", name)
`

func main() {
	// create a cxk instance
	cxkInstance := &cxk.CXK{
		Singer:  cxk.LadyGaga{},
		Dancer:  cxk.ZhaoFour{},
		NBAStar: nba.James{},
		Rapper:  cxk.ChrisWu{ElectricVoice: cxk.AutoTune{}},
	}

	// inject codes to every interface call
	err := svcw.Gen(cxkInstance, interruptionCode, true)
	if err != nil {
		panic(err)
	}
	// only with build tag `svcw` will replace service dependencies
	wrapflask.Wrap(cxkInstance)

	// call the service
	cxkInstance.Rapper.Rap(ctx)
	cxkInstance.Dancer.Dance(ctx)
	cxkInstance.NBAStar.PlayBall(ctx)
	cxkInstance.Rapper.Rap(ctx)
}
