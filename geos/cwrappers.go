package geos

// Created mechanically from C API header - DO NOT EDIT

/*
#include <geos_c.h>
*/
import "C"

import (
	"unsafe"
)

func cinitGEOS(notice_function C.GEOSMessageHandler, error_function C.GEOSMessageHandler) C.GEOSContextHandle_t {
	return C.initGEOS_r(notice_function, error_function)
}

func (context *GeosContext) cfinishGEOS() {
	cfinishGEOS(context)
}

func cfinishGEOS(context *GeosContext) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.finishGEOS_r(context.handle)
}

func cGEOSversion() *C.char {
	return C.GEOSversion()
}

func (context *GeosContext) cGEOSGeomFromWKT(wkt *C.char) *C.GEOSGeometry {
	return cGEOSGeomFromWKT(context, wkt)
}

func cGEOSGeomFromWKT(context *GeosContext, wkt *C.char) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomFromWKT_r(context.handle, wkt)
}

func (context *GeosContext) cGEOSGeomToWKT(g *C.GEOSGeometry) *C.char {
	return cGEOSGeomToWKT(context, g)
}

func cGEOSGeomToWKT(context *GeosContext, g *C.GEOSGeometry) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomToWKT_r(context.handle, g)
}

func (context *GeosContext) cGEOS_getWKBOutputDims() C.int {
	return cGEOS_getWKBOutputDims(context)
}

func cGEOS_getWKBOutputDims(context *GeosContext) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOS_getWKBOutputDims_r(context.handle)
}

func cGEOS_setWKBOutputDims(newDims C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBOutputDims_r(handle, newDims)
}

func cGEOS_getWKBByteOrder() C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_getWKBByteOrder_r(handle)
}

func cGEOS_setWKBByteOrder(byteOrder C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOS_setWKBByteOrder_r(handle, byteOrder)
}

func cGEOSGeomFromWKB_buf(wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromWKB_buf_r(handle, wkb, size)
}

func cGEOSGeomToWKB_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToWKB_buf_r(handle, g, size)
}

func cGEOSGeomFromHEX_buf(hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomFromHEX_buf_r(handle, hex, size)
}

func cGEOSGeomToHEX_buf(g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSGeomToHEX_buf_r(handle, g, size)
}

func (context *GeosContext) cGEOSCoordSeq_create(size C.uint, dims C.uint) *C.GEOSCoordSequence {
	return cGEOSCoordSeq_create(context, size, dims)
}

func cGEOSCoordSeq_create(context *GeosContext, size C.uint, dims C.uint) *C.GEOSCoordSequence {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_create_r(context.handle, size, dims)
}

func (context *GeosContext) cGEOSCoordSeq_clone(s *C.GEOSCoordSequence) *C.GEOSCoordSequence {
	return cGEOSCoordSeq_clone(context, s)
}

func cGEOSCoordSeq_clone(context *GeosContext, s *C.GEOSCoordSequence) *C.GEOSCoordSequence {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_clone_r(context.handle, s)
}

func (context *GeosContext) cGEOSCoordSeq_destroy(s *C.GEOSCoordSequence) {
	cGEOSCoordSeq_destroy(context, s)
}

func cGEOSCoordSeq_destroy(context *GeosContext, s *C.GEOSCoordSequence) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSCoordSeq_destroy_r(context.handle, s)
}

func (context *GeosContext) cGEOSCoordSeq_setX(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return cGEOSCoordSeq_setX(context, s, idx, val)
}

func cGEOSCoordSeq_setX(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_setX_r(context.handle, s, idx, val)
}

func (context *GeosContext) cGEOSCoordSeq_setY(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return cGEOSCoordSeq_setY(context, s, idx, val)
}

func cGEOSCoordSeq_setY(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_setY_r(context.handle, s, idx, val)
}

func (context *GeosContext) cGEOSCoordSeq_setZ(s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	return cGEOSCoordSeq_setZ(context, s, idx, val)
}

func cGEOSCoordSeq_setZ(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_setZ_r(context.handle, s, idx, val)
}

func cGEOSCoordSeq_setOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSCoordSeq_setOrdinate_r(handle, s, idx, dim, val)
}

func (context *GeosContext) cGEOSCoordSeq_getX(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return cGEOSCoordSeq_getX(context, s, idx, val)
}

func cGEOSCoordSeq_getX(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getX_r(context.handle, s, idx, val)
}

func (context *GeosContext) cGEOSCoordSeq_getY(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return cGEOSCoordSeq_getY(context, s, idx, val)
}

func cGEOSCoordSeq_getY(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getY_r(context.handle, s, idx, val)
}

func (context *GeosContext) cGEOSCoordSeq_getZ(s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	return cGEOSCoordSeq_getZ(context, s, idx, val)
}

func cGEOSCoordSeq_getZ(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, val *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getZ_r(context.handle, s, idx, val)
}

func (context *GeosContext) cGEOSCoordSeq_getOrdinate(s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val *C.double) C.int {
	return cGEOSCoordSeq_getOrdinate(context, s, idx, dim, val)
}

func cGEOSCoordSeq_getOrdinate(context *GeosContext, s *C.GEOSCoordSequence, idx C.uint, dim C.uint, val *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getOrdinate_r(context.handle, s, idx, dim, val)
}

func (context *GeosContext) cGEOSCoordSeq_getSize(s *C.GEOSCoordSequence, size *C.uint) C.int {
	return cGEOSCoordSeq_getSize(context, s, size)
}

