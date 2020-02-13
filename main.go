package main

import "path/filepath"
import "fmt"
import "io/ioutil"
import "log"
import "os"

import "github.com/blevesearch/bleve"
import "github.com/blevesearch/bleve/analysis/lang/en"
import "github.com/blevesearch/bleve/analysis/analyzer/keyword"
import "github.com/blevesearch/bleve/search"
import "github.com/asaskevich/govalidator"

import "github.com/arraypad/bleve-sort/types"

type Document struct {
	Id        string
	Title     string
	TitleSort string
}

const (
	simpleSort  = true
	noSort      = false
	useProto    = false
	searchField = "TitleSort"
)

func getDoc(id, title string) interface{} {
	if useProto {
		return types.Document{Id: id, Title: title, TitleSort: title}
	}

	return Document{Id: id, Title: title, TitleSort: title}
}

func fieldName(name string) string {
	if useProto {
		return govalidator.CamelCaseToUnderscore(name)
	}

	return name
}

func main() {
	// create temp dir to ensure we have a fresh index
	dir, err := ioutil.TempDir("", "bleve-sort")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	path := filepath.Join(dir, "test.bleve")

	// create mappings & index
	textMapping := bleve.NewTextFieldMapping()
	textMapping.Analyzer = en.AnalyzerName
	kwMapping := bleve.NewTextFieldMapping()
	kwMapping.Analyzer = keyword.Name

	dm := bleve.NewDocumentMapping()
	dm.AddFieldMappingsAt(fieldName("Title"), textMapping)
	dm.AddFieldMappingsAt(fieldName("TitleSort"), kwMapping)

	im := bleve.NewIndexMapping()
	im.AddDocumentMapping("doc", dm)
	im.DefaultAnalyzer = "en"
	im.DefaultMapping = dm

	index, err := bleve.New(path, im)
	if err != nil {
		log.Fatal(err)
	}

	// index some data
	docs := [][]string{
		[]string{"foo:  4", "Xylophones make music"},
		[]string{"bar:  2", "Rhythm is a funny word"},
		[]string{"barp: 1", "Dogs can smell thunder"},
		[]string{"qaz:  3", "Silly little duckling"},
	}

	for _, d := range docs {
		err = index.Index(d[0], getDoc(d[0], d[1]))
		if err != nil {
			log.Fatal(err)
		}
	}

	// search
	req := bleve.NewSearchRequest(bleve.NewMatchAllQuery())

	if simpleSort {
		req.SortBy([]string{fieldName(searchField)})
	} else if !noSort {
		req.SortByCustom(search.SortOrder{
			&search.SortField{
				Field: fieldName(searchField),
				Desc:  false,
				Type:  search.SortFieldAsString,
			},
		})
	}

	res, err := index.Search(req)
	if err != nil {
		log.Fatal(err)
	}

	for _, hit := range res.Hits {
		fmt.Println(hit.ID)
	}
}
