package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceLengthConverter defined resource for the terraform plugin
func ResourceLengthConverter() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_length` converts from the given length type to the wanted length type",
		Create:        ConvertFunc(GetLengthTypes()),
		Read:          ReadNil,
		Delete:        schema.RemoveFromState,
		SchemaVersion: 1,
		Schema:        ObjectSchema(),
	}
}

// GetLengthTypes creating the length types needed for conversion
func GetLengthTypes() func() []Object {
	var m []Object

	return func() []Object {
		// Metric values
		m = append(m, Object{"Meter", "m", *value})
		m = append(m, Object{"Decameter", "dam", m[0].Unit * 10})
		m = append(m, Object{"Hectometer", "hm", m[1].Unit * 10})
		m = append(m, Object{"Kilometer", "km", m[2].Unit * 100})
		m = append(m, Object{"Decimeter", "dm", m[0].Unit / 10})
		m = append(m, Object{"Centimeter", "cm", m[4].Unit / 10})
		m = append(m, Object{"Millimeter", "mm", m[5].Unit / 10})
		m = append(m, Object{"Micrometer", "µm", m[6].Unit / 1000})
		m = append(m, Object{"Nanometer", "mm", m[7].Unit / 1000})

		// Imperial values
		m = append(m, Object{"Inch", "\"\"", m[0].Unit * 0.0254})
		m = append(m, Object{"Feet", "\"", m[8].Unit * 12})
		m = append(m, Object{"Yard", "yd", m[9].Unit * 3})
		m = append(m, Object{"Mile", "mi", m[0].Unit * 1609.344})
		return m
	}
}
