// Copyright 2021 The CC Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cc // import "modernc.org/cc/v4"

import (
	"bytes"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"sync"
	"sync/atomic"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/pmezard/go-difflib/difflib"
	"modernc.org/ccorpus"
	"modernc.org/httpfs"
	"modernc.org/mathutil"
)

var (
	cfs         = ccorpus.FileSystem()
	cFS         = &corpusFS{cfs}
	corpus      = map[string][]byte{}
	corpusIndex []string
	re          *regexp.Regexp
	defaultCfg0 *Config
	builtin     = `
#ifdef __SIZE_TYPE__
typedef __SIZE_TYPE__ size_t;
#else
#error __SIZE_TYPE__ undefined
#endif

#ifdef __WCHAR_TYPE__
typedef __WCHAR_TYPE__ wchar_t;
#else
#error __WCHAR_TYPE__ undefined
#endif

#ifdef __PTRDIFF_TYPE__
typedef __PTRDIFF_TYPE__ ptrdiff_t;
#else
#error __PTRDIFF_TYPE__ undefined
#endif

#define __builtin_offsetof(type, member) ((size_t)&(((type*)0)->member))
#define __builtin_types_compatible_p(t1, t2) (sizeof(t1) == sizeof(t2))
#define __extension__


#ifndef __builtin_va_list
#define __builtin_va_list __builtin_va_list
typedef void *__builtin_va_list;
#endif

#ifndef __builtin_va_arg
#define __builtin_va_arg(va, type) (*(type*)__builtin_va_arg_sink(0, va))
void *__builtin_va_arg_sink(int, ...);
#endif

#ifdef __clang__
#define __builtin_bit_cast(type, arg) (*(type*)&arg)
#define __builtin_convertvector(src, type) (*(type*)&src)
#define __builtin_bit_cast(type, arg) (*(type*)&arg)
#elif defined(_WIN32) || defined(WIN32) || defined(_WIN64) || defined(WIN64)
#endif

#ifdef __UINT16_TYPE__
__UINT16_TYPE__ __builtin_bswap16 (__UINT16_TYPE__);
#endif

#ifdef __UINT32_TYPE__
__UINT32_TYPE__ __builtin_bswap32 (__UINT32_TYPE__);
#endif

#ifdef __UINT64_TYPE__
__UINT64_TYPE__ __builtin_bswap64 (__UINT64_TYPE__ x);
#endif

void __builtin_exit(int status);
int __builtin_printf(const char *format, ...);
void __ccgo_dmesg(char*, ...);

// No operations, only for tests.
#define __builtin_va_copy(dst, src)
#define __builtin_va_end(ap)
#define __builtin_va_start(ap, v)

//TODO #define __builtin_constant_p(x) __builtin_constant_p_impl(0, x)
//TODO __builtin___stpcpy_chk (check.go:1818:check:)
//TODO __builtin_alloca (check.go:1818:check:)
//TODO __builtin_apply (check.go:1818:check:)
//TODO __builtin_bcopy (check.go:1818:check:)
//TODO __builtin_ceilf (check.go:1818:check:)
//TODO __builtin_choose_expr (check.go:1818:check:)
//TODO __builtin_conjf (check.go:1818:check:)
//TODO __builtin_ffs (check.go:1818:check:)
//TODO __builtin_fprintf (check.go:1818:check:)
//TODO __builtin_fprintf_unlocked (check.go:1818:check:)
//TODO __builtin_fputc (check.go:1818:check:)
//TODO __builtin_fputs (check.go:1818:check:)
//TODO __builtin_fputs_unlocked (check.go:1818:check:)
//TODO __builtin_fwrite (check.go:1818:check:)
//TODO __builtin_huge_vall (check.go:1818:check:)
//TODO __builtin_ia32_2intersectd128 (check.go:1818:check:)
//TODO __builtin_ia32_2intersectd256 (check.go:1818:check:)
//TODO __builtin_ia32_2intersectd512 (check.go:1818:check:)
//TODO __builtin_ia32_2intersectq128 (check.go:1818:check:)
//TODO __builtin_ia32_2intersectq256 (check.go:1818:check:)
//TODO __builtin_ia32_2intersectq512 (check.go:1818:check:)
//TODO __builtin_ia32_4fmaddps (check.go:1818:check:)
//TODO __builtin_ia32_4fmaddps_mask (check.go:1818:check:)
//TODO __builtin_ia32_4fmaddss (check.go:1818:check:)
//TODO __builtin_ia32_4fmaddss_mask (check.go:1818:check:)
//TODO __builtin_ia32_4fnmaddps (check.go:1818:check:)
//TODO __builtin_ia32_4fnmaddps_mask (check.go:1818:check:)
//TODO __builtin_ia32_4fnmaddss (check.go:1818:check:)
//TODO __builtin_ia32_4fnmaddss_mask (check.go:1818:check:)
//TODO __builtin_ia32_addcarryx_u32 (check.go:1818:check:)
//TODO __builtin_ia32_addcarryx_u64 (check.go:1818:check:)
//TODO __builtin_ia32_addpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_addpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_addpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_addps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_addps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_addps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_addsd (check.go:1818:check:)
//TODO __builtin_ia32_addsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_addss (check.go:1818:check:)
//TODO __builtin_ia32_addss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_addsubpd (check.go:1818:check:)
//TODO __builtin_ia32_addsubpd256 (check.go:1818:check:)
//TODO __builtin_ia32_addsubps (check.go:1818:check:)
//TODO __builtin_ia32_addsubps256 (check.go:1818:check:)
//TODO __builtin_ia32_aesdec128 (check.go:1818:check:)
//TODO __builtin_ia32_aesdeclast128 (check.go:1818:check:)
//TODO __builtin_ia32_aesenc128 (check.go:1818:check:)
//TODO __builtin_ia32_aesenclast128 (check.go:1818:check:)
//TODO __builtin_ia32_aesimc128 (check.go:1818:check:)
//TODO __builtin_ia32_andnotsi256 (check.go:1818:check:)
//TODO __builtin_ia32_andnpd (check.go:1818:check:)
//TODO __builtin_ia32_andnpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_andnpd256 (check.go:1818:check:)
//TODO __builtin_ia32_andnpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_andnpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_andnps (check.go:1818:check:)
//TODO __builtin_ia32_andnps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_andnps256 (check.go:1818:check:)
//TODO __builtin_ia32_andnps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_andnps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_andpd (check.go:1818:check:)
//TODO __builtin_ia32_andpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_andpd256 (check.go:1818:check:)
//TODO __builtin_ia32_andpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_andpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_andps (check.go:1818:check:)
//TODO __builtin_ia32_andps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_andps256 (check.go:1818:check:)
//TODO __builtin_ia32_andps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_andps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_bextr_u32 (check.go:1818:check:)
//TODO __builtin_ia32_bextr_u64 (check.go:1818:check:)
//TODO __builtin_ia32_blendmd_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_blendmpd_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_blendmps_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_blendmq_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_blendvpd (check.go:1818:check:)
//TODO __builtin_ia32_blendvpd256 (check.go:1818:check:)
//TODO __builtin_ia32_blendvps (check.go:1818:check:)
//TODO __builtin_ia32_blendvps256 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf32x2_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf32x2_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf32x4_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf32x4_512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf32x8_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf64x2_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf64x2_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastf64x4_512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x2_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x2_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x2_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x4_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x4_512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti32x8_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti64x2_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti64x2_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcasti64x4_512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmb128 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmb256 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmb512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmw128 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmw256 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastmw512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastsd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastsd512 (check.go:1818:check:)
//TODO __builtin_ia32_broadcastss128_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastss256_mask (check.go:1818:check:)
//TODO __builtin_ia32_broadcastss512 (check.go:1818:check:)
//TODO __builtin_ia32_bzhi_di (check.go:1818:check:)
//TODO __builtin_ia32_bzhi_si (check.go:1818:check:)
//TODO __builtin_ia32_cldemote (check.go:1818:check:)
//TODO __builtin_ia32_clflush (check.go:1818:check:)
//TODO __builtin_ia32_clflushopt (check.go:1818:check:)
//TODO __builtin_ia32_clrssbsy (check.go:1818:check:)
//TODO __builtin_ia32_clwb (check.go:1818:check:)
//TODO __builtin_ia32_cmpb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpeqpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpeqps (check.go:1818:check:)
//TODO __builtin_ia32_cmpeqsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpeqss (check.go:1818:check:)
//TODO __builtin_ia32_cmpgepd (check.go:1818:check:)
//TODO __builtin_ia32_cmpgeps (check.go:1818:check:)
//TODO __builtin_ia32_cmpgtpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpgtps (check.go:1818:check:)
//TODO __builtin_ia32_cmplepd (check.go:1818:check:)
//TODO __builtin_ia32_cmpleps (check.go:1818:check:)
//TODO __builtin_ia32_cmplesd (check.go:1818:check:)
//TODO __builtin_ia32_cmpless (check.go:1818:check:)
//TODO __builtin_ia32_cmpltpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpltps (check.go:1818:check:)
//TODO __builtin_ia32_cmpltsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpltss (check.go:1818:check:)
//TODO __builtin_ia32_cmpneqpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpneqps (check.go:1818:check:)
//TODO __builtin_ia32_cmpneqsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpneqss (check.go:1818:check:)
//TODO __builtin_ia32_cmpngepd (check.go:1818:check:)
//TODO __builtin_ia32_cmpngeps (check.go:1818:check:)
//TODO __builtin_ia32_cmpngtpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpngtps (check.go:1818:check:)
//TODO __builtin_ia32_cmpnlepd (check.go:1818:check:)
//TODO __builtin_ia32_cmpnleps (check.go:1818:check:)
//TODO __builtin_ia32_cmpnlesd (check.go:1818:check:)
//TODO __builtin_ia32_cmpnless (check.go:1818:check:)
//TODO __builtin_ia32_cmpnltpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpnltps (check.go:1818:check:)
//TODO __builtin_ia32_cmpnltsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpnltss (check.go:1818:check:)
//TODO __builtin_ia32_cmpordpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpordps (check.go:1818:check:)
//TODO __builtin_ia32_cmpordsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpordss (check.go:1818:check:)
//TODO __builtin_ia32_cmppd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpunordpd (check.go:1818:check:)
//TODO __builtin_ia32_cmpunordps (check.go:1818:check:)
//TODO __builtin_ia32_cmpunordsd (check.go:1818:check:)
//TODO __builtin_ia32_cmpunordss (check.go:1818:check:)
//TODO __builtin_ia32_cmpw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cmpw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_comieq (check.go:1818:check:)
//TODO __builtin_ia32_comige (check.go:1818:check:)
//TODO __builtin_ia32_comigt (check.go:1818:check:)
//TODO __builtin_ia32_comile (check.go:1818:check:)
//TODO __builtin_ia32_comilt (check.go:1818:check:)
//TODO __builtin_ia32_comineq (check.go:1818:check:)
//TODO __builtin_ia32_comisdeq (check.go:1818:check:)
//TODO __builtin_ia32_comisdge (check.go:1818:check:)
//TODO __builtin_ia32_comisdgt (check.go:1818:check:)
//TODO __builtin_ia32_comisdle (check.go:1818:check:)
//TODO __builtin_ia32_comisdlt (check.go:1818:check:)
//TODO __builtin_ia32_comisdneq (check.go:1818:check:)
//TODO __builtin_ia32_compressdf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressdf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressdf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressdi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressdi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressdi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresshi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresshi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresshi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compresssi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoredi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoresi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_compressstoreuqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_crc32di (check.go:1818:check:)
//TODO __builtin_ia32_crc32hi (check.go:1818:check:)
//TODO __builtin_ia32_crc32qi (check.go:1818:check:)
//TODO __builtin_ia32_crc32si (check.go:1818:check:)
//TODO __builtin_ia32_cvtb2mask128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtb2mask256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtb2mask512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtd2mask128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtd2mask256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtd2mask512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2pd (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2pd256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2ps (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2ps256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtdq2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2b128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2b256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2b512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2d128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2d256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2d512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2q128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2q256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2q512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2w128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2w256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtmask2w512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v16hi (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v16hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v32hi (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v32hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v8hi (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtne2ps2bf16_v8hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v16sf (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v16sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v16sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v4sf (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v4sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v4sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v8sf (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v8sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtneps2bf16_v8sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2dq (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2dq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2dq256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2dq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2dq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2pi (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2ps (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2ps256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2ps_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2qq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2qq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2qq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2udq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2udq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2udq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2uqq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2uqq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpd2uqq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtpi2pd (check.go:1818:check:)
//TODO __builtin_ia32_cvtpi2ps (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2dq (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2dq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2dq256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2dq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2dq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pd (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pd256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2pi (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2qq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2qq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2qq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2udq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2udq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2udq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2uqq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2uqq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtps2uqq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtq2mask128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtq2mask256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtq2mask512 (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtqq2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtsd2si (check.go:1818:check:)
//TODO __builtin_ia32_cvtsd2si64 (check.go:1818:check:)
//TODO __builtin_ia32_cvtsd2ss (check.go:1818:check:)
//TODO __builtin_ia32_cvtsi2sd (check.go:1818:check:)
//TODO __builtin_ia32_cvtsi2ss (check.go:1818:check:)
//TODO __builtin_ia32_cvtsi642sd (check.go:1818:check:)
//TODO __builtin_ia32_cvtsi642ss (check.go:1818:check:)
//TODO __builtin_ia32_cvtss2sd (check.go:1818:check:)
//TODO __builtin_ia32_cvtss2si (check.go:1818:check:)
//TODO __builtin_ia32_cvtss2si64 (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2dq (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2dq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2dq256 (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2dq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2dq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2pi (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2qq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2qq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2qq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2udq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2udq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2udq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2uqq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2uqq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttpd2uqq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2dq (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2dq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2dq256 (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2dq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2dq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2pi (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2qq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2qq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2qq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2udq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2udq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2udq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2uqq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2uqq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttps2uqq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvttsd2si (check.go:1818:check:)
//TODO __builtin_ia32_cvttsd2si64 (check.go:1818:check:)
//TODO __builtin_ia32_cvttss2si (check.go:1818:check:)
//TODO __builtin_ia32_cvttss2si64 (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtudq2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtuqq2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_cvtusi2sd32 (check.go:1818:check:)
//TODO __builtin_ia32_cvtusi2sd64 (check.go:1818:check:)
//TODO __builtin_ia32_cvtusi2ss32 (check.go:1818:check:)
//TODO __builtin_ia32_cvtusi2ss64 (check.go:1818:check:)
//TODO __builtin_ia32_cvtw2mask128 (check.go:1818:check:)
//TODO __builtin_ia32_cvtw2mask256 (check.go:1818:check:)
//TODO __builtin_ia32_cvtw2mask512 (check.go:1818:check:)
//TODO __builtin_ia32_directstoreu_u32 (check.go:1818:check:)
//TODO __builtin_ia32_directstoreu_u64 (check.go:1818:check:)
//TODO __builtin_ia32_divpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_divpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_divpd_mask (check.go:1818:check:)
//TODO __builtin_ia32_divps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_divps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_divps_mask (check.go:1818:check:)
//TODO __builtin_ia32_divsd (check.go:1818:check:)
//TODO __builtin_ia32_divsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_divss (check.go:1818:check:)
//TODO __builtin_ia32_divss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v16sf (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v16sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v16sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v4sf (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v4sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v4sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v8sf (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v8sf_mask (check.go:1818:check:)
//TODO __builtin_ia32_dpbf16ps_v8sf_maskz (check.go:1818:check:)
//TODO __builtin_ia32_emms (check.go:1818:check:)
//TODO __builtin_ia32_enqcmd (check.go:1818:check:)
//TODO __builtin_ia32_enqcmds (check.go:1818:check:)
//TODO __builtin_ia32_expanddf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddf128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expanddf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddf256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expanddf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddf512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expanddi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expanddi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expanddi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expanddi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandhi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandhi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandhi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddf512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloaddi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadhi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadqi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsf512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandloadsi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandqi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandqi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandqi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsf128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsf128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsf256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsf512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_expandsi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_expandsi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_extractf32x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_extractf64x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_extracti32x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_fxrstor (check.go:1818:check:)
//TODO __builtin_ia32_fxrstor64 (check.go:1818:check:)
//TODO __builtin_ia32_fxsave (check.go:1818:check:)
//TODO __builtin_ia32_fxsave64 (check.go:1818:check:)
//TODO __builtin_ia32_getexppd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_getexppd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_getexpps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_getexpps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_haddpd (check.go:1818:check:)
//TODO __builtin_ia32_haddpd256 (check.go:1818:check:)
//TODO __builtin_ia32_haddps (check.go:1818:check:)
//TODO __builtin_ia32_haddps256 (check.go:1818:check:)
//TODO __builtin_ia32_hsubpd (check.go:1818:check:)
//TODO __builtin_ia32_hsubpd256 (check.go:1818:check:)
//TODO __builtin_ia32_hsubps (check.go:1818:check:)
//TODO __builtin_ia32_hsubps256 (check.go:1818:check:)
//TODO __builtin_ia32_incsspq (check.go:1818:check:)
//TODO __builtin_ia32_insertf32x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_insertf64x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_inserti32x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_inserti64x4_mask (check.go:1818:check:)
//TODO __builtin_ia32_kadddi (check.go:1818:check:)
//TODO __builtin_ia32_kaddhi (check.go:1818:check:)
//TODO __builtin_ia32_kaddqi (check.go:1818:check:)
//TODO __builtin_ia32_kaddsi (check.go:1818:check:)
//TODO __builtin_ia32_kanddi (check.go:1818:check:)
//TODO __builtin_ia32_kandhi (check.go:1818:check:)
//TODO __builtin_ia32_kandndi (check.go:1818:check:)
//TODO __builtin_ia32_kandnhi (check.go:1818:check:)
//TODO __builtin_ia32_kandnqi (check.go:1818:check:)
//TODO __builtin_ia32_kandnsi (check.go:1818:check:)
//TODO __builtin_ia32_kandqi (check.go:1818:check:)
//TODO __builtin_ia32_kandsi (check.go:1818:check:)
//TODO __builtin_ia32_kmovb (check.go:1818:check:)
//TODO __builtin_ia32_kmovd (check.go:1818:check:)
//TODO __builtin_ia32_kmovq (check.go:1818:check:)
//TODO __builtin_ia32_kmovw (check.go:1818:check:)
//TODO __builtin_ia32_knotdi (check.go:1818:check:)
//TODO __builtin_ia32_knothi (check.go:1818:check:)
//TODO __builtin_ia32_knotqi (check.go:1818:check:)
//TODO __builtin_ia32_knotsi (check.go:1818:check:)
//TODO __builtin_ia32_kordi (check.go:1818:check:)
//TODO __builtin_ia32_korhi (check.go:1818:check:)
//TODO __builtin_ia32_korqi (check.go:1818:check:)
//TODO __builtin_ia32_korsi (check.go:1818:check:)
//TODO __builtin_ia32_kortestcdi (check.go:1818:check:)
//TODO __builtin_ia32_kortestchi (check.go:1818:check:)
//TODO __builtin_ia32_kortestcqi (check.go:1818:check:)
//TODO __builtin_ia32_kortestcsi (check.go:1818:check:)
//TODO __builtin_ia32_kortestzdi (check.go:1818:check:)
//TODO __builtin_ia32_kortestzhi (check.go:1818:check:)
//TODO __builtin_ia32_kortestzqi (check.go:1818:check:)
//TODO __builtin_ia32_kortestzsi (check.go:1818:check:)
//TODO __builtin_ia32_ktestcdi (check.go:1818:check:)
//TODO __builtin_ia32_ktestchi (check.go:1818:check:)
//TODO __builtin_ia32_ktestcqi (check.go:1818:check:)
//TODO __builtin_ia32_ktestcsi (check.go:1818:check:)
//TODO __builtin_ia32_ktestzdi (check.go:1818:check:)
//TODO __builtin_ia32_ktestzhi (check.go:1818:check:)
//TODO __builtin_ia32_ktestzqi (check.go:1818:check:)
//TODO __builtin_ia32_ktestzsi (check.go:1818:check:)
//TODO __builtin_ia32_kunpckdi (check.go:1818:check:)
//TODO __builtin_ia32_kunpckhi (check.go:1818:check:)
//TODO __builtin_ia32_kunpcksi (check.go:1818:check:)
//TODO __builtin_ia32_kxnordi (check.go:1818:check:)
//TODO __builtin_ia32_kxnorhi (check.go:1818:check:)
//TODO __builtin_ia32_kxnorqi (check.go:1818:check:)
//TODO __builtin_ia32_kxnorsi (check.go:1818:check:)
//TODO __builtin_ia32_kxordi (check.go:1818:check:)
//TODO __builtin_ia32_kxorhi (check.go:1818:check:)
//TODO __builtin_ia32_kxorqi (check.go:1818:check:)
//TODO __builtin_ia32_kxorsi (check.go:1818:check:)
//TODO __builtin_ia32_lddqu (check.go:1818:check:)
//TODO __builtin_ia32_lddqu256 (check.go:1818:check:)
//TODO __builtin_ia32_ldmxcsr (check.go:1818:check:)
//TODO __builtin_ia32_lfence (check.go:1818:check:)
//TODO __builtin_ia32_loadapd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadapd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadapd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadaps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadaps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadaps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqudi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqudi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqudi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddquqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqusi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqusi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loaddqusi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadhpd (check.go:1818:check:)
//TODO __builtin_ia32_loadhps (check.go:1818:check:)
//TODO __builtin_ia32_loadlpd (check.go:1818:check:)
//TODO __builtin_ia32_loadlps (check.go:1818:check:)
//TODO __builtin_ia32_loadsd_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadss_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadupd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadupd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadupd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadups128_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadups256_mask (check.go:1818:check:)
//TODO __builtin_ia32_loadups512_mask (check.go:1818:check:)
//TODO __builtin_ia32_lzcnt_u16 (check.go:1818:check:)
//TODO __builtin_ia32_lzcnt_u32 (check.go:1818:check:)
//TODO __builtin_ia32_lzcnt_u64 (check.go:1818:check:)
//TODO __builtin_ia32_maskloadd (check.go:1818:check:)
//TODO __builtin_ia32_maskloadd256 (check.go:1818:check:)
//TODO __builtin_ia32_maskloadpd (check.go:1818:check:)
//TODO __builtin_ia32_maskloadpd256 (check.go:1818:check:)
//TODO __builtin_ia32_maskloadps (check.go:1818:check:)
//TODO __builtin_ia32_maskloadps256 (check.go:1818:check:)
//TODO __builtin_ia32_maskloadq (check.go:1818:check:)
//TODO __builtin_ia32_maskloadq256 (check.go:1818:check:)
//TODO __builtin_ia32_maskmovdqu (check.go:1818:check:)
//TODO __builtin_ia32_maskstored (check.go:1818:check:)
//TODO __builtin_ia32_maskstored256 (check.go:1818:check:)
//TODO __builtin_ia32_maskstorepd (check.go:1818:check:)
//TODO __builtin_ia32_maskstorepd256 (check.go:1818:check:)
//TODO __builtin_ia32_maskstoreps (check.go:1818:check:)
//TODO __builtin_ia32_maskstoreps256 (check.go:1818:check:)
//TODO __builtin_ia32_maskstoreq (check.go:1818:check:)
//TODO __builtin_ia32_maskstoreq256 (check.go:1818:check:)
//TODO __builtin_ia32_maxpd (check.go:1818:check:)
//TODO __builtin_ia32_maxpd256 (check.go:1818:check:)
//TODO __builtin_ia32_maxpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxpd_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxps (check.go:1818:check:)
//TODO __builtin_ia32_maxps256 (check.go:1818:check:)
//TODO __builtin_ia32_maxps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxps_mask (check.go:1818:check:)
//TODO __builtin_ia32_maxsd (check.go:1818:check:)
//TODO __builtin_ia32_maxsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_maxss (check.go:1818:check:)
//TODO __builtin_ia32_maxss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_mfence (check.go:1818:check:)
//TODO __builtin_ia32_minpd (check.go:1818:check:)
//TODO __builtin_ia32_minpd256 (check.go:1818:check:)
//TODO __builtin_ia32_minpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_minpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_minpd_mask (check.go:1818:check:)
//TODO __builtin_ia32_minps (check.go:1818:check:)
//TODO __builtin_ia32_minps256 (check.go:1818:check:)
//TODO __builtin_ia32_minps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_minps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_minps_mask (check.go:1818:check:)
//TODO __builtin_ia32_minsd (check.go:1818:check:)
//TODO __builtin_ia32_minsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_minss (check.go:1818:check:)
//TODO __builtin_ia32_minss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_monitor (check.go:1818:check:)
//TODO __builtin_ia32_movapd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movapd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movapd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movaps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movaps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movaps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movddup128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movddup256 (check.go:1818:check:)
//TODO __builtin_ia32_movddup256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movddup512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdir64b (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32load128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32load256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32load512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32store128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32store256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa32store512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64load128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64load256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64load512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64store128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64store256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdqa64store512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movdquqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movesd_mask (check.go:1818:check:)
//TODO __builtin_ia32_movess_mask (check.go:1818:check:)
//TODO __builtin_ia32_movhlps (check.go:1818:check:)
//TODO __builtin_ia32_movlhps (check.go:1818:check:)
//TODO __builtin_ia32_movmskpd (check.go:1818:check:)
//TODO __builtin_ia32_movmskpd256 (check.go:1818:check:)
//TODO __builtin_ia32_movmskps (check.go:1818:check:)
//TODO __builtin_ia32_movmskps256 (check.go:1818:check:)
//TODO __builtin_ia32_movntdq (check.go:1818:check:)
//TODO __builtin_ia32_movntdq256 (check.go:1818:check:)
//TODO __builtin_ia32_movntdq512 (check.go:1818:check:)
//TODO __builtin_ia32_movntdqa (check.go:1818:check:)
//TODO __builtin_ia32_movntdqa256 (check.go:1818:check:)
//TODO __builtin_ia32_movntdqa512 (check.go:1818:check:)
//TODO __builtin_ia32_movnti (check.go:1818:check:)
//TODO __builtin_ia32_movnti64 (check.go:1818:check:)
//TODO __builtin_ia32_movntpd (check.go:1818:check:)
//TODO __builtin_ia32_movntpd256 (check.go:1818:check:)
//TODO __builtin_ia32_movntpd512 (check.go:1818:check:)
//TODO __builtin_ia32_movntps (check.go:1818:check:)
//TODO __builtin_ia32_movntps256 (check.go:1818:check:)
//TODO __builtin_ia32_movntps512 (check.go:1818:check:)
//TODO __builtin_ia32_movntq (check.go:1818:check:)
//TODO __builtin_ia32_movq128 (check.go:1818:check:)
//TODO __builtin_ia32_movsd (check.go:1818:check:)
//TODO __builtin_ia32_movshdup (check.go:1818:check:)
//TODO __builtin_ia32_movshdup128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movshdup256 (check.go:1818:check:)
//TODO __builtin_ia32_movshdup256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movshdup512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movsldup (check.go:1818:check:)
//TODO __builtin_ia32_movsldup128_mask (check.go:1818:check:)
//TODO __builtin_ia32_movsldup256 (check.go:1818:check:)
//TODO __builtin_ia32_movsldup256_mask (check.go:1818:check:)
//TODO __builtin_ia32_movsldup512_mask (check.go:1818:check:)
//TODO __builtin_ia32_movss (check.go:1818:check:)
//TODO __builtin_ia32_mulpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulpd_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulps_mask (check.go:1818:check:)
//TODO __builtin_ia32_mulsd (check.go:1818:check:)
//TODO __builtin_ia32_mulsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_mulss (check.go:1818:check:)
//TODO __builtin_ia32_mulss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_mwait (check.go:1818:check:)
//TODO __builtin_ia32_orpd (check.go:1818:check:)
//TODO __builtin_ia32_orpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_orpd256 (check.go:1818:check:)
//TODO __builtin_ia32_orpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_orpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_orps (check.go:1818:check:)
//TODO __builtin_ia32_orps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_orps256 (check.go:1818:check:)
//TODO __builtin_ia32_orps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_orps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsb (check.go:1818:check:)
//TODO __builtin_ia32_pabsb128 (check.go:1818:check:)
//TODO __builtin_ia32_pabsb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsb256 (check.go:1818:check:)
//TODO __builtin_ia32_pabsb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsd (check.go:1818:check:)
//TODO __builtin_ia32_pabsd128 (check.go:1818:check:)
//TODO __builtin_ia32_pabsd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsd256 (check.go:1818:check:)
//TODO __builtin_ia32_pabsd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsw (check.go:1818:check:)
//TODO __builtin_ia32_pabsw128 (check.go:1818:check:)
//TODO __builtin_ia32_pabsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsw256 (check.go:1818:check:)
//TODO __builtin_ia32_pabsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pabsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_packssdw (check.go:1818:check:)
//TODO __builtin_ia32_packssdw128 (check.go:1818:check:)
//TODO __builtin_ia32_packssdw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_packssdw256 (check.go:1818:check:)
//TODO __builtin_ia32_packssdw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_packssdw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_packsswb (check.go:1818:check:)
//TODO __builtin_ia32_packsswb128 (check.go:1818:check:)
//TODO __builtin_ia32_packsswb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_packsswb256 (check.go:1818:check:)
//TODO __builtin_ia32_packsswb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_packsswb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_packusdw128 (check.go:1818:check:)
//TODO __builtin_ia32_packusdw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_packusdw256 (check.go:1818:check:)
//TODO __builtin_ia32_packusdw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_packusdw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_packuswb (check.go:1818:check:)
//TODO __builtin_ia32_packuswb128 (check.go:1818:check:)
//TODO __builtin_ia32_packuswb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_packuswb256 (check.go:1818:check:)
//TODO __builtin_ia32_packuswb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_packuswb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddb (check.go:1818:check:)
//TODO __builtin_ia32_paddb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddd (check.go:1818:check:)
//TODO __builtin_ia32_paddd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddq (check.go:1818:check:)
//TODO __builtin_ia32_paddq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsb (check.go:1818:check:)
//TODO __builtin_ia32_paddsb128 (check.go:1818:check:)
//TODO __builtin_ia32_paddsb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsb256 (check.go:1818:check:)
//TODO __builtin_ia32_paddsb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsw (check.go:1818:check:)
//TODO __builtin_ia32_paddsw128 (check.go:1818:check:)
//TODO __builtin_ia32_paddsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsw256 (check.go:1818:check:)
//TODO __builtin_ia32_paddsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusb (check.go:1818:check:)
//TODO __builtin_ia32_paddusb128 (check.go:1818:check:)
//TODO __builtin_ia32_paddusb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusb256 (check.go:1818:check:)
//TODO __builtin_ia32_paddusb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusw (check.go:1818:check:)
//TODO __builtin_ia32_paddusw128 (check.go:1818:check:)
//TODO __builtin_ia32_paddusw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusw256 (check.go:1818:check:)
//TODO __builtin_ia32_paddusw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddusw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddw (check.go:1818:check:)
//TODO __builtin_ia32_paddw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_paddw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pand (check.go:1818:check:)
//TODO __builtin_ia32_pandd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandn (check.go:1818:check:)
//TODO __builtin_ia32_pandn128 (check.go:1818:check:)
//TODO __builtin_ia32_pandnd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandnd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandnd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandnq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandnq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandnq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pandq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pause (check.go:1818:check:)
//TODO __builtin_ia32_pavgb (check.go:1818:check:)
//TODO __builtin_ia32_pavgb128 (check.go:1818:check:)
//TODO __builtin_ia32_pavgb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pavgb256 (check.go:1818:check:)
//TODO __builtin_ia32_pavgb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pavgb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pavgw (check.go:1818:check:)
//TODO __builtin_ia32_pavgw128 (check.go:1818:check:)
//TODO __builtin_ia32_pavgw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pavgw256 (check.go:1818:check:)
//TODO __builtin_ia32_pavgw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pavgw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pblendvb128 (check.go:1818:check:)
//TODO __builtin_ia32_pblendvb256 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb128 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb128_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb256 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb256_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb512_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd128 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd128_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd256 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd256_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd512 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastd512_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq128 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq128_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq256 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq256_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq512 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastq512_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw128 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw128_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw256 (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw256_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw512_gpr_mask (check.go:1818:check:)
//TODO __builtin_ia32_pbroadcastw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqb (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqd (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqw (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpeqw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtb (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtd (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtw (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pcmpgtw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pd256_pd (check.go:1818:check:)
//TODO __builtin_ia32_pd512_256pd (check.go:1818:check:)
//TODO __builtin_ia32_pd512_pd (check.go:1818:check:)
//TODO __builtin_ia32_pd_pd256 (check.go:1818:check:)
//TODO __builtin_ia32_pdep_di (check.go:1818:check:)
//TODO __builtin_ia32_pdep_si (check.go:1818:check:)
//TODO __builtin_ia32_permvardf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvardf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvardi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvardi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarsf256 (check.go:1818:check:)
//TODO __builtin_ia32_permvarsf256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarsf512_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarsi256 (check.go:1818:check:)
//TODO __builtin_ia32_permvarsi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_permvarsi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pext_di (check.go:1818:check:)
//TODO __builtin_ia32_pext_si (check.go:1818:check:)
//TODO __builtin_ia32_phaddd (check.go:1818:check:)
//TODO __builtin_ia32_phaddd128 (check.go:1818:check:)
//TODO __builtin_ia32_phaddd256 (check.go:1818:check:)
//TODO __builtin_ia32_phaddsw (check.go:1818:check:)
//TODO __builtin_ia32_phaddsw128 (check.go:1818:check:)
//TODO __builtin_ia32_phaddsw256 (check.go:1818:check:)
//TODO __builtin_ia32_phaddw (check.go:1818:check:)
//TODO __builtin_ia32_phaddw128 (check.go:1818:check:)
//TODO __builtin_ia32_phaddw256 (check.go:1818:check:)
//TODO __builtin_ia32_phminposuw128 (check.go:1818:check:)
//TODO __builtin_ia32_phsubd (check.go:1818:check:)
//TODO __builtin_ia32_phsubd128 (check.go:1818:check:)
//TODO __builtin_ia32_phsubd256 (check.go:1818:check:)
//TODO __builtin_ia32_phsubsw (check.go:1818:check:)
//TODO __builtin_ia32_phsubsw128 (check.go:1818:check:)
//TODO __builtin_ia32_phsubsw256 (check.go:1818:check:)
//TODO __builtin_ia32_phsubw (check.go:1818:check:)
//TODO __builtin_ia32_phsubw128 (check.go:1818:check:)
//TODO __builtin_ia32_phsubw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaddubsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaddwd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsb128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsb256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxub512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxud128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxud128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxud256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxud256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxud512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmaxuw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsb128 (check.go:1818:check:)
//TODO __builtin_ia32_pminsb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsb256 (check.go:1818:check:)
//TODO __builtin_ia32_pminsb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsd128 (check.go:1818:check:)
//TODO __builtin_ia32_pminsd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsd256 (check.go:1818:check:)
//TODO __builtin_ia32_pminsd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsw (check.go:1818:check:)
//TODO __builtin_ia32_pminsw128 (check.go:1818:check:)
//TODO __builtin_ia32_pminsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsw256 (check.go:1818:check:)
//TODO __builtin_ia32_pminsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminub (check.go:1818:check:)
//TODO __builtin_ia32_pminub128 (check.go:1818:check:)
//TODO __builtin_ia32_pminub128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminub256 (check.go:1818:check:)
//TODO __builtin_ia32_pminub256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminub512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminud128 (check.go:1818:check:)
//TODO __builtin_ia32_pminud128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminud256 (check.go:1818:check:)
//TODO __builtin_ia32_pminud256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminud512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuw128 (check.go:1818:check:)
//TODO __builtin_ia32_pminuw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuw256 (check.go:1818:check:)
//TODO __builtin_ia32_pminuw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pminuw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovdw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovmskb (check.go:1818:check:)
//TODO __builtin_ia32_pmovmskb128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovmskb256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqd512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovqw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsdw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqd512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsqw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovswb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxbw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxdq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxdq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxdq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxdq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxdq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovsxwq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusdw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqd512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovusqw512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovuswb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb128mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb256mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovwb512mem_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxbw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxdq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxdq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxdq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxdq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxdq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwd128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwd256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmovzxwq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuldq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmuldq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuldq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmuldq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuldq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhrsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhuw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw128 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw256 (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulhw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulld128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulld256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmulld512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullw (check.go:1818:check:)
//TODO __builtin_ia32_pmullw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmullw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq128 (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq256 (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pmuludq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_por (check.go:1818:check:)
//TODO __builtin_ia32_pord128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pord256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pord512_mask (check.go:1818:check:)
//TODO __builtin_ia32_porq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_porq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_porq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_prolvq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_prorvq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ps256_ps (check.go:1818:check:)
//TODO __builtin_ia32_ps512_256ps (check.go:1818:check:)
//TODO __builtin_ia32_ps512_ps (check.go:1818:check:)
//TODO __builtin_ia32_ps_ps256 (check.go:1818:check:)
//TODO __builtin_ia32_psadbw (check.go:1818:check:)
//TODO __builtin_ia32_psadbw128 (check.go:1818:check:)
//TODO __builtin_ia32_psadbw256 (check.go:1818:check:)
//TODO __builtin_ia32_psadbw512 (check.go:1818:check:)
//TODO __builtin_ia32_pshufb (check.go:1818:check:)
//TODO __builtin_ia32_pshufb128 (check.go:1818:check:)
//TODO __builtin_ia32_pshufb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pshufb256 (check.go:1818:check:)
//TODO __builtin_ia32_pshufb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pshufb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psignb (check.go:1818:check:)
//TODO __builtin_ia32_psignb128 (check.go:1818:check:)
//TODO __builtin_ia32_psignb256 (check.go:1818:check:)
//TODO __builtin_ia32_psignd (check.go:1818:check:)
//TODO __builtin_ia32_psignd128 (check.go:1818:check:)
//TODO __builtin_ia32_psignd256 (check.go:1818:check:)
//TODO __builtin_ia32_psignw (check.go:1818:check:)
//TODO __builtin_ia32_psignw128 (check.go:1818:check:)
//TODO __builtin_ia32_psignw256 (check.go:1818:check:)
//TODO __builtin_ia32_pslld (check.go:1818:check:)
//TODO __builtin_ia32_pslld128 (check.go:1818:check:)
//TODO __builtin_ia32_pslld128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pslld256 (check.go:1818:check:)
//TODO __builtin_ia32_pslld256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pslld512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pslldi (check.go:1818:check:)
//TODO __builtin_ia32_pslldi128 (check.go:1818:check:)
//TODO __builtin_ia32_pslldi256 (check.go:1818:check:)
//TODO __builtin_ia32_psllq (check.go:1818:check:)
//TODO __builtin_ia32_psllq128 (check.go:1818:check:)
//TODO __builtin_ia32_psllq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllq256 (check.go:1818:check:)
//TODO __builtin_ia32_psllq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllqi (check.go:1818:check:)
//TODO __builtin_ia32_psllqi128 (check.go:1818:check:)
//TODO __builtin_ia32_psllqi256 (check.go:1818:check:)
//TODO __builtin_ia32_psllv16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv2di (check.go:1818:check:)
//TODO __builtin_ia32_psllv2di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv4di (check.go:1818:check:)
//TODO __builtin_ia32_psllv4di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv4si (check.go:1818:check:)
//TODO __builtin_ia32_psllv4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllv8si (check.go:1818:check:)
//TODO __builtin_ia32_psllv8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllw (check.go:1818:check:)
//TODO __builtin_ia32_psllw128 (check.go:1818:check:)
//TODO __builtin_ia32_psllw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllw256 (check.go:1818:check:)
//TODO __builtin_ia32_psllw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psllwi (check.go:1818:check:)
//TODO __builtin_ia32_psllwi128 (check.go:1818:check:)
//TODO __builtin_ia32_psllwi256 (check.go:1818:check:)
//TODO __builtin_ia32_psrad (check.go:1818:check:)
//TODO __builtin_ia32_psrad128 (check.go:1818:check:)
//TODO __builtin_ia32_psrad128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrad256 (check.go:1818:check:)
//TODO __builtin_ia32_psrad256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrad512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psradi (check.go:1818:check:)
//TODO __builtin_ia32_psradi128 (check.go:1818:check:)
//TODO __builtin_ia32_psradi256 (check.go:1818:check:)
//TODO __builtin_ia32_psraq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psraq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psraq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav4si (check.go:1818:check:)
//TODO __builtin_ia32_psrav4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrav8si (check.go:1818:check:)
//TODO __builtin_ia32_psrav8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psravq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psravq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psraw (check.go:1818:check:)
//TODO __builtin_ia32_psraw128 (check.go:1818:check:)
//TODO __builtin_ia32_psraw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psraw256 (check.go:1818:check:)
//TODO __builtin_ia32_psraw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psraw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrawi (check.go:1818:check:)
//TODO __builtin_ia32_psrawi128 (check.go:1818:check:)
//TODO __builtin_ia32_psrawi256 (check.go:1818:check:)
//TODO __builtin_ia32_psrld (check.go:1818:check:)
//TODO __builtin_ia32_psrld128 (check.go:1818:check:)
//TODO __builtin_ia32_psrld128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrld256 (check.go:1818:check:)
//TODO __builtin_ia32_psrld256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrld512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrldi (check.go:1818:check:)
//TODO __builtin_ia32_psrldi128 (check.go:1818:check:)
//TODO __builtin_ia32_psrldi256 (check.go:1818:check:)
//TODO __builtin_ia32_psrlq (check.go:1818:check:)
//TODO __builtin_ia32_psrlq128 (check.go:1818:check:)
//TODO __builtin_ia32_psrlq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlq256 (check.go:1818:check:)
//TODO __builtin_ia32_psrlq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlqi (check.go:1818:check:)
//TODO __builtin_ia32_psrlqi128 (check.go:1818:check:)
//TODO __builtin_ia32_psrlqi256 (check.go:1818:check:)
//TODO __builtin_ia32_psrlv16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv2di (check.go:1818:check:)
//TODO __builtin_ia32_psrlv2di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv4di (check.go:1818:check:)
//TODO __builtin_ia32_psrlv4di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv4si (check.go:1818:check:)
//TODO __builtin_ia32_psrlv4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlv8si (check.go:1818:check:)
//TODO __builtin_ia32_psrlv8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlw (check.go:1818:check:)
//TODO __builtin_ia32_psrlw128 (check.go:1818:check:)
//TODO __builtin_ia32_psrlw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlw256 (check.go:1818:check:)
//TODO __builtin_ia32_psrlw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psrlwi (check.go:1818:check:)
//TODO __builtin_ia32_psrlwi128 (check.go:1818:check:)
//TODO __builtin_ia32_psrlwi256 (check.go:1818:check:)
//TODO __builtin_ia32_psubb (check.go:1818:check:)
//TODO __builtin_ia32_psubb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubd (check.go:1818:check:)
//TODO __builtin_ia32_psubd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubq (check.go:1818:check:)
//TODO __builtin_ia32_psubq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsb (check.go:1818:check:)
//TODO __builtin_ia32_psubsb128 (check.go:1818:check:)
//TODO __builtin_ia32_psubsb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsb256 (check.go:1818:check:)
//TODO __builtin_ia32_psubsb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsw (check.go:1818:check:)
//TODO __builtin_ia32_psubsw128 (check.go:1818:check:)
//TODO __builtin_ia32_psubsw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsw256 (check.go:1818:check:)
//TODO __builtin_ia32_psubsw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubsw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusb (check.go:1818:check:)
//TODO __builtin_ia32_psubusb128 (check.go:1818:check:)
//TODO __builtin_ia32_psubusb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusb256 (check.go:1818:check:)
//TODO __builtin_ia32_psubusb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusw (check.go:1818:check:)
//TODO __builtin_ia32_psubusw128 (check.go:1818:check:)
//TODO __builtin_ia32_psubusw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusw256 (check.go:1818:check:)
//TODO __builtin_ia32_psubusw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubusw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubw (check.go:1818:check:)
//TODO __builtin_ia32_psubw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_psubw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ptestc128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestc256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmb128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmb256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmb512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmd128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmd256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmd512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmq128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmq256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmq512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmw128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmw256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestmw512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmb128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmb256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmb512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmd128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmd256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmd512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmq128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmq256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmq512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmw128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmw256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnmw512 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnzc128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestnzc256 (check.go:1818:check:)
//TODO __builtin_ia32_ptestz128 (check.go:1818:check:)
//TODO __builtin_ia32_ptestz256 (check.go:1818:check:)
//TODO __builtin_ia32_ptwrite32 (check.go:1818:check:)
//TODO __builtin_ia32_ptwrite64 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw128 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw256 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhbw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq128 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq256 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhdq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhqdq128 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhqdq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhqdq256 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhqdq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhqdq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd128 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd256 (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckhwd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw128 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw256 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklbw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq128 (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq256 (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpckldq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklqdq128 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklqdq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklqdq256 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklqdq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklqdq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd128 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd256 (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_punpcklwd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxor (check.go:1818:check:)
//TODO __builtin_ia32_pxord128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxord256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxord512_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxorq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxorq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_pxorq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14sd (check.go:1818:check:)
//TODO __builtin_ia32_rcp14sd_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcp14ss (check.go:1818:check:)
//TODO __builtin_ia32_rcp14ss_mask (check.go:1818:check:)
//TODO __builtin_ia32_rcpps (check.go:1818:check:)
//TODO __builtin_ia32_rcpps256 (check.go:1818:check:)
//TODO __builtin_ia32_rcpss (check.go:1818:check:)
//TODO __builtin_ia32_rdfsbase32 (check.go:1818:check:)
//TODO __builtin_ia32_rdfsbase64 (check.go:1818:check:)
//TODO __builtin_ia32_rdgsbase32 (check.go:1818:check:)
//TODO __builtin_ia32_rdgsbase64 (check.go:1818:check:)
//TODO __builtin_ia32_rdpid (check.go:1818:check:)
//TODO __builtin_ia32_rdpkru (check.go:1818:check:)
//TODO __builtin_ia32_rdrand16_step (check.go:1818:check:)
//TODO __builtin_ia32_rdrand32_step (check.go:1818:check:)
//TODO __builtin_ia32_rdrand64_step (check.go:1818:check:)
//TODO __builtin_ia32_rdseed_di_step (check.go:1818:check:)
//TODO __builtin_ia32_rdseed_hi_step (check.go:1818:check:)
//TODO __builtin_ia32_rdseed_si_step (check.go:1818:check:)
//TODO __builtin_ia32_rdsspq (check.go:1818:check:)
//TODO __builtin_ia32_rndscalepd_mask (check.go:1818:check:)
//TODO __builtin_ia32_rndscaleps_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14pd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14pd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14pd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14ps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14sd (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14sd_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14ss (check.go:1818:check:)
//TODO __builtin_ia32_rsqrt14ss_mask (check.go:1818:check:)
//TODO __builtin_ia32_rsqrtps (check.go:1818:check:)
//TODO __builtin_ia32_rsqrtps256 (check.go:1818:check:)
//TODO __builtin_ia32_rsqrtss (check.go:1818:check:)
//TODO __builtin_ia32_rstorssp (check.go:1818:check:)
//TODO __builtin_ia32_saveprevssp (check.go:1818:check:)
//TODO __builtin_ia32_sbb_u32 (check.go:1818:check:)
//TODO __builtin_ia32_sbb_u64 (check.go:1818:check:)
//TODO __builtin_ia32_scalefpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_scalefsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_scalefss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_setssbsy (check.go:1818:check:)
//TODO __builtin_ia32_sfence (check.go:1818:check:)
//TODO __builtin_ia32_sha1msg1 (check.go:1818:check:)
//TODO __builtin_ia32_sha1msg2 (check.go:1818:check:)
//TODO __builtin_ia32_sha1nexte (check.go:1818:check:)
//TODO __builtin_ia32_sha256msg1 (check.go:1818:check:)
//TODO __builtin_ia32_sha256msg2 (check.go:1818:check:)
//TODO __builtin_ia32_sha256rnds2 (check.go:1818:check:)
//TODO __builtin_ia32_shufpd (check.go:1818:check:)
//TODO __builtin_ia32_shufps (check.go:1818:check:)
//TODO __builtin_ia32_si256_si (check.go:1818:check:)
//TODO __builtin_ia32_si512_256si (check.go:1818:check:)
//TODO __builtin_ia32_si512_si (check.go:1818:check:)
//TODO __builtin_ia32_si_si256 (check.go:1818:check:)
//TODO __builtin_ia32_sqrtpd (check.go:1818:check:)
//TODO __builtin_ia32_sqrtpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtpd256 (check.go:1818:check:)
//TODO __builtin_ia32_sqrtpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtps (check.go:1818:check:)
//TODO __builtin_ia32_sqrtps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtps256 (check.go:1818:check:)
//TODO __builtin_ia32_sqrtps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_sqrtsd (check.go:1818:check:)
//TODO __builtin_ia32_sqrtss (check.go:1818:check:)
//TODO __builtin_ia32_stmxcsr (check.go:1818:check:)
//TODO __builtin_ia32_storeapd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeapd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeapd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeaps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeaps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeaps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqudi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqudi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqudi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedquqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqusi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqusi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storedqusi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storehps (check.go:1818:check:)
//TODO __builtin_ia32_storelps (check.go:1818:check:)
//TODO __builtin_ia32_storesd_mask (check.go:1818:check:)
//TODO __builtin_ia32_storess_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeupd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeupd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeupd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeups128_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeups256_mask (check.go:1818:check:)
//TODO __builtin_ia32_storeups512_mask (check.go:1818:check:)
//TODO __builtin_ia32_subpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_subpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_subpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_subps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_subps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_subps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_subsd (check.go:1818:check:)
//TODO __builtin_ia32_subsd_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_subss (check.go:1818:check:)
//TODO __builtin_ia32_subss_mask_round (check.go:1818:check:)
//TODO __builtin_ia32_tpause (check.go:1818:check:)
//TODO __builtin_ia32_tzcnt_u16 (check.go:1818:check:)
//TODO __builtin_ia32_tzcnt_u32 (check.go:1818:check:)
//TODO __builtin_ia32_tzcnt_u64 (check.go:1818:check:)
//TODO __builtin_ia32_ucmpb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpw128_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpw256_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucmpw512_mask (check.go:1818:check:)
//TODO __builtin_ia32_ucomieq (check.go:1818:check:)
//TODO __builtin_ia32_ucomige (check.go:1818:check:)
//TODO __builtin_ia32_ucomigt (check.go:1818:check:)
//TODO __builtin_ia32_ucomile (check.go:1818:check:)
//TODO __builtin_ia32_ucomilt (check.go:1818:check:)
//TODO __builtin_ia32_ucomineq (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdeq (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdge (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdgt (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdle (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdlt (check.go:1818:check:)
//TODO __builtin_ia32_ucomisdneq (check.go:1818:check:)
//TODO __builtin_ia32_umonitor (check.go:1818:check:)
//TODO __builtin_ia32_umwait (check.go:1818:check:)
//TODO __builtin_ia32_unpckhpd (check.go:1818:check:)
//TODO __builtin_ia32_unpckhpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpckhpd256 (check.go:1818:check:)
//TODO __builtin_ia32_unpckhpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpckhpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpckhps (check.go:1818:check:)
//TODO __builtin_ia32_unpckhps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpckhps256 (check.go:1818:check:)
//TODO __builtin_ia32_unpckhps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpckhps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklpd (check.go:1818:check:)
//TODO __builtin_ia32_unpcklpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklpd256 (check.go:1818:check:)
//TODO __builtin_ia32_unpcklpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklps (check.go:1818:check:)
//TODO __builtin_ia32_unpcklps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklps256 (check.go:1818:check:)
//TODO __builtin_ia32_unpcklps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_unpcklps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vaesdec_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesdec_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesdeclast_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesdeclast_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesenc_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesenc_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesenclast_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vaesenclast_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastf128_pd256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastf128_ps256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastsd256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastsd_pd256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastsi256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastss (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastss256 (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastss_ps (check.go:1818:check:)
//TODO __builtin_ia32_vbroadcastss_ps256 (check.go:1818:check:)
//TODO __builtin_ia32_vcvtph2ps (check.go:1818:check:)
//TODO __builtin_ia32_vcvtph2ps256 (check.go:1818:check:)
//TODO __builtin_ia32_vcvtph2ps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vcvtph2ps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vcvtph2ps_mask (check.go:1818:check:)
//TODO __builtin_ia32_vcvtsd2usi32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvtsd2usi64 (check.go:1818:check:)
//TODO __builtin_ia32_vcvtss2usi32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvtss2usi64 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttsd2si32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttsd2si64 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttsd2usi32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttsd2usi64 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttss2si32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttss2si64 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttss2usi32 (check.go:1818:check:)
//TODO __builtin_ia32_vcvttss2usi64 (check.go:1818:check:)
//TODO __builtin_ia32_vec_ext_v2si (check.go:1818:check:)
//TODO __builtin_ia32_vec_ext_v4sf (check.go:1818:check:)
//TODO __builtin_ia32_vec_ext_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vec_init_v2si (check.go:1818:check:)
//TODO __builtin_ia32_vec_init_v4hi (check.go:1818:check:)
//TODO __builtin_ia32_vec_init_v8qi (check.go:1818:check:)
//TODO __builtin_ia32_vextractf128_pd256 (check.go:1818:check:)
//TODO __builtin_ia32_vextractf128_ps256 (check.go:1818:check:)
//TODO __builtin_ia32_vextractf128_si256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsd3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsd3_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsd3_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsd3_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddss3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddss3_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddss3_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddss3_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmaddsubps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubaddps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps256 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubsd3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubsd3_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubss3 (check.go:1818:check:)
//TODO __builtin_ia32_vfmsubss3_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps256 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddsd3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmaddss3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps128_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps256 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps256_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps512_mask3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubsd3 (check.go:1818:check:)
//TODO __builtin_ia32_vfnmsubss3 (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v16qi (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v16qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v32qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vgf2p8mulb_v64qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vinsertf128_pd256 (check.go:1818:check:)
//TODO __builtin_ia32_vinsertf128_ps256 (check.go:1818:check:)
//TODO __builtin_ia32_vinsertf128_si256 (check.go:1818:check:)
//TODO __builtin_ia32_vp4dpwssd (check.go:1818:check:)
//TODO __builtin_ia32_vp4dpwssd_mask (check.go:1818:check:)
//TODO __builtin_ia32_vp4dpwssds (check.go:1818:check:)
//TODO __builtin_ia32_vp4dpwssds_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictdi_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictdi_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictdi_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictsi_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictsi_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpconflictsi_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusd_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpbusds_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssd_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpdpwssds_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2vard128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2vard256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2vard512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermi2varqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarpd (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarpd_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarps (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarps256 (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermilvarps_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2vard512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varhi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varpd512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varps512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varq512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpermt2varqi512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntd_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntd_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntd_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntq_128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntq_256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vplzcntq_512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52huq512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq128_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq256_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmadd52luq512_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpmultishiftqb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmultishiftqb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpmultishiftqb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v16qi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v16qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v32qi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v32qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v64qi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountb_v64qi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountd_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v2di (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v2di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v4di (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v4di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v8di (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountq_v8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v16hi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v32hi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v8hi (check.go:1818:check:)
//TODO __builtin_ia32_vpopcountw_v8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v2di (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v2di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v2di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v32hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v32hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4di (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8di (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshldv_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16si (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v16si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v2di (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v2di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v2di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v32hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v32hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v32hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4di (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4si (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v4si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8di (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8di_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8di_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8hi (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8hi_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8hi_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8si (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8si_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshrdv_v8si_maskz (check.go:1818:check:)
//TODO __builtin_ia32_vpshufbitqmb128_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshufbitqmb256_mask (check.go:1818:check:)
//TODO __builtin_ia32_vpshufbitqmb512_mask (check.go:1818:check:)
//TODO __builtin_ia32_vtestcpd (check.go:1818:check:)
//TODO __builtin_ia32_vtestcpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vtestcps (check.go:1818:check:)
//TODO __builtin_ia32_vtestcps256 (check.go:1818:check:)
//TODO __builtin_ia32_vtestnzcpd (check.go:1818:check:)
//TODO __builtin_ia32_vtestnzcpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vtestnzcps (check.go:1818:check:)
//TODO __builtin_ia32_vtestnzcps256 (check.go:1818:check:)
//TODO __builtin_ia32_vtestzpd (check.go:1818:check:)
//TODO __builtin_ia32_vtestzpd256 (check.go:1818:check:)
//TODO __builtin_ia32_vtestzps (check.go:1818:check:)
//TODO __builtin_ia32_vtestzps256 (check.go:1818:check:)
//TODO __builtin_ia32_vzeroall (check.go:1818:check:)
//TODO __builtin_ia32_vzeroupper (check.go:1818:check:)
//TODO __builtin_ia32_wbinvd (check.go:1818:check:)
//TODO __builtin_ia32_wbnoinvd (check.go:1818:check:)
//TODO __builtin_ia32_wrfsbase32 (check.go:1818:check:)
//TODO __builtin_ia32_wrfsbase64 (check.go:1818:check:)
//TODO __builtin_ia32_wrgsbase32 (check.go:1818:check:)
//TODO __builtin_ia32_wrgsbase64 (check.go:1818:check:)
//TODO __builtin_ia32_wrpkru (check.go:1818:check:)
//TODO __builtin_ia32_wrssd (check.go:1818:check:)
//TODO __builtin_ia32_wrssq (check.go:1818:check:)
//TODO __builtin_ia32_wrussd (check.go:1818:check:)
//TODO __builtin_ia32_wrussq (check.go:1818:check:)
//TODO __builtin_ia32_xbegin (check.go:1818:check:)
//TODO __builtin_ia32_xend (check.go:1818:check:)
//TODO __builtin_ia32_xgetbv (check.go:1818:check:)
//TODO __builtin_ia32_xorpd (check.go:1818:check:)
//TODO __builtin_ia32_xorpd128_mask (check.go:1818:check:)
//TODO __builtin_ia32_xorpd256 (check.go:1818:check:)
//TODO __builtin_ia32_xorpd256_mask (check.go:1818:check:)
//TODO __builtin_ia32_xorpd512_mask (check.go:1818:check:)
//TODO __builtin_ia32_xorps (check.go:1818:check:)
//TODO __builtin_ia32_xorps128_mask (check.go:1818:check:)
//TODO __builtin_ia32_xorps256 (check.go:1818:check:)
//TODO __builtin_ia32_xorps256_mask (check.go:1818:check:)
//TODO __builtin_ia32_xorps512_mask (check.go:1818:check:)
//TODO __builtin_ia32_xrstor (check.go:1818:check:)
//TODO __builtin_ia32_xrstor64 (check.go:1818:check:)
//TODO __builtin_ia32_xrstors (check.go:1818:check:)
//TODO __builtin_ia32_xrstors64 (check.go:1818:check:)
//TODO __builtin_ia32_xsave (check.go:1818:check:)
//TODO __builtin_ia32_xsave64 (check.go:1818:check:)
//TODO __builtin_ia32_xsavec (check.go:1818:check:)
//TODO __builtin_ia32_xsavec64 (check.go:1818:check:)
//TODO __builtin_ia32_xsaveopt (check.go:1818:check:)
//TODO __builtin_ia32_xsaveopt64 (check.go:1818:check:)
//TODO __builtin_ia32_xsaves (check.go:1818:check:)
//TODO __builtin_ia32_xsaves64 (check.go:1818:check:)
//TODO __builtin_ia32_xsetbv (check.go:1818:check:)
//TODO __builtin_ia32_xtest (check.go:1818:check:)
//TODO __builtin_infl (check.go:1818:check:)
//TODO __builtin_isgreater (check.go:1818:check:)
//TODO __builtin_isgreaterequal (check.go:1818:check:)
//TODO __builtin_isinf (check.go:1818:check:)
//TODO __builtin_isinff (check.go:1818:check:)
//TODO __builtin_isinfl (check.go:1818:check:)
//TODO __builtin_isless (check.go:1818:check:)
//TODO __builtin_islessequal (check.go:1818:check:)
//TODO __builtin_islessgreater (check.go:1818:check:)
//TODO __builtin_longjmp (check.go:1818:check:)
//TODO __builtin_memchr (check.go:1818:check:)
//TODO __builtin_memmove (check.go:1818:check:)
//TODO __builtin_mempcpy (check.go:1818:check:)
//TODO __builtin_modfl (check.go:1818:check:)
//TODO __builtin_popcountll (check.go:1818:check:)
//TODO __builtin_printf_unlocked (check.go:1818:check:)
//TODO __builtin_putchar (check.go:1818:check:)
//TODO __builtin_puts (check.go:1818:check:)
//TODO __builtin_return_address (check.go:1818:check:)
//TODO __builtin_setjmp (check.go:1818:check:)
//TODO __builtin_shuffle (check.go:1818:check:)
//TODO __builtin_signbit (check.go:1818:check:)
//TODO __builtin_stack_restore (check.go:1818:check:)
//TODO __builtin_stpcpy (check.go:1818:check:)
//TODO __builtin_stpncpy (check.go:1818:check:)
//TODO __builtin_strncasecmp (check.go:1818:check:)
//TODO __builtin_strncat (check.go:1818:check:)
//TODO __builtin_strncmp (check.go:1818:check:)
//TODO __builtin_strncpy (check.go:1818:check:)
//TODO __builtin_strndup (check.go:1818:check:)
//TODO __builtin_va_arg_pack (check.go:1818:check:)

char *__builtin_strchr(const char *s, int c);
char *__builtin_strcpy(char *dest, const char *src);
double __builtin_copysign ( double x, double y );
double __builtin_copysignl (long double x, long double y );
double __builtin_huge_val (void);
double __builtin_inf (void);
double __builtin_nan (const char *str);
float __builtin_copysignf ( float x, float y );
float __builtin_huge_valf (void);
float __builtin_inff (void);
float __builtin_nanf (const char *str);
int __builtin___snprintf_chk (char *s, size_t maxlen, int flag, size_t os, const char *fmt, ...);
int __builtin_abs(int j);
int __builtin_add_overflow();
int __builtin_clz (unsigned);
int __builtin_isunordered(double x, double y);
int __builtin_memcmp(const void *s1, const void *s2, size_t n);
int __builtin_mul_overflow(); //TODO
int __builtin_popcount (unsigned int x);
int __builtin_strcmp(const char *s1, const char *s2);
int __builtin_sub_overflow(); //TODO 
long __builtin_expect (long exp, long c);
size_t __builtin_object_size (void * ptr, int type);
size_t __builtin_strlen(const char *s);
void *__builtin___memcpy_chk (void *dest, const void *src, size_t n, size_t os);
void *__builtin_malloc(size_t size);
void *__builtin_memcpy(void *dest, const void *src, size_t n);
void *__builtin_memset(void *s, int c, size_t n);
void __builtin_abort(void);
void __builtin_bzero(void *s, size_t n);
void __builtin_free(void *ptr);
void __builtin_prefetch (const void *addr, ...);
void __builtin_trap (void);
void __builtin_unreachable (void);

`

	oTrace = flag.Bool("trc", false, "Print tested paths.")
)

