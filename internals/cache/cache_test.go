package cache_test

import (
	"testing"

	"time"

	"bytes"

	"github.com/phr3nzy/rescounts-api/internals/cache"
)

func parse(s string) time.Duration {
	d, _ := time.ParseDuration(s)
	return d
}

func TestGetEmpty(t *testing.T) {
	c := cache.NewCache()
	content := c.Get("a_random_key")

	assertContentEquals(t, content, []byte(""))
}

func TestGetValue(t *testing.T) {
	c := cache.NewCache()
	c.Set("a_random_key", []byte("123456"), parse("5s"))
	content := c.Get("a_random_key")

	assertContentEquals(t, content, []byte("123456"))
}

func TestGetExpiredValue(t *testing.T) {
	c := cache.NewCache()
	c.Set("a_random_key", []byte("123456"), parse("1s"))
	time.Sleep(parse("1s200ms"))
	content := c.Get("a_random_key")

	assertContentEquals(t, content, []byte(""))
}

func assertContentEquals(t *testing.T, content, expected []byte) {
	if !bytes.Equal(content, expected) {
		t.Errorf("Content should '%s', but was '%s'.", expected, content)
	}
}
