package unitconverter

import (
	"fmt"
	"math"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var value *float64

// Object that is used to build up a conversion type
type Object struct {
	Name, NameShort string
	Unit            float64
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
		"output": {
			Description: "The converted input.",
			Type:        schema.TypeFloat,
			Computed:    true,
		},
	}
}

// FindObjectByName function to find the object in the array and return the index
func FindObjectByName(what string, array []Object) (idx Object) {
	for i, v := range array {
		if strings.EqualFold(what, v.Name) || what == v.NameShort {
			return array[i]
		}
	}
	return Object{"null", "n", -1}
}

// ConvertFunc function to convert the data
func ConvertFunc(Types func() []Object) func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {

		// Readying the needed variables
		input := d.Get("input").(float64)
		value = &input
		original := FindObjectByName(d.Get("original").(string), Types())
		wanted := FindObjectByName(d.Get("wanted").(string), Types())

		if original.Name == "null" || wanted.Name == "Name" {
			return fmt.Errorf("Unable to find the conversion type. Please make sure you are using the correct resource and type")
		}
		output := (original.Unit * input) / wanted.Unit
		d.Set("output", math.Round(output*1000)/1000)
		d.SetId(fmt.Sprintf("%.2f", output))
		return nil
	}
}
