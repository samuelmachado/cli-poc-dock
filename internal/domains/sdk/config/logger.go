package config

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/merci-app/dock-sdk-backend"
)

type loggingEvent struct {
	URL          string    `json:"url"`
	RequestID    string    `json:"request_id"`
	CurrentTime  time.Time `json:"current_time"`
	StatusCode   int       `json:"status_code,omitempty"`
	ErrorMessage string    `json:"error_message,omitempty"`
	EventLevel   string    `json:"event_level"`
}

type logger struct {
	beforeRequest func(req *http.Request)
	afterRequest  func(req *http.Request, res *http.Response, err error)
}

func (l *logger) BeforeRequest(req *http.Request) {
	l.beforeRequest(req)
}

func (l *logger) AfterRequest(req *http.Request, res *http.Response, err error) {
	l.afterRequest(req, res, err)
}

func loggerHooks() dock.Hooks {
	l := &logger{}

	f, err := os.OpenFile("logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	log.SetOutput(f)

	l.beforeRequest = func(req *http.Request) {}
	l.afterRequest = func(req *http.Request, res *http.Response, err error) {
		reqID := req.Context().Value("request_id").(string)

		e := &loggingEvent{
			URL:         req.URL.String(),
			CurrentTime: time.Now(),
			RequestID:   reqID,
		}

		if err != nil {
			e.ErrorMessage = err.Error()
			e.EventLevel = "error"
		} else {
			e.StatusCode = res.StatusCode

			if e.StatusCode >= 400 {
				e.EventLevel = "warning"
			} else {
				e.EventLevel = "info"
			}
		}

		logMsgBytes, err := json.Marshal(e)
		if err != nil {
			log.Fatalf("error marshalling event struct: %v", err)
		}

		log.Println(string(logMsgBytes))
	}

	return l
}
