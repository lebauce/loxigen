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
::

package goloxi

import (
	"bytes"
	"encoding/binary"
)

type Encoder struct {
	buffer *bytes.Buffer
}

func NewEncoder() *Encoder {
	return &Encoder{
		buffer: new(bytes.Buffer), 
	}
}

func (e *Encoder) PutChar(c byte) {
	e.buffer.WriteByte(c)
}

func (e *Encoder) PutUint8(i uint8) {
	e.buffer.WriteByte(i)
}

func (e *Encoder) PutUint16(i uint16) {
	var tmp [2]byte
	binary.BigEndian.PutUint16(tmp[0:2], i)
	e.buffer.Write(tmp[:])
}

func (e *Encoder) PutUint32(i uint32) {
	var tmp [4]byte
	binary.BigEndian.PutUint32(tmp[0:4], i)
	e.buffer.Write(tmp[:])
}

func (e *Encoder) PutUint64(i uint64) {
	var tmp [8]byte
	binary.BigEndian.PutUint64(tmp[0:8], i)
	e.buffer.Write(tmp[:])
}

func (e *Encoder) PutUint128(i Uint128) {
	var tmp [16]byte
	binary.BigEndian.PutUint64(tmp[0:8], i.Hi)
	binary.BigEndian.PutUint64(tmp[8:16], i.Lo)
	e.buffer.Write(tmp[:])
}

func (e *Encoder) Write(b []byte) {
	e.buffer.Write(b)
}

func (e *Encoder) Bytes() []byte {
	return e.buffer.Bytes()
}

func (e *Encoder) SkipAlign() {
	length := len(e.buffer.Bytes())
	e.Write(bytes.Repeat([]byte{0}, (length+7)/8*8 - length))
}