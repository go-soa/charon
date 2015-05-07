package routing

import (
	"errors"
	"net/url"
)

var (
	// ErrTemplateURLFuncInvalidNumberOfArguments  ...
	ErrTemplateURLFuncInvalidNumberOfArguments = errors.New("routing: template url func invalid number of argumetns")
	// ErrTemplateURLFuncInvalidArgument  ...
	ErrTemplateURLFuncInvalidArgument = errors.New("routing: template url func invalid argument")
)

// URLTemplateFunc ...
func URLTemplateFunc(urlGenerator URLGenerator) func(pathName string, values ...interface{}) (*url.URL, error) {
	return func(pathName string, values ...interface{}) (*url.URL, error) {
		routeName, dict, err := parseTemplateParams(pathName, values)
		if err != nil {
			return nil, err
		}

		return urlGenerator.Generate(RouteName(routeName), dict)
	}
}

// URLAbsTemplateFunc ...
func URLAbsTemplateFunc(urlGenerator URLGenerator) func(pathName string, values ...interface{}) (*url.URL, error) {
	return func(pathName string, values ...interface{}) (*url.URL, error) {
		routeName, dict, err := parseTemplateParams(pathName, values)
		if err != nil {
			return nil, err
		}

		return urlGenerator.GenerateAbs(RouteName(routeName), dict)
	}
}

func parseTemplateParams(pathName string, values []interface{}) (string, map[string]interface{}, error) {
	if len(values)%2 != 0 && len(values) >= 2 {
		return "", nil, ErrTemplateURLFuncInvalidNumberOfArguments
	}

	dict := make(map[string]interface{}, len(values)/2)
	if len(values) > 2 {
		for i := 0; i < len(values); i += 2 {
			key, ok := values[i].(string)
			if !ok {
				return "", nil, ErrTemplateURLFuncInvalidArgument
			}
			dict[key] = values[i+1]
		}
	}

	return pathName, dict, nil
}