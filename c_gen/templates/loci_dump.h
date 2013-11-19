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
:: include('_copyright.c')

#if !defined(_LOCI_DUMP_H_)
#define _LOCI_DUMP_H_

#include <loci/loci_base.h>
#include <loci/of_match.h>
#include <stdio.h>

/* g++ requires this to pick up PRI, etc.
 * See  http://gcc.gnu.org/ml/gcc-help/2006-10/msg00223.html
 */
#if !defined(__STDC_FORMAT_MACROS)
#define __STDC_FORMAT_MACROS
#endif
#include <inttypes.h>

typedef int (*loci_obj_dump_f)(loci_writer_f writer, void *cookie, of_object_t *obj);

/****************************************************************
 *
 * Per-datatype dump macros
 *
 ****************************************************************/
#define LOCI_DUMP_u8(writer, cookie, val) writer(cookie, "%u", (val))
#define LOCI_DUMP_u16(writer, cookie, val) writer(cookie, "%u (0x%x)", (val), (val))
#define LOCI_DUMP_u32(writer, cookie, val) writer(cookie, "%u (0x%x)", (val), (val))
#define LOCI_DUMP_u64(writer, cookie, val) writer(cookie, "%" PRIu64 "(0x%" PRIx64 ")", (val), (val))

/* @todo Add checks for special port numbers */
#define LOCI_DUMP_port_no(writer, cookie, val) LOCI_DUMP_u32(writer, cookie, val)
#define LOCI_DUMP_fm_cmd(writer, cookie, val) LOCI_DUMP_u16(writer, cookie, val)

/* @todo Decode wildcards */
#define LOCI_DUMP_wc_bmap(writer, cookie, val)         \
    writer(cookie, "0x%" PRIx64, (val))
#define LOCI_DUMP_match_bmap(writer, cookie, val)      \
    writer(cookie, "0x%" PRIx64, (val))

/* @todo Dump first N bytes of data */
#define LOCI_DUMP_octets(writer, cookie, val)                                      \
    writer(cookie, "%d bytes at location %p", (val).bytes, (val).data)

#define LOCI_DUMP_mac(writer, cookie, val)                                 \
    writer(cookie, "%02x:%02x:%02x:%02x:%02x:%02x",            \
               (val).addr[0], (val).addr[1], (val).addr[2],     \
               (val).addr[3], (val).addr[4], (val).addr[5])

#define LOCI_DUMP_ipv4(writer, cookie, val)                                        \
    writer(cookie, "%d.%d.%d.%d", val >> 24, (val >> 16) & 0xff,       \
               (val >> 8) & 0xff, val & 0xff)

#define LOCI_DUMP_ipv6(writer, cookie, val)                                        \
    writer(cookie, "%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x:%02x%02x", \
               (val).addr[0], (val).addr[1], (val).addr[2], (val).addr[3], \
               (val).addr[4], (val).addr[5], (val).addr[6], (val).addr[7], \
               (val).addr[8], (val).addr[9], (val).addr[10], (val).addr[11], \
               (val).addr[12], (val).addr[13], (val).addr[14], (val).addr[15])

#define LOCI_DUMP_string(writer, cookie, val) writer(cookie, "%s", val)

#define LOCI_DUMP_port_name(writer, cookie, val) LOCI_DUMP_string(writer, cookie, val)
#define LOCI_DUMP_tab_name(writer, cookie, val) LOCI_DUMP_string(writer, cookie, val)
#define LOCI_DUMP_desc_str(writer, cookie, val) LOCI_DUMP_string(writer, cookie, val)
#define LOCI_DUMP_ser_num(writer, cookie, val) LOCI_DUMP_string(writer, cookie, val)

int loci_dump_match(loci_writer_f writer, void* cookie, of_match_t *match);
#define LOCI_DUMP_match(writer, cookie, val) loci_dump_match(writer, cookie, &val)

#define LOCI_DUMP_bitmap_128(writer, cookie, val) writer(cookie, "%" PRIx64 "%" PRIx64, (val).hi, (val).lo)

/**
 * Generic version for any object
 */
int of_object_dump(loci_writer_f writer, void *cookie, of_object_t *obj);
#endif /* _LOCI_DUMP_H_ */