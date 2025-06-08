package clock

import "time"

type Clock struct{}

func (s *Clock) Now() time.Time {
	return time.Now()
}
