package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataLengthConverter defined resource for the terraform plugin
func DataLengthConverter() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_length` converts from the given length type to the wanted length type",
		Create:        nil,
		Read:          ConvertFunc(GetLengthTypes()),
		Delete:        nil,
		SchemaVersion: 1,
		Schema:        ObjectSchema(),
	}
}

// GetLengthTypes creating the length types needed for conversion
func GetLengthTypes() func(float64, bool) []Object {

	return func(value float64, sublist bool) []Object {
		var m []Object

		// Metric values
		m = append(m, Object{"Nanometer", "mm", value})
		m = append(m, Object{"Micrometer", "µm", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Millimeter", "mm", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Centimeter", "cm", m[len(m)-1].Unit * 10})
		m = append(m, Object{"Decimeter", "dm", m[len(m)-1].Unit * 10})
		m = append(m, Object{"Meter", "m", m[len(m)-1].Unit * 10})
		m = append(m, Object{"Decameter", "dam", m[len(m)-1].Unit * 10})
		m = append(m, Object{"Hectometer", "hm", m[len(m)-1].Unit * 10})
		m = append(m, Object{"Kilometer", "km", m[len(m)-1].Unit * 10})

		// Imperial values
		m = append(m, Object{"Inch", "\"\"", m[5].Unit * 39.37})
		m = append(m, Object{"Feet", "\"", m[len(m)-1].Unit * 12})
		m = append(m, Object{"Mile", "mi", m[len(m)-1].Unit * 5280})
		m = append(m, Object{"Yard", "yd", m[len(m)-1].Unit * 1760})
		return m
	}
}
