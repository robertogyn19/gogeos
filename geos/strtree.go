package geos

/*
#include "geos.h"
*/
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
	"sync"
)

type STRTreeCallbackFunc func(item []byte)
type STRTreeItem interface {
	Parse() []byte
}

type StrTree struct {
	r  *C.GEOSSTRtree
	cb C.GEOSQueryCallback
}

func NewSTRTree(capacity int) *StrTree {
	r := cGEOSSTRtree_create(geosGlobalContext, C.size_t(capacity))

	if r == nil {
		return nil
	}

	strtree := &StrTree{r, NewCallback()}
	runtime.SetFinalizer(strtree, (*StrTree).destroy)

	return strtree
}

func (tree *StrTree) Insert(g *Geometry, item STRTreeItem) {
	bstr := item.Parse()
	cstr := C.CString(string(bstr))
	p := (*C.void)(unsafe.Pointer(cstr))

	cGEOSSTRtree_insert(geosGlobalContext, tree.r, g.g, p)
}

func (tree *StrTree) Query(g *Geometry, cb STRTreeCallbackFunc) {
	cbid := register(cb)
	ccbid := C.int(cbid)
	cGEOSSTRtree_query(geosGlobalContext, tree.r, g.g, tree.cb, (*C.void)(unsafe.Pointer(&ccbid)))
	defer unregister(cbid)
}

func (tree *StrTree) Iterate(cb STRTreeCallbackFunc) {
	cbid := register(cb)
	ccbid := C.int(cbid)
	cGEOSSTRtree_iterate(geosGlobalContext, tree.r, tree.cb, (*C.void)(unsafe.Pointer(&ccbid)))
	defer unregister(cbid)
}

func (tree *StrTree) destroy() {
	cGEOSSTRtree_destroy(geosGlobalContext, tree.r)
	tree.r = nil
}

func NewCallback() C.GEOSQueryCallback {
	return NewCustomCallback(unsafe.Pointer(C.strTreeQueryCallbackGo))
}

func NewCustomCallback(pointer unsafe.Pointer) C.GEOSQueryCallback {
	return (C.GEOSQueryCallback)(pointer)
}

/* ------------------------------------------------------------------------- */
/*                                                                           */
/*                            Callback functions                             */
/*                                                                           */
/* ------------------------------------------------------------------------- */

var mu sync.Mutex
var fns = make(map[int]STRTreeCallbackFunc)
var cbIndex int

//export StrTreeQueryCallbackGo
func StrTreeQueryCallbackGo(item unsafe.Pointer, data unsafe.Pointer) {
	s := itemBytes(item)

	if data == nil {
		fmt.Println("Callback data not found. Forgot to register a callback?")
		return
	}

	ccbid := *(*C.int)(data)
	cbid := int(ccbid)
	cb := lookup(cbid)
	cb(s)
}

func itemBytes(item unsafe.Pointer) []byte {
	value1 := (*C.char)(item)
	return []byte(C.GoString(value1))
}

func register(fn STRTreeCallbackFunc) int {
	mu.Lock()
	defer mu.Unlock()
	cbIndex++

	for fns[cbIndex] != nil {
		cbIndex++
	}

	fns[cbIndex] = fn
	return cbIndex
}

func lookup(i int) STRTreeCallbackFunc {
	mu.Lock()
	defer mu.Unlock()
	return fns[i]
}

func unregister(i int) {
	mu.Lock()
	defer mu.Unlock()
	delete(fns, i)
}