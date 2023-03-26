package database

import (
	"database/sql"
	"encoding/json"
	"github.com/pelletier/go-toml"
	"golang_server/gorm_models/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"strconv"
	"time"
)

var Db = Database{}

func InitDatabase() error {
	var dbCfg dbConfig
	err := dbCfg.init()
	if err != nil {
		return err
	}
	connectionStr := dbCfg.toString()
	log.Printf("Connection to db: %s", connectionStr)
	err = connectAndPing(connectionStr)
	if err != nil {
		return err
	}
	err = Db.connect(connectionStr, os.Getenv("DB_HOST") != "")
	if err != nil {
		return err
	}
	return nil
}

type dbConfig struct {
	host   string
	port   int64
	user   string
	pass   string
	dbname string
}

func (c *dbConfig) init() error {
	if dbHost := os.Getenv("DB_HOST"); dbHost == "" {
		// no environment configuration specified so using fallback
		config, err := toml.LoadFile("sqlboiler.toml")
		if err != nil {
			return err
		}
		dbTree := config.Get("psql").(*toml.Tree)
		c.host = dbTree.Get("host").(string)
		c.port = dbTree.Get("port").(int64)
		c.user = dbTree.Get("user").(string)
		c.pass = dbTree.Get("pass").(string)
		c.dbname = dbTree.Get("dbname").(string)
	} else {
		// environment configuration so using them
		c.host = dbHost
		c.port, _ = strconv.ParseInt(os.Getenv("DB_PORT"), 10, 64)
		c.user = os.Getenv("DB_USER")
		c.pass = os.Getenv("DB_PASSWORD")
		c.dbname = os.Getenv("DB_NAME")
	}

	return nil
}

func (c *dbConfig) toString() string {
	return "postgres://" + c.user + ":" + c.pass + "@" + c.host + ":" + strconv.FormatInt(c.port, 10) + "/" + c.dbname + "?sslmode=disable"
}

type Database struct {
	Instance *gorm.DB
}

func (d *Database) connect(connectionString string, makeMigrationsAndSeeding bool) error {
	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionString,
	}), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		return err
	}
	confDb, err := gormDb.DB()
	if err != nil {
		return err
	}
	confDb.SetMaxIdleConns(10)
	confDb.SetMaxOpenConns(100)
	confDb.SetConnMaxLifetime(time.Hour)
	d.Instance = gormDb
	if makeMigrationsAndSeeding {
		gormDb.AutoMigrate(&model.Event{})
		gormDb.AutoMigrate(&model.News{})

		gormDb.AutoMigrate(&model.Image{})
		gormDb.AutoMigrate(&model.LocationImage{})

		gormDb.AutoMigrate(&model.Training{})
		gormDb.AutoMigrate(&model.TrainingDay{})
		gormDb.AutoMigrate(&model.TrainingTrainer{})

		gormDb.AutoMigrate(&model.ServerConfig{})
		gormDb.AutoMigrate(&model.ClientConfig{})

		gormDb.AutoMigrate(&model.AuthenticationToken{})
		gormDb.AutoMigrate(&model.PushToken{})

		gormDb.AutoMigrate(&model.Todo{})
		gormDb.AutoMigrate(&model.Registration{})
		gormDb.AutoMigrate(&model.Feedback{})
		gormDb.AutoMigrate(&model.Location{})
		gormDb.AutoMigrate(&model.Trainer{})
		gormDb.AutoMigrate(&model.Credit{})

		gormDb.AutoMigrate(&model.NewsletterParticipation{})

		gormDb.AutoMigrate(&model.User{})
		gormDb.AutoMigrate(&model.UserCredit{})
		gormDb.AutoMigrate(&model.UserEvent{})
		gormDb.AutoMigrate(&model.UserTraining{})
		gormDb.AutoMigrate(&model.UserTrainingInternal{})

		userSeeds, _ := os.ReadFile("seeds/users.json")
		imageSeeds, _ := os.ReadFile("seeds/images.json")
		var uData []model.User
		var iData []model.Image
		_ = json.Unmarshal(userSeeds, &uData)
		_ = json.Unmarshal(imageSeeds, &iData)
		gormDb.Create(uData)
		gormDb.Create(iData)
	}
	return nil
}

func connectAndPing(connectionString string) error {
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	defer database.Close()
	err = database.Ping()
	if err != nil {
		return err
	}
	return nil
}
