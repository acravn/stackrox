package datastore

import (
	pgStore "github.com/stackrox/rox/central/activecomponent/datastore/internal/store/postgres"
	"github.com/stackrox/rox/central/activecomponent/datastore/search"
	"github.com/stackrox/rox/central/globaldb"
	"github.com/stackrox/rox/pkg/sync"
)

var (
	once sync.Once

	ds DataStore
)

func initialize() {
	storage := pgStore.New(globaldb.GetPostgres())
	indexer := pgStore.NewIndexer(globaldb.GetPostgres())
	searcher := search.NewV2(storage, indexer)
	ds = New(storage, indexer, searcher)
}

// Singleton provides the interface for non-service external interaction.
func Singleton() DataStore {
	once.Do(initialize)
	return ds
}
