package usecase

import (
	"fmt"

	"github.com/rianocta97/pointmanagerpackage/repository"
)

type PointUsecase interface {
	AddPoint(point int) int
	GetBalance(data interface{})
}

type PointUsecaseImpl struct {
	pointRepo repository.PointRepo
}

func (p *PointUsecaseImpl) AddPoint(point int) int {
	point += p.pointRepo.Add()

	return point
}

func (p *PointUsecaseImpl) GetBalance(m interface{}) {
	fmt.Println(m)
	data := m.(map[string]interface{})
	fmt.Println(data)
}

func InitPointUsecase() PointUsecase {
	return &PointUsecaseImpl{
		pointRepo: repository.InitPointRepo(),
	}
}