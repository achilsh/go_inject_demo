package ProviderDemo

import "github.com/google/wire"

type Xyz struct {
	X int
}

func ProviderXyz() (y Xyz, err error) {
	y = Xyz{
		X: 100,
	}
	err = nil
	return
}

type Abc struct {
	A string
	Y int
}

func ProviderAbc(y Xyz) (b Abc, err error) {
	b = Abc{
		A: "aaaaa",
		Y: y.X,
	}
	err = nil
	return
}

type IJK struct {
	II Abc
}

func ProvideAbcAgain(a Abc) (b IJK, err error) {
	b = IJK{
		II: Abc{
			A: a.A,
			Y: a.Y,
		},
	}
	err = nil
	return
}

//
var SetDemoOne = wire.NewSet(ProviderXyz, ProviderAbc, ProvideAbcAgain)
