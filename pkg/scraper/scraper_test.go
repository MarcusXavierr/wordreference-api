package scraper_test

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/MarcusXavierr/wordreference-api/pkg/scraper"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

func TestReadPageContent(t *testing.T) {
	html := getParsedHTML(t, "dumb_down_definition.html")

	got, err := scraper.ParsePage(html)

	if err != nil {
		t.Fatal("Error while parsing the html")
	}
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

func getParsedHTML(t testing.TB, filename string) *html.Node {
	t.Helper()

	reader := getFileContent(filename, t)

	html, err := htmlquery.Parse(reader)
	if err != nil {
		t.Fatal("could not parse html")
	}

	return html
}

func getFileContent(filename string, t testing.TB) *strings.Reader {
	t.Helper()

	path := filepath.Join("testdata", filename)
	content, err := ioutil.ReadFile(path)

	if err != nil {
		t.Fatalf("error reading testfile %s: %s", filename, err)
	}

	return strings.NewReader(string(content))
}
