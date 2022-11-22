package CustomModel

import (
	"fmt"
	"time"
)

type CustomDate struct {
	time.Time
}

func (t *CustomDate) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
func (t *CustomDate) ToDateString() string {
	y, m, d := t.Date()
	return fmt.Sprintf("%v-%v-%v", y, int(m), d)
}
