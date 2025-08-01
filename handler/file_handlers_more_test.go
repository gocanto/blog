package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	pkghttp "github.com/oullin/pkg/http"
)

type testData struct {
	Version string `json:"version"`
	Data    any    `json:"data"`
}

func writeJSON(t *testing.T, v any) string {
	f, err := os.CreateTemp("", "data.json")
	if err != nil {
		t.Fatalf("tmp: %v", err)
	}
	enc := json.NewEncoder(f)
	if err := enc.Encode(v); err != nil {
		t.Fatalf("encode: %v", err)
	}
	f.Close()
	return f.Name()
}

func runFileHandlerTest(t *testing.T, makeFn func(string) interface {
	Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
}, path string) {
	file := writeJSON(t, testData{Version: "v1", Data: []map[string]string{{"id": "1"}}})
	defer os.Remove(file)
	h := makeFn(file)

	req := httptest.NewRequest("GET", path, nil)
	rec := httptest.NewRecorder()
	if err := h.Handle(rec, req); err != nil {
		t.Fatalf("err: %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Fatalf("status %d", rec.Code)
	}

	req2 := httptest.NewRequest("GET", path, nil)
	req2.Header.Set("If-None-Match", "\"v1\"")
	rec2 := httptest.NewRecorder()
	if err := h.Handle(rec2, req2); err != nil {
		t.Fatalf("err: %v", err)
	}
	if rec2.Code != http.StatusNotModified {
		t.Fatalf("status %d", rec2.Code)
	}

	badF, _ := os.CreateTemp("", "bad.json")
	badF.WriteString("{")
	badF.Close()
	defer os.Remove(badF.Name())
	bad := makeFn(badF.Name())
	rec3 := httptest.NewRecorder()
	req3 := httptest.NewRequest("GET", path, nil)
	if bad.Handle(rec3, req3) == nil {
		t.Fatalf("expected error")
	}
}

func TestAdditionalFileHandlers(t *testing.T) {
	runFileHandlerTest(t, func(p string) interface {
		Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
	} {
		h := MakeEducationHandler(p)
		return h
	}, "/education")

	runFileHandlerTest(t, func(p string) interface {
		Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
	} {
		h := MakeExperienceHandler(p)
		return h
	}, "/experience")

	runFileHandlerTest(t, func(p string) interface {
		Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
	} {
		h := MakeProjectsHandler(p)
		return h
	}, "/projects")

	runFileHandlerTest(t, func(p string) interface {
		Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
	} {
		h := MakeRecommendationsHandler(p)
		return h
	}, "/recommendations")

	runFileHandlerTest(t, func(p string) interface {
		Handle(http.ResponseWriter, *http.Request) *pkghttp.ApiError
	} {
		h := MakeTalksHandler(p)
		return h
	}, "/talks")
}
