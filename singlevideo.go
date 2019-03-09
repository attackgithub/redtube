package redtube

type RedtubeSingleVideo struct {
	Video RedtubeVideoAlone `json:"video,omitempty"`
}

type RedtubeVideoAlone struct {
	ID           string  `json:"video_id,omitempty"`
	Duration     string  `json:"duration,omitempty"`
	Views        float64 `json:"views,string"`
	Rating       string  `json:"rating,omitempty"`
	Ratings      string  `json:"ratings,omitempty"`
	Title        string  `json:"title,omitempty"`
	URL          string  `json:"url,omitempty"`
	EmbedURL     string  `json:"embed_url,omitempty"`
	DefaultThumb string  `json:"default_thumb,omitempty"`
	Thumb        string  `json:"thumb,omitempty"`
	PublishDate  string  `json:"publish_date,omitempty"`
	Thumbs       []RedtubeThumbAlone
	Tags         []string
}

type RedtubeThumbAlone struct {
	Size   string  `json:"size,omitempty"`
	Width  float64 `json:"width,omitempty"`
	Height float64 `json:"height,omitempty"`
	Src    string  `json:"src,omitempty"`
}
