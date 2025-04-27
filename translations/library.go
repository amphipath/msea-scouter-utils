package translations

type (
	TranslationKey map[Language]string

	Language string

	TranslationLibrary struct {
		Categories map[string]map[string]TranslationKey `json:"keys"`
	}

	PossibleValues struct {
		ItemEquipmentPart         map[string]string `json:"item_equipment_part,omitempty"`
		ItemEquipmentSlot         map[string]string `json:"item_equipment_slot,omitempty"`
		PotentialOption           map[string]string `json:"potential_option,omitempty"`
		AdditionalPotentialOption map[string]string `json:"additional_potential_option,omitempty"`
		CharacterClass            map[string]string `json:"character_class,omitempty"`
		Ability                   map[string]string `json:"ability,omitempty"`
		LinkSkills                map[string]string `json:"link_skills,omitempty"`
	}
)

const (
	LanguageMSEA Language = "msea"
	LanguageKMS  Language = "kms"

	CategoryJob            = "jobs"
	CategoryJobAbbreviated = "jobs_short"
)

func (t *TranslationLibrary) AddKey(category, key string, lang Language, value string) {
	if t.Categories == nil {
		t.Categories = make(map[string]map[string]TranslationKey)
	}
	if c, ok := t.Categories[category]; !ok {
		t.Categories[category] = map[string]TranslationKey{
			key: map[Language]string{},
		}
	} else if _, ok := c[key]; !ok {
		t.Categories[category][key] = map[Language]string{}
	}

	t.Categories[category][key][lang] = value
}

func (t *TranslationLibrary) GetKey(category, key string, lang Language) *string {
	if t.Categories != nil {
		if c := t.Categories[category]; c != nil {
			if k := c[key]; k != nil {
				s := k.Translate(lang)
				return &s
			}
		}
	}
	return nil
}

func (t *TranslationLibrary) AddKeyIfAbsent(category, key string, lang Language, value string) {
	if t.GetKey(category, key, lang) == nil {
		t.AddKey(category, key, lang, value)
	}
}

func (t TranslationKey) Translate(lang Language) string {
	return t[lang]
}
