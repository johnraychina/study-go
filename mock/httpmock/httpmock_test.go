package httpmock

import (
	"github.com/jarcoal/httpmock"
	"testing"
)

// see https://github.com/jarcoal/httpmock

func TestFetchArticles(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	// Exact URL match
	httpmock.RegisterResponder("GET", "https://api.mybiz.com/articles",
		httpmock.NewStringResponder(200, `[{"id": 1, "name": "My Great Article"}]`))

	// Regexp match (could use httpmock.RegisterRegexpResponder instead)
	httpmock.RegisterResponder("GET", `=~^https://api\.mybiz\.com/articles/id/\d+\z`,
		httpmock.NewStringResponder(200, `{"id": 1, "name": "My Great Article"}`))

	// do stuff that makes a request to articles

	// get count info
	httpmock.GetTotalCallCount()

	// get the amount of calls for the registered responder
	info := httpmock.GetCallCountInfo()
	info["GET https://api.mybiz.com/articles"]               // number of GET calls made to https://api.mybiz.com/articles
	info["GET https://api.mybiz.com/articles/id/12"]         // number of GET calls made to https://api.mybiz.com/articles/id/12
	info[`GET =~^https://api\.mybiz\.com/articles/id/\d+\z`] // number of GET calls made to https://api.mybiz.com/articles/id/<any-number>
}
