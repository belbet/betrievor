package retrievor

import (
	"testing"
	"time"
)

func TestConvertDate1(t *testing.T) {
	i := "28/11"
	e, _ := time.Parse("2006-01-02", "2020-11-28")
	r, err := convertDate(i, time.Now())
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDate2(t *testing.T) {
	i := "28/12"
	e, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDate(i, time.Now())
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDateNextYear(t *testing.T) {
	i := "01/01"
	e, _ := time.Parse("2006-01-02", "2021-01-01")
	now, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDate(i, now)
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDateErrorInput1(t *testing.T) {
	i := "0/101"
	e := time.Time{}
	now, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDate(i, now)
	if err == nil {
		t.Errorf("Expected an error, got: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDateErrorInput2(t *testing.T) {
	i := "0101"
	e := time.Time{}
	now, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDate(i, now)
	if err == nil {
		t.Errorf("Expected an error, got: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}
