package base

//type TimeStamp int64

//func (d TimeStamp) MarshalJSON() ([]byte, error) {
//	rs := time.Unix(int64(d), 0).Format("2006-01-02 11:22")
//	js, er := json.Marshal(rs)
//	return js, er
//}
//
//func (d *TimeStamp) UnmarshalJSON(data []byte) error {
//	var rs string
//	e := json.Unmarshal(data, &rs)
//	if e != nil {
//		return e
//	}
//	t, er := time.Parse("2006-01-02 11:22", rs)
//	if er != nil {
//		return er
//	}
//	*d = TimeStamp(t.Unix())
//	return nil
//}

type CommonModel struct {
	ID        uint  `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at,omitempty"`
	UpdatedAt int64 `json:"updated_at,omitempty"`
	Deleted   int64 `json:"-" gorm:"default:0"`
}
