package horizon

import (
	"log"
	"time"

	"github.com/stellar/throttled"

	"github.com/TosinShada/monorepo/network"
	"github.com/TosinShada/monorepo/services/horizon/internal/test"
	supportLog "github.com/TosinShada/monorepo/support/log"
)

func NewTestApp() *App {
	app, err := NewApp(NewTestConfig())
	if err != nil {
		log.Fatal("cannot create app", err)
	}
	return app
}

func NewTestConfig() Config {
	return Config{
		DatabaseURL:            test.DatabaseURL(),
		StellarCoreDatabaseURL: test.StellarCoreDatabaseURL(),
		RateQuota: &throttled.RateQuota{
			MaxRate:  throttled.PerHour(1000),
			MaxBurst: 100,
		},
		ConnectionTimeout: 55 * time.Second, // Default
		LogLevel:          supportLog.InfoLevel,
		NetworkPassphrase: network.TestNetworkPassphrase,
	}
}

func NewRequestHelper(app *App) test.RequestHelper {
	return test.NewRequestHelper(app.webServer.Router.Mux)
}
