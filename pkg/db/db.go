package db

import (
	"context"
	"database/sql"
	"fmt"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	"github.com/Nulandmori/micorservices-pattern/pkg/env"
	"github.com/Nulandmori/micorservices-pattern/services/item/ent"
	goauth "golang.org/x/oauth2/google"
)

var (
	dbUser                 = env.MustGetEnv("DB_USER")
	dbPwd                  = env.MustGetEnv("DB_PASS")
	instanceConnectionName = env.MustGetEnv("INSTANCE_CONNECTION_NAME")
	dbName                 = env.MustGetEnv("DB_NAME")
)

func NewClient() (*ent.Client, error) {
	ctx := context.Background()

	err := initProxy(ctx)
	if err != nil {
		panic(err)
	}

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s sslmode=disable", dbUser, dbPwd, dbName, instanceConnectionName)
	fmt.Print(dbURI)

	dbConn, err := sql.Open("cloudsqlpostgres", dbURI)
	if err != nil {
		panic(err)
	}

	drv := entsql.OpenDB(dialect.Postgres, dbConn)
	defer drv.Close()

	opts := []ent.Option{ent.Driver(drv)}

	return ent.NewClient(opts...), nil
}

func initProxy(ctx context.Context) error {
	const SQLScope = "https://www.googleapis.com/auth/sqlservice.admin"

	client, err := goauth.DefaultClient(ctx, SQLScope)
	if err != nil {
		return err
	}

	proxy.Init(client, nil, nil)

	return nil
}
