package newsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type SourceCountry string

const (
	CountryAU SourceCountry = "au"
	CountryDE SourceCountry = "de"
	CountryGB SourceCountry = "gb"
	CountryIN SourceCountry = "in"
	CountryIT SourceCountry = "it"
	CountryUS SourceCountry = "us"
)

type SourceLanguage string

const (
	LangEN SourceLanguage = "en"
	LangDE SourceLanguage = "de"
	LangFR SourceLanguage = "fr"
)

type SourceCategory string

const (
	CatBusiness         SourceCategory = "business"
	CatEntertainment    SourceCategory = "entertainment"
	CatGaming           SourceCategory = "gaming"
	CatGeneral          SourceCategory = "general"
	CatMusic            SourceCategory = "music"
	CatScienceAndNature SourceCategory = "science-and-nature"
	CatSport            SourceCategory = "sport"
	CatTechnology       SourceCategory = "technology"
)

type SortByOption string

const (
	SortTop     SortByOption = "top"
	SortLatest  SortByOption = "latest"
	SortPopular SortByOption = "popular"
)

type NewsAPIClient struct {
	Key     string
	Client  *http.Client
	BaseUrl string
}

func NewClient(key string, client *http.Client) *NewsAPIClient {
	if client == nil {
		client = &http.Client{}
	}
	return &NewsAPIClient{
		Key:     key,
		Client:  client,
		BaseUrl: "https://newsapi.org",
	}
}

func (c *NewsAPIClient) get(u string, query url.Values) (*http.Response, error) {
	u = c.BaseUrl + u
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.URL.RawQuery = query.Encode()
	return c.Client.Do(req)
}

type SourcesOptions struct {
	Language SourceLanguage
	Category SourceCategory
	Country  SourceCountry
}

type SourcesResponse struct {
	Status  string
	Sources []Source
}

type UrlsToLogos struct {
	Small  string
	Medium string
	Large  string
}

type Source struct {
	Id               string
	Name             string
	Description      string
	Url              string
	Category         string
	Language         string
	Country          string
	UrlsToLogos      UrlsToLogos
	SortBysAvailable []string
}

func (c *NewsAPIClient) Sources(options *SourcesOptions) (*SourcesResponse, error) {
	query := url.Values{}

	if options.Language != "" {
		query.Add("language", string(options.Language))
	}

	if options.Category != "" {
		query.Add("category", string(options.Category))
	}

	if options.Country != "" {
		query.Add("country", string(options.Country))
	}

	resp, err := c.get("/v1/sources", query)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var sourcesResponse SourcesResponse
	err = json.NewDecoder(resp.Body).Decode(&sourcesResponse)
	if err != nil {
		return nil, err
	}
	return &sourcesResponse, nil
}

type ArticleOptions struct {
	SortBy SortByOption
}

type Article struct {
	Author      string
	Title       string
	Description string
	Url         string
	UrlToImage  string
	PublishedAt string
}

type ArticlesResponse struct {
	Status   string
	Source   string
	SortBy   string
	Articles []Article
}

func (c *NewsAPIClient) Articles(source string, options *ArticleOptions) (*ArticlesResponse, error) {
	query := url.Values{}
	query.Add("source", source)
	if options.SortBy != "" {
		query.Add("sortBy", string(options.SortBy))
	}
	resp, err := c.get("/v1/articles", query)

	defer resp.Body.Close()

	var articlesResponse ArticlesResponse
	err = json.NewDecoder(resp.Body).Decode(&articlesResponse)
	if err != nil {
		return nil, err
	}
	return &articlesResponse, nil
}