func init() {
	flag.BoolVar(&traceFails, "trcfails", false, "")
	isTesting = true
	var err error
	if defaultCfg0, err = NewConfig(runtime.GOOS, runtime.GOARCH); err != nil {
		panic(errorf("NewConfig: %v", err))
	}

	var chars int
	if err := walk("/", func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		f, err := cfs.Open(pth)
		if err != nil {
			return errorf("%v: %v", pth, err)
		}

		b, err := io.ReadAll(f)
		if err != nil {
			return errorf("%v: %v", pth, err)
		}

		switch filepath.Ext(pth) {
		case ".c", ".h":
			if len(b) != 0 && b[len(b)-1] != '\n' {
				b = append(b, '\n')
			}
		}
		chars += len(b)
		corpus[pth] = b
		corpusIndex = append(corpusIndex, pth)
		return nil
	}); err != nil {
		panic(err)
	}
}

type corpusFS struct {
	*httpfs.FileSystem
}

func (c *corpusFS) Open(name string) (fs.File, error) {
	name = filepath.ToSlash(name)
	if !strings.HasPrefix(name, "/") {
		name = "/" + name
	}
	f, err := c.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}

	return fs.File(f), nil
}

func walk(dir string, f func(pth string, fi os.FileInfo) error) error {
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	root, err := cfs.Open(dir)
	if err != nil {
		return err
	}

	fi, err := root.Stat()
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("%s: not a directory", fi.Name())
	}

	fis, err := root.Readdir(-1)
	if err != nil {
		return err
	}

	for _, v := range fis {
		switch {
		case v.IsDir():
			if err = walk(v.Name(), f); err != nil {
				return err
			}
		default:
			if err = f(v.Name(), v); err != nil {
				return err
			}
		}
	}
	return nil
}

