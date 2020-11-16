package mongodb

import (
	"context"

	"github.com/mooxun/emgo-web/pkg/conf"
	"github.com/mooxun/emgo-web/pkg/logger"
	"github.com/qiniu/qmgo"
)

var mg *qmgo.Database

func Init() {
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: conf.App.Mongodb.Uri})
	if err != nil {
		logger.Error(err)
	}
	mg = client.Database(conf.App.Mongodb.Database)

	logger.Infof("Successfully connected and pinged.")
}

// mongodb 数据集合
func Coll(collection string) *qmgo.Collection {
	return mg.Collection(collection)
}