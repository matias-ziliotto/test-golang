package validation

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

/*
	Params:
		-struct to validate
		-slice of required fields to validate
	Return:
		-bool: if everything was validated ok or not
		-string: required field error message
*/
func ValidateRequiredData(structRequest interface{}, requiredFields []string) (bool, string) {
	structTypeOf := reflect.TypeOf(structRequest)
	structValueOf := reflect.ValueOf(structRequest)

	for _, field := range requiredFields {
		fieldByName, fieldFound := structTypeOf.FieldByName(field) // Try to get the field

		if !fieldFound {
			return false, "field '" + strings.ToLower(field) + "'not found"
		}

		// Get the type of the field
		typeOfVField := fieldByName.Type.Kind()
		// Get value of the field as string format
		valueOfField := fmt.Sprintf("%v", structValueOf.FieldByName(field).Interface())

		if !validateRequiredField(typeOfVField, valueOfField) {
			return false, "field '" + strings.ToLower(field) + "' is required"
		}
	}

	return true, ""
}

/*
	Params:
		-field type: "string", "int", etc.
		-value of field in string format: "this is a string", "2005.50", "true", etc.
	Return:
		-bool: if field was validated ok or not
*/
func validateRequiredField(fieldType reflect.Kind, value string) bool {
	switch fieldType {
	case reflect.String:
		if value != "" {
			return true
		}
	case reflect.Float32, reflect.Float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err == nil && floatVal != 0 {
			return true
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		intVal, err := strconv.Atoi(value)
		if err == nil && intVal != 0 {
			return true
		}
	case reflect.Bool:
		if value != "" && (value == "true" || value == "false") {
			return true
		}
	default:
		return false
	}

	return false
}
