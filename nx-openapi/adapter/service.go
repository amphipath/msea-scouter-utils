package adapter

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"

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
	u, _ := url.Parse(s.baseURL)
	u = u.JoinPath("v1/id")
	q := u.Query()
	q.Add(getIdIGNParam, ign)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add(headerApiKey, s.apiKey)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		k := string(b)
		return nil, errors.New(k)
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
	req.Header.Add(headerApiKey, s.apiKey)

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode >= http.StatusBadRequest {
		b, _ := io.ReadAll(res.Body)
		k := string(b)
		return nil, errors.New(k)
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
