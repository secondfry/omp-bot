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

	HasBefore(idx uint64) bool
	HasAfter(idx uint64) bool
}

type DummyTheServiceService struct{}

func NewDummyTheServiceService() *DummyTheServiceService {
	return &DummyTheServiceService{}
}

func (service *DummyTheServiceService) Describe(theservice_id uint64) (*TheService, error) {
	check, err := service.CheckId(theservice_id)

	if !check || err != nil {
		return nil, err
	}

	return &allEntities[theservice_id], nil
}

func (service *DummyTheServiceService) List(cursor uint64, limit uint64) []TheService {
	end := cursor + limit
	size := uint64(len(allEntities))

	if end > size {
		end = size
	}

	return allEntities[cursor:end]
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
	check, err := service.CheckId(theservice_id)

	if !check || err != nil {
		return false, err
	}

	allEntities = append(allEntities[:theservice_id], allEntities[theservice_id+1:]...)
	return true, nil
}

func (service *DummyTheServiceService) CheckId(theservice_id uint64) (bool, error) {
	if theservice_id >= uint64(len(allEntities)) {
		return false, errors.New("index out of range")
	}

	return true, nil
}

func (service *DummyTheServiceService) HasBefore(idx uint64) bool {
	return idx > 0
}

func (service *DummyTheServiceService) HasAfter(idx uint64) bool {
	return idx < uint64(len(allEntities))
}
