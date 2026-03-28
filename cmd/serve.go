package cmd

import (
	"fmt"
	"os"
	"sysagent/config"
	"sysagent/infra/db"
	"sysagent/repo"
	"sysagent/rest"
	productHandler "sysagent/rest/handlers/product"
	userHandler "sysagent/rest/handlers/user"
	middleware "sysagent/rest/middlewares"
	"sysagent/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	middlewares := middleware.NewMiddlewares(cnf)

	//repos
	productRepo := repo.NewProductRepo(dbCon)
	userRepo := repo.NewUserRepo(dbCon)

	//domains
	usrSvc := user.NewService(userRepo)

	productHandler := productHandler.NewHandler(middlewares, productRepo)
	userHandler := userHandler.NewHandler(cnf, usrSvc)

	server := rest.NewServer(
		cnf,
		productHandler,
		userHandler,
	)
	server.Start()
}
