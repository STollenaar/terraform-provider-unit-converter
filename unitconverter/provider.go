package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProviders map[string]*schema.Provider
var testAccProvider *schema.Provider

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"unitconverter_time":   ResourceTimeConverter(),
			"unitconverter_length": ResourceLengthConverter(),
			"unitconverter_byte":   ResourceByteConverter(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
