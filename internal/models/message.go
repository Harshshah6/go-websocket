package models

type Message struct {
	Name      string   `json:"name"`
	Message   string   `json:"message"`
	Timestamp string   `json:"timestamp"`
	System    bool     `json:"system"`
	Users     []string `json:"users,omitempty"`
}
