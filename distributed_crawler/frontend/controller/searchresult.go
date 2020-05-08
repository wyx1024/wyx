package distributed_controller

import (
	"context"
	"crawler/distributed_crawler/config"
	"crawler/distributed_crawler/engine"
	model "crawler/distributed_crawler/frontend/model"
	view "crawler/distributed_crawler/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/olivere/elastic/v7"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

func CreateSearchResultHandler(filename string) (SearchResultHandler, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return SearchResultHandler{}, err
	}
	view := view.CreateSearchView(filename)
	return SearchResultHandler{
		view:   view,
		client: client,
	}, nil
}
func (s SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	from, _ := strconv.Atoi(r.FormValue("from"))
	q := strings.TrimSpace(r.FormValue("q"))

	page, err := s.getSearchItem(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = s.view.Reader(w, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func (s SearchResultHandler) getSearchItem(q string, from int) (model.SearchResult, error) {
	page := model.SearchResult{}
	resp, err := s.client.Search(config.ELASTIC_INDEX).Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).From(from).Do(context.Background())
	if err != nil {
		return model.SearchResult{}, err
	}
	page.Query = q
	page.Start = from
	page.Hits = resp.TotalHits()
	page.Items = resp.Each(reflect.TypeOf(engine.Item{}))
	page.NextFrom = from + len(page.Items)
	page.PrevFrom = from - len(page.Items)
	return page, nil
}
func rewriteQueryString(q string) string {
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	return re.ReplaceAllString(q, "Payload.$1:")
}
