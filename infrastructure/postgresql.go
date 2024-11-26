package infrastructure

import (
	"fmt"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	infrastructureconfiguration "panel-subs/infrastructure/configuration"
	admininfrastructurerepositorymodel "panel-subs/pkg/admin/infrastructure/repository/model"
	announcementsinfrastructurerepositorymodel "panel-subs/pkg/announcements/infrastructure/repository/model"
	configinfrastructurerepositorymodel "panel-subs/pkg/config/infrastructure/repository/model"
)

var DB *gorm.DB

func InitializePostgresql() {
	conn := infrastructureconfiguration.PostgresqlConn
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		if err := db.Use(otelgorm.NewPlugin(otelgorm.WithDBName(infrastructureconfiguration.DatabaseName))); err != nil {
			log.Fatal(err)
		} else {
			DB = db
			fmt.Printf("Connecting To PostgreSQL : %+v \n", DB)
		}
	}

	fmt.Println("=== Start Migrating ===")
	MigrateDB()
	fmt.Println("=== Finish Migrating ===")
}

func ClosePostgresql() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}

func MigrateDB() {
	// Note
	// ID format : [Module Name][Running Number]. ex : PRODUCTCOUPON00001

	m := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "ADMIN00001",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time

				return tx.Migrator().CreateTable(
					&admininfrastructurerepositorymodel.Admin{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "Announcements00001",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time

				return tx.Migrator().CreateTable(
					&announcementsinfrastructurerepositorymodel.Announcements{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
		{
			ID: "Config00001",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time

				return tx.Migrator().CreateTable(
					&configinfrastructurerepositorymodel.Config{},
				)
			},
			Rollback: func(tx *gorm.DB) error {
				return nil
			},
		},
	})

	if err := m.Migrate(); err != nil {
		fmt.Printf("Migration failed: %v", err)
	}
	fmt.Println("Migration did run successfully")
}
