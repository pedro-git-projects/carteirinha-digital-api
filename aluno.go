package main

import "encoding/json"

type Aluno struct {
	username string
	hash     string
}

func (a *Aluno) MarshalJSON() ([]byte, error) {
	data := map[string]string{
		"username": a.username,
		"password": a.hash,
	}
	return json.Marshal(data)
}
