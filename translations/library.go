package translations

type (
	TranslationKey struct {
		MSEA string `json:"msea"`
		KMS  string `json:"kms"`
	}
	TranslationLibrary struct {
		Keys map[string]TranslationKey `json:"keys"`
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
