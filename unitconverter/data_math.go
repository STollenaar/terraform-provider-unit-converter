package unitconverter

import (
	"errors"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataMath do simple math between mutliple resources of the same kind
func DataMath() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_math` does simple mathematical operations between 2 units",
		Create:        nil,
		Read:          PrepareMath(),
		Delete:        nil,
		SchemaVersion: 1,
		Schema:        mathFields(),
	}
}

// PrepareMath do the actual math for the resources
func PrepareMath() func(d *schema.ResourceData, meta interface{}) error {
	return func(d *schema.ResourceData, meta interface{}) error {
		operations := d.Get("operation").([]interface{})
		allTypes := allTypes()
		for index, operation := range operations {
			op := operation.(map[string]interface{})

			firstValue := op["first"].([]interface{})[0].(map[string]interface{})
			secondValue := op["second"].([]interface{})[0].(map[string]interface{})

			operant := op["operant"].(string)
			var baseUnit, unitFirst, unitSecond *Object
			var unitListFirst, unitListSecond []Object
			var errorUnit error

			// Finding the unit that is used
			for _, convertType := range allTypes {

				// Either getting the previous computed value and unit or getting the first unit
				if firstValue["previous"].(bool) && index > 0 {
					opPrevious := operations[index-1].(map[string]interface{})

					previousValue := opPrevious["first"].([]interface{})[0].(map[string]interface{})
					firstValue["value"] = d.Get("result").(float64)
					firstValue["unit"] = previousValue["unit"]

					unitListFirst = convertType(d.Get("result").(float64), false)
					unitFirst, errorUnit = FindObjectValueByName(previousValue["unit"].(string), unitListFirst)
				} else {
					unitListFirst = convertType(firstValue["value"].(float64), false)
					unitFirst, errorUnit = FindObjectValueByName(firstValue["unit"].(string), unitListFirst)
				}

				// Either getting the previous computed value and unit or getting the second unit
				if secondValue["previous"].(bool) && index > 0 {
					opPrevious := operations[index-1].(map[string]interface{})

					previousValue := opPrevious["second"].([]interface{})[0].(map[string]interface{})
					secondValue["value"] = d.Get("result").(float64)
					secondValue["unit"] = previousValue["unit"]

					unitListSecond = convertType(d.Get("result").(float64), false)
					unitSecond, errorUnit = FindObjectValueByName(previousValue["unit"].(string), unitListSecond)
				} else {
					unitListSecond = convertType(secondValue["value"].(float64), false)
					unitSecond, errorUnit = FindObjectValueByName(secondValue["unit"].(string), unitListSecond)
				}

				if errorUnit == nil {
					baseUnit = &unitListFirst[0]
					break
				}
			}
			if unitFirst == nil || unitSecond == nil {
				return errors.New("One or more unit types not supported. Please make sure they are of the same and use long name if still encountering this issue")
			}

			// Doing the math
			result, errorUnit := doMath(operant, unitFirst.Unit, unitSecond.Unit)
			if errorUnit != nil {
				return errorUnit
			}

			result = ConvertFuncMath(result, baseUnit.Unit, unitFirst.Unit)
			d.Set("result", result)
			d.SetId(fmt.Sprintf("%.2f", result))
		}

		return nil
	}
}

// doMath doing the actual operation
func doMath(operant string, firstValue float64, secondValue float64) (float64, error) {
	switch operant {
	case "+":
		return firstValue + secondValue, nil
	case "-":
		return firstValue - secondValue, nil
	case "*":
		return firstValue * secondValue, nil
	case "/":
		return firstValue / secondValue, nil
	default:
		return 0, errors.New("Unsupported operator")
	}
}

// allTypes making a slice of all the supported unit types
func allTypes() (types [](func(float64, bool) []Object)) {
	bytes := GetByteTypes()
	lengths := GetLengthTypes()
	times := GetTimeTypes()
	types = [](func(float64, bool) []Object){bytes, lengths, times}
	return types
}

// mathFields the schema used for the math resource
func mathFields() map[string]*schema.Schema {
	s := map[string]*schema.Schema{
		"operation": {
			Type:        schema.TypeList,
			Optional:    true,
			ForceNew:    true,
			Description: "List of environment variables to set in the container. Cannot be updated.",
			Elem: &schema.Resource{
				Schema: map[string]*schema.Schema{
					"first": {
						Type:        schema.TypeList,
						Required:    true,
						Description: "The first argument of the operation",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"value": {
									Type:     schema.TypeFloat,
									Optional: true,
									ForceNew: true,
								},
								"unit": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
								"previous": {
									Type:     schema.TypeBool,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
					"operant": {
						Type:        schema.TypeString,
						Required:    true,
						ForceNew:    true,
						Description: "The kind of operation",
					},
					"second": {
						Type:        schema.TypeList,
						Required:    true,
						ForceNew:    true,
						Description: "The first argument of the operation",
						Elem: &schema.Resource{
							Schema: map[string]*schema.Schema{
								"value": {
									Type:     schema.TypeFloat,
									Optional: true,
									ForceNew: true,
								},
								"unit": {
									Type:     schema.TypeString,
									Optional: true,
									ForceNew: true,
								},
								"previous": {
									Type:     schema.TypeBool,
									Optional: true,
									ForceNew: true,
								},
							},
						},
					},
				},
			},
		},
		"result": {
			Description: "The computed input.",
			Type:        schema.TypeFloat,
			Computed:    true,
		},
	}
	return s
}