// Produce the AST used in examples documentation.
func exampleAST(rule int, src string) (r interface{}) {
	defer func() {
		if err := recover(); err != nil {
			r = fmt.Sprintf("%v (%v:)", err, origin(5))
			trc("%v\n%s", r, debug.Stack())
		}
	}()

	src = strings.Replace(src, "\\n", "\n", -1)
	cfg := &Config{}
	ast, _ := Parse(cfg, []Source{{Name: "example.c", Value: src}})
	if ast == nil {
		return "FAIL"
	}

	pc, _, _, _ := runtime.Caller(1)
	typ := runtime.FuncForPC(pc - 1).Name()
	i := strings.LastIndexByte(typ, '.')
	typ = typ[i+1+len("Example"):]
	i = strings.LastIndexByte(typ, '_')
	typ = typ[:i]
	var node Node
	depth := mathutil.MaxInt
	findNode(typ, ast.TranslationUnit, 0, &node, &depth)
	return node
}

func TestMain(m *testing.M) {
	oRE := flag.String("re", "", "")
	flag.Parse()
	if *oRE != "" {
		re = regexp.MustCompile(*oRE)
	}
	os.Exit(m.Run())
}

func TestScannerSource(t *testing.T) {
	const fn = "all_test.go"
	exp, err := os.ReadFile(fn)
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.Open(fn)
	if err != nil {
		t.Fatal(err)
	}

	testScannerSource(t, fn, f, exp, false)
	testScannerSource(t, fn, exp, exp, false)
	testScannerSource(t, fn, string(exp), exp, false)
	testScannerSource(t, fn, bytes.NewReader(exp), exp, false)
	testScannerSource(t, fn, nil, exp, false)
	testScannerSource(t, fn, 42, nil, true)
}

