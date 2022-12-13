package repository

import config "github.com/rianocta97/pointmanagerpackage/config"

type PointRepo interface {
	Add() int
}

type PointRepoImpl struct {
	pointDatabase *config.SomeDatabase
}

func (p *PointRepoImpl) Add() int {
	// pointNow := p.pointAdder

	// return pointNow

	pointIncrease := p.pointDatabase.Point
	return pointIncrease
}

func InitPointRepo() PointRepo {
	config.Init()

	return &PointRepoImpl{
		pointDatabase: config.GetDatabase(),
	}
}