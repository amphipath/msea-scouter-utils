package types

import "time"

type (
	CharacterEquipmentResponse struct {
		Date              string      `json:"date"`
		Gender            string      `json:"character_gender"`
		Class             string      `json:"character_class"`
		PresetNo          int         `json:"preset_no"`
		ItemEquipment     []Equipment `json:"item_equipment"`
		Preset1           []Equipment `json:"item_equipment_preset_1"`
		Preset2           []Equipment `json:"item_equipment_preset_2"`
		Preset3           []Equipment `json:"item_equipment_preset_3"`
		Title             Title       `json:"title"`
		DragonEquipment   []Equipment `json:"dragon_equipment"`
		MechanicEquipment []Equipment `json:"mechanic-equipment"`
	}

	Equipment struct {
		Part                           string         `json:"item_equipment_part"`
		Slot                           string         `json:"item_equipment_slot"`
		Name                           string         `json:"item_name"`
		Icon                           string         `json:"item_icon"`
		Description                    string         `json:"item_description"`
		AnvilName                      string         `json:"item_shape_name"`
		AnvilIcon                      string         `json:"item_shape_icon"`
		RequiredGender                 *string        `json:"item_gender"`
		TotalStats                     EquipmentStats `json:"item_total_option"`
		BaseStats                      EquipmentStats `json:"item_base_option"`
		PotentialGrade                 string         `json:"potential_option_grade,omitempty"`
		AdditionalPotentialOptionGrade string         `json:"additional_potential_option_grade,omitempty"`
		PotentialOptionFlag            string         `json:"potential_option_flag"`
		PotentialLine1                 string         `json:"potential_option_1"`
		PotentialLine2                 string         `json:"potential_option_2"`
		PotentialLine3                 string         `json:"potential_option_3"`
		AdditionalPotentialOptionFlag  string         `json:"additional_potential_option_flag"`
		AdditionalPotentialLine1       string         `json:"additional_potential_option_1"`
		AdditionalPotentialLine2       string         `json:"additional_potential_option_2"`
		AdditionalPotentialLine3       string         `json:"additional_potential_option_3"`
		EquipmentLevelIncrease         int            `json:"equipment_level_increase"`
		ItemExceptionalEnhancement     EquipmentStats `json:"item_exceptional_option"`
		ItemFlameStats                 EquipmentStats `json:"item_add_option"`
		GrowthEXP                      int            `json:"growth_exp"`
		GrowthLevel                    int            `json:"growth_level"`
		ScrollUsedCount                string         `json:"scroll_upgrade"`
		ScissorsCount                  *string        `json:"cuttable_count"`
		GoldenHammerFlag               *string        `json:"golden_hammer_flag"`
		CleanSlateUsableCount          *string        `json:"scroll_resilience_count"`
		ScrollUsableCount              *string        `json:"scroll_upgradeable_count"`
		SoulName                       *string        `json:"soul_name"`
		SoulStats                      *string        `json:"soul_option"`
		ItemScrolledStats              EquipmentStats `json:"item_etc_option"`
		Starforce                      *string        `json:"starforce"`
		StarforceScrollFlag            *string        `json:"starforce_scroll_flag"`
		StarforceStats                 EquipmentStats `json:"item_starforce_option"`
		SpecialRingLevel               int            `json:"special_ring_level"`
		ExpiryDate                     *time.Time     `json:"date_expire,omitempty"`
	}

	EquipmentStats struct {
		STR                    string  `json:"str"`
		DEX                    string  `json:"dex"`
		INT                    string  `json:"int"`
		LUK                    string  `json:"luk"`
		MaxHP                  string  `json:"max_hp"`
		MaxMP                  string  `json:"max_mp"`
		WeaponAttack           string  `json:"attack_power"`
		MagicAttack            string  `json:"magic_power"`
		Armor                  string  `json:"armor"`
		Speed                  string  `json:"speed"`
		Jump                   string  `json:"jump"`
		BossDamage             string  `json:"boss_damage"`
		IgnoreDEF              string  `json:"ignore_monster_armor"`
		AllStatPercent         string  `json:"all_stat"`
		Damage                 string  `json:"damage"`
		EquipmentLevelDecrease int     `json:"equipment_level_decrease"`
		MaxHPPercent           *string `json:"max_hp_rate,omitempty"`
		MaxMPPercent           *string `json:"max_mp_rate,omitempty"`
		BaseEquipmentLevel     *int    `json:"base_equipment_level,omitempty"`
		ExceptionalUpgrade     *int    `json:"exceptional_upgrade,omitempty"`
	}

	Title struct {
		Name           string     `json:"title_name"`
		Icon           string     `json:"title_icon"`
		Description    string     `json:"title_description"`
		ExpiryDate     *time.Time `json:"date_expire,omitempty"`
		StatExpiryDate *time.Time `json:"date_option_expire,omitempty"`
	}
)
