package entity

type Day struct {
	ID       uint   `json:"id" gorm:"autoIncreament;primaryKey"`
	StudioID uint   `json:"-"`
	VenueID  uint   `json:"-"`
	Day      string `json:"day"`
	Time     []Time `json:"time" goem:"foreignKey:DayID"`
}

type Time struct {
	DayID  uint   `json:"-"`
	Status bool   `json:"status"`
	Time   string `json:"time"`
}
