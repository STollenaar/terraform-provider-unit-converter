package unitconverter

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataMathDownStep do simple math between mutliple resources of the same kind
func DataMathDownStep() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_math_downstep` converts the units value to a smaller one until its a whole number",
		Create:        nil,
		Read:          StepDown(),
		Delete:        nil,
		SchemaVersion: 1,
		Schema:        stepFields(),
	}
}

// StepDown does the step down
func StepDown() func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		value := d.Get("input_value").(float64)
		unit := d.Get("input_unit").(string)

		resultValue, resultUnit, errorUnit := doDownstep(value, unit)
		d.Set("output_value", resultValue)
		d.Set("output_unit", resultUnit)
		d.SetId(fmt.Sprintf("%.2f", value))
		return errorUnit
	}
}

func doDownstep(startValue float64, startUnit string) (float64, string, error) {
	allTypes := allTypes()

	var unitIndex int
	var unitList []Object
	var errorUnit error
	// Finding the unit that is used
	for _, convertType := range allTypes {
		unitList = convertType(startValue, false)
		unitIndex, errorUnit = FindObjectIndexByName(startUnit, unitList)

		if errorUnit == nil {
			break
		}
	}
	startObject, errorUnit := FindObjectValueByName(startUnit, unitList)
	unitObject := unitList[unitIndex]
	for !isIntegral(startValue) && unitIndex > 0 {
		unitObject = unitList[unitIndex]
		startValue = ConvertFuncMath(startValue, startObject.Unit, unitObject.Unit)
		unitIndex--
	}
	if !isIntegral(unitObject.Unit) {
		errorUnit = errors.New("Unable to step down unit")
	}
	return startValue, unitObject.Name, errorUnit
}

func isIntegral(val float64) bool {
	return val == float64(int(val))
}

// stepFields the schema used for the math resource
func stepFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"input_value": {
			Type:     schema.TypeFloat,
			Required: true,
			ForceNew: true,
		},
		"input_unit": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"output_value": {
			Type:     schema.TypeFloat,
			Computed: true,
		},
		"output_unit": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
	}
	return s
}
