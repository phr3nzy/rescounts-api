package posts_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2/utils"
	"github.com/phr3nzy/rescounts-api/http/server"
)

func TestFetchMultiplePostsWithCaching(t *testing.T) {
	app := server.Bootstrap()
	response, err := app.Test(httptest.NewRequest("GET", "/api/v1/posts?tags=science,tech", nil), -1)

	utils.AssertEqual(t, nil, err, "app.Test")
	utils.AssertEqual(t, 200, response.StatusCode, "Status Code")
}
