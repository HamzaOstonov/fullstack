package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Testvar struct {
	ID          uint64 `gorm:"primary_key;auto_increment" json:"id"`
	IdxID       uint64 `gorm:"size:255;not null;unique" json:"idx_id"`
	VariantText string `gorm:"size:255;not null;unique" json:"variant_text"`
	VariantCode string `gorm:"size:255;not null;unique" json:"variant_code"`
}

func (p *Testvar) Prepare() {
	p.ID = 0
	p.VariantText = html.EscapeString(strings.TrimSpace(p.VariantText))
}

func (p *Testvar) Validate() error {

	if p.VariantText == "" {
		return errors.New("Required VariantText")
	}

	return nil
}

func (p *Testvar) FindAllTestvars(db *gorm.DB) (*[]Testvar, error) {
	var err error
	tests := []Testvar{}
	err = db.Debug().Model(&Testvar{}).Limit(100).Find(&tests).Error
	if err != nil {
		return &[]Testvar{}, err
	}
	return &tests, err
}

func (p *Testvar) Find2(db *gorm.DB, hh int) (*[]Testvar, error) {
	var err error
	posts := []Testvar{}
	//err = db.Debug().Model(&Post{}).Limit(100).Find(&posts).Error
	err = db.Debug().Model(&Testvar{}).Where("id = ?", hh).Take(&Testvar{}).Error

	if err != nil {
		return &[]Testvar{}, err
	}
	return &posts, err

}
