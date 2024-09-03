package is

import "time"

func Time(s string, layout string) bool {
	_, err := time.Parse(layout, s)
	return err == nil
}

func Duration(s string) bool {
	_, err := time.ParseDuration(s)
	return err == nil
}

func Timezone(s string) bool {
	_, err := time.LoadLocation(s)
	return err == nil
}
