package courses

import (
	"courses/app/model"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type CoursesRepo interface {
	Insert(payload *model.Course) error
	Update(payload *model.Course) error
	GetCourseID(id int) (model.Course, error)
	GetCourses(params model.ParamsCourse) ([]model.Course, error)
	Total() (int, error)
	TotalFree() (int, error)

	GetCategoryID(id int) (model.Category, error)
	GetCategory() ([]model.Category, error)
}

type coursesRepo struct {
	db *gorm.DB
}

func NewCoursesRepo(db *gorm.DB) CoursesRepo {
	return &coursesRepo{
		db: db,
	}
}

func (c *coursesRepo) GetCategoryID(id int) (model.Category, error) {
	res := model.Category{}
	query := c.db.Raw("SELECT id, name, created_at FROM category WHERE id = ?", id).Scan(&res)
	if query.Error != nil {
		return res, query.Error
	}

	return res, nil
}

func (c *coursesRepo) GetCategory() ([]model.Category, error) {
	res := []model.Category{}
	query := c.db.Raw("SELECT id, name, created_at FROM category").Scan(&res)
	if query.Error != nil {
		return res, query.Error
	}

	return res, nil
}

func (c *coursesRepo) Insert(payload *model.Course) error {
	row, err := c.db.DB().Exec("INSERT INTO courses(category_id,name,price,image_url,created_at) VALUES (?,?,?,?,?)", payload.Category.ID, payload.Name, payload.Price, payload.ImageURL, time.Now())
	if err != nil {
		return err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return err
	}

	payload.ID = int(id)

	return nil
}

func (c *coursesRepo) Update(payload *model.Course) error {
	query := c.db.Exec("UPDATE courses SET category_id = ?, name = ?, price =? WHERE id = ?", payload.Category.ID, payload.Name, payload.Price, payload.ID)
	if query.Error != nil {
		return query.Error
	}

	return nil
}

func (c *coursesRepo) GetCourseID(id int) (model.Course, error) {
	res := model.Course{}
	query := c.db.Raw("SELECT id, category_id, name, price, qty, created_at FROM courses WHERE id = ?", id).Scan(&res)
	if query.Error != nil {
		return res, query.Error
	}

	return res, nil
}

func (c *coursesRepo) GetCourses(params model.ParamsCourse) ([]model.Course, error) {
	limit := 10
	if params.Limit > 0 {
		limit = params.Limit
	}

	offset := 0
	if params.Page > 0 {
		offset = (params.Page - 1) * limit
	}

	queryLimit := fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset)

	queryParam := ""

	if params.Name != "" && queryParam == "" {
		queryParam = fmt.Sprintf("%s%s%s%s", queryParam, " WHERE courses.name = '", params.Name, "'")
	} else if params.Name != "" && queryParam != "" {
		queryParam = fmt.Sprintf("%s%s%s%s", queryParam, " AND courses.name = '", params.Name, "'")
	}

	if params.Price > 0 && queryParam == "" {
		queryParam = fmt.Sprintf("%s%s%d", queryParam, " WHERE courses.price = ", params.Price)
	} else if params.Price > 0 && queryParam != "" {
		queryParam = fmt.Sprintf("%s%s%d", queryParam, " AND courses.price = ", params.Price)
	}

	if params.Category != "" && queryParam == "" {
		queryParam = fmt.Sprintf("%s%s%s%s", queryParam, " WHERE category.name = '", params.Category, "'")
	} else if params.Category != "" && queryParam != "" {
		queryParam = fmt.Sprintf("%s%s%s%s", queryParam, " AND category.name = '", params.Category, "'")
	}

	queryOrder := " ORDER BY courses.id DESC"
	if len(params.SortBy) == 2 {
		queryOrder = fmt.Sprintf(" ORDER BY courses.%s %s", params.SortBy[0], params.SortBy[1])
	}

	query := "SELECT * FROM courses " + queryParam + queryOrder + queryLimit

	log.Print(query)
	res := []model.Course{}
	q := c.db.Raw(query).Scan(&res)
	if q.Error != nil {
		return res, q.Error
	}

	return res, nil
}

func (c *coursesRepo) Total() (int, error) {
	total := 0
	err := c.db.DB().QueryRow("SELECT count(*) as total FROM courses").Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}

func (c *coursesRepo) TotalFree() (int, error) {
	total := 0
	err := c.db.DB().QueryRow("SELECT count(*) FROM courses WHERE price = ?", 0).Scan(&total)
	if err != nil {
		return total, err
	}

	return total, nil
}
