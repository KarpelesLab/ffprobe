package ffprobe

type Stream struct {
	Index          int       `json:"index"` // stream index
	CodecName      string    `json:"codec_name"`
	CodecLongName  string    `json:"codec_long_name"`
	Profile        string    `json:"profile"`
	CodecType      string    `json:"codec_type"`       // video
	CodecTagString string    `json:"codec_tag_string"` // avc1
	CodecTag       string    `json:"codec_tag"`        // 0x31637661
	Width          int       `json:"width,omitempty"`
	Height         int       `json:"height,omitempty"`
	CodedWidth     int       `json:"coded_width"`
	CodedHeight    int       `json:"coded_height"`
	ClosedCaptions int       `json:"closed_captions"`
	FrameRate      FrameRate `json:"r_frame_rate"`
	SampleRate     int       `json:"sample_rate,string"`
	Duration       float64   `json:"duration,string"`
	StartTime      float64   `json:"start_time,string"`
	StartPTS       int       `json:"start_pts"`

	Disposition *struct {
		Default int `json:"default"`
		Dub     int `json:"dub"`
	} `json:"disposition,omitempty"`

	Tags map[string]string `json:"tags"`
}
