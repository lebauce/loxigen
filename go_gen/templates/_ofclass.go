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
:: from loxi_ir import *
:: import go_gen.oftype
:: import go_gen.util as util
:: import loxi_utils.loxi_utils as loxi_utils
::
type ${ofclass.goname} struct {
:: # Embed superclass
:: if ofclass.superclass:
	*${util.go_ident(ofclass.superclass.name)}
:: #endif
::
:: # Create struct properties
:: for member in ofclass.unherited_members:
::     if type(member) != OFPadMember:
::         if ofclass.superclass and ofclass.superclass.member_by_name(member.name):
::             continue
::         #endif
::
::         oftype = go_gen.oftype.lookup_type_data(member.oftype, version)
::         if not oftype and loxi_utils.oftype_is_list(member.oftype):
::             oftype = "[]goloxi.Serializable" # util.go_ident(loxi_utils.oftype_list_elem(member.oftype))
::         elif oftype != None:
::             oftype = oftype.name
::         else:
::             raise Exception("Could not determine type for %s in %s" % (member.name, ofclass.name))
::         #endif
::
	${member.goname} ${oftype}
::     #endif
:: #endfor
}

:: base_length = ofclass.embedded_length
:: base_offset = 0
:: for member in ofclass.unherited_members:
::     if type(member) != OFPadMember:
::         base_offset = member.offset
::         break
::     #endif
:: #endfor
::
:: if ofclass.superclass:
::     base_length -= (base_offset if len(ofclass.unherited_members) else ofclass.superclass.embedded_length)
:: #endif
::
:: type_members = [m for m in ofclass.unherited_members if type(m) == OFTypeMember]
::
:: include('_serialize.go', ofclass=ofclass, members=ofclass.unherited_members, type_members=type_members,
::                          base_length=base_length)

:: include('_decode.go', ofclass=ofclass, members=ofclass.unherited_members, base_length=base_length, base_offset=base_offset)
::
:: include('_constructor.go', ofclass=ofclass, type_members=type_members)