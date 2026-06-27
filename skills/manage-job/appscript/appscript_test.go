package appscript

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

// mockAppsScript spins up an httptest.Server that mimics the Apps Script web app:
// GET returns JSON directly; POST returns 302 with Location pointing to /redirect,
// which returns the actual response body — same flow as the real deployment.
func mockAppsScript(t *testing.T, getHandler http.HandlerFunc) *httptest.Server {
	t.Helper()

	mux := http.NewServeMux()

	mux.HandleFunc("/exec", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			if getHandler != nil {
				getHandler(w, r)
			} else {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"status":"success","rows":[]}`))
			}
			return
		}

		// POST: read body, return 302 to /redirect with the body echoed via query param
		body, _ := io.ReadAll(r.Body)
		scheme := "http"
		redirectURL := scheme + "://" + r.Host + "/redirect?payload=" + url.QueryEscape(string(body))
		http.Redirect(w, r, redirectURL, http.StatusFound)
	})

	mux.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		payload := r.URL.Query().Get("payload")
		if payload == "" {
			w.Write([]byte(`{"status":"success"}`))
			return
		}
		w.Write([]byte(payload))
	})

	server := httptest.NewServer(mux)
	t.Cleanup(server.Close)
	return server
}

func TestGet(t *testing.T) {
	server := mockAppsScript(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("industry") != "Tech" {
			t.Errorf("expected industry=Tech query param, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success","rows":[{"companyName":"Acme"}]}`))
	})

	app := &AppScript{URL: server.URL + "/exec"}
	params := url.Values{}
	params.Set("industry", "Tech")

	result, err := app.Get(params)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["status"] != "success" {
		t.Errorf("expected status=success, got %v", resp["status"])
	}
}

func TestGet_NoParams(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	result, err := app.Get(nil)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if !strings.Contains(result, `"status":"success"`) {
		t.Errorf("unexpected response: %s", result)
	}
}

func TestCreate(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	entry := map[string]interface{}{
		"companyName": "Acme Corp",
		"link":        "https://example.com/job",
		"industry":    "Tech",
		"status":      "Not Started",
	}

	result, err := app.Create(entry)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "create" {
		t.Errorf("expected action=create in echoed payload, got %v", resp["action"])
	}
	if resp["companyName"] != "Acme Corp" {
		t.Errorf("expected companyName=Acme Corp, got %v", resp["companyName"])
	}
}

func TestPatch(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{"companyName": "Acme Corp"}
	update := map[string]interface{}{"status": "Interview!"}

	result, err := app.Patch(matchBy, update)
	if err != nil {
		t.Fatalf("Patch failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "patch" {
		t.Errorf("expected action=patch, got %v", resp["action"])
	}
	matchByResp, ok := resp["matchBy"].(map[string]interface{})
	if !ok {
		t.Fatalf("matchBy not a map: %T", resp["matchBy"])
	}
	if matchByResp["companyName"] != "Acme Corp" {
		t.Errorf("expected matchBy.companyName=Acme Corp, got %v", matchByResp["companyName"])
	}
	updateResp, ok := resp["update"].(map[string]interface{})
	if !ok {
		t.Fatalf("update not a map: %T", resp["update"])
	}
	if updateResp["status"] != "Interview!" {
		t.Errorf("expected update.status=Interview!, got %v", updateResp["status"])
	}
}

func TestDelete(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{"companyName": "Acme Corp"}

	result, err := app.Delete(matchBy)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "delete" {
		t.Errorf("expected action=delete, got %v", resp["action"])
	}
	matchByResp, ok := resp["matchBy"].(map[string]interface{})
	if !ok {
		t.Fatalf("matchBy not a map: %T", resp["matchBy"])
	}
	if matchByResp["companyName"] != "Acme Corp" {
		t.Errorf("expected matchBy.companyName=Acme Corp, got %v", matchByResp["companyName"])
	}
}

func TestPostFollowRedirect_Non302(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"bad request"}`))
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err == nil {
		t.Fatal("expected error for non-302 response, got nil")
	}
	if !strings.Contains(err.Error(), "expected 302") {
		t.Errorf("expected 'expected 302' in error, got: %v", err)
	}
}

func TestPostFollowRedirect_NoLocationHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err == nil {
		t.Fatal("expected error for missing Location header, got nil")
	}
	if !strings.Contains(err.Error(), "no redirect Location header") {
		t.Errorf("expected 'no redirect Location header' in error, got: %v", err)
	}
}

func TestGet_MultipleParams(t *testing.T) {
	server := mockAppsScript(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("industry") != "Tech" || q.Get("status") != "Applied Only" || q.Get("page") != "1" {
			t.Errorf("unexpected query params: %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success"}`))
	})

	app := &AppScript{URL: server.URL + "/exec"}
	params := url.Values{}
	params.Set("industry", "Tech")
	params.Set("status", "Applied Only")
	params.Set("page", "1")

	_, err := app.Get(params)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
}

