package database

import (
	"github.com/almeidacavalcante/go-encoder/domain"
	gorm "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
	"log"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.Debug = true
	dbInstance.AutoMigrateDb = true

	connection, err := dbInstance.Connect()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	return connection
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env == "test" {
		d.Db, err = gorm.Open(d.DbTypeTest, d.DsnTest)
	} else {
		d.Db, err = gorm.Open(d.DbTypeTest, d.Dsn)
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.LogMode(true)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
		d.Db.Model(domain.Job{}).AddForeignKey("video_id", "videos(id)", "CASCADE", "CASCADE")

	}

	return d.Db, nil
}
