package cxk

import (
	"context"

	"github.com/Just-maple/svcw/example/nba"
)

type (
	CXK struct {
		Singer  Singer
		Dancer  Dancer
		NBAStar nba.Basketball
		Rapper  Rapper
	}
	Singer interface {
		Sing(ctx context.Context)
	}
	Dancer interface {
		Dance(ctx context.Context)
	}

	Rapper interface {
		Rap(ctx context.Context)
	}
	ElectricVoice interface {
		AutoTune(ctx context.Context)
	}

	LadyGaga struct{}

	ChrisWu struct {
		ElectricVoice ElectricVoice
	}
	ZhaoFour struct{}
	AutoTune struct{}
)

func (a AutoTune) AutoTune(ctx context.Context) {}

func (z ZhaoFour) Dance(ctx context.Context) {}

func (c ChrisWu) Rap(ctx context.Context) {}

func (l LadyGaga) Sing(ctx context.Context) {}

var (
	_ Singer        = LadyGaga{}
	_ Rapper        = ChrisWu{}
	_ Dancer        = ZhaoFour{}
	_ ElectricVoice = AutoTune{}
)
