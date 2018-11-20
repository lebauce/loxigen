:: # Copyright 2013, Big Switch Networks, Inc.
:: # Copyright 2018, Red Hat, Inc.
:: #
:: # LoxiGen is licensed under the Eclipse Public License, version 1.0 (EPL), with
:: # the following special exception:
:: #
:: # LOXI Exception
:: #
:: # As a special exception to the terms of the EPL, you may distribute libraries
:: # generated by LoxiGen (LoxiGen Libraries) under the terms of your choice, provided
:: # that copyright and licensing notices generated by LoxiGen are not altered or removed
:: # from the LoxiGen Libraries and the notice provided below is (i) included in
:: # the LoxiGen Libraries, if distributed in source code form and (ii) included in any
:: # documentation for the LoxiGen Libraries, if distributed in binary form.
:: #
:: # Notice: "Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler."
:: #
:: # You may not use this file except in compliance with the EPL or LOXI Exception. You may obtain
:: # a copy of the EPL at:
:: #
:: # http://www.eclipse.org/legal/epl-v10.html
:: #
:: # Unless required by applicable law or agreed to in writing, software
:: # distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
:: # WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
:: # EPL for the specific language governing permissions and limitations
:: # under the EPL.

package ${package}

import (
	"encoding/binary"

	"github.com/google/gopacket"
	"github.com/skydive-project/goloxi"
)

// TODO: set real types
:: if version.wire_version >= 3:
type OXM = Oxm
:: #endif
type uint128 = goloxi.Uint128
type Checksum128 [16]byte
type Bitmap128 uint128
type Bitmap512 struct {
	a, b, c, d uint128
}
type Unimplemented struct {}
type BSNVport uint16
type ControllerURI uint16

func (h *Header) MessageType() uint8 {
	return h.Type
}

:: fake_types = ["Checksum128", "Bitmap128", "Bitmap512", "BSNVport", "ControllerURI"]
:: for fake_type in fake_types:
func (self *${fake_type}) Decode(data []byte) error {
	return nil
}

func (self *${fake_type}) Serialize(encoder *goloxi.Encoder) error {
	return nil
}
:: #endfor

:: import loxi_ir.ir_offset as ir_offset
:: import go_gen.util as util
:: import loxi_globals
::
:: for name, v in ir_offset.of_mixed_types.items():
::     if version.wire_version in v:
::         goname = util.go_ident(name[:-2])
::         gotype = base_type = v[version.wire_version][:-2]
::         if gotype.startswith("of_"):
::             gotype = util.go_ident(gotype)
::         #endif
::         if goname != gotype:
::             ofproto = loxi_globals.ir[version]
::             member_ofclass = ofproto.class_by_name(base_type)
::             if not member_ofclass:
type ${goname} ${gotype}

func (self *${goname}) Serialize(encoder *goloxi.Encoder) error {
	encoder.Put${util.go_ident(gotype)}(${gotype}(*self))
	return nil
}

func (self *${goname}) Decode(data []byte) error {
::                 if base_type == "uint8":
	*self = ${goname}(data[0])
::                 else:
	*self = ${goname}(binary.BigEndian.${util.go_ident(gotype)}(data[:]))
::                 #endif
	return nil
}
::             else:
type ${goname} = ${gotype}
::             #endif
::         #endif
::     #endif
:: #endfor

func DecodeMessage(data []byte) (goloxi.Message, error) {
	header, err := decodeHeader(data)
	if err != nil {
		return nil, err
	}

	return header.(goloxi.Message), nil
}