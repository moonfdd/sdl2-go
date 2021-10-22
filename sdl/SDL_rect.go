package sdl

import (
	"sdl2-go/common"
	"unsafe"
)

/**
 * The structure that defines a point (integer)
 *
 * \sa SDL_EnclosePoints
 * \sa SDL_PointInRect
 */
type SDL_Point struct {
	x common.FInt
	y common.FInt
}

/**
 * The structure that defines a point (floating point)
 *
 * \sa SDL_EnclosePoints
 * \sa SDL_PointInRect
 */
type SDL_FPoint struct {
	X common.FFloat
	Y common.FFloat
}

/**
 * A rectangle, with the origin at the upper left (integer).
 *
 * \sa SDL_RectEmpty
 * \sa SDL_RectEquals
 * \sa SDL_HasIntersection
 * \sa SDL_IntersectRect
 * \sa SDL_UnionRect
 * \sa SDL_EnclosePoints
 */
type SDL_Rect struct {
	X, Y common.FInt
	W, H common.FInt
}

/**
 * A rectangle, with the origin at the upper left (floating point).
 */
type SDL_FRect struct {
	x common.FFloat
	y common.FFloat
	w common.FFloat
	h common.FFloat
}

///**
// * Returns true if point resides inside a rectangle.
// */
//SDL_FORCE_INLINE SDL_bool SDL_PointInRect(const SDL_Point *p, const SDL_Rect *r)
//{
//return ( (p->x >= r->x) && (p->x < (r->x + r->w)) &&
//(p->y >= r->y) && (p->y < (r->y + r->h)) ) ? SDL_TRUE : SDL_FALSE;
//}
//
///**
// * Returns true if the rectangle has no area.
// */
//SDL_FORCE_INLINE SDL_bool SDL_RectEmpty(const SDL_Rect *r)
//{
//return ((!r) || (r->w <= 0) || (r->h <= 0)) ? SDL_TRUE : SDL_FALSE;
//}
//
///**
// * Returns true if the two rectangles are equal.
// */
//SDL_FORCE_INLINE SDL_bool SDL_RectEquals(const SDL_Rect *a, const SDL_Rect *b)
//{
//return (a && b && (a->x == b->x) && (a->y == b->y) &&
//(a->w == b->w) && (a->h == b->h)) ? SDL_TRUE : SDL_FALSE;
//}

/**
 * Determine whether two rectangles intersect.
 *
 * If either pointer is NULL the function will return SDL_FALSE.
 *
 * \param A an SDL_Rect structure representing the first rectangle
 * \param B an SDL_Rect structure representing the second rectangle
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_IntersectRect
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasIntersection(const SDL_Rect * A,
//const SDL_Rect * B);
func SDL_HasIntersection(A, B *SDL_Rect) (res SDL_bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_HasIntersection").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
	)
	if t == 0 {

	}
	res = SDL_bool(t)
	return
}

/**
 * Calculate the intersection of two rectangles.
 *
 * If `result` is NULL then this function will return SDL_FALSE.
 *
 * \param A an SDL_Rect structure representing the first rectangle
 * \param B an SDL_Rect structure representing the second rectangle
 * \param result an SDL_Rect structure filled in with the intersection of
 *               rectangles `A` and `B`
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_HasIntersection
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IntersectRect(const SDL_Rect * A,
//const SDL_Rect * B,
//SDL_Rect * result);
func SDL_IntersectRect(A, B, result *SDL_Rect) (res SDL_bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_IntersectRect").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	if t == 0 {

	}
	res = SDL_bool(t)
	return
}

/**
 * Calculate the union of two rectangles.
 *
 * \param A an SDL_Rect structure representing the first rectangle
 * \param B an SDL_Rect structure representing the second rectangle
 * \param result an SDL_Rect structure filled in with the union of rectangles
 *               `A` and `B`
 */
//extern DECLSPEC void SDLCALL SDL_UnionRect(const SDL_Rect * A,
//const SDL_Rect * B,
//SDL_Rect * result);
func SDL_UnionRect(A, B, result *SDL_Rect) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UnionRect").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	if t == 0 {

	}
	return
}

/**
 * Calculate a minimal rectangle enclosing a set of points.
 *
 * If `clip` is not NULL then only points inside of the clipping rectangle are
 * considered.
 *
 * \param points an array of SDL_Point structures representing points to be
 *               enclosed
 * \param count the number of structures in the `points` array
 * \param clip an SDL_Rect used for clipping or NULL to enclose all points
 * \param result an SDL_Rect structure filled in with the minimal enclosing
 *               rectangle
 * \returns SDL_TRUE if any points were enclosed or SDL_FALSE if all the
 *          points were outside of the clipping rectangle.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_EnclosePoints(const SDL_Point * points,
//int count,
//const SDL_Rect * clip,
//SDL_Rect * result);
func SDL_EnclosePoints(points *SDL_Point, count common.FInt,
	clip *SDL_Rect,
	result *SDL_Rect) (res SDL_bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_EnclosePoints").Call(
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(unsafe.Pointer(clip)),
		uintptr(unsafe.Pointer(result)),
	)
	if t == 0 {

	}
	res = SDL_bool(t)
	return
}

/**
 * Calculate the intersection of a rectangle and line segment.
 *
 * This function is used to clip a line segment to a rectangle. A line segment
 * contained entirely within the rectangle or that does not intersect will
 * remain unchanged. A line segment that crosses the rectangle at either or
 * both ends will be clipped to the boundary of the rectangle and the new
 * coordinates saved in `X1`, `Y1`, `X2`, and/or `Y2` as necessary.
 *
 * \param rect an SDL_Rect structure representing the rectangle to intersect
 * \param X1 a pointer to the starting X-coordinate of the line
 * \param Y1 a pointer to the starting Y-coordinate of the line
 * \param X2 a pointer to the ending X-coordinate of the line
 * \param Y2 a pointer to the ending Y-coordinate of the line
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IntersectRectAndLine(const SDL_Rect *
//rect, int *X1,
//int *Y1, int *X2,
//int *Y2);
func (rect *SDL_Rect) SDL_IntersectRectAndLine(X1 *common.FInt,
	Y1 *common.FInt, X2 *common.FInt,
	Y2 *common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_IntersectRectAndLine").Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(X1)),
		uintptr(unsafe.Pointer(Y1)),
		uintptr(unsafe.Pointer(X2)),
		uintptr(unsafe.Pointer(Y2)),
	)
	if t == 0 {

	}
	return
}
