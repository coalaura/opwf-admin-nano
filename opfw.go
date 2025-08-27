package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func (d *Database) String() string {
	var name string

	if d.Name != nil {
		name = strings.TrimSpace(string(*d.Name))

		if name != "" {
			name = " - " + name
		}
	}

	return fmt.Sprintf("%d%s", d.ID, name)
}

func LoadDatabases(config *Config) ([]*Database, error) {
	req, err := http.NewRequest("GET", "https://op-framework.com/api/fivem/servers/databases", nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", config.Token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var databases []*Database

	err = json.NewDecoder(resp.Body).Decode(&databases)
	if err != nil {
		return nil, err
	}

	sort.Slice(databases, func(i, j int) bool {
		return databases[i].ID < databases[j].ID
	})

	return databases, nil
}
