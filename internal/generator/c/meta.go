/*
 * Copyright 2019 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cgenerator

import (
	"github.com/objectbox/objectbox-go/internal/generator/fbsparser/reflection"
	"github.com/objectbox/objectbox-go/internal/generator/model"
)

type fbsProperty struct {
	mProp    *model.Property
	fbsField *reflection.Field
}

// Merge implements model.PropertyMeta interface
func (mp *fbsProperty) Merge(property *model.Property) model.PropertyMeta {
	return &fbsProperty{property, mp.fbsField}
}

// CppType returns C++ variable name with reserved keywords suffixed by an underscore
func (mp *fbsProperty) CppName() string {
	if reservedKeywords[mp.mProp.Name] {
		return mp.mProp.Name + "_"
	}
	return mp.mProp.Name
}

// CppType returns C++ type name
func (mp *fbsProperty) CppType() string {
	var fbsType = mp.fbsField.Type(nil)
	var baseType = fbsType.BaseType()
	var cppType = fbsTypeToCppType[baseType]
	if baseType == reflection.BaseTypeVector {
		cppType = cppType + "<" + fbsTypeToCppType[fbsType.Element()] + ">"
	}
	return cppType
}

// FbOffsetFactory returns an offset factory used to build flatbuffers if this property is a complex type.
// See also FbOffsetType().
func (mp *fbsProperty) FbOffsetFactory() string {
	switch mp.mProp.Type {
	case model.PropertyTypeString:
		return "CreateString"
	case model.PropertyTypeByteVector:
		return "CreateVector"
	case model.PropertyTypeStringVector:
		return "CreateVectorOfStrings"
	}
	return ""
}

// FbOffsetType returns a type used to read flatbuffers if this property is a complex type.
// See also FbOffsetFactory().
func (mp *fbsProperty) FbOffsetType() string {
	switch mp.mProp.Type {
	case model.PropertyTypeString:
		return "flatbuffers::Vector<char>"
	case model.PropertyTypeByteVector:
		return "flatbuffers::Vector<" + fbsTypeToCppType[mp.fbsField.Type(nil).Element()] + ">"
	case model.PropertyTypeStringVector:
		return "" // NOTE custom handling in the template
	}
	return ""
}

// FbDefaultValue returns a default value for scalars
func (mp *fbsProperty) FbDefaultValue() string {
	if mp.mProp.Type == model.PropertyTypeFloat || mp.mProp.Type == model.PropertyTypeDouble {
		return "0.0"
	}
	return "0"
}