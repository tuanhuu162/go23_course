package source

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tuanhuu162/go23_course/ex2/constants"
)

func SimpleBubbleSort[V string | float32](element []V, is_desc bool) []V {
	var compare_function func(a, b V) bool
	if is_desc {
		compare_function = func(a, b V) bool {
			return a < b
		}
	} else {
		compare_function = func(a, b V) bool {
			return a > b
		}
	}
	for i := 0; i < len(element); i++ {
		for j := i + 1; j < len(element); j++ {
			if compare_function(element[i], element[j]) {
				element[i], element[j] = element[j], element[i]
			}
		}
	}
	return element
}

type SortingFactory struct {
}

type SortingFactoryInterface interface {
	SortElement(elemen []string, is_desc bool) string
}

type NumericSorting struct {
}

func (s *NumericSorting) SortElement(element []string, is_desc bool) string {
	var numberic_element []float32
	for _, value := range element {
		float, err := strconv.ParseFloat(value, 32)
		if err != nil {
			fmt.Println("Cannot parse float/int:", value)
			return ""
		}
		numberic_element = append(numberic_element, float32(float))
	}
	sorted_element := SimpleBubbleSort(numberic_element, is_desc)
	var tmp_element []string
	for _, value := range sorted_element {
		tmp_element = append(tmp_element, strconv.FormatFloat(float64(value), 'f', -1, 32))
	}
	return strings.Join(tmp_element, " ")
}

type StringSorting struct {
}

func (s *StringSorting) SortElement(element []string, is_desc bool) string {
	sorted_element := SimpleBubbleSort(element, is_desc)
	return strings.Join(sorted_element, " ")
}

type MixedSorting struct {
}

func (m *MixedSorting) SortElement(element []string, is_desc bool) string {
	var numeric_element []string
	var string_element []string

	for _, value := range element {
		if _, err := strconv.ParseFloat(value, 32); err == nil {
			numeric_element = append(numeric_element, value)
		} else {
			string_element = append(string_element, value)
		}
	}
	numberSorting := NumericSorting{}
	var sorted_float string = numberSorting.SortElement(numeric_element, is_desc)
	stringSorting := StringSorting{}
	var sorted_string string = stringSorting.SortElement(string_element, is_desc)
	return sorted_float + " " + sorted_string

}

func (s *SortingFactory) GetSorter(sortType string) SortingFactoryInterface {
	switch constants.SortingType(sortType) {
	case constants.NumericSorting:
		return &NumericSorting{}
	case constants.StringSorting:
		return &StringSorting{}
	case constants.MixedSorting:
		return &MixedSorting{}
	default:
		fmt.Println("Unsupport sort type :", sortType)
		return nil
	}
}

func PerformSort(sortType string, element []string, is_desc bool) {
	sortingFactory := SortingFactory{}
	sorter := sortingFactory.GetSorter(sortType)
	var finalValue string = sorter.SortElement(element, is_desc)
	if finalValue == "" {
		return
	}
	fmt.Println("Output:", finalValue)
}
