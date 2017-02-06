package newsapi

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

var (
	mux    *http.ServeMux
	client *NewsAPIClient
	server *httptest.Server
)

type values map[string]string

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient("0000000000000000000", nil)
	client.BaseUrl = server.URL
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testFormValues(t *testing.T, r *http.Request, values values) {
	want := url.Values{}
	for k, v := range values {
		want.Add(k, v)
	}

	r.ParseForm()
	if got := r.Form; !reflect.DeepEqual(want, got) {
		t.Errorf("Request Parameters: %v, want", got, want)
	}

}

func TestSources(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sources", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"language": "en",
			"country":  "gb",
			"category": "business",
		})
		fmt.Fprint(w, `{"status":"ok","sources":[{"id":"abc-news-au","name":"ABC News (AU)","description":"Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.","url":"http://www.abc.net.au/news","category":"general","language":"en","country":"au","urlsToLogos":{"small":"http://i.newsapi.org/abc-news-au-s.png","medium":"http://i.newsapi.org/abc-news-au-m.png","large":"http://i.newsapi.org/abc-news-au-l.png"},"sortBysAvailable":["top"]}]}`)
	})
	sourcesResponse, err := client.Sources(&SourcesOptions{Language: LangEN, Country: CountryGB, Category: CatBusiness})
	if err != nil {
		t.Errorf("Sources returned err: %v", err)
	}

	want := &SourcesResponse{
		Status: "ok",
		Sources: []Source{
			{
				Id:          "abc-news-au",
				Name:        "ABC News (AU)",
				Description: "Australia's most trusted source of local, national and world news. Comprehensive, independent, in-depth analysis, the latest business, sport, weather and more.",
				Url:         "http://www.abc.net.au/news",
				Category:    "general",
				Language:    "en",
				Country:     "au",
				UrlsToLogos: UrlsToLogos{
					Small:  "http://i.newsapi.org/abc-news-au-s.png",
					Medium: "http://i.newsapi.org/abc-news-au-m.png",
					Large:  "http://i.newsapi.org/abc-news-au-l.png",
				},
				SortBysAvailable: []string{"top"},
			},
		},
	}

	if !reflect.DeepEqual(sourcesResponse, want) {
		t.Errorf("Sources returned %v, got %v", sourcesResponse, want)
	}
}

func TestArticles(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/articles", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testFormValues(t, r, values{
			"sortBy": "top",
			"source": "bbc-news",
		})
		fmt.Fprint(w, `{"status":"ok","source":"bbc-news","sortBy":"top","articles":[{"author":"https://www.facebook.com/bbcnews","title":"Speaker Bercow: Trump should not speak in Parliament","description":"The Speaker says it is important to oppose racism and sexism, as he sets out his stance to MPs.","url":"http://www.bbc.co.uk/news/uk-politics-38884604","urlToImage":"http://ichef.bbci.co.uk/news/1024/cpsprodpb/6F0E/production/_94003482_de30-1.jpg","publishedAt":"2017-02-06T20:16:52Z"}]}`)
	})

	articlesResponse, err := client.Articles("bbc-news", &ArticleOptions{SortBy: SortTop})
	if err != nil {
		t.Errorf("Articles returned err: %v", err)
	}

	want := &ArticlesResponse{
		Status: "ok",
		Source: "bbc-news",
		SortBy: "top",
		Articles: []Article{
			{
				Author:      "https://www.facebook.com/bbcnews",
				Title:       "Speaker Bercow: Trump should not speak in Parliament",
				Description: "The Speaker says it is important to oppose racism and sexism, as he sets out his stance to MPs.",
				Url:         "http://www.bbc.co.uk/news/uk-politics-38884604",
				UrlToImage:  "http://ichef.bbci.co.uk/news/1024/cpsprodpb/6F0E/production/_94003482_de30-1.jpg",
				PublishedAt: "2017-02-06T20:16:52Z",
			},
		},
	}

	if !reflect.DeepEqual(articlesResponse, want) {
		t.Errorf("Articles returned %v, got %v", articlesResponse, want)
	}

}
