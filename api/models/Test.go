package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Test struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title       string `gorm:"size:255;not null;unique" json:"title"`
	Answer_code string `gorm:"not null" json:"Trueansw"`
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

func (p *Test) FindAllTests(db *gorm.DB) (*[]Test, error) {
	var err error
	tests := []Test{}
	err = db.Debug().Model(&Test{}).Limit(100).Find(&tests).Error
	if err != nil {
		return &[]Test{}, err
	}
	return &tests, err
}
