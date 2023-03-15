package middleware

import (
	"context"
	"database/sql"
	"github.com/pelletier/go-toml"
	"log"
	"strconv"
)

var db = database{}
var Ctx = context.Background()

func InitDatabase() error {
	var dbCfg dbConfig
	err := dbCfg.init()
	if err != nil {
		return err
	}
	err = db.connect(dbCfg.toString())
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
	return nil
}

func (c *dbConfig) toString() string {
	return "postgres://" + c.user + ":" + c.pass + "@" + c.host + ":" + strconv.FormatInt(c.port, 10) + "/" + c.dbname + "?sslmode=disable"
}

type database struct {
	DbInstance *sql.DB
}

func (d *database) connect(connectionString string) error {
	database, err := sql.Open("postgres", connectionString)
	if err != nil {
		return err
	}
	go func() {
		err := database.Ping()
		if err != nil {
			log.Panic(err)
		}
	}()
	d.DbInstance = database
	return nil
}

func (d *database) isConnected() bool {
	if err := d.DbInstance.Ping(); err != nil {
		return false
	}
	return true
}
