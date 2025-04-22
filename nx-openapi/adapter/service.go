package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"

	"github.com/amphipath/msea-scouter-utils/nx-openapi/types"
)

type OpenAPIService interface {
	GetCharacterEquipment(ocid string) (*GetCharacterEquipmentResponse, error)
	GetCharacterOCID(ign string) (*GetCharacterIDResponse, error)

	SetCharacter(IGN string)
	GetSetCharacterEquipment() (*GetCharacterEquipmentResponse, error)
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

	service struct {
		ocid    string
		baseURL string
		apiKey  string
		client  *http.Client
	}
)

const (
	headerApiKey  = "x-nxopen-api-key"
	getIdIGNParam = "character_name"
	ocidParam     = "ocid"
)

func NewService(baseURL, apiKey string) OpenAPIService {
	return &service{
		baseURL: baseURL,
		apiKey:  apiKey,
		client:  http.DefaultClient,
	}
}

func (s *service) GetCharacterOCID(ign string) (*GetCharacterIDResponse, error) {
	req, err := http.NewRequest(http.MethodGet, path.Join(s.baseURL, fmt.Sprintf("/v1/id?%s=%s", getIdIGNParam, ign)), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerApiKey, s.apiKey)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer := []byte{}
	b.Read(buffer)

	r := GetCharacterIDResponse{}

	err = json.Unmarshal(buffer, &r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (s *service) SetCharacter(ign string) {
	r, err := s.GetCharacterOCID(ign)
	if err == nil && r != nil && len(r.OCID) == 0 {
		s.ocid = r.OCID
	}
}

func (s *service) GetCharacterEquipment(ocid string) (*GetCharacterEquipmentResponse, error) {
	req, err := http.NewRequest(http.MethodGet, path.Join(s.baseURL, fmt.Sprintf("/v1/character/item-equipment?%s=%s", ocidParam, ocid)), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerApiKey, s.apiKey)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	b := res.Body
	defer b.Close()
	buffer := []byte{}
	b.Read(buffer)

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
