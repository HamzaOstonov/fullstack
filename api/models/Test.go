package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Test struct {
	ID    uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:255;not null;unique" json:"title"`
}

func (p *Test) Prepare() {
	p.ID = 0
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
}

func (p *Test) Validate() error {

	if p.Title == "" {
		return errors.New("Required Title")
	}

	return nil
}

func (p *Post) SaveTest(db *gorm.DB) (*Post, error) {
	var err error
	err = db.Debug().Model(&Post{}).Create(&p).Error
	if err != nil {
		return &Post{}, err
	}
	if p.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", p.AuthorID).Take(&p.Author).Error
		if err != nil {
			return &Post{}, err
		}
	}
	return p, nil
}

func (p *Test) FindAllTests(db *gorm.DB) (*[]Test, error) {
	var err error
	tests := []Test{}
	err = db.Debug().Model(&Test{}).Limit(100).Find(&tests).Error
	if err != nil {
		return &[]Test{}, err
	}
	return &tests, err
}
