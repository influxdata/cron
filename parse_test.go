package cron

import (
	"time"
)

func ceilTime(t time.Time, d time.Duration) time.Time {
	if t.Truncate(d).Equal(t) {
		return t
	}
	return t.Add(d).Truncate(d)
}

// func Test_parse(t *testing.T) {
// 	type args struct {
// 	}
// 	tests := []struct {
// 		name    string
// 		s  string
// 		tz *time.Location
// 		want    nextTime
// 		wantErr bool
// 	}{
// 		{name: "all",  s: "* * * * *"}
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := parse(tt.args.s, tt.args.tz)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("parse() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
