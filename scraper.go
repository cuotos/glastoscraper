package glastoscraper

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

const (
	urlStringFormat = "https://www.glastonburyfestivals.co.uk/line-up/line-up-%d/?artist"
)

type scraper struct {
	url *url.URL
}

func New(year int) (*scraper, error) {
	s := &scraper{}

	createdUrlString := fmt.Sprintf(urlStringFormat, year)
	parsedURL, err := url.Parse(createdUrlString)
	if err != nil {
		return nil, err
	}

	s.url = parsedURL

	return s, nil
}

func (s *scraper) GetAllPerformances() ([]Artist, error) {
	artists := []Artist{}

	c := colly.NewCollector(colly.AllowURLRevisit())

	c.OnHTML("#main > div.col_5.lineup > div.inner > ul > li", func(h *colly.HTMLElement) {

		artists = append(artists, Artist{
			Title:  strings.TrimSpace(h.ChildText(".title")),
			Stage:  strings.TrimSpace(h.ChildText(".stage")),
			Day:    ParseDay(strings.TrimSpace(h.ChildText(".day"))),
			DayRaw: strings.TrimSpace(h.ChildText(".day")),
			Time:   strings.TrimSpace(h.ChildText(".end")),
		})
	})

	err := c.Visit(s.url.String())
	if err != nil {
		return artists, err
	}

	return artists, nil
}

func (s *scraper) GetAllArtists() ([]string, error) {
	artistNamesOnly := []string{}

	allPerformances, err := s.GetAllPerformances()
	if err != nil {
		return artistNamesOnly, err
	}

	tempMap := make(map[string]any)

	for _, p := range allPerformances {
		tempMap[p.Title] = struct{}{}
	}

	for k := range tempMap {
		if k != "" {
			artistNamesOnly = append(artistNamesOnly, k)
		}
	}

	return artistNamesOnly, nil
}
