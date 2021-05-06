package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{
			"unitconverter_time":      DataTimeConverter(),
			"unitconverter_length":    DataLengthConverter(),
			"unitconverter_byte":      DataByteConverter(),
			"unitconverter_math":      DataMath(),
			"unitconverter_step_down": DataMathDownStep(),
		},
	}
}
