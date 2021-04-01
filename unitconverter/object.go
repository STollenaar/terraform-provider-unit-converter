package unitconverter

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Object that is used to build up a conversion type
type Object struct {
	Name      string  `json:"Name"`
	NameShort string  `json:"NameShort"`
	Unit      float64 `json:"Unit"`
}

// ReadNil generic function to return nil for the read operation
func ReadNil(d *schema.ResourceData, meta interface{}) error {
	return nil
}

// ObjectSchema schema used for the conversion functions
func ObjectSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"wanted": {
			Description: "The wanted type.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"original": {
			Description: "The orignal type.",
			Type:        schema.TypeString,
			Required:    true,
			ForceNew:    true,
		},
		"input": {
			Description: "The value to convert.",
			Type:        schema.TypeFloat,
			Required:    true,
			ForceNew:    true,
		},
		"sublist_only": {
			Description: "An optional value to only use a small sublist",
			Type:        schema.TypeBool,
			Optional:    true,
			Default:     false,
			ForceNew:    true,
		},
		"output": {
			Description: "The converted input.",
			Type:        schema.TypeFloat,
			Computed:    true,
		},
	}
}

// FindObjectValueByName function to find the object in the array and return the index
func FindObjectValueByName(what string, array []Object) (*Object, error) {
	for i, v := range array {
		if strings.EqualFold(what, v.Name) || v.NameShort == what {
			return &array[i], nil
		}
	}
	return nil, errors.New("Can't find unit in objects array")

}

// FindObjectIndexByName function to find the object in the array and return the index
func FindObjectIndexByName(what string, array []Object) (int, error) {
	for i, v := range array {
		if strings.EqualFold(what, v.Name) || v.NameShort == what {
			return i, nil
		}
	}
	return -1, errors.New("Can't find unit in objects array")

}

// ConvertFunc function to convert the data
func ConvertFunc(Types func(float64, bool) []Object) func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {

		// Readying the needed variables
		input := d.Get("input").(float64)

		original, unitError := FindObjectValueByName(d.Get("original").(string), Types(input, d.Get("sublist_only").(bool)))
		wanted, unitError := FindObjectValueByName(d.Get("wanted").(string), Types(input, d.Get("sublist_only").(bool)))

		if unitError != nil {
			return fmt.Errorf("Unable to find the conversion type. Please make sure you are using the correct resource and type")
		}
		output := ConvertFuncMath(input, original.Unit, wanted.Unit)
		d.Set("output", output)
		d.SetId(fmt.Sprintf("%.2f", output))
		return nil
	}
}

// ConvertFuncMath function to do the actual math
func ConvertFuncMath(input float64, orginalUnit float64, wantedUnit float64) float64 {
	output := (orginalUnit * input) / wantedUnit
	return math.Round(output*1000) / 1000
}
