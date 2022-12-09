package ingest

import (
	"context"

	"github.com/TosinShada/monorepo/services/horizon/internal/db2/history"
	"github.com/TosinShada/monorepo/services/horizon/internal/ingest/processors"
	"github.com/stretchr/testify/mock"
)

type MockFilters struct {
	mock.Mock
}

func (m *MockFilters) GetFilters(filterQ history.QFilter, ctx context.Context) []processors.LedgerTransactionFilterer {
	return []processors.LedgerTransactionFilterer{}
}
