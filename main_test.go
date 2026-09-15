package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
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

func TestDaysHandler_Integration(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(daysHandler))
	defer server.Close()

	t.Run("without date param", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/days")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200, got %d", resp.StatusCode)
		}

		var body response
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if body.Days < 0 || body.Days > 366 {
			t.Errorf("unexpected days value: %d", body.Days)
		}
	})

	t.Run("with valid date param", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/days?date=2026-07-01")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("expected status 200, got %d", resp.StatusCode)
		}

		var body response
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if body.Days != 184 {
			t.Errorf("expected 184 days, got %d", body.Days)
		}
		if body.Date != "2026-07-01" {
			t.Errorf("expected date 2026-07-01, got %s", body.Date)
		}
	})

	t.Run("with invalid date param", func(t *testing.T) {
		resp, err := http.Get(server.URL + "/days?date=not-a-date")
		if err != nil {
			t.Fatalf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusBadRequest {
			t.Fatalf("expected status 400, got %d", resp.StatusCode)
		}

		var body errorResponse
		if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
			t.Fatalf("failed to decode response: %v", err)
		}

		if body.Error == "" {
			t.Error("expected error message, got empty string")
		}
	})
}
