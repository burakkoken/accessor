/*
Copyright © 2021 Accessor Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"github.com/procyon-projects/marker"
)

var (
	accessorInterfaces []marker.InterfaceType
)

// Register your marker definitions.
func RegisterDefinitions(registry *marker.Registry) error {
	markers := []struct {
		Name   string
		Level  marker.TargetLevel
		Output interface{}
	}{
		{Name: MarkerAccessor, Level: marker.InterfaceTypeLevel, Output: &AccessorMarker{}},
		{Name: MarkerAccessorMapping, Level: marker.InterfaceMethodLevel, Output: &AccessorMappingMarker{}},
	}

	for _, m := range markers {
		err := registry.Register(m.Name, PkgId, m.Level, m.Output)
		if err != nil {
			return err
		}
	}

	return nil
}

// Process your markers.
func ProcessMarkers(collector *marker.Collector, pkgs []*marker.Package) error {
	marker.EachFile(collector, pkgs, func(file *marker.File, err error) {
		findAccessorInterfaces(file.InterfaceTypes)
	})
	return nil
}

func findAccessorInterfaces(interfaceTypes []marker.InterfaceType) {
	for _, interfaceType := range interfaceTypes {
		markerValues := interfaceType.Markers

		if markerValues == nil {
			return
		}

		markers, ok := markerValues[MarkerAccessor]

		if !ok {
			return
		}

		for _, candidateMarker := range markers {
			switch candidateMarker.(type) {
			case AccessorMarker:
				accessorInterfaces = append(accessorInterfaces, interfaceType)
			}
		}

	}
}
