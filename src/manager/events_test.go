package manager

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makeTestServer(t *testing.T, path string, status int, body interface{}) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != path {
			http.NotFound(w, r)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		if body != nil {
			json.NewEncoder(w).Encode(body)
		}
	}))
}

func TestNew(t *testing.T) {
	a := New("http://example.com")
	if a.URL != "http://example.com" {
		t.Errorf("New URL = %q, want %q", a.URL, "http://example.com")
	}
}

func TestGetNextEvent_Success(t *testing.T) {
	expected := Event{
		Name:     "LAN Party",
		Capacity: 100,
		Slug:     "lan-party",
	}
	srv := makeTestServer(t, "/events/next", http.StatusOK, expected)
	defer srv.Close()

	a := New(srv.URL)
	event, err := a.GetNextEvent()
	if err != nil {
		t.Fatalf("GetNextEvent returned error: %v", err)
	}
	if event.Name != expected.Name {
		t.Errorf("event.Name = %q, want %q", event.Name, expected.Name)
	}
	if event.Capacity != expected.Capacity {
		t.Errorf("event.Capacity = %d, want %d", event.Capacity, expected.Capacity)
	}
}

func TestGetNextEvent_NonOKStatus(t *testing.T) {
	srv := makeTestServer(t, "/events/next", http.StatusNotFound, nil)
	defer srv.Close()

	a := New(srv.URL)
	_, err := a.GetNextEvent()
	if err == nil {
		t.Error("expected error on non-200 status, got nil")
	}
}

func TestGetNextEvent_Unreachable(t *testing.T) {
	a := New("http://127.0.0.1:0")
	_, err := a.GetNextEvent()
	if err == nil {
		t.Error("expected error for unreachable server, got nil")
	}
}

func TestGetNextEventParticipants_Success(t *testing.T) {
	expected := Event{
		Name: "LAN Party",
		Participants: []EventParticipant{
			{ID: 1, Username: "player1", Seat: "A1"},
			{ID: 2, Username: "player2", Seat: "A2"},
		},
	}
	srv := makeTestServer(t, "/events/next", http.StatusOK, expected)
	defer srv.Close()

	a := New(srv.URL)
	participants, err := a.GetNextEventParticipants()
	if err != nil {
		t.Fatalf("GetNextEventParticipants returned error: %v", err)
	}
	if len(participants) != 2 {
		t.Fatalf("got %d participants, want 2", len(participants))
	}
	if participants[0].Username != "player1" {
		t.Errorf("participants[0].Username = %q, want %q", participants[0].Username, "player1")
	}
}

func TestGetUpcomingEvents_Success(t *testing.T) {
	expected := []Event{
		{Name: "Event A", Slug: "event-a"},
		{Name: "Event B", Slug: "event-b"},
	}
	srv := makeTestServer(t, "/events/upcoming", http.StatusOK, expected)
	defer srv.Close()

	a := New(srv.URL)
	events, err := a.GetUpcomingEvents()
	if err != nil {
		t.Fatalf("GetUpcomingEvents returned error: %v", err)
	}
	if len(events) != 2 {
		t.Fatalf("got %d events, want 2", len(events))
	}
	if events[0].Name != "Event A" {
		t.Errorf("events[0].Name = %q, want %q", events[0].Name, "Event A")
	}
}

func TestGetUpcomingEvents_NonOKStatus(t *testing.T) {
	srv := makeTestServer(t, "/events/upcoming", http.StatusServiceUnavailable, nil)
	defer srv.Close()

	a := New(srv.URL)
	_, err := a.GetUpcomingEvents()
	if err == nil {
		t.Error("expected error on non-200 status, got nil")
	}
}
