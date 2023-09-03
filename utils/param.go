package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetParamValue(c echo.Context, key string, typeName string, defaultValue interface{}) interface{} {
	value := c.Param(key)
	convertedValue, err := convertTypeValue(typeName, value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err))
	}
	if convertedValue == nil {
		return defaultValue
	}
	return convertedValue
}

func GetQueryParam(c echo.Context, key string, typeName string, defaultValue interface{}) interface{} {
	value := c.QueryParam(key)
	convertedValue, err := convertTypeValue(typeName, value)
	if err != nil {
		return c.JSON(http.StatusBadRequest, NewError(err))
	}
	if convertedValue == nil {
		return defaultValue
	}
	return convertedValue
}

func convertTypeValue(typeName string, value string) (interface{}, error) {
	if value == "" {
		return nil, nil
	}
	switch typeName {
	case "string":
		return value, nil
	case "integer":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		return intValue, nil
	default:
		return nil, errors.New("Unsupported type")
	}
}
