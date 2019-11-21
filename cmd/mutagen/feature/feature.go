package feature

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"
)

// IsEnabled returns true if the named feature should be enabled in this version of Docker Desktop.
func IsEnabled(name string) bool {
	features, err := fetch()
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range features {
		if f.ID == name {
			return f.Enabled
		}
	}
	return false
}

// Implementation details follow:

type feature struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Enabled bool   `json:"enabled"`
}

func hTTP() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, _, _ string) (net.Conn, error) {
				return dial(ctx)
			},
		},
	}
}

func fetch() ([]feature, error) {
	resp, err := hTTP().Get("http://localhost/features")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, err
	}
	var features []feature
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&features); err != nil {
		return nil, err
	}
	return features, nil
}
