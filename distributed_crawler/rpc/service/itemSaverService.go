package service

import (
	"context"
	"crawler/distributed_crawler/engine"
	"crawler/distributed_crawler/model"
	"log"

	"github.com/olivere/elastic/v7"
)

type ItemSaverService struct {
	Client *elastic.Client
	Index  string
}

func (i ItemSaverService) Saver(item engine.Item, resule *bool) error {
	_, err := i.Client.Index().Index(i.Index).Id(item.Id).BodyJson(item).Do(context.Background())
	car, err := model.FormJsonObj(item.Payload)
	if err != nil {
		return err
	}
	log.Printf("%v", car)
	if err == nil {
		*resule = true
	}
	return err
}
