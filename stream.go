package ffprobe

type Stream struct {
	Index          int      `json:"index"` // stream index
	CodecName      string   `json:"codec_name"`
	CodecLongName  string   `json:"codec_long_name"`
	Profile        string   `json:"profile"`
	CodecType      string   `json:"codec_type"`       // video
	CodecTagString string   `json:"codec_tag_string"` // avc1
	CodecTag       string   `json:"codec_tag"`        // 0x31637661
	Width          int      `json:"width,omitempty"`
	Height         int      `json:"height,omitempty"`
	CodedWidth     int      `json:"coded_width"`
	CodedHeight    int      `json:"coded_height"`
	ClosedCaptions int      `json:"closed_captions"`
	FrameRate      Fraction `json:"r_frame_rate"`
	AvgFrameRate   Fraction `json:"avg_frame_rate"`
	TimeBase       Fraction `json:"time_base"`
	SampleRate     int      `json:"sample_rate,string"`
	Duration       float64  `json:"duration,string"`
	DurationTs     uint64   `json:"duration_ts"`
	StartTime      float64  `json:"start_time,string"` // can be negative
	StartPTS       int      `json:"start_pts"`
	ExtradataSize  int      `json:"extradata_size"`

	Disposition *struct {
		Default         int `json:"default"`
		Dub             int `json:"dub"`
		Original        int `json:"original"`
		Comment         int `json:"comment"`
		Lyrics          int `json:"lyrics"`
		Karaoke         int `json:"karaoke"`
		Forced          int `json:"forced"`
		HearingImpaired int `json:"hearing_impaired"`
		VisualImpaired  int `json:"visual_impaired"`
		CleanEffects    int `json:"clean_effects"`
		AttachedPic     int `json:"attached_pic"`
		TimedThumbnails int `json:"timed_thumbnails"`
		Captions        int `json:"captions"`
		Descriptions    int `json:"descriptions"`
		Metadata        int `json:"metadata"`
		Dependent       int `json:"dependent"`
		StillImage      int `json:"still_image"`
	} `json:"disposition,omitempty"`

	Tags map[string]string `json:"tags"`
}
