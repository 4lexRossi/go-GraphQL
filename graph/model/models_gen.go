// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Courses     []*Course `json:"courses"`
}

type Course struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description,omitempty"`
	Category    *Category `json:"category"`
}

type NewCategory struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
}

type NewCourse struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	CategoryID    string  `json:"categoryId"`
}
