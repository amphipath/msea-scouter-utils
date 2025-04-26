package adapter

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/amphipath/msea-scouter-utils/nx-openapi/types"
)

type OpenAPIService interface {
	GetCharacterEquipment(ocid string) (*GetCharacterEquipmentResponse, error)
	GetCharacterOCID(ign string) (*GetCharacterIDResponse, error)
	GetCharacterAbility(ocid string) (*GetCharacterAbilityResponse, error)
	GetCharacterLinkSkill(ocid string) (*GetCharacterLinkSkillResponse, error)

	SetCharacter(IGN string)
	GetSetCharacterEquipment() (*GetCharacterEquipmentResponse, error)
	GetSetCharacterAbility() (*GetCharacterAbilityResponse, error)
	GetSetCharacterLinkSkill() (*GetCharacterLinkSkillResponse, error)
}

type (
	GetCharacterEquipmentResponse struct {
		Date              string            `json:"date"`
		Gender            string            `json:"character_gender"`
		Class             string            `json:"character_class"`
		PresetNo          int               `json:"preset_no"`
		ItemEquipment     []types.Equipment `json:"item_equipment"`
		Preset1           []types.Equipment `json:"item_equipment_preset_1"`
		Preset2           []types.Equipment `json:"item_equipment_preset_2"`
		Preset3           []types.Equipment `json:"item_equipment_preset_3"`
		Title             types.Title       `json:"title"`
		DragonEquipment   []types.Equipment `json:"dragon_equipment"`
		MechanicEquipment []types.Equipment `json:"mechanic-equipment"`
	}

	GetCharacterIDResponse struct {
		OCID string `json:"ocid"`
	}

	GetCharacterAbilityResponse struct {
		Date         *string              `json:"date,omitempty"`
		AbilityRank  string               `json:"ability_grade"`
		AbilityInfo  []types.AbilityInfo  `json:"ability_info"`
		HonourEXP    int                  `json:"remain_fame"`
		PresetNumber int                  `json:"preset_no"`
		Preset1      *types.AbilityPreset `json:"ability_preset_1"`
		Preset2      *types.AbilityPreset `json:"ability_preset_2"`
		Preset3      *types.AbilityPreset `json:"ability_preset_3"`
	}

	GetCharacterLinkSkillResponse struct {
		Date                  *string       `json:"date,omitempty"`
		CharacterClass        string        `json:"character_class"`
		LinkSkills            []types.Skill `json:"character_link_skill"`
		LinkSkillPreset1      []types.Skill `json:"character_link_skill_preset_1"`
		LinkSkillPreset2      []types.Skill `json:"character_link_skill_preset_2"`
		LinkSkillPreset3      []types.Skill `json:"character_link_skill_preset_3"`
		OwnedLinkSkill        *types.Skill  `json:"character_owned_link_skill"`
		OwnedLinkSkillPreset1 *types.Skill  `json:"character_owned_link_skill_preset_1"`
		OwnedLinkSkillPreset2 *types.Skill  `json:"character_owned_link_skill_preset_2"`
		OwnedLinkSkillPreset3 *types.Skill  `json:"character_owned_link_skill_preset_3"`
	}

	service struct {
		ocid    string
		baseURL string
		client  *http.Client
	}
)

const (
	getIdIGNParam = "character_name"
)

func NewService(baseURL, apiKey string) OpenAPIService {
	s := &service{
		baseURL: baseURL,
	}
	c := &http.Client{}
	AddMiddlewaresToClient(
		c,
		Convert400ResponseToError(),
		APIKeyHeaderMiddleware(apiKey),
		ThrottleMiddleware(20),
		RetryMiddleware(3),
	)
	s.client = c
	return s
}

func (s *service) GetCharacterOCID(ign string) (*GetCharacterIDResponse, error) {
	u, _ := url.Parse(s.baseURL)
	u = u.JoinPath("v1/id")
	q := u.Query()
	q.Add(getIdIGNParam, ign)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer, _ := io.ReadAll(b)
	r := GetCharacterIDResponse{}
	err = json.Unmarshal(buffer, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *service) SetCharacter(ign string) {
	r, err := s.GetCharacterOCID(ign)
	if err == nil && r != nil {
		s.ocid = r.OCID
	}
}

func (s *service) GetCharacterEquipment(ocid string) (*GetCharacterEquipmentResponse, error) {
	u, _ := url.Parse(s.baseURL)
	u = u.JoinPath("v1/character/item-equipment")
	q := u.Query()
	q.Add(ocidParam, ocid)
	u.RawQuery = q.Encode()

	k := u.String()

	req, err := http.NewRequest(http.MethodGet, k, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer, _ := io.ReadAll(b)

	r := GetCharacterEquipmentResponse{}
	err = json.Unmarshal(buffer, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *service) GetSetCharacterEquipment() (*GetCharacterEquipmentResponse, error) {
	return s.GetCharacterEquipment(s.ocid)
}

func (s *service) GetCharacterAbility(ocid string) (*GetCharacterAbilityResponse, error) {
	u, _ := url.Parse(s.baseURL)
	u = u.JoinPath("v1/character/ability")
	q := u.Query()
	q.Add(ocidParam, ocid)
	u.RawQuery = q.Encode()

	k := u.String()

	req, err := http.NewRequest(http.MethodGet, k, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer, _ := io.ReadAll(b)

	r := GetCharacterAbilityResponse{}
	err = json.Unmarshal(buffer, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *service) GetSetCharacterAbility() (*GetCharacterAbilityResponse, error) {
	return s.GetCharacterAbility(s.ocid)
}

func (s *service) GetCharacterLinkSkill(ocid string) (*GetCharacterLinkSkillResponse, error) {
	u, _ := url.Parse(s.baseURL)
	u = u.JoinPath("v1/character/link-skill")
	q := u.Query()
	q.Add(ocidParam, ocid)
	u.RawQuery = q.Encode()

	k := u.String()

	req, err := http.NewRequest(http.MethodGet, k, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer, _ := io.ReadAll(b)

	r := GetCharacterLinkSkillResponse{}
	err = json.Unmarshal(buffer, &r)
	if err != nil {
		return nil, err
	}
	return &r, nil
}

func (s *service) GetSetCharacterLinkSkill() (*GetCharacterLinkSkillResponse, error) {
	return s.GetCharacterLinkSkill(s.ocid)
}
