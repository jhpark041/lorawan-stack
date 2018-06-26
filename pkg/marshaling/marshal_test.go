// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package marshaling_test

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"testing"

	"github.com/kr/pretty"
	"github.com/smartystreets/assertions"
	"github.com/smartystreets/assertions/should"
	. "go.thethings.network/lorawan-stack/pkg/marshaling"
)

func TestFlattened(t *testing.T) {
	for _, tc := range []struct {
		in  map[string]interface{}
		out map[string]interface{}
	}{
		{
			map[string]interface{}{
				"foo": map[string]interface{}{
					"bar": os.Stdout,
					"baz": map[string]string{"test": "foo"},
					"recursive": map[string]interface{}{
						"hello": struct{ hi string }{"hello"},
					},
				},
				"42": map[string]interface{}{
					"foo": 42,
					"baz": "baz",
				},
			},
			map[string]interface{}{
				"foo.bar":             os.Stdout,
				"foo.baz":             map[string]string{"test": "foo"},
				"foo.recursive.hello": struct{ hi string }{"hello"},
				"42.foo":              42,
				"42.baz":              "baz",
			},
		},
	} {
		assertions.New(t).So(Flattened(tc.in), should.Resemble, tc.out)
	}
}

func TestFlattenedValue(t *testing.T) {
	for _, tc := range []struct {
		in  map[string]reflect.Value
		out map[string]reflect.Value
	}{
		{
			map[string]reflect.Value{
				"foo": reflect.ValueOf(map[string]reflect.Value{
					"bar": reflect.ValueOf(os.Stdout),
					"baz": reflect.ValueOf(map[string]string{"test": "foo"}),
					"recursive": reflect.ValueOf(map[string]reflect.Value{
						"hello": reflect.ValueOf(struct{ hi string }{"hello"}),
					}),
				}),
				"42": reflect.ValueOf(map[string]reflect.Value{
					"foo": reflect.ValueOf(42),
					"baz": reflect.ValueOf("baz"),
				}),
			},
			map[string]reflect.Value{
				"foo.bar":             reflect.ValueOf(os.Stdout),
				"foo.baz":             reflect.ValueOf(map[string]string{"test": "foo"}),
				"foo.recursive.hello": reflect.ValueOf(struct{ hi string }{"hello"}),
				"42.foo":              reflect.ValueOf(42),
				"42.baz":              reflect.ValueOf("baz"),
			},
		},
	} {
		a := assertions.New(t)
		ret := FlattenedValue(tc.in)
		if !a.So(ret, should.HaveLength, len(tc.out)) {
			return
		}
		for k, v := range tc.out {
			a.So(ret[k].Interface(), should.Resemble, v.Interface())
			a.So(ret[k].Type(), should.Equal, v.Type())
		}
	}
}

func TestToBytes(t *testing.T) {
	for i, tc := range ByteValues {
		t.Run(fmt.Sprintf("%d/%+v", i, tc.value), func(t *testing.T) {
			a := assertions.New(t)

			b, err := ToBytes(tc.value)
			if !a.So(err, should.BeNil) || !a.So(b, should.HaveLength, len(tc.bytes)) {
				return
			}

			if len(b) >= 2 && Encoding(b[1]) == GobEncoding {
				if !a.So(b[2], should.Resemble, tc.bytes[2]) {
					t.Log("Encoding type mismatch")
					return
				}

				// Gob encoding output is not deterministic, hence we need
				// to check if value obtained by decoding the bytes
				// using gob resembles the original value.
				typ := reflect.TypeOf(tc.value)

				ve := reflect.New(typ)
				if err := gob.NewDecoder(bytes.NewBuffer(b[2:])).DecodeValue(ve); err != nil {
					panic("Failed to decode testcase bytes from gob")
				}

				va := reflect.New(typ)
				err = gob.NewDecoder(bytes.NewBuffer(tc.bytes[2:])).DecodeValue(va)
				a.So(err, should.BeNil)
				if !a.So(va.Interface(), should.Resemble, ve.Interface()) {
					t.Log(pretty.Sprint("Value:", tc.value))
				}
			} else {
				if !a.So(b, should.Resemble, tc.bytes) {
					t.Log(pretty.Sprint("Value:", tc.value))
				}
			}
		})
	}
}

func TestMarshalMap(t *testing.T) {
	for i, v := range structValues {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assertions.New(t)

			m, err := MarshalMap(v.unmarshaled)
			a.So(err, should.BeNil)
			if !a.So(m, should.Resemble, v.marshaled) {
				pretty.Ldiff(t, m, v.marshaled)
			}
		})
	}
}

func TestMarshalByteMap(t *testing.T) {
	for i, v := range structValues {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			a := assertions.New(t)

			m, err := MarshalByteMap(v.unmarshaled)
			a.So(err, should.BeNil)
			if !a.So(m, should.Resemble, v.bytes) {
				pretty.Ldiff(t, m, v.bytes)
			}
		})
	}
}

func TestOrderPreservation(t *testing.T) {
	for i, tc := range []struct {
		values []interface{}
	}{
		{
			values: []interface{}{byte(0), 0, int8(0), uintptr(0), int8(22), uint8(24), int8(42), byte(43)},
		},
		{
			values: []interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 20, 40, 66, 120, 420, 1244, 4154, 0x4444},
		},
		{
			values: []interface{}{int32(41), uint32(42)},
		},
		{
			values: []interface{}{"apple", "bar", "crowbar", "foo", "foo", "guinness", "la chouffe", "leffe", "nintendo", "vodka", "zelda"},
		},
		{
			values: []interface{}{[3]byte{0, 1, 2}, []byte{0, 1, 3}, [3]byte{0, 1, 4}, [2]byte{0, 2}, []byte{42}},
		},
		{
			values: []interface{}{1.2, 2.454532, 3.14159265358979323, 4.20, 4.5, 4.56},
		},
	} {
		t.Run(fmt.Sprintf("%d/%+v", i, tc.values), func(t *testing.T) {
			enc := make([]string, 0, len(tc.values))
			for _, v := range tc.values {
				b, err := ToBytes(v)
				if err != nil {
					t.Skip("ToBytes broken, skipping...")
				}
				enc = append(enc, string(b))
			}
			if !assertions.New(t).So(sort.StringsAreSorted(enc), should.BeTrue) {
				for i := 0; i < len(enc)-1; i++ {
					if si, sj := enc[i], enc[i+1]; si > sj {
						t.Errorf("mismatch at %d and %d: %+v(%+v) > %+v(%+v)", i, i+1, []byte(si), tc.values[i], []byte(sj), tc.values[i+1])
					}
				}
			}
		})
	}
}