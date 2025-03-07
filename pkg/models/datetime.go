package models

import (
	"fmt"
	"strings"
	"time"
)

const layout = "2006-01-02T15:04:05Z0700"

type Datetime time.Time

func (v *Datetime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)

	t, err := time.Parse(layout, s)
	if err != nil {
		return err
	}

	*v = Datetime(t)
	return nil
}

func (v Datetime) MarshalJSON() ([]byte, error) {
	s := time.Time(v).Format(layout)
	s = fmt.Sprintf(`"%s"`, s)
	return []byte(s), nil
}