func TestGet_EmptyParams(t *testing.T) {
	server := mockAppsScript(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.RawQuery != "" {
			t.Errorf("expected no query string, got %s", r.URL.RawQuery)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success"}`))
	})

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Get(url.Values{})
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
}

func TestGet_SpecialCharsInParams(t *testing.T) {
	server := mockAppsScript(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("search") != "Acme & Co" {
			t.Errorf("expected search=Acme & Co, got %s", r.URL.Query().Get("search"))
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success"}`))
	})

	app := &AppScript{URL: server.URL + "/exec"}
	params := url.Values{}
	params.Set("search", "Acme & Co")

	_, err := app.Get(params)
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
}

func TestGet_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"server error"}`))
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	result, err := app.Get(nil)
	if err != nil {
		t.Fatalf("Get should not return error for non-200, got: %v", err)
	}
	if !strings.Contains(result, "server error") {
		t.Errorf("expected error body in response, got: %s", result)
	}
}

func TestGet_InvalidURL(t *testing.T) {
	app := &AppScript{URL: "http://invalid.localhost.invalid:99999/exec"}
	_, err := app.Get(nil)
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}
}

func TestCreate_EmptyEntry(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	entry := map[string]interface{}{}

	result, err := app.Create(entry)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "create" {
		t.Errorf("expected action=create, got %v", resp["action"])
	}
}

func TestCreate_MutatesEntryMap(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	entry := map[string]interface{}{"companyName": "Acme"}

	_, _ = app.Create(entry)
	if entry["action"] != "create" {
		t.Errorf("Create should inject action=create into caller's map, got %v", entry["action"])
	}
}

func TestCreate_AllOptionalFields(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	entry := map[string]interface{}{
		"companyName": "Acme",
		"link":        "https://example.com",
		"industry":    "Tech",
		"status":      "Not Started",
		"email":       "hr@acme.com",
		"phoneNumber": "9179991234",
		"notes":       "Follow up next week",
	}

	result, err := app.Create(entry)
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["email"] != "hr@acme.com" {
		t.Errorf("expected email=hr@acme.com, got %v", resp["email"])
	}
	if resp["phoneNumber"] != "9179991234" {
		t.Errorf("expected phoneNumber=9179991234, got %v", resp["phoneNumber"])
	}
	if resp["notes"] != "Follow up next week" {
		t.Errorf("expected notes, got %v", resp["notes"])
	}
}

func TestPatch_EmptyMatchBy(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{}
	update := map[string]interface{}{"status": "Done"}

	result, err := app.Patch(matchBy, update)
	if err != nil {
		t.Fatalf("Patch failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "patch" {
		t.Errorf("expected action=patch, got %v", resp["action"])
	}
}

func TestPatch_EmptyUpdate(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{"companyName": "Acme"}
	update := map[string]interface{}{}

	result, err := app.Patch(matchBy, update)
	if err != nil {
		t.Fatalf("Patch failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "patch" {
		t.Errorf("expected action=patch, got %v", resp["action"])
	}
}

func TestPatch_MultipleUpdateFields(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{"companyName": "Acme", "link": "https://example.com"}
	update := map[string]interface{}{"status": "Done", "notes": "Rejected", "email": "new@acme.com"}

	result, err := app.Patch(matchBy, update)
	if err != nil {
		t.Fatalf("Patch failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	updateResp := resp["update"].(map[string]interface{})
	if len(updateResp) != 3 {
		t.Errorf("expected 3 update fields, got %d", len(updateResp))
	}
	matchByResp := resp["matchBy"].(map[string]interface{})
	if len(matchByResp) != 2 {
		t.Errorf("expected 2 matchBy fields, got %d", len(matchByResp))
	}
}

func TestDelete_EmptyMatchBy(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{}

	result, err := app.Delete(matchBy)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if resp["action"] != "delete" {
		t.Errorf("expected action=delete, got %v", resp["action"])
	}
}

func TestDelete_MultipleMatchByFields(t *testing.T) {
	server := mockAppsScript(t, nil)

	app := &AppScript{URL: server.URL + "/exec"}
	matchBy := map[string]interface{}{
		"companyName": "Acme",
		"link":        "https://example.com",
		"status":      "Done",
	}

	result, err := app.Delete(matchBy)
	if err != nil {
		t.Fatalf("Delete failed: %v", err)
	}

	var resp map[string]interface{}
	if err := json.Unmarshal([]byte(result), &resp); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	matchByResp := resp["matchBy"].(map[string]interface{})
	if len(matchByResp) != 3 {
		t.Errorf("expected 3 matchBy fields, got %d", len(matchByResp))
	}
}

func TestPostFollowRedirect_Server500(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error":"internal"}`))
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err == nil {
		t.Fatal("expected error for 500 response, got nil")
	}
	if !strings.Contains(err.Error(), "expected 302, got 500") {
		t.Errorf("expected 'expected 302, got 500' in error, got: %v", err)
	}
}

