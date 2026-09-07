package main

import (
	"testing"
	"time"
)

func TestDaysUntilNewYear(t *testing.T) {
	cases := []struct {
		name string
		now  time.Time
		want int
	}{
		{
			name: "start of year",
			now:  time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC),
			want: 365,
		},
		{
			name: "mid year",
			now:  time.Date(2026, time.July, 1, 0, 0, 0, 0, time.UTC),
			want: 184,
		},
		{
			name: "day before new year",
			now:  time.Date(2026, time.December, 31, 0, 0, 0, 0, time.UTC),
			want: 1,
		},
		{
			name: "leap year day",
			now:  time.Date(2024, time.February, 28, 0, 0, 0, 0, time.UTC),
			want: 308,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := daysUntilNewYear(c.now)
			if got != c.want {
				t.Errorf("daysUntilNewYear(%v) = %d, want %d", c.now, got, c.want)
			}
		})
	}
}
