package media

type VideoHeatmap struct {
	EndTime   float32 `json:"end_time"`
	StartTime float32 `json:"start_time"`
	Value     float64 `json:"value"` // Updated type to float64 to handle decimal values
}

type UserInput struct {
	From     float32 `json:"from,omitempty"` // Use pointers to make fields optional
	To       float32 `json:"to,omitempty"`
	Duration float32 `json:"duration,omitempty"`
}

type VideoMetadata struct {
	Title          string         `json:"title"`
	Duration       float32        `json:"duration"`        // Assuming duration is in seconds
	DurationString string         `json:"duration_string"` // This might need to be calculated separately if not provided directly
	Heatmap        []VideoHeatmap `json:"heatmap"`
	OriginalUrl    string         `json:"original_url"`
}

type CombinedData struct {
	VideoMetadata VideoMetadata `json:"fetched"`
	UserInput     UserInput     `json:"userInput"`
}
