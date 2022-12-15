package pointmanagerpackage

import (
	"encoding/json"
	"fmt"

	"github.com/rianocta97/pointmanagerpackage/usecase"
)

type pointManagerImpl struct {
	pointUsecase usecase.PointUsecase
}

func (p *pointManagerImpl) increasePoint(point int) int {
	fmt.Printf("Point awal: %d\n", point)

	point = p.pointUsecase.AddPoint(point)

	fmt.Printf("Point baru: %d\n", point)

	return point
}

// func (p *pointManagerImpl) getBalance(data interface{}) (map[string]interface{}, error) {
// 	m := make(map[string]interface{})
// 	if err := json.Unmarshal(data.([]byte), &m); err != nil {
// 		return nil, err
// 	}
// }

func (p *pointManagerImpl) getStatement() {}

func (p *pointManagerImpl) genericRedeem() {}

func initPointManager() *pointManagerImpl {
	pm := new(pointManagerImpl)
	pm.pointUsecase = usecase.InitPointUsecase()
	return pm
}

func PointManagerIncreasePoint(p int) int {
	pm := initPointManager()
	p = pm.increasePoint(p)
	return p
	// pm.pointUsecase.AddPoint()
}

func GetBalance(data interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data.([]byte), &m); err != nil {
		return nil, err
	}

	pu := usecase.InitPointUsecase()
	pu.GetBalance(m)
	return nil, nil
}