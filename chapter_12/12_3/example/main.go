// 12.3 Рекурсивный вывод значения
package main

import (
	"reflect"
	"fmt"
	"strconv"
)

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T): \n", name, x)
	display(name, reflect.ValueOf(x))
}

func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" + strconv.FormatUint(uint64(v.Pointer()), 16)
	default:
		return v.Type().String() + " value"
	}
}

type Movie struct {
	Title, Subtitle string
	Year int
	Color bool
	Actor map[string]string
	Oscars []string
	Sequel *string
}

func main() {
	strangelove := Movie{
		Title: "Dr. Strangelove",
		Subtitle: "How I learned to stop worrying and love the bomb",
		Year: 1964,
		Color: false,
		Actor: map[string]string{
			"Dr. Strangelove": "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley": "Peter Sellers",
			"Gen. Buck Turgidson": "George C. Scott",
			"Brig. Gen. Jack D. Ripper": "Sterling Hayden",
			"Maj. T.J. King Kong": "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	
	Display("stranglove", strangelove)
	// stranglove.Color = false
 	// stranglove.Actor["Maj. T.J. King Kong"] = "Slim Pickens"
 	// stranglove.Actor["Dr. Strangelove"] = "Peter Sellers"
 	// stranglove.Actor["Grp. Capt. Lionel Mandrake"] = "Peter Sellers"
 	// stranglove.Actor["Pres. Merkin Muffley"] = "Peter Sellers"
 	// stranglove.Actor["Gen. Buck Turgidson"] = "George C. Scott"
 	// stranglove.Actor["Brig. Gen. Jack D. Ripper"] = "Sterling Hayden"
 	// stranglove.Oscars[0] = "Best Actor (Nomin.)"
 	// stranglove.Oscars[1] = "Best Adapted Screenplay (Nomin.)"
 	// stranglove.Oscars[2] = "Best Director (Nomin.)"
 	// stranglove.Oscars[3] = "Best Picture (Nomin.)"
	// stranglove.Sequel = nil
}

func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type(). Field(i) .Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path, formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
		display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf(" %s = nil \n ", path)
		} else {
			fmt.Printf(" %s.type = %s\n", path, v.Elem().Type())
			display(path + " .value ", v.Elem())
		}
	default:	
		fmt.Printf(" %s = %s\n", path , formatAtom(v))
	}
}

