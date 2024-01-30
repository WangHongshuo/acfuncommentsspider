package articleslistspider

import (
	"github.com/WangHongshuo/acfuncommentsspider-go/internal/logger"
	"github.com/asynkron/protoactor-go/actor"
)

const actorName = "ArticlesListExec"

var log = logger.NewLogger(actorName)

type ArticlesListExecutor struct {
	pid      *actor.PID
	children []*actor.PID
}

func (a *ArticlesListExecutor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		a.init(ctx)
	default:
		log.Infof("%v recv msg: %T\n", ctx.Self().Id, msg)
	}
}
