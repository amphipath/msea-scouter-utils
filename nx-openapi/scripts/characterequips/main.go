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
}

//go:embed dictionary.json
var rawDict []byte

func main() {
	apiKey := os.Getenv("NXOPENAPIKEY")
	baseUrl := "https://open.api.nexon.com/maplestorysea"

	s := adapter.NewService(baseUrl, apiKey)

	data := PossibleValues{}
	json.Unmarshal(rawDict, &data)

	igns := resources.LoadIGNs()

	for _, ign := range igns {
		s.SetCharacter(ign)

		r, e := s.GetSetCharacterEquipment()
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
	}

	b, _ := json.MarshalIndent(data, "", "  ")
	os.WriteFile("./output.json", b, 0644)
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
