package repository

import (
	"courses/app/repository/courses"
	"courses/app/repository/users"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	User   users.UsersRepo
	Course courses.CoursesRepo
}

func Init(db *gorm.DB) *Repository {
	repo := &Repository{
		User: users.NewUsersRepo(
			db,
		),
		Course: courses.NewCoursesRepo(
			db,
		),
	}
	return repo
}
