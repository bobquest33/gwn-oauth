package main

import (
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/helderfarias/gwn-oauth/endpoint"
	. "github.com/helderfarias/gwn-oauth/log"
	"github.com/helderfarias/gwn-oauth/middleware"
	"github.com/helderfarias/gwn-oauth/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

func main() {
	Logger.Info("Lendo configurações de sistema")
	config := loadConfig()

	Logger.Info("Estabelecendo conexão com banco de dados")
	db := sqlx.MustOpen("postgres", config.Database.Datasource)
	defer db.Close()

	Logger.Info("Criando pool de conexões: Min: %d, Max: %d", config.Database.Pool.Min, config.Database.Pool.Max)
	db.Ping()
	db.SetMaxIdleConns(config.Database.Pool.Min)
	db.SetMaxOpenConns(config.Database.Pool.Max)

	Logger.Info("Registrando middlewares")
	router := gin.Default()
	router.Use(middleware.DataBase(db))

	Logger.Info("Registrando endpoints")
	endpoint.RegisterEndpoints(router)
	router.Run(":8003")
}

func loadConfig() *util.ResourceConfig {
	config := util.GetOpt("CONFIG_FILE", "application.yml")

	data, err := ioutil.ReadFile(config)
	if err != nil {
		panic(err)
	}

	var cfg util.ResourceConfig
	err = yaml.Unmarshal([]byte(data), &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
