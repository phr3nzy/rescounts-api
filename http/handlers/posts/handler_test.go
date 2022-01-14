package posts_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/phr3nzy/rescounts-api/http/server"
)

const (
	TIMEOUT = -1 // for infinite.
)

func TestFetchMultiplePostsWithCaching(t *testing.T) {
	app := server.Bootstrap()
	response, err := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=science,tech", nil), TIMEOUT)

	utils.AssertEqual(t, nil, err, "test with defaults")
	utils.AssertEqual(t, 200, response.StatusCode, "status code")
}

func TestFetchMultiplePostsWithCachingSortingAndOrdering(t *testing.T) {
	app := server.Bootstrap()

	response, err := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=design&sortBy=likes&direction=desc", nil), TIMEOUT)

	utils.AssertEqual(t, nil, err, "test with sorting and ordering")
	utils.AssertEqual(t, 200, response.StatusCode, "status code")
}

func TestErrorResponses(t *testing.T) {
	app := server.Bootstrap()

	missingTagsQueryResponse, err1 := app.Test(httptest.NewRequest("GET", "/api/v1/posts", nil), TIMEOUT)
	missingSortByQueryResponse, err2 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=science&direction=asc", nil), TIMEOUT)
	missingDirectionQueryResponse, err3 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=design&sortBy=likes", nil), TIMEOUT)
	invalidTagsQueryResponse, err4 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=", nil), TIMEOUT)
	invalidSortByQueryResponse, err5 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=tech&sortBy=invalid_sortBy&direction=asc", nil), TIMEOUT)
	invalidDirectionQueryResponse, err6 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=history&sortBy=likes&direction=invalid_direction", nil), TIMEOUT)

	utils.AssertEqual(t, nil, err1, "test without tags query")
	utils.AssertEqual(t, nil, err2, "test without sortBy query")
	utils.AssertEqual(t, nil, err3, "test without direction query")
	utils.AssertEqual(t, nil, err4, "test with invalid tags query")
	utils.AssertEqual(t, nil, err5, "test with invalid sortBy query")
	utils.AssertEqual(t, nil, err6, "test with invalid direction query")

	utils.AssertEqual(t, 400, missingTagsQueryResponse.StatusCode, "status code")
	utils.AssertEqual(t, 400, missingSortByQueryResponse.StatusCode, "status code")
	utils.AssertEqual(t, 400, missingDirectionQueryResponse.StatusCode, "status code")
	utils.AssertEqual(t, 400, invalidTagsQueryResponse.StatusCode, "status code")
	utils.AssertEqual(t, 400, invalidSortByQueryResponse.StatusCode, "status code")
	utils.AssertEqual(t, 400, invalidDirectionQueryResponse.StatusCode, "status code")
}

func TestCachedResponse(t *testing.T) {
	app := server.Bootstrap()

	req1, err1 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=science,history&sortBy=likes&direction=desc", nil), TIMEOUT)
	req2, err2 := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=science,history&sortBy=likes&direction=desc", nil), TIMEOUT)

	utils.AssertEqual(t, nil, err1, "test without caching")
	utils.AssertEqual(t, nil, err2, "test with caching")

	var req1Body []byte
	var req2Body []byte

	req1.Body.Read(req1Body)
	req2.Body.Read(req2Body)

	req1.Body.Close()
	req2.Body.Close()

	utils.AssertEqual(t, req1Body, req2Body, "response body")
}
