//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//

package model

import (
	"encoding/xml"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by build-utils. DO NOT EDIT.

// The data-structure of this message
type S7ParameterReadVarResponse struct {
	NumItems uint8
	Parent   *S7Parameter
}

// The corresponding interface
type IS7ParameterReadVarResponse interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(io utils.WriteBuffer) error
	xml.Marshaler
	xml.Unmarshaler
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7ParameterReadVarResponse) ParameterType() uint8 {
	return 0x04
}

func (m *S7ParameterReadVarResponse) MessageType() uint8 {
	return 0x03
}

func (m *S7ParameterReadVarResponse) InitializeParent(parent *S7Parameter) {
}

func NewS7ParameterReadVarResponse(numItems uint8) *S7Parameter {
	child := &S7ParameterReadVarResponse{
		NumItems: numItems,
		Parent:   NewS7Parameter(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7ParameterReadVarResponse(structType interface{}) *S7ParameterReadVarResponse {
	castFunc := func(typ interface{}) *S7ParameterReadVarResponse {
		if casted, ok := typ.(S7ParameterReadVarResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*S7ParameterReadVarResponse); ok {
			return casted
		}
		if casted, ok := typ.(S7Parameter); ok {
			return CastS7ParameterReadVarResponse(casted.Child)
		}
		if casted, ok := typ.(*S7Parameter); ok {
			return CastS7ParameterReadVarResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7ParameterReadVarResponse) GetTypeName() string {
	return "S7ParameterReadVarResponse"
}

func (m *S7ParameterReadVarResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7ParameterReadVarResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Simple field (numItems)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7ParameterReadVarResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7ParameterReadVarResponseParse(io *utils.ReadBuffer) (*S7Parameter, error) {

	// Simple Field (numItems)
	numItems, _numItemsErr := io.ReadUint8(8)
	if _numItemsErr != nil {
		return nil, errors.Wrap(_numItemsErr, "Error parsing 'numItems' field")
	}

	// Create a partially initialized instance
	_child := &S7ParameterReadVarResponse{
		NumItems: numItems,
		Parent:   &S7Parameter{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7ParameterReadVarResponse) Serialize(io utils.WriteBuffer) error {
	ser := func() error {

		// Simple Field (numItems)
		numItems := uint8(m.NumItems)
		_numItemsErr := io.WriteUint8(8, (numItems))
		if _numItemsErr != nil {
			return errors.Wrap(_numItemsErr, "Error serializing 'numItems' field")
		}

		return nil
	}
	return m.Parent.SerializeParent(io, m, ser)
}

func (m *S7ParameterReadVarResponse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var token xml.Token
	var err error
	foundContent := false
	token = start
	for {
		switch token.(type) {
		case xml.StartElement:
			foundContent = true
			tok := token.(xml.StartElement)
			switch tok.Name.Local {
			case "numItems":
				var data uint8
				if err := d.DecodeElement(&data, &tok); err != nil {
					return err
				}
				m.NumItems = data
			}
		}
		token, err = d.Token()
		if err != nil {
			if err == io.EOF && foundContent {
				return nil
			}
			return err
		}
	}
}

func (m *S7ParameterReadVarResponse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if err := e.EncodeElement(m.NumItems, xml.StartElement{Name: xml.Name{Local: "numItems"}}); err != nil {
		return err
	}
	return nil
}

func (m S7ParameterReadVarResponse) String() string {
	return string(m.Box("S7ParameterReadVarResponse", utils.DefaultWidth*2))
}

func (m S7ParameterReadVarResponse) Box(name string, width int) utils.AsciiBox {
	if name == "" {
		name = "S7ParameterReadVarResponse"
	}
	boxes := make([]utils.AsciiBox, 0)
	boxes = append(boxes, utils.BoxAnything("NumItems", m.NumItems, width-2))
	return utils.BoxBox(name, utils.AlignBoxes(boxes, width-2), 0)
}
