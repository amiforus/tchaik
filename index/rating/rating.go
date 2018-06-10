// Package rating defines types and methods for setting/getting ratings for paths and
// persisting this data.
package rating

import (
	"fmt"
	"sync"

	"github.com/amiforus/tchaik/index"
)

// Value is a type which represents a rating value.
type Value uint

// None is the Value used to mark a path as having no rating.
const None Value = 0

// IsValid returns true if the Value is valid.
func (v Value) IsValid() bool {
	return 0 <= v && v <= 5
}

// Store is an interface which defines methods necessary for setting and getting ratings for
// index paths.
type Store interface {
	// Set the rating for the path.
	Set(index.Path, Value) error
	// Get the rating for the path.
	Get(index.Path) Value
}

// NewStore creates a basic implementation of a ratings store, using the given path as the
// source of data. Note: we do not enforce any locking on the underlying file, which is read
// once to initialise the store, and then overwritten after each call to Set.
func NewStore(path string) (Store, error) {
	m := make(map[string]Value)
	s, err := index.NewPersistStore(path, &m)
	if err != nil {
		return nil, err
	}

	return &store{
		m:     m,
		store: s,
	}, nil
}

type store struct {
	sync.RWMutex

	m     map[string]Value
	store index.PersistStore
}

// Set implements Store.
func (s *store) Set(p index.Path, v Value) error {
	s.Lock()
	defer s.Unlock()

	s.m[fmt.Sprintf("%v", p)] = v
	return s.store.Persist(&s.m)
}

// Get implements Store.
func (s *store) Get(p index.Path) Value {
	s.RLock()
	defer s.RUnlock()

	return s.m[fmt.Sprintf("%v", p)]
}
