package retrievor

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// MatchesProba List of all matches retrieved
type MatchesProba struct {
	competition string
	Matches     []MatchProba
}

// MatchProba Data about a specific match
type MatchProba struct {
	Date          time.Time `json:"time"`
	MatchID       string    `json:"matchID"`
	CompetitionID string    `json:"competitionID"`
	T1Id          int       `json:"t1ID"`
	T2Id          int       `json:"t2ID"`
	T1Name        string    `json:"t1Name"`
	T2Name        string    `json:"t2Name"`
	T1Proba       int       `json:"t1Proba"`
	T2Proba       int       `json:"t2Proba"`
	DrawProba     int       `json:"drawProba"`
}

const url = "https://projects.fivethirtyeight.com/soccer-predictions/"

// Returns formatted url for the competition
func getMatchesURL(competition string) string {
	return fmt.Sprintf(url + competition)
}

func convertDateToTime(date string, now time.Time) (time.Time, error) {
	nowYear := now.Year()
	t := strings.Split(date, "/")
	if len(t) != 2 {
		return time.Time{}, errors.New("Input error")
	}

	tmpM, _ := strconv.Atoi(t[0])
	tmpD, _ := strconv.Atoi(t[1])
	m := fmt.Sprintf("%02d", tmpM)
	d := fmt.Sprintf("%02d", tmpD)
	if now.Month() == 12 && m == "01" {
		// If we are in december and upcoming match month is January, use current Year + 1
		nowYear++
	}
	s := strconv.Itoa(nowYear) + "-" + m + "-" + d
	r, err := time.Parse("2006-01-02", s)
	if err != nil {
		return time.Time{}, err
	}
	return r, nil
}

// ParsePageFTE Entrypoint of the module. Takes a string that represent the competition you want to retrieve
func (r *MatchesProba) ParsePageFTE(competition string) error {
	var url = getMatchesURL(competition)
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return err
	}
	r.competition = competition
	r.parseProbaFTE(doc)
	if err != nil {
		log.Println("Error during parsing")
		return err
	}
	log.Println(fmt.Sprintf("%s : Parsing done !", url))

	return nil
}

func (r *MatchesProba) parseProbaFTE(doc *goquery.Document) {
	doc.Find(".upcoming .initial-view table").Children().Each(func(i int, s *goquery.Selection) {
		var err error
		m := MatchProba{}
		d := s.Find(".date").Text()
		m.Date, err = convertDateToTime(d, time.Now())
		if err != nil {
			log.Println("Date parsing error: ", err)
		}
		m.CompetitionID = r.competition
		s.Children().Each(func(i int, s2 *goquery.Selection) {
			a, e := s2.Attr("class")
			if e {
				if a == "match-top" {
					if err != nil {
						log.Println(err)
						return
					}
					m.T1Name = s2.Find(".team .team-div .name").Text()
					m.T1Id = getTeamIDFromMap(m.T1Name)
					proba := strings.Split(s2.Find(".prob").Text(), "%")
					m.T1Proba, err = strconv.Atoi(proba[0])
					if err != nil {
						log.Println(err)
						return
					}
					m.DrawProba, err = strconv.Atoi(proba[1])
					if err != nil {
						log.Println(err)
						return
					}
				} else if a == "match-bottom" {
					m.T2Name = s2.Find(".team .team-div .name").Text()
					m.T2Id = getTeamIDFromMap(m.T2Name)
					proba := strings.Split(s2.Find(".prob").Text(), "%")
					m.T2Proba, _ = strconv.Atoi(proba[0])
				}
			}
		})
		m.setMatchID()
		r.Matches = append(r.Matches, m)
	})
	fmt.Println(r)
}

func (m *MatchProba) setMatchID() {
	t0 := time.Time{}
	if m.T1Id == 0 || m.T2Id == 0 || m.Date == t0 {
		log.Println("One parameter missing to build match ID for match", m)
		return
	}
	y := m.Date.Year()
	mo := m.Date.Month()
	d := m.Date.Day()
	m.MatchID = fmt.Sprintf("%v-%02d-%02d-%d-%d", y, mo, d, m.T1Id, m.T2Id)
}

func getTeamIDFromMap(n string) int {
	for _, v := range teamMap {
		if v.CompetitorFTEName == n {
			return v.CompetitorID
		}
	}
	return 0
}
