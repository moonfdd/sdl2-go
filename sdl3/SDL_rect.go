package sdl3

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/**
 *  \file SDL_rect.h
 *
 *  Header file for SDL_rect definition and management functions.
 */

// #ifndef SDL_rect_h_
// #define SDL_rect_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>
// #include <SDL3/SDL_pixels.h>
// #include <SDL3/SDL_rwops.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 * The structure that defines a point (integer)
 *
 * \sa SDL_GetRectEnclosingPoints
 * \sa SDL_PointInRect
 */
type SDL_Point struct {
	X ffcommon.FInt
	Y ffcommon.FInt
}

/**
 * The structure that defines a point (floating point)
 *
 * \sa SDL_GetRectEnclosingPointsFloat
 * \sa SDL_PointInRectFloat
 */
type SDL_FPoint struct {
	X ffcommon.FFloat
	Y ffcommon.FFloat
}

/**
 * A rectangle, with the origin at the upper left (integer).
 *
 * \sa SDL_RectEmpty
 * \sa SDL_RectsEqual
 * \sa SDL_HasRectIntersection
 * \sa SDL_GetRectIntersection
 * \sa SDL_GetRectAndLineIntersection
 * \sa SDL_GetRectUnion
 * \sa SDL_GetRectEnclosingPoints
 */
type SDL_Rect struct {
	X, Y ffcommon.FInt
	W, H ffcommon.FInt
}

/**
 * A rectangle, with the origin at the upper left (floating point).
 *
 * \sa SDL_RectEmptyFloat
 * \sa SDL_RectsEqualFloat
 * \sa SDL_RectsEqualEpsilon
 * \sa SDL_HasRectIntersectionFloat
 * \sa SDL_GetRectIntersectionFloat
 * \sa SDL_GetRectAndLineIntersectionFloat
 * \sa SDL_GetRectUnionFloat
 * \sa SDL_GetRectEnclosingPointsFloat
 * \sa SDL_PointInRectFloat
 */
type SDL_FRect struct {
	X ffcommon.FFloat
	Y ffcommon.FFloat
	W ffcommon.FFloat
	H ffcommon.FFloat
}

/**
 * Returns true if point resides inside a rectangle.
 */
// SDL_FORCE_INLINE SDL_bool SDL_PointInRect(const SDL_Point *p, const SDL_Rect *r)
// {
//     return ( (p->x >= r->x) && (p->x < (r->x + r->w)) &&
//              (p->y >= r->y) && (p->y < (r->y + r->h)) ) ? SDL_TRUE : SDL_FALSE;
// }
func SDL_PointInRect(p *SDL_Point, r *SDL_Rect) (res bool) {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		res = true
	}
	return
}

/**
 * Returns true if the rectangle has no area.
 */
// SDL_FORCE_INLINE SDL_bool SDL_RectEmpty(const SDL_Rect *r)
// {
//     return ((!r) || (r->w <= 0) || (r->h <= 0)) ? SDL_TRUE : SDL_FALSE;
// }
func SDL_RectEmpty(r *SDL_Rect) (res bool) {
	if (r == nil) || (r.W <= 0) || (r.H <= 0) {
		res = true
	}
	return
}

/**
 * Returns true if the two rectangles are equal.
 */
// SDL_FORCE_INLINE SDL_bool SDL_RectsEqual(const SDL_Rect *a, const SDL_Rect *b)
// {
//     return (a && b && (a->x == b->x) && (a->y == b->y) &&
//             (a->w == b->w) && (a->h == b->h)) ? SDL_TRUE : SDL_FALSE;
// }
func SDL_RectsEqual(a, b *SDL_Rect) (res bool) {
	if a != nil && b != nil && (a.X == b.X) && (a.Y == b.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		res = true
	}
	return
}

