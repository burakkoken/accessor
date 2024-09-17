package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	MarkerAccessor        = "accessor"
	MarkerAccessorMapping = "accessor:mapping"
)

type AccessorMarker struct {
	Value string `marker:"Value,useValueSyntax"`
	Url   string `marker:"Url"`
}

func (a AccessorMarker) Validate() error {
	if strings.TrimSpace(a.Value) == "" {
		return errors.New("'Value' cannot be empty or nil")
	}

	if strings.TrimSpace(a.Url) == "" {
		return errors.New("'Url' cannot be empty or nil")
	}

	return nil
}

type AccessorMappingMarker struct {
	Value  string `marker:"Value,useValueSyntax"`
	Method string `marker:"Method"`
}

func (a AccessorMappingMarker) Validate() error {
	if strings.TrimSpace(a.Value) == "" {
		return errors.New("'Value' cannot be empty or nil")
	}

	if strings.TrimSpace(a.Method) == "" {
		return errors.New("'Url' cannot be empty or nil")
	}

	matched := false
	methods := []string{"GET", "POST", "HEAD", "OPTIONS", "PUT", "PATCH", "DELETE", "TRACE"}

	for _, method := range methods {
		if strings.TrimSpace(a.Method) == method {
			matched = true
		}
	}

	if !matched {
		return fmt.Errorf("invalid 'Method' argument value. Here is the list of valid values %s", strings.Join(methods, ", "))
	}

	return nil
}
