package additionalproperties

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"encoding/json"
	"github.com/Azure/go-autorest/autorest"
)

// The package's fully qualified name.
const fqdn = "tests/generated/additionalproperties"

// CatAPTrue ...
type CatAPTrue struct {
	autorest.Response `json:"-"`
	Friendly          *bool `json:"friendly,omitempty"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	ID                   *int32                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	// Status - READ-ONLY
	Status *bool `json:"status,omitempty"`
}

// MarshalJSON is the custom marshaler for CatAPTrue.
func (cat CatAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if cat.Friendly != nil {
		objectMap["friendly"] = cat.Friendly
	}
	if cat.ID != nil {
		objectMap["id"] = cat.ID
	}
	if cat.Name != nil {
		objectMap["name"] = cat.Name
	}
	for k, v := range cat.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for CatAPTrue struct.
func (cat *CatAPTrue) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		case "friendly":
			if v != nil {
				var friendly bool
				err = json.Unmarshal(*v, &friendly)
				if err != nil {
					return err
				}
				cat.Friendly = &friendly
			}
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if cat.AdditionalProperties == nil {
					cat.AdditionalProperties = make(map[string]interface{})
				}
				cat.AdditionalProperties[k] = additionalProperties
			}
		case "id":
			if v != nil {
				var ID int32
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				cat.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				cat.Name = &name
			}
		case "status":
			if v != nil {
				var status bool
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				cat.Status = &status
			}
		}
	}

	return nil
}

// Error ...
type Error struct {
	Status  *int32  `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
}

// PetAPInProperties ...
type PetAPInProperties struct {
	autorest.Response `json:"-"`
	ID                *int32  `json:"id,omitempty"`
	Name              *string `json:"name,omitempty"`
	// Status - READ-ONLY
	Status               *bool               `json:"status,omitempty"`
	AdditionalProperties map[string]*float64 `json:"additionalProperties"`
}

// MarshalJSON is the custom marshaler for PetAPInProperties.
func (paip PetAPInProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if paip.ID != nil {
		objectMap["id"] = paip.ID
	}
	if paip.Name != nil {
		objectMap["name"] = paip.Name
	}
	if paip.AdditionalProperties != nil {
		objectMap["additionalProperties"] = paip.AdditionalProperties
	}
	return json.Marshal(objectMap)
}

// PetAPInPropertiesWithAPString ...
type PetAPInPropertiesWithAPString struct {
	autorest.Response `json:"-"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]*string `json:""`
	ID                   *int32             `json:"id,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	// Status - READ-ONLY
	Status                *bool               `json:"status,omitempty"`
	OdataLocation         *string             `json:"@odata.location,omitempty"`
	AdditionalProperties1 map[string]*float64 `json:"additionalProperties"`
}

// MarshalJSON is the custom marshaler for PetAPInPropertiesWithAPString.
func (paipwas PetAPInPropertiesWithAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if paipwas.ID != nil {
		objectMap["id"] = paipwas.ID
	}
	if paipwas.Name != nil {
		objectMap["name"] = paipwas.Name
	}
	if paipwas.OdataLocation != nil {
		objectMap["@odata.location"] = paipwas.OdataLocation
	}
	if paipwas.AdditionalProperties1 != nil {
		objectMap["additionalProperties"] = paipwas.AdditionalProperties1
	}
	for k, v := range paipwas.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PetAPInPropertiesWithAPString struct.
func (paipwas *PetAPInPropertiesWithAPString) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties string
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if paipwas.AdditionalProperties == nil {
					paipwas.AdditionalProperties = make(map[string]*string)
				}
				paipwas.AdditionalProperties[k] = &additionalProperties
			}
		case "id":
			if v != nil {
				var ID int32
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				paipwas.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				paipwas.Name = &name
			}
		case "status":
			if v != nil {
				var status bool
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				paipwas.Status = &status
			}
		case "@odata.location":
			if v != nil {
				var odataLocation string
				err = json.Unmarshal(*v, &odataLocation)
				if err != nil {
					return err
				}
				paipwas.OdataLocation = &odataLocation
			}
		case "additionalProperties":
			if v != nil {
				var additionalProperties1 map[string]*float64
				err = json.Unmarshal(*v, &additionalProperties1)
				if err != nil {
					return err
				}
				paipwas.AdditionalProperties1 = additionalProperties1
			}
		}
	}

	return nil
}

