package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceTimeConverter defined resource for the terraform plugin
func ResourceTimeConverter() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_time` converts from the given time type to the wanted time type",
		Create:        ConvertFunc(GetTimeTypes()),
		Read:          ReadNil,
		Delete:        schema.RemoveFromState,
		SchemaVersion: 1,
		Schema:        ObjectSchema(),
	}
}

// GetTimeTypes creating the time types needed for conversion
func GetTimeTypes() func() []Object {
	var m []Object

	return func() []Object {
		m = append(m, Object{"Second", "S", *value})
		m = append(m, Object{"Miliseconds", "Ms", m[0].Unit * 1000})
		m = append(m, Object{"Minute", "M", m[0].Unit * 60})
		m = append(m, Object{"Hour", "H", m[2].Unit * 60})
		m = append(m, Object{"Day", "D", m[3].Unit * 24})
		m = append(m, Object{"Week", "W", m[4].Unit * 7})
		m = append(m, Object{"Year", "Y", m[5].Unit * 52})
		return m
	}
}
