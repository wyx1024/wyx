package model

import "encoding/json"

type Car struct {
	Name         string
	Price        string //价格
	Level        string //级别
	Fuel         string //油耗
	Transmission string //变速箱
	Displacement string // 排量
	Structure    string //结构
	Guarantee    string //保修
}

func FormJsonObj(o interface{}) (Car, error) {
	var car Car
	s, err := json.Marshal(o)
	if err != nil {
		return Car{}, err
	}
	err = json.Unmarshal(s, &car)
	return car, err
}
