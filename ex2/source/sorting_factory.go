package source

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tuanhuu162/go23_course/ex2/constants"
)

func SimpleBubbleSort[V string | float32](element []V) []V {
	for i := 0; i < len(element); i++ {
		for j := i + 1; j < len(element); j++ {
			if element[i] > element[j] {
				element[i], element[j] = element[j], element[i]
			}
		}
	}
	return element
}

type SortingFactory struct {
}

type SortingFactoryInterface interface {
	SortElement(elemen []string) string
}

type NumericSorting struct {
}

func (s *NumericSorting) SortElement(element []string) string {
	var numberic_element []float32
	for _, value := range element {
		float, _ := strconv.ParseFloat(value, 32)
		numberic_element = append(numberic_element, float32(float))
	}
	sorted_element := SimpleBubbleSort(numberic_element)
	var tmp_element []string
	for _, value := range sorted_element {
		tmp_element = append(tmp_element, strconv.FormatFloat(float64(value), 'f', -1, 32))
	}
	return strings.Join(tmp_element, " ")
}

type StringSorting struct {
}

func (s *StringSorting) SortElement(element []string) string {
	sorted_element := SimpleBubbleSort(element)
	return strings.Join(sorted_element, " ")
}

type MixedSorting struct {
}

func (m *MixedSorting) SortElement(element []string) string {
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
	var sorted_float string = numberSorting.SortElement(numeric_element)
	stringSorting := StringSorting{}
	var sorted_string string = stringSorting.SortElement(string_element)
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
		panic(fmt.Errorf("Unsupport sort type : %s", sortType))
	}
}