func cGEOSCoordSeq_getSize(context *GeosContext, s *C.GEOSCoordSequence, size *C.uint) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getSize_r(context.handle, s, size)
}

func (context *GeosContext) cGEOSCoordSeq_getDimensions(s *C.GEOSCoordSequence, dims *C.uint) C.int {
	return cGEOSCoordSeq_getDimensions(context, s, dims)
}

func cGEOSCoordSeq_getDimensions(context *GeosContext, s *C.GEOSCoordSequence, dims *C.uint) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoordSeq_getDimensions_r(context.handle, s, dims)
}

func (context *GeosContext) cGEOSProject(g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	return cGEOSProject(context, g, p)
}

func cGEOSProject(context *GeosContext, g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSProject_r(context.handle, g, p)
}

func (context *GeosContext) cGEOSInterpolate(g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	return cGEOSInterpolate(context, g, d)
}

func cGEOSInterpolate(context *GeosContext, g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSInterpolate_r(context.handle, g, d)
}

func (context *GeosContext) cGEOSProjectNormalized(g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	return cGEOSProjectNormalized(context, g, p)
}

func cGEOSProjectNormalized(context *GeosContext, g *C.GEOSGeometry, p *C.GEOSGeometry) C.double {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSProjectNormalized_r(context.handle, g, p)
}

func cGEOSInterpolateNormalized(g *C.GEOSGeometry, d C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSInterpolateNormalized_r(handle, g, d)
}

func cGEOSBufferParams_create() *C.GEOSBufferParams {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_create_r(handle)
}

func cGEOSBufferParams_destroy(parms *C.GEOSBufferParams) {
	handlemu.Lock()
	defer handlemu.Unlock()
	C.GEOSBufferParams_destroy_r(handle, parms)
}

func cGEOSBufferParams_setEndCapStyle(p *C.GEOSBufferParams, style C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setEndCapStyle_r(handle, p, style)
}

func cGEOSBufferParams_setJoinStyle(p *C.GEOSBufferParams, joinStyle C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setJoinStyle_r(handle, p, joinStyle)
}

func cGEOSBufferParams_setMitreLimit(p *C.GEOSBufferParams, mitreLimit C.double) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setMitreLimit_r(handle, p, mitreLimit)
}

func cGEOSBufferParams_setQuadrantSegments(p *C.GEOSBufferParams, quadSegs C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setQuadrantSegments_r(handle, p, quadSegs)
}

func cGEOSBufferParams_setSingleSided(p *C.GEOSBufferParams, singleSided C.int) C.int {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferParams_setSingleSided_r(handle, p, singleSided)
}

func cGEOSBufferWithParams(g1 *C.GEOSGeometry, p *C.GEOSBufferParams, width C.double) *C.GEOSGeometry {
	handlemu.Lock()
	defer handlemu.Unlock()
	return C.GEOSBufferWithParams_r(handle, g1, p, width)
}

func (context *GeosContext) cGEOSBuffer(g1 *C.GEOSGeometry, width C.double, quadsegs C.int) *C.GEOSGeometry {
	return cGEOSBuffer(context, g1, width, quadsegs)
}

func cGEOSBuffer(context *GeosContext, g1 *C.GEOSGeometry, width C.double, quadsegs C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSBuffer_r(context.handle, g1, width, quadsegs)
}

func (context *GeosContext) cGEOSBufferWithStyle(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, endCapStyle C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	return cGEOSBufferWithStyle(context, g1, width, quadsegs, endCapStyle, joinStyle, mitreLimit)
}

func cGEOSBufferWithStyle(context *GeosContext, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, endCapStyle C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSBufferWithStyle_r(context.handle, g1, width, quadsegs, endCapStyle, joinStyle, mitreLimit)
}

func (context *GeosContext) cGEOSSingleSidedBuffer(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double, leftSide C.int) *C.GEOSGeometry {
	return cGEOSSingleSidedBuffer(context, g1, width, quadsegs, joinStyle, mitreLimit, leftSide)
}

func cGEOSSingleSidedBuffer(context *GeosContext, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double, leftSide C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSingleSidedBuffer_r(context.handle, g1, width, quadsegs, joinStyle, mitreLimit, leftSide)
}

func (context *GeosContext) cGEOSOffsetCurve(g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	return cGEOSOffsetCurve(context, g1, width, quadsegs, joinStyle, mitreLimit)
}

func cGEOSOffsetCurve(context *GeosContext, g1 *C.GEOSGeometry, width C.double, quadsegs C.int, joinStyle C.int, mitreLimit C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSOffsetCurve_r(context.handle, g1, width, quadsegs, joinStyle, mitreLimit)
}

func (context *GeosContext) cGEOSGeom_createPoint(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return cGEOSGeom_createPoint(context, s)
}

func cGEOSGeom_createPoint(context *GeosContext, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createPoint_r(context.handle, s)
}

func (context *GeosContext) cGEOSGeom_createEmptyPoint() *C.GEOSGeometry {
	return cGEOSGeom_createEmptyPoint(context, )
}

func cGEOSGeom_createEmptyPoint(context *GeosContext) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createEmptyPoint_r(context.handle)
}

func (context *GeosContext) cGEOSGeom_createLinearRing(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return cGEOSGeom_createLinearRing(context, s)
}

func cGEOSGeom_createLinearRing(context *GeosContext, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createLinearRing_r(context.handle, s)
}

