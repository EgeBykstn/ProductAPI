package model

import (
	"sync"
	"time"
)

type Product struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Price     int       `json:"price"`
	Color     string    `json:"color"`
	Size      int       `json:"size"`
	UpdatedAt time.Time `json:"-"`
	CreatedAt time.Time
}
type AutoInc struct {
	sync.Mutex
	id int
}
type Filter struct {
	Code string `form:"code" json:"code,omitempty"`
	Name string `form:"name" json:"name,omitempty"`
}

func (s *AutoInc) ID() (id int) {
	s.Lock()
	defer s.Unlock()

	id = s.id
	s.id++
	return id
}
