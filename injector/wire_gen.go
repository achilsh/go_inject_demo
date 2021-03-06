// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package injector

import (
	"demo/ProviderDemo"
)

// Injectors from demo.go:

func InitializeIjk() (ProviderDemo.IJK, error) {
	xyz, err := ProviderDemo.ProviderXyz()
	if err != nil {
		return ProviderDemo.IJK{}, err
	}
	abc, err := ProviderDemo.ProviderAbc(xyz)
	if err != nil {
		return ProviderDemo.IJK{}, err
	}
	ijk, err := ProviderDemo.ProvideAbcAgain(abc)
	if err != nil {
		return ProviderDemo.IJK{}, err
	}
	return ijk, nil
}
