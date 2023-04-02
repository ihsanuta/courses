package model

type Course struct {
	ID         int      `json:"id"`
	Name       string   `json:"name"`
	ImageURL   string   `json:"image_url"`
	Price      int      `json:"price"`
	Qty        int      `json:"-"`
	CategoryID int      `json:"-"`
	Category   Category `json:"category"`
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}

type PayloadCourse struct {
	Name       string `json:"name" form:"name"`
	Price      int    `json:"price" form:"price"`
	CategoryID int    `json:"category_id" form:"category_id"`
	Image      string `json:"image" form:"image"`
}

type ParamsCourse struct {
	Name     string   `json:"name" form:"name"`
	Price    int      `json:"price" form:"price"`
	Category string   `json:"category" form:"category"`
	SortBy   []string `json:"sort_by" form:"sortBy"`
	Limit    int      `json:"limit" form:"limit"`
	Page     int      `json:"page" form:"page"`
}

type Total struct {
	TotalUser        int `json:"total_users"`
	TotalCourses     int `json:"total_courses"`
	TotalCOursesFree int `json:"total_courses_free"`
}
