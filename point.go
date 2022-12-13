package pointmanagerpackage

import (
	"fmt"

	"github.com/rianocta97/pointmanagerpackage/usecase"
)

type PointManager struct {
	pointUsecase usecase.PointUsecase
}

func (p *PointManager) IncreasePoint(point int) int {
	fmt.Printf("Point awal: %d\n", point)

	point = p.pointUsecase.AddPoint(point)

	fmt.Printf("Point baru: %d\n", point)

	return point
}

func InitPointManager() *PointManager {
	pm := new(PointManager)
	pm.pointUsecase = usecase.InitPointUsecase()
	return pm
}

func PointManagerIncreasePoint(p int) int {
	pm := InitPointManager()
	p = pm.IncreasePoint(p)
	return p
}