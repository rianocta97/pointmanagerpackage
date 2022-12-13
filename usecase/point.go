package usecase

import "pointmanagerpackage/repository"

type PointUsecase interface {
	AddPoint(point int) int
}

type PointUsecaseImpl struct {
	pointRepo repository.PointRepo
}

func (p *PointUsecaseImpl) AddPoint(point int) int {
	point += p.pointRepo.Add()

	return point
}

func InitPointUsecase() PointUsecase {
	return &PointUsecaseImpl{
		pointRepo: repository.InitPointRepo(),
	}
}