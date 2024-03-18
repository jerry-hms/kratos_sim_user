package data

import (
	"gorm.io/gorm/schema"
	"kratos_sim/app/user/service/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo, NewDB)

// Data .
type Data struct {
	db  *gorm.DB
	log *log.Helper
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	logHelper := log.NewHelper(log.With(logger, "module", "user-server/data/gorm"))

	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   conf.Database.Prefix,
			SingularTable: true,
		},
	})
	if err != nil {
		logHelper.Fatalf("failed opening connection to mysql: %v", err)
	}

	return db
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		db:  db,
		log: log.NewHelper(log.With(logger, "module", "data/server")),
	}, cleanup, nil
}