func testScannerSource(t *testing.T, name string, value interface{}, exp []byte, mustFail bool) {
	ss, err := newScannerSource(Source{name, value, nil})
	if err != nil != mustFail {
		t.Fatalf("(%q, %T): %v", name, value, err)
	}

	if err != nil {
		return
	}

	if !bytes.Equal(ss.buf, exp) {
		t.Fatal("buf does not match")
	}
}

func TestToken(t *testing.T) {
	s, err := newScannerSource(Source{"test", `abc
def
 ghi
`, nil})
	// abc\ndef\n ghi\n
	//             1
	// 0123 4567 89012
	if err != nil {
		t.Fatal(err)
	}

	s.file.AddLine(4)
	s.file.AddLine(8)
	s.file.AddLine(13)
	for itest, test := range []struct {
		Token
		line int
		col  int
		ch   rune
		sep  string
		src  string
	}{
		{newToken(s, 0, 0, 0, 3), 1, 1, 0, "", "abc"},
		{newToken(s, 1, 3, 4, 3), 2, 1, 1, "\n", "def"},
		{newToken(s, 2, 7, 9, 3), 3, 2, 2, "\n ", "ghi"},
		{newToken(s, eof, 13, 13, 0), 3, 6, eof, "", ""},
	} {
		tok := test.Token
		if g, e := tok.Position().Line, test.line; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := tok.Position().Column, test.col; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := tok.Ch, test.ch; g != e {
			t.Fatal(itest, g, e)
		}
		if g, e := string(tok.Sep()), test.sep; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok.Src()), test.src; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}

		tok2 := tok
		tok2.Set([]byte("xyz0123"), []byte("456789"))
		if g, e := string(tok.Sep()), test.sep; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok.Src()), test.src; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok2.Sep()), "xyz0123"; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
		if g, e := string(tok2.Src()), "456789"; g != e {
			t.Fatalf("%v %q %q", itest, g, e)
		}
	}
}

