package main

import (
	"courses/app/repository"
	"courses/app/usecase"
	"courses/handler"
	"courses/helper/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// Business Layer
	repo *repository.Repository
	uc   *usecase.Usecase

	h handler.Handler
)

func main() {
	// konek to mysql
	db := mysql.GetMysqlConnection()

	// Business layer Initialization
	repo = repository.Init(
		db,
	)
	uc = usecase.Init(repo)
	handler.Init(uc)
}
