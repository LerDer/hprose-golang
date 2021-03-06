/*--------------------------------------------------------*\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: https://hprose.com                     |
|                                                          |
| encoding/relect.go                                       |
|                                                          |
| LastModified: Apr 18, 2020                               |
| Author: Ma Bingyao <andot@hprose.com>                    |
|                                                          |
\*________________________________________________________*/

package encoding

import (
	"reflect"
	"unsafe"
)

type eface struct {
	typ uintptr
	ptr unsafe.Pointer
}

func unpackEFace(ptr *interface{}) *eface {
	return (*eface)(unsafe.Pointer(ptr))
}

func unsafeString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}

func checkType(v interface{}) reflect.Type {
	t := reflect.TypeOf(v)
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
