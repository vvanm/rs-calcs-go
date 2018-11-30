package transcript

type Transcript struct {
	ID            string `json:"ID,omitempty"`
	CreatedBy     string `json:"createdBy,omitempty"`
	CreatedByName string `json:"createdByName,omitempty"`
	CreatedOn     int64  `json:"createdOn,omitempty"`
	//
	Title          string                `json:"title,omitempty"`
	PublishDate    int64                 `json:"publishDate,omitempty"`
	Type           string                `json:"type,omitempty"`
	TwitchVOD      string                `json:"twitchVOD,omitempty"`
	Keywords       []string              `json:"keywords,omitempty"`
	PatchNotesLink string                `json:"patchNotesLink,omitempty"`
	Mods           []string              `json:"mods,omitempty"`
	Fragments      []Transcript_Fragment `json:"fragments,omitempty"`
	//
	FragmentsCount int `json:"fragmentsCount,omitempty"`
	KeywordsCount  int `json:"keywordsCount,omitempty"`
	ModsCount      int `json:"modsCount,omitempty"`
}

type Transcript_Fragment struct {
	Title     string   `json:"title"`
	Timestamp int      `json:"timestamp"`
	Keywords  []string `json:"keywords"`
	Info      string   `json:"info"`
}
