package retrievor

import (
	"testing"
	"time"
)

func TestConvertDate1(t *testing.T) {
	i := "11/28"
	e, _ := time.Parse("2006-01-02", "2020-11-28")
	r, err := convertDateToTime(i, time.Now())
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDate2(t *testing.T) {
	i := "12/28"
	e, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDateToTime(i, time.Now())
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDate3(t *testing.T) {
	i := "10/2"
	e, _ := time.Parse("2006-01-02", "2020-10-02")
	r, err := convertDateToTime(i, time.Now())
	if err != nil {
		t.Errorf("Got an unexpected error: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestConvertDateNextYear(t *testing.T) {
	i := "1/1"
	e, _ := time.Parse("2006-01-02", "2021-01-01")
	now, _ := time.Parse("2006-01-02", "2020-12-28")
	r, err := convertDateToTime(i, now)
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
	r, err := convertDateToTime(i, now)
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
	r, err := convertDateToTime(i, now)
	if err == nil {
		t.Errorf("Expected an error, got: %v", err)
	}
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestGetCompetitorIDFromMap1(t *testing.T) {
	i := "Marseille"
	e := 1641
	r := getTeamIDFromMap(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestGetCompetitorIDFromMap2(t *testing.T) {
	i := "St Étienne"
	e := 1678
	r := getTeamIDFromMap(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestGetCompetitorIDFromMap3(t *testing.T) {
	i := "Nîmes"
	e := 1663
	r := getTeamIDFromMap(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestGetCompetitorIDFromMapNone(t *testing.T) {
	i := "IDontExists"
	e := 0
	r := getTeamIDFromMap(i)
	if r != e {
		t.Errorf("Expected %v, got %v", e, r)
	}
}

func TestSetMatchID1(t *testing.T) {
	dateString := "2020-11-29"
	d, _ := time.Parse("2006-01-02", dateString)
	i := MatchProba{T1Id: 1649, T2Id: 1682, Date: d}
	e := "2020-11-29-1649-1682"
	i.setMatchID()
	if i.MatchID != e {
		t.Errorf("Expected %v, got %v", e, i.MatchID)
	}
}

func TestSetMatchID2(t *testing.T) {
	dateString := "2021-01-02"
	d, _ := time.Parse("2006-01-02", dateString)
	i := MatchProba{T1Id: 1648, T2Id: 1684, Date: d}
	e := "2021-01-02-1648-1684"
	i.setMatchID()
	if i.MatchID != e {
		t.Errorf("Expected %v, got %v", e, i.MatchID)
	}
}

func TestSetMatchIDError(t *testing.T) {
	dateString := "2021-01-02"
	d, _ := time.Parse("2006-01-02", dateString)
	i := MatchProba{T1Id: 0, T2Id: 1684, Date: d}
	e := ""
	i.setMatchID()
	if i.MatchID != e {
		t.Errorf("Expected %v, got %v", e, i.MatchID)
	}
}
