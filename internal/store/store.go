package store

import (
	"Wb-L0/internal/model"
	"Wb-L0/internal/store/cache"
	"Wb-L0/internal/store/sqlstore"
	"log"
)

type Store struct {
	sqlStore *sqlstore.SqlStore
	cache    *cache.Cache
}

func NewStore(dbUrl string) (*Store, error) {
	log.Println("Connecting to postgres ...")
	db, err := sqlstore.New(dbUrl)
	if err != nil {
		return nil, err
	}

	log.Println("Creating cache ...")
	cch := cache.New()
	str := &Store{db, cch}
	if err = str.RestoreCache(); err != nil {
		return nil, err
	}
	return str, nil

}

func (s *Store) RestoreCache() error {
	orders, err := s.sqlStore.GetAll()
	if err != nil {
		return err
	}
	for _, order := range orders {
		s.cache.Add(&order)
	}
	return nil
}

func (s *Store) GetByUUID(uuid string) string {
	return s.cache.GetByUUID(uuid)
}

func (s *Store) Add(order *model.Order) error {
	if err := s.cache.Add(order); err != nil {
		return err
	}
	if err := s.sqlStore.Add(order); err != nil {
		return err
	}
	return nil
}

func (s *Store) Close() {
	s.sqlStore.Close()
}