// PetAPObject ...
type PetAPObject struct {
	autorest.Response `json:"-"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	ID                   *int32                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	// Status - READ-ONLY
	Status *bool `json:"status,omitempty"`
}

// MarshalJSON is the custom marshaler for PetAPObject.
func (pao PetAPObject) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pao.ID != nil {
		objectMap["id"] = pao.ID
	}
	if pao.Name != nil {
		objectMap["name"] = pao.Name
	}
	for k, v := range pao.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PetAPObject struct.
func (pao *PetAPObject) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if pao.AdditionalProperties == nil {
					pao.AdditionalProperties = make(map[string]interface{})
				}
				pao.AdditionalProperties[k] = additionalProperties
			}
		case "id":
			if v != nil {
				var ID int32
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pao.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pao.Name = &name
			}
		case "status":
			if v != nil {
				var status bool
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				pao.Status = &status
			}
		}
	}

	return nil
}

// PetAPString ...
type PetAPString struct {
	autorest.Response `json:"-"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]*string `json:""`
	ID                   *int32             `json:"id,omitempty"`
	Name                 *string            `json:"name,omitempty"`
	// Status - READ-ONLY
	Status *bool `json:"status,omitempty"`
}

// MarshalJSON is the custom marshaler for PetAPString.
func (pas PetAPString) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pas.ID != nil {
		objectMap["id"] = pas.ID
	}
	if pas.Name != nil {
		objectMap["name"] = pas.Name
	}
	for k, v := range pas.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PetAPString struct.
func (pas *PetAPString) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties string
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if pas.AdditionalProperties == nil {
					pas.AdditionalProperties = make(map[string]*string)
				}
				pas.AdditionalProperties[k] = &additionalProperties
			}
		case "id":
			if v != nil {
				var ID int32
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pas.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pas.Name = &name
			}
		case "status":
			if v != nil {
				var status bool
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				pas.Status = &status
			}
		}
	}

	return nil
}

// PetAPTrue ...
type PetAPTrue struct {
	autorest.Response `json:"-"`
	// AdditionalProperties - Unmatched properties from the message are deserialized this collection
	AdditionalProperties map[string]interface{} `json:""`
	ID                   *int32                 `json:"id,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	// Status - READ-ONLY
	Status *bool `json:"status,omitempty"`
}

// MarshalJSON is the custom marshaler for PetAPTrue.
func (pat PetAPTrue) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	if pat.ID != nil {
		objectMap["id"] = pat.ID
	}
	if pat.Name != nil {
		objectMap["name"] = pat.Name
	}
	for k, v := range pat.AdditionalProperties {
		objectMap[k] = v
	}
	return json.Marshal(objectMap)
}

// UnmarshalJSON is the custom unmarshaler for PetAPTrue struct.
func (pat *PetAPTrue) UnmarshalJSON(body []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(body, &m)
	if err != nil {
		return err
	}
	for k, v := range m {
		switch k {
		default:
			if v != nil {
				var additionalProperties interface{}
				err = json.Unmarshal(*v, &additionalProperties)
				if err != nil {
					return err
				}
				if pat.AdditionalProperties == nil {
					pat.AdditionalProperties = make(map[string]interface{})
				}
				pat.AdditionalProperties[k] = additionalProperties
			}
		case "id":
			if v != nil {
				var ID int32
				err = json.Unmarshal(*v, &ID)
				if err != nil {
					return err
				}
				pat.ID = &ID
			}
		case "name":
			if v != nil {
				var name string
				err = json.Unmarshal(*v, &name)
				if err != nil {
					return err
				}
				pat.Name = &name
			}
		case "status":
			if v != nil {
				var status bool
				err = json.Unmarshal(*v, &status)
				if err != nil {
					return err
				}
				pat.Status = &status
			}
		}
	}

	return nil
}
