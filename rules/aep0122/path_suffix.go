// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aep0122

import (
	"strings"

	"github.com/aep-dev/api-linter/lint"
	"github.com/aep-dev/api-linter/locations"
	"github.com/jhump/protoreflect/desc"
)

var pathSuffix = &lint.FieldRule{
	Name: lint.NewRuleName(122, "path-suffix"),
	OnlyIf: func(f *desc.FieldDescriptor) bool {
		return strings.HasSuffix(f.GetName(), "_path")
	},

	LintField: func(f *desc.FieldDescriptor) []lint.Problem {
		return []lint.Problem{{
			Message:    "Fields should not use the `_path` suffix.",
			Suggestion: strings.TrimSuffix(f.GetName(), "_path"),
			Descriptor: f,
			Location:   locations.DescriptorName(f),
		}}
	},
}
