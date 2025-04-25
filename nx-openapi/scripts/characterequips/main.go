package main

import (
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

func main() {
	apiKey := os.Getenv("NXOPENAPIKEY")
	baseUrl := "https://open.api.nexon.com/maplestorysea"

	s := adapter.NewService(baseUrl, apiKey)

	data := PossibleValues{map[string]string{}, map[string]string{}, map[string]string{}, map[string]string{}, map[string]string{}}

	igns := resources.LoadIGNs()

	for _, ign := range igns {
		s.SetCharacter(ign)

		r, e := s.GetSetCharacterEquipment()
		if e != nil {
			println(e.Error())
		}

		if r != nil {
			data.CharacterClass[r.Class] = ""
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
	data.ItemEquipmentPart[eq.Part] = ""
	data.ItemEquipmentSlot[eq.Slot] = ""
	data.PotentialOption[eq.PotentialLine1] = ""
	data.PotentialOption[eq.PotentialLine2] = ""
	data.PotentialOption[eq.PotentialLine3] = ""
	data.AdditionalPotentialOption[eq.AdditionalPotentialLine1] = ""
	data.AdditionalPotentialOption[eq.AdditionalPotentialLine2] = ""
	data.AdditionalPotentialOption[eq.AdditionalPotentialLine3] = ""
}
