package storages

import (
	"github.com/pkg/errors"
	"github.com/thewizardplusplus/go-exercises-backend/entities"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OpenDB ...
func OpenDB(dbDSN string, logWriter logger.Writer) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{
		Logger: logger.New(logWriter, logger.Config{
			SlowThreshold: 0, // disallow analysis of query execution speed
			Colorful:      false,
			LogLevel:      logger.Info,
		}),
	})
	if err != nil {
		return nil, errors.Wrap(err, "unable to open the DB")
	}

	if err := db.AutoMigrate(
		&entities.Task{},
		&entities.Solution{},
		&entities.User{},
	); err != nil {
		return nil, errors.Wrap(err, "unable to migrate the entities automatically")
	}

	return db, nil
}

// CloseDB ...
func CloseDB(db *gorm.DB) error {
	innerDB, err := db.DB()
	if err != nil {
		return errors.Wrap(err, "unable to get the inner DB")
	}

	if err := innerDB.Close(); err != nil {
		return errors.Wrap(err, "unable to close the inner DB")
	}

	return nil
}
