# Copyright 2013, Big Switch Networks, Inc.
# Copyright 2018, Red Hat, Inc.
#
# LoxiGen is licensed under the Eclipse Public License, version 1.0 (EPL), with
# the following special exception:
#
# LOXI Exception
#
# As a special exception to the terms of the EPL, you may distribute libraries
# generated by LoxiGen (LoxiGen Libraries) under the terms of your choice, provided
# that copyright and licensing notices generated by LoxiGen are not altered or removed
# from the LoxiGen Libraries and the notice provided below is (i) included in
# the LoxiGen Libraries, if distributed in source code form and (ii) included in any
# documentation for the LoxiGen Libraries, if distributed in binary form.
#
# Notice: "Copyright 2013, Big Switch Networks, Inc. This library was generated by the LoxiGen Compiler."
#
# You may not use this file except in compliance with the EPL or LOXI Exception. You may obtain
# a copy of the EPL at:
#
# http://www.eclipse.org/legal/epl-v10.html
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# EPL for the specific language governing permissions and limitations
# under the EPL.

from collections import defaultdict
import subprocess
import os
import loxi_globals
import template_utils
import loxi_utils.loxi_utils as utils
import util
import loxi_ir 

# Return the filename and struct names for the generated Go struct
def generate_goname(ofclass):
    name = util.go_ident(ofclass.name[3:])
    if ofclass.is_message:
        return "message", name
    elif ofclass.is_oxm:
        return "oxm", name
    elif ofclass.is_oxs:
        return "oxs", name
    elif ofclass.is_action:
        return "action", name
    elif ofclass.is_action_id:
        return "action_id", name
    elif ofclass.is_instruction:
        return "instruction", name
    return 'common', name

def oftype_unherited_members(ofclass):
    unherited = []
    for i, member in enumerate(ofclass.members):
        if hasattr(member, "name") and ofclass.superclass and ofclass.superclass.member_by_name(member.name):
            continue
        if hasattr(member, "oftype") and member.oftype == "of_octets_t" and ofclass.discriminator and i == len(ofclass.members) - 1:
            continue
        unherited.append(member)
    return unherited

# Create intermediate representation, extended from the LOXI IR
def build_ofclasses(version):
    superclasses = {}
    field_lengths = {}
    modules = defaultdict(list)

    # for ofclass in loxi_globals.ir[version].classes:
    #     if ofclass.superclass:
    #         subclasses = ofclass.superclass.__dict__.setdefault("subclasses", [])
    #         subclasses.append(ofclass)

    for ofclass in loxi_globals.ir[version].classes:
        module_name, ofclass.goname = generate_goname(ofclass)
        ofclass.field_lengths = {}
        modules[module_name].append(ofclass)
        offset = 0
        for m in ofclass.members:
            if type(m) != loxi_ir.OFPadMember:
                m.goname = util.go_ident(m.name)
                if type(m) == loxi_ir.OFTypeMember and ofclass.superclass:
                    superclass = ofclass.superclass
                    member = superclass.member_by_name(m.name)
                    if member == None:
                        # openflow_input/ovs_tcp_flags specifies a value for experimenter_id
                        # that is defined nowhere (bug ?)
                        continue
                    discriminator_values = member.__dict__.setdefault("values", {})
                    discriminator_values[m.value] = ofclass

            if type(m) == loxi_ir.OFFieldLengthMember:
                ofclass.field_lengths[m.name] = m.field_name

        ofclass.unherited_members = oftype_unherited_members(ofclass)

        ofclass.embedded_length = ofclass.base_length
        if ofclass.virtual and len(ofclass.members):
            member = ofclass.members[-1]
            if type(member) == loxi_ir.OFPadMember:
                ofclass.embedded_length -= member.pad_length

    return modules

def codegen(install_dir):
    def render(name, template_name=None, **ctx):
        if template_name is None:
            template_name = os.path.basename(name)
        with template_utils.open_output(install_dir, name) as out:
            util.render_template(out, template_name, **ctx)
        subprocess.call(["go", "fmt", os.path.join(install_dir, name)])
        subprocess.call(["goimports", "-w", os.path.join(install_dir, name)])

    render('globals.go', versions=loxi_globals.OFVersions.all_supported)

    for version in loxi_globals.OFVersions.all_supported:
        subdir = 'of' + version.version.replace('.', '')
        modules = build_ofclasses(version)

        render(os.path.join(subdir, 'types.go'), package=subdir, version=version)

        render(os.path.join(subdir, 'const.go'), package=subdir, version=version,
               enums=loxi_globals.ir[version].enums)

        for name, ofclasses in modules.items():
            render(os.path.join(subdir, name + '.go'), template_name='module.go',
                   package=subdir, version=version, ofclasses=ofclasses,
                   subdir=subdir)