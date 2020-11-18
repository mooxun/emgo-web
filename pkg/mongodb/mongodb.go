package mongodb

import (
	"context"
	"fmt"

	"github.com/mooxun/emgo-web/pkg/cfg"
	"github.com/mooxun/emgo-web/pkg/logger"
	"github.com/qiniu/qmgo"
)

var mg *qmgo.Database

func Init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: cfg.App.Mongodb.Uri})
	if err != nil {
		logger.Error(err)
	}
	fmt.Println(cfg.App.Mongodb.Database)
	mg = client.Database(cfg.App.Mongodb.Database)

	logger.Infof("Successfully connected and pinged.")
}

// mongodb 数据集合
func Coll(collection string) *qmgo.Collection {
	return mg.Collection(collection)
}