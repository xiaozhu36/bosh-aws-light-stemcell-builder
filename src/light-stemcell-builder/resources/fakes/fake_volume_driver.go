// This file was generated by counterfeiter
package fakes

import (
	"light-stemcell-builder/resources"
	"sync"
)

type FakeVolumeDriver struct {
	CreateStub        func(resources.VolumeDriverConfig) (resources.Volume, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 resources.VolumeDriverConfig
	}
	createReturns struct {
		result1 resources.Volume
		result2 error
	}
	DeleteStub        func(resources.Volume) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 resources.Volume
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeVolumeDriver) Create(arg1 resources.VolumeDriverConfig) (resources.Volume, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 resources.VolumeDriverConfig
	}{arg1})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeVolumeDriver) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumeDriver) CreateArgsForCall(i int) resources.VolumeDriverConfig {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].arg1
}

func (fake *FakeVolumeDriver) CreateReturns(result1 resources.Volume, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 resources.Volume
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeDriver) Delete(arg1 resources.Volume) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 resources.Volume
	}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeVolumeDriver) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeVolumeDriver) DeleteArgsForCall(i int) resources.Volume {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].arg1
}

func (fake *FakeVolumeDriver) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ resources.VolumeDriver = new(FakeVolumeDriver)