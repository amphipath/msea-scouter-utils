package types

type (
	AbilityInfo struct {
		LineNumber string `json:"ability_no"`
		Rank       string `json:"ability_grade"`
		Value      string `json:"ability_value"`
	}

	AbilityPreset struct {
		Grade       string        `json:"ability_preset_grade"`
		AbilityInfo []AbilityInfo `json:"ability_info"`
	}
)
