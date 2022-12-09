package ticker

import (
	"github.com/TosinShada/monorepo/services/ticker/internal/gql"
	"github.com/TosinShada/monorepo/services/ticker/internal/tickerdb"
	hlog "github.com/TosinShada/monorepo/support/log"
)

func StartGraphQLServer(s *tickerdb.TickerSession, l *hlog.Entry, port string) {
	graphql := gql.New(s, l)

	graphql.Serve(port)
}