/**
 * Determine whether two rectangles intersect.
 *
 * If either pointer is NULL the function will return SDL_FALSE.
 *
 * \param A an SDL_Rect structure representing the first rectangle
 * \param B an SDL_Rect structure representing the second rectangle
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetRectIntersection
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_HasRectIntersection(const SDL_Rect * A,
//                                                      const SDL_Rect * B);
func SDL_HasRectIntersection(A, B *SDL_Rect) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasRectIntersection").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
	)
	res = ffcommon.GoBool(t)
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
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_HasRectIntersection
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectIntersection(const SDL_Rect * A,
//                                                    const SDL_Rect * B,
//                                                    SDL_Rect * result);
func SDL_GetRectIntersection(A, B, result *SDL_Rect) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectIntersection").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Calculate the union of two rectangles.
 *
 * \param A an SDL_Rect structure representing the first rectangle
 * \param B an SDL_Rect structure representing the second rectangle
 * \param result an SDL_Rect structure filled in with the union of rectangles
 *               `A` and `B`
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_GetRectUnion(const SDL_Rect * A,
//                                            const SDL_Rect * B,
//                                            SDL_Rect * result);
func SDL_GetRectUnion(A, B, result *SDL_Rect) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectUnion").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.FInt(t)
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
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectEnclosingPoints(const SDL_Point * points,
//                                                    int count,
//                                                    const SDL_Rect * clip,
//                                                    SDL_Rect * result);
func (points *SDL_Point) SDL_GetRectEnclosingPoints(
	count ffcommon.FInt,
	clip *SDL_Rect,
	result *SDL_Rect,
) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectEnclosingPoints").Call(
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(unsafe.Pointer(clip)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.GoBool(t)
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
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectAndLineIntersection(const SDL_Rect *
//                                                           rect, int *X1,
//                                                           int *Y1, int *X2,
//                                                           int *Y2);

func (rect *SDL_Rect) SDL_GetRectAndLineIntersection(X1, Y1, X2, Y2 *ffcommon.FInt) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectAndLineIntersection").Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(X1)),
		uintptr(unsafe.Pointer(Y1)),
		uintptr(unsafe.Pointer(X2)),
		uintptr(unsafe.Pointer(Y2)),
	)
	res = ffcommon.GoBool(t)
	return
}

/* SDL_FRect versions... */

/**
 * Returns true if point resides inside a rectangle.
 */
// SDL_FORCE_INLINE SDL_bool SDL_PointInRectFloat(const SDL_FPoint *p, const SDL_FRect *r)
// {
//     return ( (p->x >= r->x) && (p->x < (r->x + r->w)) &&
//              (p->y >= r->y) && (p->y < (r->y + r->h)) ) ? SDL_TRUE : SDL_FALSE;
// }
func SDL_PointInRectFloat(p, r *SDL_FRect) (res bool) {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		res = true
	}
	return
}

/**
 * Returns true if the rectangle has no area.
 */
// SDL_FORCE_INLINE SDL_bool SDL_RectEmptyFloat(const SDL_FRect *r)
// {
//     return ((!r) || (r->w <= 0.0f) || (r->h <= 0.0f)) ? SDL_TRUE : SDL_FALSE;
// }
func SDL_RectEmptyFloat(r *SDL_FRect) (res bool) {
	if r != nil || r.W <= 0.0 || r.H <= 0.0 {
		res = true
	}
	return
}

/**
 * Returns true if the two rectangles are equal, within some given epsilon.
 *
 * \since This function is available since SDL 2.0.22.
 */
// SDL_FORCE_INLINE SDL_bool SDL_RectsEqualEpsilon(const SDL_FRect *a, const SDL_FRect *b, const float epsilon)
// {
//     return (a && b && ((a == b) ||
//             ((SDL_fabsf(a->x - b->x) <= epsilon) &&
//             (SDL_fabsf(a->y - b->y) <= epsilon) &&
//             (SDL_fabsf(a->w - b->w) <= epsilon) &&
//             (SDL_fabsf(a->h - b->h) <= epsilon))))
//             ? SDL_TRUE : SDL_FALSE;
// }
func SDL_RectsEqualEpsilon(a, b *SDL_FRect, epsilon ffcommon.FFloat) (res bool) {
	if a != nil && b != nil && ((a == b) ||
		((SDL_fabsf(a.X-b.X) <= epsilon) &&
			(SDL_fabsf(a.Y-b.Y) <= epsilon) &&
			(SDL_fabsf(a.W-b.W) <= epsilon) &&
			(SDL_fabsf(a.H-b.H) <= epsilon))) {
		res = true
	}
	return
}

func SDL_fabsf(a ffcommon.FFloat) (res ffcommon.FFloat) {
	if a < 0 {
		res = -a
	} else {
		res = a
	}
	return
}

/**
 * Returns true if the two rectangles are equal, using a default epsilon.
 *
 * \since This function is available since SDL 2.0.22.
 */
// SDL_FORCE_INLINE SDL_bool SDL_RectsEqualFloat(const SDL_FRect *a, const SDL_FRect *b)
// {
//     return SDL_RectsEqualEpsilon(a, b, SDL_FLT_EPSILON);
// }

