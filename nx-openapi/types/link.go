package types

type (
	Skill struct {
		SkillName       string  `json:"skill_name"`
		Description     string  `json:"skill_description"`
		Level           int     `json:"skill_level"`
		Effect          string  `json:"skill_effect"`
		Icon            string  `json:"skill_icon"`
		NextLevelEffect *string `json:"skill_effect_next"`
	}
)
