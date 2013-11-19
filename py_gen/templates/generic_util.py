:: # Copyright 2013, Big Switch Networks, Inc.
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
:: include('_copyright.py')
"""
Utility functions independent of the protocol version
"""

:: include('_autogen.py')

import loxi
import struct

def unpack_list(reader, deserializer):
    """
    The deserializer function should take an OFReader and return the new object.
    """
    entries = []
    while not reader.is_empty():
        entries.append(deserializer(reader))
    return entries

def unpack_list_lv16(reader, deserializer):
    """
    The deserializer function should take an OFReader and return the new object.
    """
    def wrapper(reader):
        length, = reader.peek('!H')
        return deserializer(reader.slice(length))
    return unpack_list(reader, wrapper)

def unpack_list_tlv16(reader, deserializer):
    """
    The deserializer function should take an OFReader and an integer type
    and return the new object.
    """
    def wrapper(reader):
        typ, length, = reader.peek('!HH')
        return deserializer(reader.slice(length), typ)
    return unpack_list(reader, wrapper)

def pad_to(alignment, length):
    """
    Return a string of zero bytes that will pad a string of length 'length' to
    a multiple of 'alignment'.
    """
    return "\x00" * ((length + alignment - 1)/alignment*alignment - length)

class OFReader(object):
    """
    Cursor over a read-only buffer

    OpenFlow messages are best thought of as a sequence of elements of
    variable size, rather than a C-style struct with fixed offsets and
    known field lengths. This class supports efficiently reading
    fields sequentially and is intended to be used recursively by the
    parsers of child objects which will implicitly update the offset.
    """
    def __init__(self, buf):
        self.buf = buf
        self.offset = 0

    def read(self, fmt):
        st = struct.Struct(fmt)
        if self.offset + st.size > len(self.buf):
            raise loxi.ProtocolError("Buffer too short")
        result = st.unpack_from(self.buf, self.offset)
        self.offset += st.size
        return result

    def read_all(self):
        buf = buffer(self.buf, self.offset)
        self.offset += len(buf)
        return str(buf)

    def peek(self, fmt):
        st = struct.Struct(fmt)
        if self.offset + st.size > len(self.buf):
            raise loxi.ProtocolError("Buffer too short")
        result = st.unpack_from(self.buf, self.offset)
        return result

    def skip(self, length):
        if self.offset + length > len(self.buf):
            raise loxi.ProtocolError("Buffer too short")
        self.offset += length

    def skip_align(self):
        new_offset = (self.offset + 7) / 8 * 8
        if new_offset > len(self.buf):
            raise loxi.ProtocolError("Buffer too short")
        self.offset = new_offset

    def is_empty(self):
        return self.offset == len(self.buf)

    # Used when parsing variable length objects which have external length
    # fields (e.g. the actions list in an OF 1.0 packet-out message).
    def slice(self, length):
        if self.offset + length > len(self.buf):
            raise loxi.ProtocolError("Buffer too short")
        buf = OFReader(buffer(self.buf, self.offset, length))
        self.offset += length
        return buf