type parallel struct {
	limit chan struct{}
	sync.Mutex
	wg sync.WaitGroup
}

func newParallel() *parallel {
	return &parallel{
		limit: make(chan struct{}, runtime.GOMAXPROCS(0)),
	}
}

func (p *parallel) exec(run func()) {
	p.limit <- struct{}{}
	p.wg.Add(1)

	go func() {
		defer func() {
			p.wg.Done()
			<-p.limit
		}()

		run()
	}()
}

var tokSink []Token

func TestScanner(t *testing.T) {
	defer func() { tokSink = nil }()

	var files, tokens, chars int64
	var m0, m runtime.MemStats
	p := newParallel()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		path := path
		switch filepath.Ext(path) {
		case ".c", ".h":
			files++
			p.exec(func() {
				var err error
				var chars0, tokens0 int64
				var toks []Token

				defer func() {
					p.Lock()
					chars += chars0
					tokens += tokens0
					tokSink = append(tokSink, toks...)
					if err != nil {
						t.Error(err)
					}
					p.Unlock()
				}()

				buf := corpus[path]
				chars0 += int64(len(buf))
				var s *scanner
				if s, err = newScanner(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					s.close()
					err = fmt.Errorf(msg, args...)
				}); err != nil {
					err = fmt.Errorf("%v: %v", path, err)
					return
				}

				for {
					tok := s.cppScan()
					if tok.Ch == eof {
						return
					}

					toks = append(toks, tok)
					tokens0++
				}
			})
		}
	}
	p.wg.Wait()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	t.Logf("files %v; tokens %v; bytes %v; heap %v; alloc %v", h(files), h(tokens), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func h(v interface{}) string {
	switch x := v.(type) {
	case int64:
		return humanize.Comma(x)
	case uint64:
		if x <= math.MaxInt64 {
			return humanize.Comma(int64(x))
		}
	}
	return fmt.Sprint(v)
}

func BenchmarkScanner(b *testing.B) {
	debug.FreeOSMemory()
	for i := 0; i < b.N; i++ {
		var chars int64
		for _, path := range corpusIndex {
			switch filepath.Ext(path) {
			case ".c", ".h":
				buf := corpus[path]
				chars += int64(len(buf))
				var s *scanner
				var err error
				if s, err = newScanner(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					s.close()
					b.Fatalf(msg, args...)
				}); err != nil {
					b.Fatal(path, err)
				}
				for {
					tok := s.cppScan()
					if tok.Ch == eof {
						break
					}
				}
			}
		}
		b.SetBytes(chars)
	}
}

