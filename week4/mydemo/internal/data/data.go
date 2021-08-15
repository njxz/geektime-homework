package data

import (
	"context"
	"database/sql"
	entsql "entgo.io/ent/dialect/sql"
	"fmt"
	"geektime/internal/conf"
	"geektime/internal/data/ent"
	"geektime/internal/data/ent/user"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

type Data struct {
	// TODO wrapped database client
	Db *ent.Client
}

func (d Data) CreateId(id int64, name string) error {
	u, err := d.Db.User.Create().SetID(id).SetName(name).Save(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(u)
	return nil
}
func (d Data) SelectName(id int64) (string, error) {
	u, err := d.Db.User.Query().Where(user.ID(id)).Only(context.Background())
	if err != nil {
		return "", err
	}
	return u.Name, nil
}

// NewData .
func NewData(c *conf.MydemoDatabase) (*Data, func(), error) {
	drv, err := sql.Open("mysql", c.Mysql)
	if err != nil {
		return nil, func() {}, err
	}
	drive := entsql.OpenDB("mysql", drv)
	db := ent.NewClient(ent.Driver(drive))
	data := Data{Db: db}
	cleanup := func() {
		db.Close()
		//log.NewHelper(logger).Info("closing the data resources")
	}
	if err := db.Schema.Create(context.Background()); err != nil {
		return nil, func() {}, err
	}
	return &data, cleanup, nil
}
