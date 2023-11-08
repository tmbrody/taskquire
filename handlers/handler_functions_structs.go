package handlers

import (
	"encoding/xml"
	"strings"
)

type UserList struct {
	Users []string
}

type TeamList struct {
	Teams []string
}

type TaskList struct {
	Tasks []string
}

func (u UserList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	users := strings.Join(u.Users, ", ")
	return e.EncodeElement(users, start)
}

func (t TeamList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	teams := strings.Join(t.Teams, ", ")
	return e.EncodeElement(teams, start)
}

func (t TaskList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	tasks := strings.Join(t.Tasks, ", ")
	return e.EncodeElement(tasks, start)
}
