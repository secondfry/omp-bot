package theservice

import (
	"errors"
)

type TheServiceService interface {
	Describe(theservice_id uint64) (*TheService, error)
	List(cursor uint64, limit uint64) []TheService
	Create(TheService) (uint64, error)
	Update(theservice_id uint64, theservice TheService) error
	Remove(theservice_id uint64) (bool, error)
}

type DummyTheServiceService struct{}

func NewDummyTheServiceService() *DummyTheServiceService {
	return &DummyTheServiceService{}
}

func (service *DummyTheServiceService) Describe(theservice_id uint64) (*TheService, error) {
	if theservice_id >= uint64(len(allEntities)) {
		return nil, errors.New("index out of range")
	}

	return &allEntities[theservice_id], nil
}

func (service *DummyTheServiceService) List(cursor uint64, limit uint64) []TheService {
	return allEntities[cursor : cursor+limit]
}

func (service *DummyTheServiceService) Create(theservice TheService) (uint64, error) {
	allEntities = append(allEntities, theservice)
	return uint64(len(allEntities) - 1), nil
}

func (service *DummyTheServiceService) Update(theservice_id uint64, theservice TheService) error {
	allEntities[theservice_id] = theservice
	return nil
}

func (service *DummyTheServiceService) Remove(theservice_id uint64) (bool, error) {
	allEntities = append(allEntities[:theservice_id], allEntities[theservice_id+1:]...)
	return true, nil
}