func SDL_RectsEqualFloat(a, b *SDL_FRect) (res bool) {
	res = SDL_RectsEqualEpsilon(a, b, SDL_FLT_EPSILON)
	return

}

const SDL_FLT_EPSILON = 1.1920928955078125e-07

/**
 * Determine whether two rectangles intersect with float precision.
 *
 * If either pointer is NULL the function will return SDL_FALSE.
 *
 * \param A an SDL_FRect structure representing the first rectangle
 * \param B an SDL_FRect structure representing the second rectangle
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetRectIntersection
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_HasRectIntersectionFloat(const SDL_FRect * A,
//                                                       const SDL_FRect * B);
func SDL_HasRectIntersectionFloat(A, B *SDL_Rect) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasRectIntersectionFloat").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Calculate the intersection of two rectangles with float precision.
 *
 * If `result` is NULL then this function will return SDL_FALSE.
 *
 * \param A an SDL_FRect structure representing the first rectangle
 * \param B an SDL_FRect structure representing the second rectangle
 * \param result an SDL_FRect structure filled in with the intersection of
 *               rectangles `A` and `B`
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_HasRectIntersectionFloat
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectIntersectionFloat(const SDL_FRect * A,
//                                                     const SDL_FRect * B,
//                                                     SDL_FRect * result);
func SDL_GetRectIntersectionFloat(A, B, result *SDL_Rect) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectIntersectionFloat").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Calculate the union of two rectangles with float precision.
 *
 * \param A an SDL_FRect structure representing the first rectangle
 * \param B an SDL_FRect structure representing the second rectangle
 * \param result an SDL_FRect structure filled in with the union of rectangles
 *               `A` and `B`
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_GetRectUnionFloat(const SDL_FRect * A,
//                                             const SDL_FRect * B,
//                                             SDL_FRect * result);
func SDL_GetRectUnionFloat(A, B, result *SDL_Rect) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectUnionFloat").Call(
		uintptr(unsafe.Pointer(A)),
		uintptr(unsafe.Pointer(B)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Calculate a minimal rectangle enclosing a set of points with float
 * precision.
 *
 * If `clip` is not NULL then only points inside of the clipping rectangle are
 * considered.
 *
 * \param points an array of SDL_FPoint structures representing points to be
 *               enclosed
 * \param count the number of structures in the `points` array
 * \param clip an SDL_FRect used for clipping or NULL to enclose all points
 * \param result an SDL_FRect structure filled in with the minimal enclosing
 *               rectangle
 * \returns SDL_TRUE if any points were enclosed or SDL_FALSE if all the
 *          points were outside of the clipping rectangle.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectEnclosingPointsFloat(const SDL_FPoint * points,
//                                                     int count,
//                                                     const SDL_FRect * clip,
//                                                     SDL_FRect * result);
func SDL_GetRectEnclosingPointsFloat(points *SDL_FPoint, count ffcommon.FInt, clip, result *SDL_FRect) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectEnclosingPointsFloat").Call(
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
		uintptr(unsafe.Pointer(clip)),
		uintptr(unsafe.Pointer(result)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Calculate the intersection of a rectangle and line segment with float
 * precision.
 *
 * This function is used to clip a line segment to a rectangle. A line segment
 * contained entirely within the rectangle or that does not intersect will
 * remain unchanged. A line segment that crosses the rectangle at either or
 * both ends will be clipped to the boundary of the rectangle and the new
 * coordinates saved in `X1`, `Y1`, `X2`, and/or `Y2` as necessary.
 *
 * \param rect an SDL_FRect structure representing the rectangle to intersect
 * \param X1 a pointer to the starting X-coordinate of the line
 * \param Y1 a pointer to the starting Y-coordinate of the line
 * \param X2 a pointer to the ending X-coordinate of the line
 * \param Y2 a pointer to the ending Y-coordinate of the line
 * \returns SDL_TRUE if there is an intersection, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetRectAndLineIntersectionFloat(const SDL_FRect *
//                                                            rect, float *X1,
//                                                            float *Y1, float *X2,
//                                                            float *Y2);
func (rect *SDL_FRect) SDL_GetRectAndLineIntersectionFloat(X1, Y1, X2, Y2 *ffcommon.FFloat) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRectAndLineIntersectionFloat").Call(
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(X1)),
		uintptr(unsafe.Pointer(Y1)),
		uintptr(unsafe.Pointer(X2)),
		uintptr(unsafe.Pointer(Y2)),
	)
	res = ffcommon.GoBool(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_rect_h_ */