func (context *GeosContext) cGEOSGeom_createLineString(s *C.GEOSCoordSequence) *C.GEOSGeometry {
	return cGEOSGeom_createLineString(context, s)
}

func cGEOSGeom_createLineString(context *GeosContext, s *C.GEOSCoordSequence) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createLineString_r(context.handle, s)
}

func (context *GeosContext) cGEOSGeom_createEmptyLineString() *C.GEOSGeometry {
	return cGEOSGeom_createEmptyLineString(context, )
}

func cGEOSGeom_createEmptyLineString(context *GeosContext) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createEmptyLineString_r(context.handle)
}

func (context *GeosContext) cGEOSGeom_createEmptyPolygon() *C.GEOSGeometry {
	return cGEOSGeom_createEmptyPolygon(context, )
}

func cGEOSGeom_createEmptyPolygon(context *GeosContext) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createEmptyPolygon_r(context.handle)
}

func (context *GeosContext) cGEOSGeom_createPolygon(shell *C.GEOSGeometry, holes **C.GEOSGeometry, nholes C.uint) *C.GEOSGeometry {
	return cGEOSGeom_createPolygon(context, shell, holes, nholes)
}

func cGEOSGeom_createPolygon(context *GeosContext, shell *C.GEOSGeometry, holes **C.GEOSGeometry, nholes C.uint) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createPolygon_r(context.handle, shell, holes, nholes)
}

func (context *GeosContext) cGEOSGeom_createCollection(_type C.int, geoms **C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return cGEOSGeom_createCollection(context, _type, geoms, ngeoms)
}

func cGEOSGeom_createCollection(context *GeosContext, _type C.int, geoms **C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createCollection_r(context.handle, _type, geoms, ngeoms)
}

func (context *GeosContext) cGEOSGeom_createEmptyCollection(_type C.int) *C.GEOSGeometry {
	return cGEOSGeom_createEmptyCollection(context, _type)
}

func cGEOSGeom_createEmptyCollection(context *GeosContext, _type C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_createEmptyCollection_r(context.handle, _type)
}

func (context *GeosContext) cGEOSGeom_clone(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGeom_clone(context, g)
}

func cGEOSGeom_clone(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_clone_r(context.handle, g)
}

func (context *GeosContext) cGEOSGeom_destroy(g *C.GEOSGeometry) {
	cGEOSGeom_destroy(context, g)
}

func cGEOSGeom_destroy(context *GeosContext, g *C.GEOSGeometry) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSGeom_destroy_r(context.handle, g)
}

func (context *GeosContext) cGEOSEnvelope(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSEnvelope(context, g1)
}

func cGEOSEnvelope(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSEnvelope_r(context.handle, g1)
}

func (context *GeosContext) cGEOSIntersection(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSIntersection(context, g1, g2)
}

func cGEOSIntersection(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSIntersection_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSConvexHull(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSConvexHull(context, g1)
}

func cGEOSConvexHull(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSConvexHull_r(context.handle, g1)
}

func (context *GeosContext) cGEOSDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSDifference(context, g1, g2)
}

