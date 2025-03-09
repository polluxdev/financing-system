package domain

type Tenor struct {
	ID    string `db:"id"`
	Value uint8  `db:"value"`
}

func (Tenor) TableName() string {
	return "tenors"
}
