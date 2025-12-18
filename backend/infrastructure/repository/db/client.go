package db

import (
	"context"
	"sync"
	"time"

	log "github.com/dulisoft/spirit/core/log/zapx"
	"github.com/dulisoft/spirit/core/store/database"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/uptrace/opentelemetry-go-extra/otelgorm"

	"gorm.io/gorm"
)

var (
	once sync.Once
)

// Data .
type Data struct {
	DB *gorm.DB
}

func NewDB(d *Data) *gorm.DB {
	return d.DB
}

func (d *Data) WithContext(ctx context.Context) *gorm.DB {
	return d.DB.WithContext(ctx)
}

// NewData .
func NewData(dbOpts *database.Options) (*Data, func(), error) {
	var err error
	var client *gorm.DB
	once.Do(func() {
		dbOpts.MaxConnectionLifeTime = dbOpts.MaxConnectionLifeTime * time.Second
		client, err = database.New(dbOpts)

	})
	if err != nil {
		log.Errorf("open mysql failed, err: %v", err)
		return nil, nil, err
	}

	// 数据库添加otelgorm插件
	if err = client.Use(otelgorm.NewPlugin()); err != nil {
		log.Errorf("init db otelgorm, err: %v\n", err.Error())
		return nil, nil, err
	}

	return &Data{
			DB: client,
		}, func() {
			log.Info("closing the data resources")
		}, nil
}
