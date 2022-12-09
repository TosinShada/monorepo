package actions

import "github.com/TosinShada/monorepo/services/horizon/internal/corestate"

type CoreStateGetter interface {
	GetCoreState() corestate.State
}
