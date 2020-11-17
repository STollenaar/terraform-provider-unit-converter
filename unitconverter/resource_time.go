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
func GetTimeTypes() []Object {
	m := []Object{
		{"Year", "Y", 1, 52},
		{"Week", "W", 52, 7},
		{"Day", "D", 7, 24},
		{"Hour", "H", 24, 60},
		{"Minute", "M", 60, 60},
		{"Second", "S", 60, 1000},
		{"Milisecond", "Ms", 1000, 1000000},
	}
	return m
}
