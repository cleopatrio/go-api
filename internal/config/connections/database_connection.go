package connections

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dock-tech/notes-api/internal/config/properties"
	"github.com/dock-tech/notes-api/internal/domain/interfaces"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb(cache interfaces.Cache, secret interfaces.Secret) *gorm.DB {
	ctx := context.Background()
	b, err := cache.Get(ctx, properties.CacheDbSecretKey)
	if err != nil || b == nil {
		b, err = secret.Get(ctx, properties.GetSecretDatabase())
		if err != nil {
			panic(err)
		}
		err = cache.Set(ctx, properties.CacheDbSecretKey, b, properties.CacheExpiration)
		if err != nil {
			fmt.Printf("Error setting cache: %s \n", err.Error())
		}
	}

	var secretData map[string]any
	err = json.Unmarshal(b, &secretData)
	if err != nil {
		fmt.Printf("Error parsing secret: %s \n", err.Error())
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%v sslmode=%s connect_timeout=%s application_name=%s",
		secretData["host"], secretData["username"],
		secretData["password"], secretData["dbname"],
		secretData["port"],
		properties.GetDatabaseSslMode(),
		properties.GetDatabaseTimeout(),
		properties.ApplicationRepositoryName)

	cfg := &gorm.Config{}
	dialect := postgres.Open(dsn)

	db, err := gorm.Open(dialect, cfg)
	if err != nil {
		panic("failed to connect to database. err: " + err.Error())
	}

	database, err := db.DB()
	if err != nil {
		panic("failed to connect to database. err: " + err.Error())
	}

	database.SetMaxOpenConns(properties.GetDatabaseMaxOpenConnections())
	database.SetMaxIdleConns(properties.GetDatabaseMaxIdleConnections())
	return db
}
