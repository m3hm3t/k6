/*
 */
// nolint:ireturn
package data

import (
	"github.com/dop251/goja"
)

type ImmutableArrayBuffer struct {
	arr []byte
}

// Len returns the current immutable array buffer length.
func (w WrappedImmutableArrayBuffer) Len() int {
	return len(w.arr)
}

// Get an item at index idx.
// // Note that idx may be any integer, negative or beyond the current length.
func (w WrappedImmutableArrayBuffer) Get(index int) goja.Value {
	if index < 0 || index >= len(w.arr) {
		return goja.Undefined()
	}

	// TODO: do I need to freeze the value I return?

	return w.rt.ToValue(w.arr[index])
}

// Set an item at index idx.
// // Note that idx may be any integer, negative or beyond the current length.
// // The expected behaviour when it's beyond length is that the array's length is increased to accommodate
// // the item. All elements in the 'new' section of the array should be zeroed.
func (w WrappedImmutableArrayBuffer) Set(idx int, val goja.Value) bool {
	panic(w.rt.NewTypeError("SharedArray is immutable")) // this is specifically a type error
}

// // SetLen is called when the array's 'length' property is changed. If the length is increased all elements in the
// // 'new' section of the array should be zeroed.
func (w WrappedImmutableArrayBuffer) SetLen(l int) bool {
	panic(w.rt.NewTypeError("SharedArray is immutable")) // this is specifically a type error
}

func (w WrappedImmutableArrayBuffer) Bytes() []byte {
	return w.arr
}

type Wrapper interface {
	Wrap(rt *goja.Runtime) goja.Value
}

func (iab ImmutableArrayBuffer) Wrap(rt *goja.Runtime) goja.Value {
	return rt.NewDynamicArray(WrappedImmutableArrayBuffer{
		ImmutableArrayBuffer: iab,
		rt:                   rt,
	})
}

type WrappedImmutableArrayBuffer struct {
	ImmutableArrayBuffer

	rt *goja.Runtime
}
