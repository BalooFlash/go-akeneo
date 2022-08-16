package routing

import "testing"

const BaseUri = "http://test.com/"

func TestGenerate(t *testing.T) {
	t.Run("it generates uri without uriParams", func(t *testing.T) {
		t.Parallel()

		u := NewUriGenerator(BaseUri)
		got := u.Generate("/api/products", nil, nil)
		exp := "http://test.com/api/products"

		if got != exp {
			t.Errorf("expected '%s', got '%s'", exp, got)
		}
	})

	t.Run("it generates uri with uriParams", func(t *testing.T) {
		t.Parallel()

		u := NewUriGenerator(BaseUri)
		got := u.Generate("/api/product/%s", []string{"prod123"}, nil)
		exp := "http://test.com/api/product/prod123"

		if got != exp {
			t.Errorf("expected '%s', got '%s'", exp, got)
		}
	})

	t.Run("it generates uri with uriParams and queryParams", func(t *testing.T) {
		t.Parallel()

		queryParams := map[string]string{"one": "1", "two": "2"}

		u := NewUriGenerator(BaseUri)
		got := u.Generate("/api/product/%s", []string{"prod123"}, queryParams)
		exp := "http://test.com/api/product/prod123?one=1&two=2"

		if got != exp {
			t.Errorf("expected '%s', got '%s'", exp, got)
		}
	})

	t.Run("it generates uri without uriParams and with queryParams", func(t *testing.T) {
		t.Parallel()

		queryParams := map[string]string{"one": "1", "two": "2"}

		u := NewUriGenerator(BaseUri)
		got := u.Generate("/api/products", nil, queryParams)
		exp := "http://test.com/api/products?one=1&two=2"

		if got != exp {
			t.Errorf("expected '%s', got '%s'", exp, got)
		}
	})
}
