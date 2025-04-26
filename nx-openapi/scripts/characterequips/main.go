package main

import (
	_ "embed"
	"encoding/json"
	"os"

	"github.com/amphipath/msea-scouter-utils/nx-openapi/adapter"
	"github.com/amphipath/msea-scouter-utils/nx-openapi/types"
	"github.com/amphipath/msea-scouter-utils/resources"
)

type PossibleValues struct {
	ItemEquipmentPart         map[string]string `json:"item_equipment_part,omitempty"`
	ItemEquipmentSlot         map[string]string `json:"item_equipment_slot,omitempty"`
	PotentialOption           map[string]string `json:"potential_option,omitempty"`
	AdditionalPotentialOption map[string]string `json:"additional_potential_option,omitempty"`
	CharacterClass            map[string]string `json:"character_class,omitempty"`
	Ability                   map[string]string `json:"ability,omitempty"`
	LinkSkills                map[string]string `json:"link_skills,omitempty"`
}

//go:embed dictionary.json
var rawDict []byte

//go:embed kmsDictionary.json
var kmsDict []byte

func main() {
	apiKey := os.Getenv("MSEAAPIKEY")
	kmsApiKey := os.Getenv("KMSAPIKEY")
	baseUrl := "https://open.api.nexon.com/maplestorysea"
	kmsBaseURL := "https://open.api.nexon.com/maplestory"
	s := adapter.NewService(baseUrl, apiKey)
	kmsService := adapter.NewService(kmsBaseURL, kmsApiKey)

	data := PossibleValues{}
	json.Unmarshal(rawDict, &data)

	kmsData := PossibleValues{}
	json.Unmarshal(kmsDict, &kmsData)

	igns := resources.LoadIGNs()
	kmsIgns := resources.LoadKMSIGNs()

	populateDict(s, igns, data)
	populateDict(kmsService, kmsIgns, kmsData)

	b, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile("./output.json", b, 0644)

	b, _ = json.MarshalIndent(kmsData, "", "  ")
	os.WriteFile("./kmsOutput.json", b, 0644)
}

func logItem(data PossibleValues, eq types.Equipment) {
	if _, ok := data.ItemEquipmentPart[eq.Part]; !ok {
		data.ItemEquipmentPart[eq.Part] = ""
	}
	if _, ok := data.ItemEquipmentSlot[eq.Slot]; !ok {
		data.ItemEquipmentSlot[eq.Slot] = ""
	}
	if _, ok := data.PotentialOption[eq.PotentialLine1]; !ok {
		data.PotentialOption[eq.PotentialLine1] = ""
	}
	if _, ok := data.PotentialOption[eq.PotentialLine2]; !ok {
		data.PotentialOption[eq.PotentialLine2] = ""
	}
	if _, ok := data.PotentialOption[eq.PotentialLine3]; !ok {
		data.PotentialOption[eq.PotentialLine3] = ""
	}
	if _, ok := data.AdditionalPotentialOption[eq.AdditionalPotentialLine1]; !ok {
		data.AdditionalPotentialOption[eq.AdditionalPotentialLine1] = ""
	}
	if _, ok := data.AdditionalPotentialOption[eq.AdditionalPotentialLine2]; !ok {
		data.AdditionalPotentialOption[eq.AdditionalPotentialLine2] = ""
	}
	if _, ok := data.AdditionalPotentialOption[eq.AdditionalPotentialLine3]; !ok {
		data.AdditionalPotentialOption[eq.AdditionalPotentialLine3] = ""
	}
}

func populateDict(svc adapter.OpenAPIService, igns []string, data PossibleValues) {
	for _, ign := range igns {
		println(ign)
		svc.SetCharacter(ign)

		r, e := svc.GetSetCharacterEquipment()
		if e != nil {
			println(e.Error())
		}

		if r != nil {
			if _, ok := data.CharacterClass[r.Class]; !ok {
				data.CharacterClass[r.Class] = ""
			}

			for _, eq := range r.ItemEquipment {
				logItem(data, eq)
			}
			for _, eq := range r.Preset1 {
				logItem(data, eq)
			}
			for _, eq := range r.Preset2 {
				logItem(data, eq)
			}
			for _, eq := range r.Preset3 {
				logItem(data, eq)
			}
		}

		abilRes, abilErr := svc.GetSetCharacterAbility()
		if abilErr != nil {
			println(abilErr.Error())
		}

		if abilRes != nil {
			for _, abil := range abilRes.AbilityInfo {
				if len(abil.Value) > 0 {
					if _, ok := data.Ability[abil.Value]; !ok {
						data.Ability[abil.Value] = ""
					}
				}
			}
			if abilRes.Preset1 != nil {
				for _, abil := range abilRes.Preset1.AbilityInfo {
					if len(abil.Value) > 0 {
						if _, ok := data.Ability[abil.Value]; !ok {
							data.Ability[abil.Value] = ""
						}
					}
				}
			}
			if abilRes.Preset2 != nil {
				for _, abil := range abilRes.Preset2.AbilityInfo {
					if len(abil.Value) > 0 {
						if _, ok := data.Ability[abil.Value]; !ok {
							data.Ability[abil.Value] = ""
						}
					}
				}
			}
			if abilRes.Preset3 != nil {
				for _, abil := range abilRes.Preset3.AbilityInfo {
					if len(abil.Value) > 0 {
						if _, ok := data.Ability[abil.Value]; !ok {
							data.Ability[abil.Value] = ""
						}
					}
				}
			}
		}

		linkRes, linkErr := svc.GetSetCharacterLinkSkill()
		if linkErr != nil {
			println(linkErr.Error())
		}

		if linkRes != nil {
			if linkRes.OwnedLinkSkill != nil {
				if _, ok := data.LinkSkills[linkRes.OwnedLinkSkill.SkillName]; !ok {
					data.LinkSkills[linkRes.OwnedLinkSkill.SkillName] = ""
				}
			}
		}
	}
}
