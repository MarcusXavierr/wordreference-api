package scraper_test

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/MarcusXavierr/wordreference-api/pkg/scraper"
	"github.com/antchfx/htmlquery"
)

func TestReadPageContent(t *testing.T) {
	filename := "dumb_down_definition.html"
	path := filepath.Join("testdata", filename)

	content, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("error reading testfile %s: %s", filename, err)
	}

	html, _ := htmlquery.Parse(strings.NewReader(string(content)))
	got := scraper.ParsePage(html)
	want := scraper.Page{
		Title: "dumb down",
		Definitions: []scraper.Definition{
			{Word: "dumb down"},
			{Word: "dumb [sth] down"},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}

}
