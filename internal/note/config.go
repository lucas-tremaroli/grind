package note

const (
	// Field focus constants
	FieldFilename = iota
	FieldContent
)

const (
	// UI dimensions and limits
	FilenameWidth     = 50
	FilenameCharLimit = 20
	ContentWidth      = 80
	ContentHeight     = 5
)

// Config holds configuration for the note editor
type Config struct {
	FilenameWidth     int
	FilenameCharLimit int
	ContentWidth      int
	ContentHeight     int
}

// DefaultConfig returns the default configuration
func DefaultConfig() Config {
	return Config{
		FilenameWidth:     FilenameWidth,
		FilenameCharLimit: FilenameCharLimit,
		ContentWidth:      ContentWidth,
		ContentHeight:     ContentHeight,
	}
}