func TestPostFollowRedirect_RedirectTargetError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/exec" {
			http.Redirect(w, r, "http://invalid.localhost.invalid:99999/redirect", http.StatusFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err == nil {
		t.Fatal("expected error for unreachable redirect target, got nil")
	}
}

func TestPostFollowRedirect_VerifyContentType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/exec" {
			if r.Header.Get("Content-Type") != "application/json" {
				t.Errorf("expected Content-Type=application/json, got %s", r.Header.Get("Content-Type"))
			}
			if r.Method != "POST" {
				t.Errorf("expected POST, got %s", r.Method)
			}
			http.Redirect(w, r, "http://"+r.Host+"/redirect", http.StatusFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"success"}`))
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err != nil {
		t.Fatalf("Create failed: %v", err)
	}
}

func TestPostFollowRedirect_RedirectReturnsErrorJSON(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/exec" {
			http.Redirect(w, r, "http://"+r.Host+"/redirect", http.StatusFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"error":"row not found"}`))
	}))
	t.Cleanup(server.Close)

	app := &AppScript{URL: server.URL + "/exec"}
	result, err := app.Patch(
		map[string]interface{}{"companyName": "Ghost"},
		map[string]interface{}{"status": "Done"},
	)
	if err != nil {
		t.Fatalf("Patch failed: %v", err)
	}
	if !strings.Contains(result, `"error"`) {
		t.Errorf("expected error in response body, got: %s", result)
	}
}

func TestPostFollowRedirect_InvalidURL(t *testing.T) {
	app := &AppScript{URL: "http://invalid.localhost.invalid:99999/exec"}
	_, err := app.Create(map[string]interface{}{"companyName": "Test"})
	if err == nil {
		t.Fatal("expected error for invalid URL, got nil")
	}
}
