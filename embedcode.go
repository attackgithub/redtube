package redtube

type RedtubeEmbedCode struct {
	Embed RedtubeCode `json:"embed,omitempty"`
}

type RedtubeCode struct {
	Code string `json:"code,omitempty"`
}
