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
	Matches []MatchProba
}

// MatchProba Data about a specific match
type MatchProba struct {
	Date            time.Time `json:"time"`
	CompetitionID   string    `json:"competitionID"`
	CompetitionName string    `json:"competitionName"`
	T1Id            string    `json:"t1ID"`
	T2Id            string    `json:"t2ID"`
	T1Name          string    `json:"t1Name"`
	T2Name          string    `json:"t2Name"`
	T1Proba         int       `json:"t1Proba"`
	T2Proba         int       `json:"t2Proba"`
	DrawProba       int       `json:"drawProba"`
}

const url = "https://projects.fivethirtyeight.com/soccer-predictions/"

// Returns formatted url for the competition
func getMatchesURL(competition string) string {
	return fmt.Sprintf(url + competition)
}

func convertDate(date string, now time.Time) (time.Time, error) {
	nowYear := now.Year()
	t := strings.Split(date, "/")
	if len(t) != 2 {
		return time.Time{}, errors.New("Input error")
	}
	d := t[0]
	m := t[1]
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
	r.parseProbaFTE(doc)
	log.Println(fmt.Sprintf("%s : Parsing done !", url))

	return nil
}

func (r *MatchesProba) parseProbaFTE(doc *goquery.Document) {
	// var competitionID string
	// var competitionName string
	doc.Find(".upcoming .initial-view table").Children().Each(func(i int, s *goquery.Selection) {
		m := MatchProba{}
		s.Children().Each(func(i int, s2 *goquery.Selection) {
			a, e := s2.Attr("class")
			if e {
				var err error
				if a == "match-top" {
					d := s.Find(".date").Text()
					m.Date, err = convertDate(d, time.Now())
					if err != nil {
						log.Println(err)
						return
					}
					m.T1Name = s2.Find(".team .team-div .name").Text()
					// TODO: Match team name with Winamax team id
					m.T1Id = m.T1Name
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
					m.T2Id = s2.Find(".team .team-div .name").Text()
					m.T2Name = m.T2Id
					proba := strings.Split(s2.Find(".prob").Text(), "%")
					m.T2Proba, _ = strconv.Atoi(proba[0])
					r.Matches = append(r.Matches, m)
				}
			}
		})
		fmt.Println(m)
	})
	fmt.Println(r)
}
