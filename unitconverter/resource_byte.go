package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceByteConverter defined resource for the terraform plugin
func ResourceByteConverter() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_byte` converts from the given byte type to the wanted byte type",
		Create:        ConvertFunc(GetByteTypes()),
		Read:          ReadNil,
		Delete:        schema.RemoveFromState,
		SchemaVersion: 1,
		Schema:        ObjectSchema(),
	}
}

// GetByteTypes creating the bit types needed for conversion
func GetByteTypes() func() []Object {

	var m []Object

	return func() []Object {

		// Bits
		m = append(m, Object{"Bit", "b", *value})
		m = append(m, Object{"Kilobit", "Kb", m[0].Unit * 1000})
		m = append(m, Object{"Megabit", "Mb", m[1].Unit * 1000})
		m = append(m, Object{"Gigabit", "Gb", m[2].Unit * 1000})
		m = append(m, Object{"Terabit", "Tb", m[3].Unit * 1000})
		m = append(m, Object{"Petabit", "Pb", m[4].Unit * 1000})

		// BiBit
		m = append(m, Object{"Kibibit", "Kib", m[0].Unit * 1024})
		m = append(m, Object{"Mebibit", "Mib", m[5].Unit * 1024})
		m = append(m, Object{"Gibibit", "Gib", m[6].Unit * 1024})
		m = append(m, Object{"Tebibit", "Tib", m[7].Unit * 1024})
		m = append(m, Object{"Pebibit", "Pib", m[8].Unit * 1024})

		// Bytes
		m = append(m, Object{"Byte", "B", m[0].Unit * 8})
		m = append(m, Object{"Kilobyte", "KB", m[10].Unit * 1000})
		m = append(m, Object{"Megabyte", "MB", m[11].Unit * 1000})
		m = append(m, Object{"Gigabyte", "GB", m[12].Unit * 1000})
		m = append(m, Object{"Terabyte", "TB", m[13].Unit * 1000})
		m = append(m, Object{"Petabyte", "PB", m[14].Unit * 1000})

		// BiByte
		m = append(m, Object{"Kibibyte", "KiB", m[9].Unit * 1024})
		m = append(m, Object{"Mebibyte", "MiB", m[15].Unit * 1024})
		m = append(m, Object{"Gibibyte", "GiB", m[16].Unit * 1024})
		m = append(m, Object{"Tebibyte", "TiB", m[17].Unit * 1024})
		m = append(m, Object{"Pebibyte", "PiB", m[18].Unit * 1024})

		return m
	}
}
