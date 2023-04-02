package courses

import (
	"courses/app/model"
	"courses/app/repository/courses"
)

type CoursesUsecase interface {
	Insert(payload model.PayloadCourse) (model.Course, error)
	Update(payload model.PayloadCourse, id int) (model.Course, error)
	GetByID(id int) (model.Course, error)
	GetList(params model.ParamsCourse) ([]model.Course, error)
	Total() (int, int, error)

	GetCategoryCourse() ([]model.Category, error)
}

type coursesUsecase struct {
	courses courses.CoursesRepo
}

func NewCoursesUsecase(courses courses.CoursesRepo) CoursesUsecase {
	return &coursesUsecase{
		courses: courses,
	}
}

func (c *coursesUsecase) Insert(payload model.PayloadCourse) (model.Course, error) {
	res := model.Course{}
	category, err := c.courses.GetCategoryID(payload.CategoryID)
	if err != nil {
		return res, err
	}

	res.Name = payload.Name
	res.Price = payload.Price
	res.Category = category
	res.ImageURL = payload.Image

	err = c.courses.Insert(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c *coursesUsecase) Update(payload model.PayloadCourse, id int) (model.Course, error) {
	res, err := c.courses.GetCourseID(id)
	if err != nil {
		return res, err
	}

	category, err := c.courses.GetCategoryID(payload.CategoryID)
	if err != nil {
		return res, err
	}

	res.Name = payload.Name
	res.Price = payload.Price
	res.Category = category

	err = c.courses.Update(&res)
	if err != nil {
		return res, err
	}

	return res, err
}

func (c *coursesUsecase) GetByID(id int) (model.Course, error) {
	return c.courses.GetCourseID(id)
}

func (c *coursesUsecase) GetList(params model.ParamsCourse) ([]model.Course, error) {
	res := []model.Course{}
	data, err := c.courses.GetCourses(params)
	if err != nil {
		return res, nil
	}

	for _, d := range data {
		category, _ := c.courses.GetCategoryID(d.CategoryID)
		course := model.Course{
			ID:       d.ID,
			Name:     d.Name,
			Price:    d.Price,
			ImageURL: d.ImageURL,
			Category: category,
		}

		res = append(res, course)
	}
	return res, nil
}

func (c *coursesUsecase) Total() (int, int, error) {
	courses, err := c.courses.Total()
	if err != nil {
		return 0, 0, err
	}

	free, err := c.courses.TotalFree()
	if err != nil {
		return 0, 0, err
	}

	return courses, free, nil
}

func (c *coursesUsecase) GetCategoryCourse() ([]model.Category, error) {
	return c.courses.GetCategory()
}
