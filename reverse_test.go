package goreverse

import (
	"reflect"
	"testing"
)

func TestNewURLReverse(t *testing.T) {
	want := URLReverse{
		urls: make(map[string]string),
	}
	got := NewURLReverse()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}

func TestURLReverse_Add(t *testing.T) {
	t.Run("URL", func(t *testing.T) {
		want := NewURLReverse()
		want.urls["url-name"] = "/url"
		got := NewURLReverse()
		got.Add("url-name", "/url")
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want %q, got %q", want, got)
		}
	})
	t.Run("Duplicate", func(t *testing.T) {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("wanted panic, no panic triggered")
			}
		}()
		got := NewURLReverse()
		got.Add("url-name", "/url")
		got.Add("url-name", "/url")
	})
	t.Run("Reassign", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("wanted panic, no panic triggered")
			}
		}()
		got := NewURLReverse()
		got.Add("url-name", "/url")
		got.Add("url-name", "/urls")
	})
}

func TestURLReverse_Get(t *testing.T) {
	urlReverse := NewURLReverse()
	url := "/url"
	urlReverse.Add("url-name", url)
	want := url
	got := urlReverse.Get("url-name")
	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