var (
	cppParseBlacklist = map[string]struct{}{
		"/github.com/vnmakarov/mir/c-tests/new/endif.c": {}, // 1:1: unexpected #endif
	}
	astSink []group
)

func TestCPPParse0(t *testing.T) {
	defer func() { astSink = nil }()

	var files, lines, chars int64
	var m0, m runtime.MemStats
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		if _, ok := cppParseBlacklist[path]; ok {
			continue
		}

		switch filepath.Ext(path) {
		case ".c", ".h":
			buf := corpus[path]
			chars += int64(len(buf))
			var p *cppParser
			var err error
			if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
				p.close()
				t.Fatalf(msg, args...)
			}); err != nil {
				t.Fatal(path, err)
			}

			files++
			ast := p.preprocessingFile()
			if len(ast) == 0 {
				t.Fatalf("%v: empty AST", path)
			}

			eol := ast[len(ast)-1]
			x, ok := eol.(eofLine)
			if !ok {
				t.Fatalf("%v: AST not terminated: %T", p.pos(), eol)
			}

			eof := Token(x)
			lines += int64(eof.Position().Line)
			astSink = append(astSink, ast)
		}
	}
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	astSink = nil
	t.Logf("files %v; lines %v bytes %v; heap %v; alloc %v", h(files), h(lines), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func TestCPPParse(t *testing.T) {
	defer func() { astSink = nil }()

	var files, lines, chars int64
	var m0, m runtime.MemStats
	p := newParallel()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m0)
	for _, path := range corpusIndex {
		path := path
		if _, ok := cppParseBlacklist[path]; ok {
			continue
		}

		switch filepath.Ext(path) {
		case ".c", ".h":
			files++
			p.exec(func() {
				buf := corpus[path]
				var err error
				var ast group
				var eof Token

				defer func() {
					p.Lock()
					chars += int64(len(buf))
					lines += int64(eof.Position().Line)
					astSink = append(astSink, ast)
					if err != nil {
						t.Error(err)
					}
					p.Unlock()
				}()

				var p *cppParser
				if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					p.close()
					err = fmt.Errorf(msg, args...)
				}); err != nil {
					t.Fatal(path, err)
				}

				if ast = p.preprocessingFile(); len(ast) == 0 {
					t.Fatalf("%v: empty AST", path)
				}

				eol := ast[len(ast)-1]
				x, ok := eol.(eofLine)
				if !ok {
					err = fmt.Errorf("%v: AST not terminated: %T", p.pos(), eol)
					return
				}

				eof = Token(x)
			})
		}
	}
	p.wg.Wait()
	debug.FreeOSMemory()
	runtime.ReadMemStats(&m)
	astSink = nil
	t.Logf("files %v; lines %v bytes %v; heap %v; alloc %v", h(files), h(lines), h(chars), h(m.HeapAlloc-m0.HeapAlloc), h(m.TotalAlloc-m0.TotalAlloc))
}

