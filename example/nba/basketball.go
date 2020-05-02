package nba

import "context"

func (j James) PlayBall(ctx context.Context) {}

type (
	Basketball interface {
		PlayBall(ctx context.Context)
	}

	James struct{}
)

var _ Basketball = James{}
