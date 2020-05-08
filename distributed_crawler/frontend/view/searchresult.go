package distributed_view

import (
	model "crawler/distributed_crawler/frontend/model"
	"html/template"
	"io"
)

type SearchResultView struct {
	templeate *template.Template
}

func CreateSearchView(filename string) SearchResultView {
	template := template.Must(template.ParseFiles(filename))
	return SearchResultView{
		templeate: template,
	}
}
func (s SearchResultView) Reader(w io.Writer, data model.SearchResult) error {
	return s.templeate.Execute(w, data)
}
