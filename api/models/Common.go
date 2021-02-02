package models

type Answer struct {
	VariantText string `gorm:"size:255;not null;unique" json:"answer"`
}

type Question struct {
	ID       uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Text     string    `gorm:"size:255;not null;unique" json:"text"`
	Type     string    `gorm:"size:255;not null;unique" json:"type"`
	Answers  []*Answer `json:"answers"`
	Javoblar []*string `json:"javoblar"`
	Javob    []string  `json:"javob"`
}

type General struct {
	Title     string      `gorm:"size:255;not null;unique" json:"title"`
	Questions []*Question `json:"questions"`
}
