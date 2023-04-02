package usecase

import (
	"courses/app/repository"
	"courses/app/usecase/courses"
	"courses/app/usecase/users"
)

type Usecase struct {
	User   users.UsersUsecase
	Course courses.CoursesUsecase
}

func Init(repository *repository.Repository) *Usecase {
	uc := &Usecase{
		User: users.NewUsersUsecase(
			repository.User,
		),
		Course: courses.NewCoursesUsecase(
			repository.Course,
		),
	}
	return uc
}
