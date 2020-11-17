package unitconverter

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Object that is used to build up a conversion type
type Object struct {
	Name, NameShort         string
	StepBigger, StepSmaller int
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
			Type:        schema.TypeInt,
			Required:    true,
			ForceNew:    true,
		},
		"output": {
			Description: "The converted input.",
			Type:        schema.TypeInt,
			Computed:    true,
		},
	}
}

// FindObjectByName function to find the object in the array and return the index
func FindObjectByName(what string, array []Object) (idx int) {
	for i, v := range array {
		if strings.EqualFold(what, v.Name) || strings.EqualFold(what, v.NameShort) {
			return i
		}
	}
	return -1
}

// ConvertFunc function to convert the data
func ConvertFunc(Types []Object) func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {

		// Readying the needed variables
		original := FindObjectByName(d.Get("original").(string), Types)
		wanted := FindObjectByName(d.Get("wanted").(string), Types)
		input := d.Get("input").(int)

		if original == -1 || wanted == -1 {
			return fmt.Errorf("Unable to find the conversion type. Please make sure you are using the correct resource and type")
		}

		// Calculating the total conversion delta
		mod := 1
		if original > wanted {
			mod = -1
		}

		total := 0
		for original != wanted {
			if mod > 0 {
				total += Types[original].StepSmaller
			} else {
				total += Types[original].StepBigger
			}
			original += mod
		}

		// Doing the actual conversion
		var output int
		if mod > 0 {
			output = input * total
		} else {
			output = input / total
		}

		d.Set("output", output)
		d.SetId(string(strconv.Itoa(output)))
		return nil
	}
}