func cGEOSDifference(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSDifference_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSSymDifference(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSSymDifference(context, g1, g2)
}

func cGEOSSymDifference(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSymDifference_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSBoundary(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSBoundary(context, g1)
}

func cGEOSBoundary(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSBoundary_r(context.handle, g1)
}

func (context *GeosContext) cGEOSUnion(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSUnion(context, g1, g2)
}

func cGEOSUnion(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSUnion_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSUnaryUnion(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSUnaryUnion(context, g)
}

func cGEOSUnaryUnion(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSUnaryUnion_r(context.handle, g)
}

func (context *GeosContext) cGEOSUnionCascaded(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSUnionCascaded(context, g1)
}

func cGEOSUnionCascaded(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSUnionCascaded_r(context.handle, g1)
}

func (context *GeosContext) cGEOSPointOnSurface(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSPointOnSurface(context, g1)
}

func cGEOSPointOnSurface(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPointOnSurface_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGetCentroid(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGetCentroid(context, g)
}

func cGEOSGetCentroid(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetCentroid_r(context.handle, g)
}

func (context *GeosContext) cGEOSPolygonize(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return cGEOSPolygonize(context, geoms, ngeoms)
}

func cGEOSPolygonize(context *GeosContext, geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPolygonize_r(context.handle, &geoms[0], ngeoms)
}

func (context *GeosContext) cGEOSPolygonizer_getCutEdges(geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	return cGEOSPolygonizer_getCutEdges(context, geoms, ngeoms)
}

func cGEOSPolygonizer_getCutEdges(context *GeosContext, geoms []*C.GEOSGeometry, ngeoms C.uint) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPolygonizer_getCutEdges_r(context.handle, &geoms[0], ngeoms)
}

func (context *GeosContext) cGEOSPolygonize_full(input *C.GEOSGeometry, cuts **C.GEOSGeometry, dangles **C.GEOSGeometry, invalidRings **C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSPolygonize_full(context, input, cuts, dangles, invalidRings)
}

func cGEOSPolygonize_full(context *GeosContext, input *C.GEOSGeometry, cuts **C.GEOSGeometry, dangles **C.GEOSGeometry, invalidRings **C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPolygonize_full_r(context.handle, input, cuts, dangles, invalidRings)
}

func (context *GeosContext) cGEOSLineMerge(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSLineMerge(context, g)
}

func cGEOSLineMerge(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSLineMerge_r(context.handle, g)
}

func (context *GeosContext) cGEOSSimplify(g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return cGEOSSimplify(context, g1, tolerance)
}

func cGEOSSimplify(context *GeosContext, g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSimplify_r(context.handle, g1, tolerance)
}

func (context *GeosContext) cGEOSTopologyPreserveSimplify(g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return cGEOSTopologyPreserveSimplify(context, g1, tolerance)
}

func cGEOSTopologyPreserveSimplify(context *GeosContext, g1 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSTopologyPreserveSimplify_r(context.handle, g1, tolerance)
}

func (context *GeosContext) cGEOSGeom_extractUniquePoints(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGeom_extractUniquePoints(context, g)
}

func cGEOSGeom_extractUniquePoints(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_extractUniquePoints_r(context.handle, g)
}

func (context *GeosContext) cGEOSSharedPaths(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSSharedPaths(context, g1, g2)
}

func cGEOSSharedPaths(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSharedPaths_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSSnap(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	return cGEOSSnap(context, g1, g2, tolerance)
}

func cGEOSSnap(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSnap_r(context.handle, g1, g2, tolerance)
}

func (context *GeosContext) cGEOSDisjoint(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSDisjoint(context, g1, g2)
}

func cGEOSDisjoint(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSDisjoint_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSTouches(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSTouches(context, g1, g2)
}

func cGEOSTouches(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSTouches_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSIntersects(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSIntersects(context, g1, g2)
}

func cGEOSIntersects(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSIntersects_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSCrosses(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSCrosses(context, g1, g2)
}

func cGEOSCrosses(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCrosses_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSWithin(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSWithin(context, g1, g2)
}

func cGEOSWithin(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWithin_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSContains(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSContains(context, g1, g2)
}

func cGEOSContains(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSContains_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSOverlaps(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSOverlaps(context, g1, g2)
}

func cGEOSOverlaps(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSOverlaps_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSEquals(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSEquals(context, g1, g2)
}

func cGEOSEquals(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSEquals_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSEqualsExact(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) C.char {
	return cGEOSEqualsExact(context, g1, g2, tolerance)
}

func cGEOSEqualsExact(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, tolerance C.double) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSEqualsExact_r(context.handle, g1, g2, tolerance)
}

func (context *GeosContext) cGEOSCovers(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSCovers(context, g1, g2)
}

func cGEOSCovers(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCovers_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSCoveredBy(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSCoveredBy(context, g1, g2)
}

func cGEOSCoveredBy(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSCoveredBy_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSPrepare(g *C.GEOSGeometry) *C.GEOSPreparedGeometry {
	return cGEOSPrepare(context, g)
}

func cGEOSPrepare(context *GeosContext, g *C.GEOSGeometry) *C.GEOSPreparedGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPrepare_r(context.handle, g)
}

func (context *GeosContext) cGEOSPreparedGeom_destroy(g *C.GEOSPreparedGeometry) {
	 cGEOSPreparedGeom_destroy(context, g)
}

func cGEOSPreparedGeom_destroy(context *GeosContext, g *C.GEOSPreparedGeometry) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSPreparedGeom_destroy_r(context.handle, g)
}

func (context *GeosContext) cGEOSPreparedContains(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedContains(context, pg1, g2)
}

func cGEOSPreparedContains(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedContains_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedContainsProperly(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedContainsProperly(context, pg1, g2)
}

func cGEOSPreparedContainsProperly(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedContainsProperly_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedCoveredBy(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedCoveredBy(context, pg1, g2)
}

func cGEOSPreparedCoveredBy(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedCoveredBy_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedCovers(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedCovers(context, pg1, g2)
}

func cGEOSPreparedCovers(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedCovers_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedCrosses(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedCrosses(context, pg1, g2)
}

func cGEOSPreparedCrosses(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedCrosses_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedDisjoint(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedDisjoint(context, pg1, g2)
}

func cGEOSPreparedDisjoint(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedDisjoint_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedIntersects(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedIntersects(context, pg1, g2)
}

func cGEOSPreparedIntersects(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedIntersects_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedOverlaps(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedOverlaps(context, pg1, g2)
}

func cGEOSPreparedOverlaps(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedOverlaps_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedTouches(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedTouches(context, pg1, g2)
}

func cGEOSPreparedTouches(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedTouches_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSPreparedWithin(pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	return cGEOSPreparedWithin(context, pg1, g2)
}

func cGEOSPreparedWithin(context *GeosContext, pg1 *C.GEOSPreparedGeometry, g2 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSPreparedWithin_r(context.handle, pg1, g2)
}

func (context *GeosContext) cGEOSSTRtree_create(nodeCapacity C.size_t) *C.GEOSSTRtree {
	return cGEOSSTRtree_create(context, nodeCapacity)
}

func cGEOSSTRtree_create(context *GeosContext, nodeCapacity C.size_t) *C.GEOSSTRtree {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSTRtree_create_r(context.handle, nodeCapacity)
}

func (context *GeosContext) cGEOSSTRtree_insert(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) {
	cGEOSSTRtree_insert(context, tree, g, item)
}

func cGEOSSTRtree_insert(context *GeosContext, tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSSTRtree_insert_r(context.handle, tree, g, unsafe.Pointer(item))
}

func (context *GeosContext) cGEOSSTRtree_query(tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
	cGEOSSTRtree_query(context, tree, g, callback, userdata)
}

func cGEOSSTRtree_query(context *GeosContext, tree *C.GEOSSTRtree, g *C.GEOSGeometry, callback C.GEOSQueryCallback, userdata *C.void) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSSTRtree_query_r(context.handle, tree, g, callback, unsafe.Pointer(userdata))
}

func (context *GeosContext) cGEOSSTRtree_iterate(tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
	cGEOSSTRtree_iterate(context, tree, callback, userdata)
}

func cGEOSSTRtree_iterate(context *GeosContext, tree *C.GEOSSTRtree, callback C.GEOSQueryCallback, userdata *C.void) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSSTRtree_iterate_r(context.handle, tree, callback, unsafe.Pointer(userdata))
}

func (context *GeosContext) cGEOSSTRtree_remove(tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
	return cGEOSSTRtree_remove(context, tree, g, item)
}

func cGEOSSTRtree_remove(context *GeosContext, tree *C.GEOSSTRtree, g *C.GEOSGeometry, item *C.void) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSSTRtree_remove_r(context.handle, tree, g, unsafe.Pointer(item))
}

func (context *GeosContext) cGEOSSTRtree_destroy(tree *C.GEOSSTRtree) {
	cGEOSSTRtree_destroy(context, tree)
}

func cGEOSSTRtree_destroy(context *GeosContext, tree *C.GEOSSTRtree) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSSTRtree_destroy_r(context.handle, tree)
}

func (context *GeosContext) cGEOSisEmpty(g1 *C.GEOSGeometry) C.char {
	return cGEOSisEmpty(context, g1)
}

func cGEOSisEmpty(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisEmpty_r(context.handle, g1)
}

func (context *GeosContext) cGEOSisSimple(g1 *C.GEOSGeometry) C.char {
	return cGEOSisSimple(context, g1)
}

func cGEOSisSimple(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisSimple_r(context.handle, g1)
}

func (context *GeosContext) cGEOSisRing(g1 *C.GEOSGeometry) C.char {
	return cGEOSisRing(context, g1)
}

func cGEOSisRing(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisRing_r(context.handle, g1)
}

func (context *GeosContext) cGEOSHasZ(g1 *C.GEOSGeometry) C.char {
	return cGEOSHasZ(context, g1)
}

func cGEOSHasZ(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSHasZ_r(context.handle, g1)
}

func (context *GeosContext) cGEOSisClosed(g1 *C.GEOSGeometry) C.char {
	return cGEOSisClosed(context, g1)
}

func cGEOSisClosed(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisClosed_r(context.handle, g1)
}

func (context *GeosContext) cGEOSRelatePattern(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, pat *C.char) C.char {
	return cGEOSRelatePattern(context, g1, g2, pat)
}

func cGEOSRelatePattern(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, pat *C.char) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSRelatePattern_r(context.handle, g1, g2, pat)
}

func (context *GeosContext) cGEOSRelate(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.char {
	return cGEOSRelate(context, g1, g2)
}

func cGEOSRelate(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSRelate_r(context.handle, g1, g2)
}

func (context *GeosContext) cGEOSRelatePatternMatch(mat *C.char, pat *C.char) C.char {
	return cGEOSRelatePatternMatch(context, mat, pat)
}

func cGEOSRelatePatternMatch(context *GeosContext, mat *C.char, pat *C.char) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSRelatePatternMatch_r(context.handle, mat, pat)
}

func (context *GeosContext) cGEOSRelateBoundaryNodeRule(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, bnr C.int) *C.char {
	return cGEOSRelateBoundaryNodeRule(context, g1, g2, bnr)
}

func cGEOSRelateBoundaryNodeRule(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, bnr C.int) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSRelateBoundaryNodeRule_r(context.handle, g1, g2, bnr)
}

func (context *GeosContext) cGEOSisValid(g1 *C.GEOSGeometry) C.char {
	return cGEOSisValid(context, g1)
}

func cGEOSisValid(context *GeosContext, g1 *C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisValid_r(context.handle, g1)
}

func (context *GeosContext) cGEOSisValidReason(g1 *C.GEOSGeometry) *C.char {
	return cGEOSisValidReason(context, g1)
}

func cGEOSisValidReason(context *GeosContext, g1 *C.GEOSGeometry) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisValidReason_r(context.handle, g1)
}

func (context *GeosContext) cGEOSisValidDetail(g *C.GEOSGeometry, flags C.int, reason **C.char, location **C.GEOSGeometry) C.char {
	return cGEOSisValidDetail(context, g, flags, reason, location)
}

func cGEOSisValidDetail(context *GeosContext, g *C.GEOSGeometry, flags C.int, reason **C.char, location **C.GEOSGeometry) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSisValidDetail_r(context.handle, g, flags, reason, location)
}

func (context *GeosContext) cGEOSGeomType(g1 *C.GEOSGeometry) *C.char {
	return cGEOSGeomType(context, g1)
}

func cGEOSGeomType(context *GeosContext, g1 *C.GEOSGeometry) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomType_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGeomTypeId(g1 *C.GEOSGeometry) C.int {
	return cGEOSGeomTypeId(context, g1)
}

func cGEOSGeomTypeId(context *GeosContext, g1 *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomTypeId_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGetSRID(g *C.GEOSGeometry) C.int {
	return cGEOSGetSRID(context, g)
}

func cGEOSGetSRID(context *GeosContext, g *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetSRID_r(context.handle, g)
}

func (context *GeosContext) cGEOSSetSRID(g *C.GEOSGeometry, SRID C.int) {
	cGEOSSetSRID(context, g, SRID)
}

func cGEOSSetSRID(context *GeosContext, g *C.GEOSGeometry, SRID C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSSetSRID_r(context.handle, g, SRID)
}

func (context *GeosContext) cGEOSGetNumGeometries(g *C.GEOSGeometry) C.int {
	return cGEOSGetNumGeometries(context, g)
}

func cGEOSGetNumGeometries(context *GeosContext, g *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetNumGeometries_r(context.handle, g)
}

func (context *GeosContext) cGEOSGetGeometryN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return cGEOSGetGeometryN(context, g, n)
}

func cGEOSGetGeometryN(context *GeosContext, g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetGeometryN_r(context.handle, g, n)
}

func (context *GeosContext) cGEOSNormalize(g1 *C.GEOSGeometry) C.int {
	return cGEOSNormalize(context, g1)
}

func cGEOSNormalize(context *GeosContext, g1 *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSNormalize_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGetNumInteriorRings(g1 *C.GEOSGeometry) C.int {
	return cGEOSGetNumInteriorRings(context, g1)
}

func cGEOSGetNumInteriorRings(context *GeosContext, g1 *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetNumInteriorRings_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGeomGetNumPoints(g1 *C.GEOSGeometry) C.int {
	return cGEOSGeomGetNumPoints(context, g1)
}

func cGEOSGeomGetNumPoints(context *GeosContext, g1 *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetNumPoints_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGeomGetX(g1 *C.GEOSGeometry, x *C.double) C.int {
	return cGEOSGeomGetX(context, g1, x)
}

func cGEOSGeomGetX(context *GeosContext, g1 *C.GEOSGeometry, x *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetX_r(context.handle, g1, x)
}

func (context *GeosContext) cGEOSGeomGetY(g1 *C.GEOSGeometry, y *C.double) C.int {
	return cGEOSGeomGetY(context, g1, y)
}

func cGEOSGeomGetY(context *GeosContext, g1 *C.GEOSGeometry, y *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetY_r(context.handle, g1, y)
}

func (context *GeosContext) cGEOSGetInteriorRingN(g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return cGEOSGetInteriorRingN(context, g, n)
}

func cGEOSGetInteriorRingN(context *GeosContext, g *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetInteriorRingN_r(context.handle, g, n)
}

func (context *GeosContext) cGEOSGetExteriorRing(g *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGetExteriorRing(context, g)
}

func cGEOSGetExteriorRing(context *GeosContext, g *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetExteriorRing_r(context.handle, g)
}

func (context *GeosContext) cGEOSGetNumCoordinates(g1 *C.GEOSGeometry) C.int {
	return cGEOSGetNumCoordinates(context, g1)
}

func cGEOSGetNumCoordinates(context *GeosContext, g1 *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGetNumCoordinates_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGeom_getCoordSeq(g *C.GEOSGeometry) *C.GEOSCoordSequence {
	return cGEOSGeom_getCoordSeq(context, g)
}

func cGEOSGeom_getCoordSeq(context *GeosContext, g *C.GEOSGeometry) *C.GEOSCoordSequence {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_getCoordSeq_r(context.handle, g)
}

func (context *GeosContext) cGEOSGeom_getDimensions(g *C.GEOSGeometry) C.int {
	return cGEOSGeom_getDimensions(context, g)
}

func cGEOSGeom_getDimensions(context *GeosContext, g *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_getDimensions_r(context.handle, g)
}

func (context *GeosContext) cGEOSGeom_getCoordinateDimension(g *C.GEOSGeometry) C.int {
	return cGEOSGeom_getCoordinateDimension(context, g)
}

func cGEOSGeom_getCoordinateDimension(context *GeosContext, g *C.GEOSGeometry) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeom_getCoordinateDimension_r(context.handle, g)
}

func (context *GeosContext) cGEOSGeomGetPointN(g1 *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	return cGEOSGeomGetPointN(context, g1, n)
}

func cGEOSGeomGetPointN(context *GeosContext, g1 *C.GEOSGeometry, n C.int) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetPointN_r(context.handle, g1, n)
}

func (context *GeosContext) cGEOSGeomGetStartPoint(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGeomGetStartPoint(context, g1)
}

func cGEOSGeomGetStartPoint(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetStartPoint_r(context.handle, g1)
}

func (context *GeosContext) cGEOSGeomGetEndPoint(g1 *C.GEOSGeometry) *C.GEOSGeometry {
	return cGEOSGeomGetEndPoint(context, g1)
}

func cGEOSGeomGetEndPoint(context *GeosContext, g1 *C.GEOSGeometry) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetEndPoint_r(context.handle, g1)
}

func (context *GeosContext) cGEOSArea(g1 *C.GEOSGeometry, area *C.double) C.int {
	return cGEOSArea(context, g1, area)
}

func cGEOSArea(context *GeosContext, g1 *C.GEOSGeometry, area *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSArea_r(context.handle, g1, area)
}

func (context *GeosContext) cGEOSLength(g1 *C.GEOSGeometry, length *C.double) C.int {
	return cGEOSLength(context, g1, length)
}

func cGEOSLength(context *GeosContext, g1 *C.GEOSGeometry, length *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSLength_r(context.handle, g1, length)
}

func (context *GeosContext) cGEOSDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	return cGEOSDistance(context, g1, g2, dist)
}

func cGEOSDistance(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSDistance_r(context.handle, g1, g2, dist)
}

func (context *GeosContext) cGEOSHausdorffDistance(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	return cGEOSHausdorffDistance(context, g1, g2, dist)
}

func cGEOSHausdorffDistance(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, dist *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSHausdorffDistance_r(context.handle, g1, g2, dist)
}

func (context *GeosContext) cGEOSHausdorffDistanceDensify(g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	return cGEOSHausdorffDistanceDensify(context, g1, g2, densifyFrac, dist)
}

func cGEOSHausdorffDistanceDensify(context *GeosContext, g1 *C.GEOSGeometry, g2 *C.GEOSGeometry, densifyFrac C.double, dist *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSHausdorffDistanceDensify_r(context.handle, g1, g2, densifyFrac, dist)
}

func (context *GeosContext) cGEOSGeomGetLength(g1 *C.GEOSGeometry, length *C.double) C.int {
	return cGEOSGeomGetLength(context, g1, length)
}

func cGEOSGeomGetLength(context *GeosContext, g1 *C.GEOSGeometry, length *C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSGeomGetLength_r(context.handle, g1, length)
}

func (context *GeosContext) cGEOSOrientationIndex(Ax C.double, Ay C.double, Bx C.double, By C.double, Px C.double, Py C.double) C.int {
	return cGEOSOrientationIndex(context, Ax, Ay, Bx, By, Px, Py)
}

func cGEOSOrientationIndex(context *GeosContext, Ax C.double, Ay C.double, Bx C.double, By C.double, Px C.double, Py C.double) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSOrientationIndex_r(context.handle, Ax, Ay, Bx, By, Px, Py)
}

func (context *GeosContext) cGEOSWKTReader_create() *C.GEOSWKTReader {
	return cGEOSWKTReader_create(context, )
}

func cGEOSWKTReader_create(context *GeosContext, ) *C.GEOSWKTReader {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKTReader_create_r(context.handle)
}

func (context *GeosContext) cGEOSWKTReader_destroy(reader *C.GEOSWKTReader) {
	cGEOSWKTReader_destroy(context, reader)
}

func cGEOSWKTReader_destroy(context *GeosContext, reader *C.GEOSWKTReader) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTReader_destroy_r(context.handle, reader)
}

func (context *GeosContext) cGEOSWKTReader_read(reader *C.GEOSWKTReader, wkt *C.char) *C.GEOSGeometry {
	return cGEOSWKTReader_read(context, reader, wkt)
}

func cGEOSWKTReader_read(context *GeosContext, reader *C.GEOSWKTReader, wkt *C.char) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKTReader_read_r(context.handle, reader, wkt)
}

func (context *GeosContext) cGEOSWKTWriter_create() *C.GEOSWKTWriter {
	return cGEOSWKTWriter_create(context, )
}

func cGEOSWKTWriter_create(context *GeosContext, ) *C.GEOSWKTWriter {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKTWriter_create_r(context.handle)
}

func (context *GeosContext) cGEOSWKTWriter_destroy(writer *C.GEOSWKTWriter) {
	cGEOSWKTWriter_destroy(context, writer)
}

func cGEOSWKTWriter_destroy(context *GeosContext, writer *C.GEOSWKTWriter) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTWriter_destroy_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKTWriter_write(reader *C.GEOSWKTWriter, g *C.GEOSGeometry) *C.char {
	return cGEOSWKTWriter_write(context, reader, g)
}

func cGEOSWKTWriter_write(context *GeosContext, reader *C.GEOSWKTWriter, g *C.GEOSGeometry) *C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKTWriter_write_r(context.handle, reader, g)
}

func (context *GeosContext) cGEOSWKTWriter_setTrim(writer *C.GEOSWKTWriter, trim C.char) {
	cGEOSWKTWriter_setTrim(context, writer, trim)
}

func cGEOSWKTWriter_setTrim(context *GeosContext, writer *C.GEOSWKTWriter, trim C.char) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTWriter_setTrim_r(context.handle, writer, trim)
}

func (context *GeosContext) cGEOSWKTWriter_setRoundingPrecision(writer *C.GEOSWKTWriter, precision C.int) {
	cGEOSWKTWriter_setRoundingPrecision(context, writer, precision)
}

func cGEOSWKTWriter_setRoundingPrecision(context *GeosContext, writer *C.GEOSWKTWriter, precision C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTWriter_setRoundingPrecision_r(context.handle, writer, precision)
}

func (context *GeosContext) cGEOSWKTWriter_setOutputDimension(writer *C.GEOSWKTWriter, dim C.int) {
	cGEOSWKTWriter_setOutputDimension(context, writer, dim)
}

func cGEOSWKTWriter_setOutputDimension(context *GeosContext, writer *C.GEOSWKTWriter, dim C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTWriter_setOutputDimension_r(context.handle, writer, dim)
}

func (context *GeosContext) cGEOSWKTWriter_getOutputDimension(writer *C.GEOSWKTWriter) C.int {
	return cGEOSWKTWriter_getOutputDimension(context, writer)
}

func cGEOSWKTWriter_getOutputDimension(context *GeosContext, writer *C.GEOSWKTWriter) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKTWriter_getOutputDimension_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKTWriter_setOld3D(writer *C.GEOSWKTWriter, useOld3D C.int) {
	cGEOSWKTWriter_setOld3D(context, writer, useOld3D)
}

func cGEOSWKTWriter_setOld3D(context *GeosContext, writer *C.GEOSWKTWriter, useOld3D C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKTWriter_setOld3D_r(context.handle, writer, useOld3D)
}

func (context *GeosContext) cGEOSWKBReader_create() *C.GEOSWKBReader {
	return cGEOSWKBReader_create(context)
}

func cGEOSWKBReader_create(context *GeosContext, ) *C.GEOSWKBReader {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBReader_create_r(context.handle)
}

func (context *GeosContext) cGEOSWKBReader_destroy(reader *C.GEOSWKBReader) {
	cGEOSWKBReader_destroy(context, reader)
}

func cGEOSWKBReader_destroy(context *GeosContext, reader *C.GEOSWKBReader) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKBReader_destroy_r(context.handle, reader)
}

func (context *GeosContext) cGEOSWKBReader_read(reader *C.GEOSWKBReader, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	return cGEOSWKBReader_read(context, reader, wkb, size)
}

func cGEOSWKBReader_read(context *GeosContext, reader *C.GEOSWKBReader, wkb *C.uchar, size C.size_t) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBReader_read_r(context.handle, reader, wkb, size)
}

func (context *GeosContext) cGEOSWKBReader_readHEX(reader *C.GEOSWKBReader, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	return cGEOSWKBReader_readHEX(context, reader, hex, size)
}

func cGEOSWKBReader_readHEX(context *GeosContext, reader *C.GEOSWKBReader, hex *C.uchar, size C.size_t) *C.GEOSGeometry {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBReader_readHEX_r(context.handle, reader, hex, size)
}

func (context *GeosContext) cGEOSWKBWriter_create() *C.GEOSWKBWriter {
	return cGEOSWKBWriter_create(context, )
}

func cGEOSWKBWriter_create(context *GeosContext, ) *C.GEOSWKBWriter {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_create_r(context.handle)
}

func (context *GeosContext) cGEOSWKBWriter_destroy(writer *C.GEOSWKBWriter) {
	cGEOSWKBWriter_destroy(context, writer)
}

func cGEOSWKBWriter_destroy(context *GeosContext, writer *C.GEOSWKBWriter) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKBWriter_destroy_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKBWriter_write(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return cGEOSWKBWriter_write(context, writer, g, size)
}

func cGEOSWKBWriter_write(context *GeosContext, writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_write_r(context.handle, writer, g, size)
}

func (context *GeosContext) cGEOSWKBWriter_writeHEX(writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	return cGEOSWKBWriter_writeHEX(context, writer, g, size)
}

func cGEOSWKBWriter_writeHEX(context *GeosContext, writer *C.GEOSWKBWriter, g *C.GEOSGeometry, size *C.size_t) *C.uchar {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_writeHEX_r(context.handle, writer, g, size)
}

func (context *GeosContext) cGEOSWKBWriter_getOutputDimension(writer *C.GEOSWKBWriter) C.int {
	return cGEOSWKBWriter_getOutputDimension(context, writer)
}

func cGEOSWKBWriter_getOutputDimension(context *GeosContext, writer *C.GEOSWKBWriter) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_getOutputDimension_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKBWriter_setOutputDimension(writer *C.GEOSWKBWriter, newDimension C.int) {
	cGEOSWKBWriter_setOutputDimension(context, writer, newDimension)
}

func cGEOSWKBWriter_setOutputDimension(context *GeosContext, writer *C.GEOSWKBWriter, newDimension C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKBWriter_setOutputDimension_r(context.handle, writer, newDimension)
}

func (context *GeosContext) cGEOSWKBWriter_getByteOrder(writer *C.GEOSWKBWriter) C.int {
	return cGEOSWKBWriter_getByteOrder(context, writer)
}

func cGEOSWKBWriter_getByteOrder(context *GeosContext, writer *C.GEOSWKBWriter) C.int {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_getByteOrder_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKBWriter_setByteOrder(writer *C.GEOSWKBWriter, byteOrder C.int) {
	cGEOSWKBWriter_setByteOrder(context, writer, byteOrder)
}

func cGEOSWKBWriter_setByteOrder(context *GeosContext, writer *C.GEOSWKBWriter, byteOrder C.int) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKBWriter_setByteOrder_r(context.handle, writer, byteOrder)
}

func (context *GeosContext) cGEOSWKBWriter_getIncludeSRID(writer *C.GEOSWKBWriter) C.char {
	return cGEOSWKBWriter_getIncludeSRID(context, writer)
}

func cGEOSWKBWriter_getIncludeSRID(context *GeosContext, writer *C.GEOSWKBWriter) C.char {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	return C.GEOSWKBWriter_getIncludeSRID_r(context.handle, writer)
}

func (context *GeosContext) cGEOSWKBWriter_setIncludeSRID(writer *C.GEOSWKBWriter, writeSRID C.char) {
	cGEOSWKBWriter_setIncludeSRID(context, writer, writeSRID)
}

func cGEOSWKBWriter_setIncludeSRID(context *GeosContext, writer *C.GEOSWKBWriter, writeSRID C.char) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSWKBWriter_setIncludeSRID_r(context.handle, writer, writeSRID)
}

func (context *GeosContext) cGEOSFree(buffer *C.void) {
	cGEOSFree(context, buffer)
}

func cGEOSFree(context *GeosContext, buffer *C.void) {
	context.handlemu.Lock()
	defer context.handlemu.Unlock()
	C.GEOSFree_r(context.handle, unsafe.Pointer(buffer))
}
