package models

type Answer struct {
	VariantText string `gorm:"size:255;not null;unique" json:"answer"`
}

type Question struct {
	ID   uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Text string `gorm:"size:255;not null;unique" json:"text"`
	Type string `gorm:"size:255;not null;unique" json:"type"`
	//Answers  []*Answer `json:"answers"`
	Answers []string `json:"answers"`
	//Javob    []string  `json:"javob"`
	Answer string `json:"answer"`
}

type General struct {
	Title     string     `gorm:"size:255;not null;unique" json:"title"`
	Questions []Question `json:"questions"`
}

type Javob struct {
	Javobnum  string
	Javobtext string
	Javobsign string
}

type Savol struct {
	Savolnum   string
	Savoltext  string
	Javoblar   []Javob
	Jav1       string
	Tugrijavob string
}

type Kattatest struct {
	Mavzu    string
	Savollar []Savol
}
