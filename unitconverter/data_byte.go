package unitconverter

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// DataByteConverter defined resource for the terraform plugin
func DataByteConverter() *schema.Resource {
	return &schema.Resource{
		Description:   "The resource `unitconverter_byte` converts from the given byte type to the wanted byte type",
		Create:        nil,
		Read:          ConvertFunc(GetByteTypes()),
		Delete:        nil,
		SchemaVersion: 1,
		Schema:        ObjectSchema(),
	}
}

// GetByteTypes creating the bit types needed for conversion
func GetByteTypes() func(float64, bool) []Object {

	return func(value float64, sublist bool) []Object {
		var m []Object

		if !sublist {

			// Bits
			m = append(m, Object{"Bit", "b", value})
			m = append(m, Object{"Kilobit", "Kb", m[len(m)-1].Unit * 1000})
			m = append(m, Object{"Megabit", "Mb", m[len(m)-1].Unit * 1000})
			m = append(m, Object{"Gigabit", "Gb", m[len(m)-1].Unit * 1000})
			m = append(m, Object{"Terabit", "Tb", m[len(m)-1].Unit * 1000})
			m = append(m, Object{"Petabit", "Pb", m[len(m)-1].Unit * 1000})

			// BiBit
			m = append(m, Object{"Kibibit", "Kib", value * 1024})
			m = append(m, Object{"Mebibit", "Mib", m[len(m)-1].Unit * 1024})
			m = append(m, Object{"Gibibit", "Gib", m[len(m)-1].Unit * 1024})
			m = append(m, Object{"Tebibit", "Tib", m[len(m)-1].Unit * 1024})
			m = append(m, Object{"Pebibit", "Pib", m[len(m)-1].Unit * 1024})
		}

		// Bytes
		m = append(m, Object{"Byte", "B", value * 8})
		m = append(m, Object{"Kilobyte", "KB", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Megabyte", "MB", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Gigabyte", "GB", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Terabyte", "TB", m[len(m)-1].Unit * 1000})
		m = append(m, Object{"Petabyte", "PB", m[len(m)-1].Unit * 1000})

		// BiByte
		m = append(m, Object{"Kibibyte", "KiB", m[len(m)-6].Unit * 1024})
		m = append(m, Object{"Mebibyte", "MiB", m[len(m)-1].Unit * 1024})
		m = append(m, Object{"Gibibyte", "GiB", m[len(m)-1].Unit * 1024})
		m = append(m, Object{"Tebibyte", "TiB", m[len(m)-1].Unit * 1024})
		m = append(m, Object{"Pebibyte", "PiB", m[len(m)-1].Unit * 1024})

		return m
	}
}
