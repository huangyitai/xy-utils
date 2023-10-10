package confx

// TagKey ...
const TagKey = "confx"

// TagInfo ...
type TagInfo struct {
	Ignored bool   `tagx:"ignored"`
	Name    string `tagx:"name"`
	Format  string `tagx:"format"`
	Read    string `tagx:"read"`
	Binding bool   `tagx:"binding"`
	Path    string `tagx:"path"`
}
