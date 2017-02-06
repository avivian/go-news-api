# go-news-api #
Go Client library for accessing [newsapi.org](https://newsapi.org/)

## Installation
```bash
go get github.com/avivian/go-news-api
```

## Usage
You will need an API key from [newsapi.org](https://newsapi.org/) before you can use the client.

Once you have your key, construct a new NewAPIClient and then use it to interact with the API.

```go
apiKey := "0000000000000000"
client := newsapi.NewClient(apiKey, nil)
sourcesResponse, _ := client.Sources(nil)
articlesResponse, _ := client.Articles("bbc-news", nil)
```

### Optional Parameters
Both the `Sources` and `Articles` functions can have a optional parameters that can be passed.

```go
apiKey := "0000000000000000"
client := newsapi.NewClient(apiKey, nil)
opt := &newsapi.SourceOptions{Country: newsapi.CountryGB}
sourcesResponse, _ := client.Sources(opt)
```

## Licence
[MIT](https://github.com/avivian/go-news-api/blob/master/LICENSE)