func BenchmarkCPPParse(b *testing.B) {
	debug.FreeOSMemory()
	for i := 0; i < b.N; i++ {
		var chars int64
		for _, path := range corpusIndex {
			if _, ok := cppParseBlacklist[path]; ok {
				continue
			}

			switch filepath.Ext(path) {
			case ".c", ".h":
				buf := corpus[path]
				chars += int64(len(buf))
				var p *cppParser
				var err error
				if p, err = newCppParser(Source{path, buf, nil}, func(msg string, args ...interface{}) {
					p.close()
					b.Fatalf(msg, args...)
				}); err != nil {
					b.Fatal(path, err)
				}

				ast := p.preprocessingFile()
				if len(ast) == 0 {
					b.Fatalf("%v: empty AST", path)
				}

				eol := ast[len(ast)-1]
				if _, ok := eol.(eofLine); !ok {
					b.Fatalf("%v: AST not terminated: %T", p.pos(), eol)
				}
			}
		}
		b.SetBytes(chars)
	}
}

func defaultCfg() *Config {
	c := *defaultCfg0
	return &c
}

func TestCPPExpand(t *testing.T) {
	testCPPExpand(t, "testdata/cpp-expand/", nil, true)
}

func testCPPExpand(t *testing.T, dir string, blacklist map[string]struct{}, fakeIncludes bool) {
	var fails []string
	var files, ok, skip int
	var c *cpp
	cfg := defaultCfg()
	cfg.fakeIncludes = fakeIncludes
	cfg.PragmaHandler = func(s []Token) error {
		pragmaTestTok := Token{s: s[0].s, Ch: rune(IDENTIFIER)}
		pragmaTestTok.Set(nil, []byte("__pragma"))
		a := textLine{pragmaTestTok}
		for i, v := range s {
			if i == 0 {
				v.Set(sp, v.Src())
			}
			a = append(a, v)
		}
		nlTok := Token{s: s[0].s, Ch: '\n'}
		nlTok.Set(nil, nl)
		c.push(append(a, nlTok))
		return nil
	}
	err := filepath.Walk(filepath.FromSlash(dir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || (!strings.HasSuffix(path, ".c") && !strings.HasSuffix(path, ".h")) {
			return nil
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(path) {
				skip++
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(path)]; ok {
				skip++
				return nil
			}
		}

		if *oTrace {
			fmt.Fprintln(os.Stderr, path)
		}
		var b strings.Builder
		if c, err = newCPP(cfg, []Source{{path, nil, nil}}, nil); err != nil {
			t.Fatalf("%v: %v", path, err)
		}

		if err := preprocess(c, &b); err != nil {
			fails = append(fails, path)
			t.Fatalf("%v: %v", path, err)
		}

		if strings.Contains(filepath.ToSlash(path), "/mustfail/") {
			if err != nil {
				return nil
			}

			fails = append(fails, path)
			return fmt.Errorf("%v: unexpected success", path)
		}

		if err != nil {
			fails = append(fails, path)
			return err
		}

		expFn := path + ".expect"
		exp, err := os.ReadFile(expFn)
		if err != nil {
			fails = append(fails, path)
			t.Error(err)
		}

		g := strings.ReplaceAll(b.String(), "\r", "")
		g = strings.TrimSpace(g)
		e := strings.ReplaceAll(string(exp), "\r", "")
		e = strings.TrimSpace(e)
		if g != e {
			fails = append(fails, path)
			diff := difflib.UnifiedDiff{
				A:        difflib.SplitLines(e),
				B:        difflib.SplitLines(g),
				FromFile: expFn,
				ToFile:   path,
				Context:  0,
			}
			s, err := difflib.GetUnifiedDiffString(diff)
			if err != nil {
				t.Fatalf("%v: %v", path, err)
			}

			t.Errorf("%v\ngot\n%s\nexp\n%s\ngot\n%s\nexp\n%s", s, g, e, hex.Dump([]byte(g)), hex.Dump([]byte(e)))
			return nil
		}
		ok++
		return nil
	})
	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	if err != nil {
		t.Fatal(err)
	}
}

func TestPreprocess(t *testing.T) {
	testCPPExpand(t, "testdata/preprocess/", nil, true)
}

func TestTCCExpand(t *testing.T) {
	testCPPExpand(t, "testdata/tcc-0.9.27/tests/pp/", map[string]struct{}{
		"11.c": {}, // https://gcc.gnu.org/onlinedocs/gcc/Variadic-Macros.html#Variadic-Macros
		"16.c": {}, // We don't produce warnings on macro redefinition.
	}, true)
}

func TestInclude(t *testing.T) {
	testCPPExpand(t, "testdata/include/", nil, false)
}

