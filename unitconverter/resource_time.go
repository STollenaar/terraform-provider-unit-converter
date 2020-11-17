package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceTimeConverter() *schema.Resource {
	return &schema.Resource{
		Description: "The resource `unitconverter_time` converts from the given time type to the wanted time type",
		Create: convertTimeFunc(),
		Read: readNil,
		Delete: schema.RemoveFromState,
		SchemaVersion: 1,
		Schema: timeSchema(),
	}
}

func convertTimeFunc() func(d *schema.ResourceData, meta interface{}) error {
 return func(d *schema.ResourceData, meta interface{}) error {
	//  original := d.Get("original").(string)
	//  wanted := d.Get("wanted").(string)
	 input := d.Get("input").(int)


	d.Set("output", int(input))
	d.SetId(string(strconv.Itoa(input)))
	 return nil
 }
}

func readNil(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func timeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"wanted": {
			Description: "The wanted time type.",
			Type: schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"original": {
			Description: "The orignal time type.",
			Type: schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"input": {
			Description: "The value to convert.",
			Type: schema.TypeInt,
			Required: true,
			ForceNew: true,
		},
		"output": {
			Description: "The converted input.",
			Computed: true,
			Type:        schema.TypeInt,
		},
	}
}