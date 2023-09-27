package processors

import (
	"bytes"
	"fmt"
	"strconv"
	"time"
)

// Date encode string to timestamp.
type Date struct{}

var timestampFlag = Flag{
	Name:  "timestamp",
	Short: "t",
	Desc:  "",
	Value: false,
	Type:  FlagBool,
}

var NOW = []byte("now")

func checkTimestampFlag(f []Flag) bool {
	timestamp := false
	for _, flag := range f {
		if flag.Short == "t" {
			r, ok := flag.Value.(bool)
			if ok {
				timestamp = r
			}
		}
	}
	return timestamp
}

func (p Date) Name() string {
	return "date"
}

func (p Date) Alias() []string {
	return nil
}

func (p Date) Transform(data []byte, f ...Flag) (string, error) {
	if checkTimestampFlag(f) {
		i, err := strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return "", err
		}
		if i > 1000000000000 {
			return time.Unix(0, i*int64(time.Millisecond)).Format(time.DateTime), nil
		}
		return time.Unix(i, 0).Format(time.DateTime), nil
	}
	var t time.Time
	if bytes.Equal(data, NOW) {
		t = time.Now()
	} else {
		var err error
		t, err = time.Parse(time.DateTime, string(data))
		if err != nil {
			return "", err
		}
	}
	return strconv.FormatInt(t.Unix(), 10), nil
}

func (p Date) Flags() []Flag {
	return []Flag{timestampFlag}
}

func (p Date) Title() string {
	title := "date"
	return fmt.Sprintf("%s (%s)", title, p.Name())
}

func (p Date) Description() string {
	return "date"
}

func (p Date) FilterValue() string {
	return p.Title()
}
