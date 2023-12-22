package main

type WhisperOutput struct {
	Text     string                  `json:"text"`
	Segments []WhisperOutputSegments `json:"segments"`
}

type WhisperOutputSegments struct {
	Id               uint    `json:"id"`
	Seek             uint    `json:"seek"`
	Start            float32 `json:"start"`
	End              float32 `json:"end"`
	Text             string  `json:"text"`
	Tokens           []uint  `json:"tokens"`
	Temperature      float32 `json:"temperature"`
	AvgLogProb       float32 `json:"avg_logprob"`
	CompressionRatio float32 `json:"compression_ratio"`
	NoSpeechProb     float32 `json:"no_speech_prob"`
}