func TestTranslationPhase4(t *testing.T) {
	cfg := defaultCfg()
	cfg.SysIncludePaths = append(cfg.SysIncludePaths, "Include") // benchmarksgame
	cfg.FS = cFS
	blacklistCompCert := map[string]struct{}{}
	blacklistGCC := map[string]struct{}{
		// assertions are deprecated, not supported.
		"950919-1.c": {},
	}
	blacklictTCC := map[string]struct{}{
		// https://gcc.gnu.org/onlinedocs/gcc/Variadic-Macros.html#Variadic-Macros, not supported.
		"11.c": {},
	}
	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "linux/s390x":
		blacklistCompCert["aes.c"] = struct{}{} // Unsupported endianness.
	}
	var files, ok, skip, fails int32
	for _, v := range []struct {
		cfg       *Config
		dir       string
		blacklist map[string]struct{}
	}{
		{cfg, "CompCert-3.6/test/c", blacklistCompCert},
		{cfg, "ccgo", nil},
		{cfg, "gcc-9.1.0/gcc/testsuite/gcc.c-torture", blacklistGCC},
		{cfg, "github.com/AbsInt/CompCert/test/c", blacklistCompCert},
		{cfg, "github.com/cxgo", nil},
		{cfg, "github.com/gcc-mirror/gcc/gcc/testsuite", blacklistGCC},
		{cfg, "github.com/vnmakarov", nil},
		{cfg, "sqlite-amalgamation-3370200", nil},
		{cfg, "tcc-0.9.27/tests", blacklictTCC},
		{cfg, "benchmarksgame-team.pages.debian.net", nil},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testTranslationPhase4(t, v.cfg, "/"+v.dir, v.blacklist)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testTranslationPhase4(t *testing.T, cfg *Config, dir string, blacklist map[string]struct{}) (files, ok, skip, nfails int32) {
	tmp := t.TempDir()
	var fails []string
	p := newParallel()
	err := walk(dir, func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		switch {
		case re != nil:
			if !re.MatchString(pth) {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		}

		files++
		apth := pth
		p.exec(func() {
			if *oTrace {
				fmt.Fprintln(os.Stderr, apth)
			}
			var err error

			defer func() {
				p.Lock()

				defer p.Unlock()

				if err != nil {
					fails = append(fails, apth)
					t.Errorf("%v: %v", apth, err)
				}
			}()

			if err = Preprocess(
				cfg,
				[]Source{
					{Name: "<predefined>", Value: cfg.Predefined},
					{Name: "<builtin>", Value: builtin},
					{Name: apth, FS: cFS},
				},
				io.Discard,
			); err == nil {
				atomic.AddInt32(&ok, 1)
				return
			}

			f, err2 := cFS.Open(apth)
			if err2 != nil {
				err = errorf("", err2)
				return
			}

			defer f.Close()

			b := make([]byte, fi.Size())
			if n, _ := f.Read(b); int64(n) != fi.Size() {
				err = errorf("%v: short read", apth)
				return
			}

			fn := filepath.Join(tmp, filepath.Base(apth))
			if err2 := os.WriteFile(fn, b, 0660); err2 != nil {
				err = errorf("", err2)
				return
			}

			defer os.Remove(fn)

			cmd := exec.Command(cfg.CC, "-E", fn)
			var buf bytes.Buffer
			cmd.Stderr = &buf
			if err2 = cmd.Run(); err2 != nil {
				t.Logf("%v: skip: %v: %s %v", apth, cfg.CC, buf.Bytes(), err2)
				atomic.AddInt32(&skip, 1)
				err = nil
			}
		})
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	p.wg.Wait()
	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	return files, ok, skip, int32(len(fails))
}

// https://gitlab.com/cznic/cc/-/issues/127
func TestIssue127(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.Chdir(wd); err != nil {
			t.Fatal(err)
		}
	}()

	if err := os.Chdir(filepath.FromSlash("testdata/issue127/")); err != nil {
		t.Error(err)
		return
	}

	cd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("working directory: %s", cd)
	cfg := defaultCfg()
	cfg.IncludePaths = append(cfg.IncludePaths, "include")
	if err := Preprocess(
		cfg,
		[]Source{{Name: "main.c"}},
		io.Discard,
	); err != nil {
		t.Fatalf("%v", err)
	}
}

func TestBOM(t *testing.T) {
	for i, v := range []struct {
		src string
		err string
	}{
		{"int main() {}", ""},
		{"\xEF\xBB\xBFint main() {}", ""},
	} {
		switch _, err := Parse(defaultCfg(), []Source{{Value: v.src}}); {
		case v.err == "" && err != nil:
			t.Errorf("%v: unexpected error %v", i, err)
		case v.err != "" && err == nil:
			t.Errorf("%v: unexpected success, expected error matching %v", i, v.err)
		case v.err != "":
			if !regexp.MustCompile(v.err).MatchString(err.Error()) {
				t.Errorf("%v: error %v does not match %v", i, err, v.err)
			}
		}
	}
}

func TestStrCatSep(t *testing.T) {
	for i, v := range []struct {
		src         string
		lit         string
		sep         string
		trailingSep string
	}{
		{`int f() {  "a";}`, `"a"`, "  ", "\n"},
		{`int f() {  L"a";}`, `L"a"`, "  ", "\n"},
		{`int f() { "a" "b";}`, `"ab"`, "  ", "\n"},
		{`int f() { "a""b";}`, `"ab"`, " ", "\n"},
		{`int f() { "a";}`, `"a"`, " ", "\n"},
		{`int f() { "a"` + "\n\t" + `"b"; }`, `"ab"`, " \n\t", "\n"},
		{`int f() { /*x*/ /*y*/ "a";}`, `"a"`, " /*x*/ /*y*/ ", "\n"},
		{`int f() { /*x*/` + "\n" + `/*y*/ "a";}`, `"a"`, " /*x*/\n/*y*/ ", "\n"},
		{`int f() { //x` + "\n" + ` "a";}`, `"a"`, " //x\n ", "\n"},
		{`int f() { //x` + "\n" + `"a";}`, `"a"`, " //x\n", "\n"},
		{`int f() { L"a" L"b";}`, `L"ab"`, "  ", "\n"},
		{`int f() { ` + "\n" + ` "a";}`, `"a"`, " \n ", "\n"},
		{`int f() { ` + "\n" + `"a";}`, `"a"`, " \n", "\n"},
		{`int f() {"a" "b";}`, `"ab"`, " ", "\n"},
		{`int f() {"a"/*y*/"b";}`, `"ab"`, "/*y*/", "\n"},
		{`int f() {"a";} /*x*/ `, `"a"`, "", " /*x*/ \n"},
		{`int f() {"a";} /*x*/`, `"a"`, "", " /*x*/\n"},
		{`int f() {"a";} /*x` + "\n" + `*/ `, `"a"`, "", " /*x\n*/ \n"},
		{`int f() {"a";} `, `"a"`, "", " \n"},
		{`int f() {"a";}/*x*/`, `"a"`, "", "/*x*/\n"},
		{`int f() {"a";}` + "\n", `"a"`, "", "\n"},
		{`int f() {"a";}`, `"a"`, "", "\n"},
		{`int f() {/*x*/ /*y*/ "a";}`, `"a"`, "/*x*/ /*y*/ ", "\n"},
		{`int f() {/*x*/"a""b";}`, `"ab"`, "/*x*/", "\n"},
		{`int f() {/*x*/"a"/*y*/"b";}`, `"ab"`, "/*x*//*y*/", "\n"},
		{`int f() {/*x*/"a";}`, `"a"`, "/*x*/", "\n"},
		{`int f() {/*x*//*y*/ "a";}`, `"a"`, "/*x*//*y*/ ", "\n"},
		{`int f() {/*x*//*y*/"a";}`, `"a"`, "/*x*//*y*/", "\n"},
		{`int f() {//` + "\n" + `"a";}`, `"a"`, "//\n", "\n"},
		{`int f() {//x` + "\n" + `"a";}`, `"a"`, "//x\n", "\n"},
		{`int f() {` + "\n" + ` "a";}`, `"a"`, "\n ", "\n"},
		{`int f() {` + "\n" + `"a";}`, `"a"`, "\n", "\n"},
	} {
		ast, err := Parse(&Config{}, []Source{{Name: "test", Value: v.src}})
		if err != nil {
			t.Errorf("%v: %v", i, err)
			continue
		}

		var n Node
		depth := mathutil.MaxInt
		findNode("PrimaryExpression", ast.TranslationUnit, 0, &n, &depth)
		tok := n.(*PrimaryExpression).Token
		if g, e := string(tok.Src()), v.lit; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := string(tok.Sep()), v.sep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
		if g, e := string(ast.EOF.Sep()), v.trailingSep; g != e {
			t.Errorf("%v: %q %q", i, g, e)
		}
	}
}

func TestParserBug(t *testing.T) {
	blacklistJourdan := map[string]struct{}{
		// Type checking has to detect the fail.
		"bitfield_declaration_ambiguity.fail.c": {},
	}
	t.Run("parser/bug", func(t *testing.T) { testParserBug(t, "testdata/parser/bug", nil) })
	t.Run("jhjourdan", func(t *testing.T) { testParserBug(t, "testdata/jhjourdan", blacklistJourdan) })
}

func testParserBug(t *testing.T, dir string, blacklist map[string]struct{}) {
	tmp := t.TempDir()
	cfg := defaultCfg()
	var fails []string
	var files, ok, skip int
	err := filepath.Walk(filepath.FromSlash(dir), func(pth string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		switch {
		case re != nil:
			if !re.MatchString(pth) {
				skip++
				return nil
			}
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(pth) {
				skip++
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				skip++
				return nil
			}
		}

		if *oTrace {
			fmt.Fprintln(os.Stderr, pth)
		}

		_, err = Parse(
			cfg,
			[]Source{
				{Name: "<predefined>", Value: cfg.Predefined},
				{Name: "<builtin>", Value: builtin},
				{Name: pth},
			},
		)
		switch {
		case strings.Contains(pth, ".fail."):
			if err == nil {
				fails = append(fails, pth)
				t.Errorf("%v: missing error", pth)
			} else {
				if *oTrace {
					t.Log(err)
				}
				ok++
			}
		case err == nil:
			ok++
		default:
			cmd := exec.Command(cfg.CC, "-c", "-o", filepath.Join(tmp, "test.o"), pth)
			var buf bytes.Buffer
			cmd.Stderr = &buf
			if err2 := cmd.Run(); err2 != nil {
				t.Logf("%v: skip: %v: %s %v", pth, cfg.CC, buf.Bytes(), err2)
				skip++
				break
			}

			fails = append(fails, pth)
			t.Errorf("%v: %v", pth, err)
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	for _, v := range fails {
		t.Log(v)
	}
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
}

func TestParse(t *testing.T) {
	cfg := defaultCfg()
	cfg.SysIncludePaths = append(cfg.SysIncludePaths, "Include") // benchmarksgame
	cfg.FS = cFS
	blacklistCompCert := map[string]struct{}{}
	blacklistGCC := map[string]struct{}{
		// Assertions are deprecated, not supported.
		"950919-1.c": {},
	}
	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "linux/s390x":
		blacklistCompCert["aes.c"] = struct{}{} // Unsupported endianness.
	}
	var files, ok, skip, fails int32
	for _, v := range []struct {
		cfg       *Config
		dir       string
		blacklist map[string]struct{}
	}{
		{cfg, "CompCert-3.6/test/c", blacklistCompCert},
		{cfg, "ccgo", nil},
		{cfg, "gcc-9.1.0/gcc/testsuite/gcc.c-torture", blacklistGCC},
		{cfg, "github.com/AbsInt/CompCert/test/c", blacklistCompCert},
		{cfg, "github.com/cxgo", nil},
		{cfg, "github.com/gcc-mirror/gcc/gcc/testsuite", blacklistGCC},
		{cfg, "github.com/vnmakarov", nil},
		{cfg, "sqlite-amalgamation-3370200", nil},
		{cfg, "tcc-0.9.27/tests/tests2", nil},
		{cfg, "benchmarksgame-team.pages.debian.net", nil},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testParse(t, v.cfg, "/"+v.dir, v.blacklist)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testParse(t *testing.T, cfg *Config, dir string, blacklist map[string]struct{}) (files, ok, skip, nfails int32) {
	tmp := t.TempDir()
	var fails []string
	p := newParallel()
	err := walk(dir, func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(pth) {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		}
		apth := pth
		p.exec(func() {
			if *oTrace {
				fmt.Fprintln(os.Stderr, apth)
			}

			var err error

			defer func() {
				p.Lock()

				defer p.Unlock()

				if err != nil {
					fails = append(fails, apth)
					t.Errorf("%v: %v", apth, err)
				}

			}()

			func() {
				defer func() {
					if e := recover(); e != nil && err == nil {
						err = fmt.Errorf("%v: PANIC: %v", apth, e)
						// trc("\n%s", debug.Stack())
					}
				}()

				if _, err = Parse(
					cfg,
					[]Source{
						{Name: "<predefined>", Value: cfg.Predefined},
						{Name: "<builtin>", Value: builtin},
						{Name: apth, FS: cFS},
					},
				); err == nil {
					atomic.AddInt32(&ok, 1)
					return
				}
			}()

			if err == nil {
				return
			}

			f, err2 := cFS.Open(apth)
			if err2 != nil {
				err = errorf("", err2)
				return
			}

			defer f.Close()

			b := make([]byte, fi.Size())
			if n, _ := f.Read(b); int64(n) != fi.Size() {
				err = errorf("%v: short read", apth)
				return
			}

			fn := filepath.Join(tmp, filepath.Base(apth))
			if err2 := os.WriteFile(fn, b, 0660); err2 != nil {
				err = errorf("", err2)
				return
			}

			defer os.Remove(fn)

			cmd := exec.Command(cfg.CC, "-c", "-o", filepath.Join(tmp, "test.o"), fn)
			var buf bytes.Buffer
			cmd.Stderr = &buf
			if err2 = cmd.Run(); err2 != nil {
				t.Logf("%v: skip: %v: %s %v", apth, cfg.CC, buf.Bytes(), err2)
				atomic.AddInt32(&skip, 1)
				err = nil
				return
			}
		})
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	p.wg.Wait()
	for _, v := range fails {
		t.Log(v)
	}
	// fmt.Fprintf(os.Stderr, "%v: files %v, skip %v, ok %v, fails %v\n", dir, files, skip, ok, len(fails))
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	return files, ok, skip, int32(len(fails))
}

func TestTranslate(t *testing.T) {
	return //TODO-
	cfg := defaultCfg()
	cfg.SysIncludePaths = append(cfg.SysIncludePaths, "Include") // benchmarksgame
	cfg.FS = cFS
	blacklistCompCert := map[string]struct{}{}
	blacklistGCC := map[string]struct{}{
		// Assertions are deprecated, not supported.
		"950919-1.c": {},
	}
	switch fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH) {
	case "linux/s390x":
		blacklistCompCert["aes.c"] = struct{}{} // Unsupported endianness.
	}
	var files, ok, skip, fails int32
	for _, v := range []struct {
		cfg       *Config
		dir       string
		blacklist map[string]struct{}
	}{
		{cfg, "CompCert-3.6/test/c", blacklistCompCert},
		{cfg, "ccgo", nil},
		{cfg, "gcc-9.1.0/gcc/testsuite/gcc.c-torture", blacklistGCC},
		{cfg, "github.com/AbsInt/CompCert/test/c", blacklistCompCert},
		{cfg, "github.com/cxgo", nil},
		{cfg, "github.com/gcc-mirror/gcc/gcc/testsuite", blacklistGCC},
		{cfg, "github.com/vnmakarov", nil},
		{cfg, "sqlite-amalgamation-3370200", nil},
		{cfg, "tcc-0.9.27/tests/tests2", nil},
		{cfg, "benchmarksgame-team.pages.debian.net", nil},
	} {
		t.Run(v.dir, func(t *testing.T) {
			f, o, s, n := testTranslate(t, v.cfg, "/"+v.dir, v.blacklist)
			files += f
			ok += o
			skip += s
			fails += n
		})
	}
	t.Logf("TOTAL: files %v, skip %v, ok %v, fails %v", files, skip, ok, fails)
}

func testTranslate(t *testing.T, cfg *Config, dir string, blacklist map[string]struct{}) (files, ok, skip, nfails int32) {
	tmp := t.TempDir()
	var fails []string
	p := newParallel()
	err := walk(dir, func(pth string, fi os.FileInfo) error {
		if fi.IsDir() {
			return nil
		}

		if filepath.Ext(pth) != ".c" {
			return nil
		}

		files++
		switch {
		case re != nil:
			if !re.MatchString(pth) {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		default:
			if _, ok := blacklist[filepath.Base(pth)]; ok {
				atomic.AddInt32(&skip, 1)
				return nil
			}
		}
		apth := pth
		p.exec(func() {
			if *oTrace {
				fmt.Fprintln(os.Stderr, apth)
			}

			var err error

			defer func() {
				p.Lock()

				defer p.Unlock()

				if err != nil {
					fails = append(fails, apth)
					t.Errorf("%v: %v", apth, err)
				}

			}()

			func() {
				defer func() {
					if e := recover(); e != nil && err == nil {
						err = fmt.Errorf("%v: PANIC: %v", apth, e)
						trc("%v: PANIC: %v\n%s", apth, e, debug.Stack())
						os.Exit(1)
					}
				}()

				if _, err = Translate(
					cfg,
					[]Source{
						{Name: "<predefined>", Value: cfg.Predefined},
						{Name: "<builtin>", Value: builtin},
						{Name: apth, FS: cFS},
					},
				); err == nil {
					atomic.AddInt32(&ok, 1)
					return
				}
			}()

			if err == nil {
				return
			}

			f, err2 := cFS.Open(apth)
			if err2 != nil {
				err = errorf("", err2)
				return
			}

			defer f.Close()

			b := make([]byte, fi.Size())
			if n, _ := f.Read(b); int64(n) != fi.Size() {
				err = errorf("%v: short read", apth)
				return
			}

			fn := filepath.Join(tmp, filepath.Base(apth))
			if err2 := os.WriteFile(fn, b, 0660); err2 != nil {
				err = errorf("", err2)
				return
			}

			defer os.Remove(fn)

			cmd := exec.Command(cfg.CC, "-c", "-o", filepath.Join(tmp, "test.o"), fn)
			var buf bytes.Buffer
			cmd.Stderr = &buf
			if err2 = cmd.Run(); err2 != nil {
				t.Logf("%v: skip: %v: %s %v", apth, cfg.CC, buf.Bytes(), err2)
				atomic.AddInt32(&skip, 1)
				err = nil
				return
			}
		})
		return nil
	})
	if err != nil {
		t.Error(err)
	}

	p.wg.Wait()
	for _, v := range fails {
		t.Log(v)
	}
	// fmt.Fprintf(os.Stderr, "%v: files %v, skip %v, ok %v, fails %v\n", dir, files, skip, ok, len(fails))
	t.Logf("files %v, skip %v, ok %v, fails %v", files, skip, ok, len(fails))
	return files, ok, skip, int32(len(fails))
}
