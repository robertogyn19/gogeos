package geos

/*
#include "geos.h"
*/
import "C"

import (
	"runtime"
)

type coordSeq struct {
	c *C.GEOSCoordSequence
}

func newCoordSeq(size, dims int) *coordSeq {
	p := cGEOSCoordSeq_create(geosGlobalContext, C.uint(size), C.uint(dims))
	if p == nil {
		return nil
	}
	return coordSeqFromPtr(p)
}

func coordSeqFromPtr(ptr *C.GEOSCoordSequence) *coordSeq {
	cs := &coordSeq{c: ptr}
	runtime.SetFinalizer(cs, func(*coordSeq) {
		cGEOSCoordSeq_destroy(geosGlobalContext, ptr)
	})
	return cs
}

func coordSeqFromSlice(coords []Coord) (*coordSeq, error) {
	// XXX: handle 3-dim
	ptr := cGEOSCoordSeq_create(geosGlobalContext, C.uint(len(coords)), C.uint(2))
	if ptr == nil {
		return nil, Error()
	}
	cs := &coordSeq{c: ptr}
	for i, c := range coords {
		if err := cs.setX(i, c.X); err != nil {
			return nil, err
		}
		if err := cs.setY(i, c.Y); err != nil {
			return nil, err
		}
	}
	return cs, nil
}

func (c *coordSeq) Clone() (*coordSeq, error) {
	p := cGEOSCoordSeq_clone(geosGlobalContext, c.c)
	if p == nil {
		return nil, Error()
	}
	return coordSeqFromPtr(p), nil
}

func (c *coordSeq) setX(idx int, val float64) error {
	i := cGEOSCoordSeq_setX(geosGlobalContext, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *coordSeq) setY(idx int, val float64) error {
	i := cGEOSCoordSeq_setY(geosGlobalContext, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *coordSeq) setZ(idx int, val float64) error {
	i := cGEOSCoordSeq_setZ(geosGlobalContext, c.c, C.uint(idx), C.double(val))
	if i == 0 {
		return Error()
	}
	return nil
}

func (c *coordSeq) x(idx int) (float64, error) {
	var val C.double
	i := cGEOSCoordSeq_getX(geosGlobalContext, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *coordSeq) y(idx int) (float64, error) {
	var val C.double
	i := cGEOSCoordSeq_getY(geosGlobalContext, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *coordSeq) z(idx int) (float64, error) {
	var val C.double
	i := cGEOSCoordSeq_getZ(geosGlobalContext, c.c, C.uint(idx), &val)
	if i == 0 {
		return 0.0, Error()
	}
	return float64(val), nil
}

func (c *coordSeq) size() (int, error) {
	var val C.uint
	i := cGEOSCoordSeq_getSize(geosGlobalContext, c.c, &val)
	if i == 0 {
		return 0, Error()
	}
	return int(val), nil
}

func (c *coordSeq) dims() (int, error) {
	var val C.uint
	i := cGEOSCoordSeq_getDimensions(geosGlobalContext, c.c, &val)
	if i == 0 {
		return 0, Error()
	}
	return int(val), nil
}
