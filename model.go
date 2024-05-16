package main

type GetBooksRespDTO struct {
	Name        string     `json:"name"`
	SubjectType string     `json:"subject_type"`
	Works       []*WorkDTO `json:"works"`
}

type BookReqDTO struct {
	Subject string `json:"subject"`
}

type WorkDTO struct {
	Title        string       `json:"title"`
	CoverID      int64        `json:"cover_id"`
	EditionCount int64        `json:"edition_count"`
	Authors      []*AuthorDTO `json:"authors"`
}

type AuthorDTO struct {
	Name string `json:"name"`
}
