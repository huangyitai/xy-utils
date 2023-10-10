package bindx

// TagKey ...
const TagKey = "bindx"

// TagInfo ...
type TagInfo struct {
	Ignored bool   `tagx:"ignored"`
	Name    string `tagx:"name"`
	Order   int    `tagx:"order"`
}
