package spider

import "testing"

func TestSpider_Scan(t *testing.T) {
	if !testing.Short() {
		t.Skip("skip test. Run with -short")
	}

	const url = "https://go.dev"
	const depth = 1
	s := New()
	data, err := s.Scan(url, depth)
	if err != nil {
		t.Errorf("c.Scan(); err = %s; want nil", err)
		return
	}

	got := len(data)
	want := 1
	if got != want {
		t.Errorf("len(data) = %d; want %d", got, want)
	}
}
