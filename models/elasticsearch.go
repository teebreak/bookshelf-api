package models

type EsSearchResponse struct {
	Hits EsHits `json:"hits"`
}

type EsHits struct {
	Total EsTotal `json:"total"`
	Hits  []EsHit `json:"hits"`
}

type EsTotal struct {
	Value int `json:"value"`
}

type EsHit struct {
	Source Book `json:"_source"`
}
