package sdl

//#ifndef __gl2_h_
//const __gl2_h_
//
///* $Revision: 20555 $ on $Date:: 2013-02-12 14:32:47 -0800 #$ */
//
///*#include <GLES2/gl2platform.h>*/
//
//#ifdef __cplusplus
//extern "C" {
//#endif
//
///*
// * This document is licensed under the SGI Free Software B License Version
// * 2.0. For details, see http://oss.sgi.com/projects/FreeB/ .
// */
//
///*-------------------------------------------------------------------------
// * Data type definitions
// *-----------------------------------------------------------------------*/

//typedef void             GLvoid;
//typedef char             GLchar;
//typedef unsigned int     GLenum;
//typedef unsigned char    GLboolean;
//typedef unsigned int     GLbitfield;
//typedef khronos_int8_t   GLbyte;
//typedef short            GLshort;
//typedef int              GLint;
//typedef int              GLsizei;
//typedef khronos_uint8_t  GLubyte;
//typedef unsigned short   GLushort;
//typedef unsigned int     GLuint;
//typedef khronos_float_t  GLfloat;
//typedef khronos_float_t  GLclampf;
//typedef khronos_int32_t  GLfixed;

/* GL types for handling large vertex buffer objects */
//typedef khronos_intptr_t GLintptr;
//typedef khronos_ssize_t  GLsizeiptr;

/* OpenGL ES core versions */
const GL_ES_VERSION_2_0 = 1

/* ClearBufferMask */
const GL_DEPTH_BUFFER_BIT = 0x00000100
const GL_STENCIL_BUFFER_BIT = 0x00000400
const GL_COLOR_BUFFER_BIT = 0x00004000

/* Boolean */
const GL_FALSE = 0
const GL_TRUE = 1

/* BeginMode */
const GL_POINTS = 0x0000
const GL_LINES = 0x0001
const GL_LINE_LOOP = 0x0002
const GL_LINE_STRIP = 0x0003
const GL_TRIANGLES = 0x0004
const GL_TRIANGLE_STRIP = 0x0005
const GL_TRIANGLE_FAN = 0x0006

/* AlphaFunction (not supported in ES20) */
/*      GL_NEVER */
/*      GL_LESS */
/*      GL_EQUAL */
/*      GL_LEQUAL */
/*      GL_GREATER */
/*      GL_NOTEQUAL */
/*      GL_GEQUAL */
/*      GL_ALWAYS */

/* BlendingFactorDest */
const GL_ZERO = 0
const GL_ONE = 1
const GL_SRC_COLOR = 0x0300
const GL_ONE_MINUS_SRC_COLOR = 0x0301
const GL_SRC_ALPHA = 0x0302
const GL_ONE_MINUS_SRC_ALPHA = 0x0303
const GL_DST_ALPHA = 0x0304
const GL_ONE_MINUS_DST_ALPHA = 0x0305

/* BlendingFactorSrc */
/*      GL_ZERO */
/*      GL_ONE */
const GL_DST_COLOR = 0x0306
const GL_ONE_MINUS_DST_COLOR = 0x0307
const GL_SRC_ALPHA_SATURATE = 0x0308

/*      GL_SRC_ALPHA */
/*      GL_ONE_MINUS_SRC_ALPHA */
/*      GL_DST_ALPHA */
/*      GL_ONE_MINUS_DST_ALPHA */

/* BlendEquationSeparate */
const GL_FUNC_ADD = 0x8006
const GL_BLEND_EQUATION = 0x8009
const GL_BLEND_EQUATION_RGB = 0x8009 /* same as BLEND_EQUATION */
const GL_BLEND_EQUATION_ALPHA = 0x883D

/* BlendSubtract */
const GL_FUNC_SUBTRACT = 0x800A
const GL_FUNC_REVERSE_SUBTRACT = 0x800B

/* Separate Blend Functions */
const GL_BLEND_DST_RGB = 0x80C8
const GL_BLEND_SRC_RGB = 0x80C9
const GL_BLEND_DST_ALPHA = 0x80CA
const GL_BLEND_SRC_ALPHA = 0x80CB
const GL_CONSTANT_COLOR = 0x8001
const GL_ONE_MINUS_CONSTANT_COLOR = 0x8002
const GL_CONSTANT_ALPHA = 0x8003
const GL_ONE_MINUS_CONSTANT_ALPHA = 0x8004
const GL_BLEND_COLOR = 0x8005

/* Buffer Objects */
const GL_ARRAY_BUFFER = 0x8892
const GL_ELEMENT_ARRAY_BUFFER = 0x8893
const GL_ARRAY_BUFFER_BINDING = 0x8894
const GL_ELEMENT_ARRAY_BUFFER_BINDING = 0x8895

const GL_STREAM_DRAW = 0x88E0
const GL_STATIC_DRAW = 0x88E4
const GL_DYNAMIC_DRAW = 0x88E8

const GL_BUFFER_SIZE = 0x8764
const GL_BUFFER_USAGE = 0x8765

const GL_CURRENT_VERTEX_ATTRIB = 0x8626

/* CullFaceMode */
const GL_FRONT = 0x0404
const GL_BACK = 0x0405
const GL_FRONT_AND_BACK = 0x0408

/* DepthFunction */
/*      GL_NEVER */
/*      GL_LESS */
/*      GL_EQUAL */
/*      GL_LEQUAL */
/*      GL_GREATER */
/*      GL_NOTEQUAL */
/*      GL_GEQUAL */
/*      GL_ALWAYS */

/* EnableCap */
const GL_TEXTURE_2D = 0x0DE1
const GL_CULL_FACE = 0x0B44
const GL_BLEND = 0x0BE2
const GL_DITHER = 0x0BD0
const GL_STENCIL_TEST = 0x0B90
const GL_DEPTH_TEST = 0x0B71
const GL_SCISSOR_TEST = 0x0C11
const GL_POLYGON_OFFSET_FILL = 0x8037
const GL_SAMPLE_ALPHA_TO_COVERAGE = 0x809E
const GL_SAMPLE_COVERAGE = 0x80A0

/* ErrorCode */
const GL_NO_ERROR = 0
const GL_INVALID_ENUM = 0x0500
const GL_INVALID_VALUE = 0x0501
const GL_INVALID_OPERATION = 0x0502
const GL_OUT_OF_MEMORY = 0x0505

/* FrontFaceDirection */
const GL_CW = 0x0900
const GL_CCW = 0x0901

/* GetPName */
const GL_LINE_WIDTH = 0x0B21
const GL_ALIASED_POINT_SIZE_RANGE = 0x846D
const GL_ALIASED_LINE_WIDTH_RANGE = 0x846E
const GL_CULL_FACE_MODE = 0x0B45
const GL_FRONT_FACE = 0x0B46
const GL_DEPTH_RANGE = 0x0B70
const GL_DEPTH_WRITEMASK = 0x0B72
const GL_DEPTH_CLEAR_VALUE = 0x0B73
const GL_DEPTH_FUNC = 0x0B74
const GL_STENCIL_CLEAR_VALUE = 0x0B91
const GL_STENCIL_FUNC = 0x0B92
const GL_STENCIL_FAIL = 0x0B94
const GL_STENCIL_PASS_DEPTH_FAIL = 0x0B95
const GL_STENCIL_PASS_DEPTH_PASS = 0x0B96
const GL_STENCIL_REF = 0x0B97
const GL_STENCIL_VALUE_MASK = 0x0B93
const GL_STENCIL_WRITEMASK = 0x0B98
const GL_STENCIL_BACK_FUNC = 0x8800
const GL_STENCIL_BACK_FAIL = 0x8801
const GL_STENCIL_BACK_PASS_DEPTH_FAIL = 0x8802
const GL_STENCIL_BACK_PASS_DEPTH_PASS = 0x8803
const GL_STENCIL_BACK_REF = 0x8CA3
const GL_STENCIL_BACK_VALUE_MASK = 0x8CA4
const GL_STENCIL_BACK_WRITEMASK = 0x8CA5
const GL_VIEWPORT = 0x0BA2
const GL_SCISSOR_BOX = 0x0C10

/*      GL_SCISSOR_TEST */
const GL_COLOR_CLEAR_VALUE = 0x0C22
const GL_COLOR_WRITEMASK = 0x0C23
const GL_UNPACK_ALIGNMENT = 0x0CF5
const GL_PACK_ALIGNMENT = 0x0D05
const GL_MAX_TEXTURE_SIZE = 0x0D33
const GL_MAX_VIEWPORT_DIMS = 0x0D3A
const GL_SUBPIXEL_BITS = 0x0D50
const GL_RED_BITS = 0x0D52
const GL_GREEN_BITS = 0x0D53
const GL_BLUE_BITS = 0x0D54
const GL_ALPHA_BITS = 0x0D55
const GL_DEPTH_BITS = 0x0D56
const GL_STENCIL_BITS = 0x0D57
const GL_POLYGON_OFFSET_UNITS = 0x2A00

/*      GL_POLYGON_OFFSET_FILL */
const GL_POLYGON_OFFSET_FACTOR = 0x8038
const GL_TEXTURE_BINDING_2D = 0x8069
const GL_SAMPLE_BUFFERS = 0x80A8
const GL_SAMPLES = 0x80A9
const GL_SAMPLE_COVERAGE_VALUE = 0x80AA
const GL_SAMPLE_COVERAGE_INVERT = 0x80AB

/* GetTextureParameter */
/*      GL_TEXTURE_MAG_FILTER */
/*      GL_TEXTURE_MIN_FILTER */
/*      GL_TEXTURE_WRAP_S */
/*      GL_TEXTURE_WRAP_T */

const GL_NUM_COMPRESSED_TEXTURE_FORMATS = 0x86A2
const GL_COMPRESSED_TEXTURE_FORMATS = 0x86A3

/* HintMode */
const GL_DONT_CARE = 0x1100
const GL_FASTEST = 0x1101
const GL_NICEST = 0x1102

/* HintTarget */
const GL_GENERATE_MIPMAP_HINT = 0x8192

/* DataType */
const GL_BYTE = 0x1400
const GL_UNSIGNED_BYTE = 0x1401
const GL_SHORT = 0x1402

//const GL_UNSIGNED_SHORT          =       0x1403
const GL_INT = 0x1404

//const GL_UNSIGNED_INT            =       0x1405
const GL_FLOAT = 0x1406
const GL_FIXED = 0x140C

/* PixelFormat */
//const GL_DEPTH_COMPONENT     =           0x1902
const GL_ALPHA = 0x1906
const GL_RGB = 0x1907
const GL_RGBA = 0x1908
const GL_LUMINANCE = 0x1909
const GL_LUMINANCE_ALPHA = 0x190A

/* PixelType */
/*      GL_UNSIGNED_BYTE */
const GL_UNSIGNED_SHORT_4_4_4_4 = 0x8033
const GL_UNSIGNED_SHORT_5_5_5_1 = 0x8034
const GL_UNSIGNED_SHORT_5_6_5 = 0x8363

/* Shaders */
const GL_FRAGMENT_SHADER = 0x8B30
const GL_VERTEX_SHADER = 0x8B31
const GL_MAX_VERTEX_ATTRIBS = 0x8869
const GL_MAX_VERTEX_UNIFORM_VECTORS = 0x8DFB
const GL_MAX_VARYING_VECTORS = 0x8DFC
const GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = 0x8B4D
const GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS = 0x8B4C
const GL_MAX_TEXTURE_IMAGE_UNITS = 0x8872
const GL_MAX_FRAGMENT_UNIFORM_VECTORS = 0x8DFD
const GL_SHADER_TYPE = 0x8B4F
const GL_DELETE_STATUS = 0x8B80
const GL_LINK_STATUS = 0x8B82
const GL_VALIDATE_STATUS = 0x8B83
const GL_ATTACHED_SHADERS = 0x8B85
const GL_ACTIVE_UNIFORMS = 0x8B86
const GL_ACTIVE_UNIFORM_MAX_LENGTH = 0x8B87
const GL_ACTIVE_ATTRIBUTES = 0x8B89
const GL_ACTIVE_ATTRIBUTE_MAX_LENGTH = 0x8B8A
const GL_SHADING_LANGUAGE_VERSION = 0x8B8C
const GL_CURRENT_PROGRAM = 0x8B8D

/* StencilFunction */
const GL_NEVER = 0x0200
const GL_LESS = 0x0201
const GL_EQUAL = 0x0202
const GL_LEQUAL = 0x0203
const GL_GREATER = 0x0204
const GL_NOTEQUAL = 0x0205
const GL_GEQUAL = 0x0206
const GL_ALWAYS = 0x0207

/* StencilOp */
/*      GL_ZERO */
const GL_KEEP = 0x1E00
const GL_REPLACE = 0x1E01
const GL_INCR = 0x1E02
const GL_DECR = 0x1E03
const GL_INVERT = 0x150A
const GL_INCR_WRAP = 0x8507
const GL_DECR_WRAP = 0x8508

/* StringName */
const GL_VENDOR = 0x1F00
const GL_RENDERER = 0x1F01
const GL_VERSION = 0x1F02
const GL_EXTENSIONS = 0x1F03

/* TextureMagFilter */
const GL_NEAREST = 0x2600
const GL_LINEAR = 0x2601

/* TextureMinFilter */
/*      GL_NEAREST */
/*      GL_LINEAR */
const GL_NEAREST_MIPMAP_NEAREST = 0x2700
const GL_LINEAR_MIPMAP_NEAREST = 0x2701
const GL_NEAREST_MIPMAP_LINEAR = 0x2702
const GL_LINEAR_MIPMAP_LINEAR = 0x2703

/* TextureParameterName */
const GL_TEXTURE_MAG_FILTER = 0x2800
const GL_TEXTURE_MIN_FILTER = 0x2801
const GL_TEXTURE_WRAP_S = 0x2802
const GL_TEXTURE_WRAP_T = 0x2803

/* TextureTarget */
/*      GL_TEXTURE_2D */
const GL_TEXTURE = 0x1702

const GL_TEXTURE_CUBE_MAP = 0x8513
const GL_TEXTURE_BINDING_CUBE_MAP = 0x8514
const GL_TEXTURE_CUBE_MAP_POSITIVE_X = 0x8515
const GL_TEXTURE_CUBE_MAP_NEGATIVE_X = 0x8516
const GL_TEXTURE_CUBE_MAP_POSITIVE_Y = 0x8517
const GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x8518
const GL_TEXTURE_CUBE_MAP_POSITIVE_Z = 0x8519
const GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x851A
const GL_MAX_CUBE_MAP_TEXTURE_SIZE = 0x851C

/* TextureUnit */
const GL_TEXTURE0 = 0x84C0
const GL_TEXTURE1 = 0x84C1
const GL_TEXTURE2 = 0x84C2
const GL_TEXTURE3 = 0x84C3
const GL_TEXTURE4 = 0x84C4
const GL_TEXTURE5 = 0x84C5
const GL_TEXTURE6 = 0x84C6
const GL_TEXTURE7 = 0x84C7
const GL_TEXTURE8 = 0x84C8
const GL_TEXTURE9 = 0x84C9
const GL_TEXTURE10 = 0x84CA
const GL_TEXTURE11 = 0x84CB
const GL_TEXTURE12 = 0x84CC
const GL_TEXTURE13 = 0x84CD
const GL_TEXTURE14 = 0x84CE
const GL_TEXTURE15 = 0x84CF
const GL_TEXTURE16 = 0x84D0
const GL_TEXTURE17 = 0x84D1
const GL_TEXTURE18 = 0x84D2
const GL_TEXTURE19 = 0x84D3
const GL_TEXTURE20 = 0x84D4
const GL_TEXTURE21 = 0x84D5
const GL_TEXTURE22 = 0x84D6
const GL_TEXTURE23 = 0x84D7
const GL_TEXTURE24 = 0x84D8
const GL_TEXTURE25 = 0x84D9
const GL_TEXTURE26 = 0x84DA
const GL_TEXTURE27 = 0x84DB
const GL_TEXTURE28 = 0x84DC
const GL_TEXTURE29 = 0x84DD
const GL_TEXTURE30 = 0x84DE
const GL_TEXTURE31 = 0x84DF
const GL_ACTIVE_TEXTURE = 0x84E0

/* TextureWrapMode */
const GL_REPEAT = 0x2901
const GL_CLAMP_TO_EDGE = 0x812F
const GL_MIRRORED_REPEAT = 0x8370

/* Uniform Types */
const GL_FLOAT_VEC2 = 0x8B50
const GL_FLOAT_VEC3 = 0x8B51
const GL_FLOAT_VEC4 = 0x8B52
const GL_INT_VEC2 = 0x8B53
const GL_INT_VEC3 = 0x8B54
const GL_INT_VEC4 = 0x8B55
const GL_BOOL = 0x8B56
const GL_BOOL_VEC2 = 0x8B57
const GL_BOOL_VEC3 = 0x8B58
const GL_BOOL_VEC4 = 0x8B59
const GL_FLOAT_MAT2 = 0x8B5A
const GL_FLOAT_MAT3 = 0x8B5B
const GL_FLOAT_MAT4 = 0x8B5C
const GL_SAMPLER_2D = 0x8B5E
const GL_SAMPLER_CUBE = 0x8B60

/* Vertex Arrays */
const GL_VERTEX_ATTRIB_ARRAY_ENABLED = 0x8622
const GL_VERTEX_ATTRIB_ARRAY_SIZE = 0x8623
const GL_VERTEX_ATTRIB_ARRAY_STRIDE = 0x8624
const GL_VERTEX_ATTRIB_ARRAY_TYPE = 0x8625
const GL_VERTEX_ATTRIB_ARRAY_NORMALIZED = 0x886A
const GL_VERTEX_ATTRIB_ARRAY_POINTER = 0x8645
const GL_VERTEX_ATTRIB_ARRAY_BUFFER_BINDING = 0x889F

/* Read Format */
const GL_IMPLEMENTATION_COLOR_READ_TYPE = 0x8B9A
const GL_IMPLEMENTATION_COLOR_READ_FORMAT = 0x8B9B

/* Shader Source */
const GL_COMPILE_STATUS = 0x8B81
const GL_INFO_LOG_LENGTH = 0x8B84
const GL_SHADER_SOURCE_LENGTH = 0x8B88
const GL_SHADER_COMPILER = 0x8DFA

/* Shader Binary */
const GL_SHADER_BINARY_FORMATS = 0x8DF8
const GL_NUM_SHADER_BINARY_FORMATS = 0x8DF9

/* Shader Precision-Specified Types */
const GL_LOW_FLOAT = 0x8DF0
const GL_MEDIUM_FLOAT = 0x8DF1
const GL_HIGH_FLOAT = 0x8DF2
const GL_LOW_INT = 0x8DF3
const GL_MEDIUM_INT = 0x8DF4
const GL_HIGH_INT = 0x8DF5

/* Framebuffer Object. */
const GL_FRAMEBUFFER = 0x8D40
const GL_RENDERBUFFER = 0x8D41

const GL_RGBA4 = 0x8056
const GL_RGB5_A1 = 0x8057
const GL_RGB565 = 0x8D62

//const GL_DEPTH_COMPONENT16   =           0x81A5
const GL_STENCIL_INDEX8 = 0x8D48

const GL_RENDERBUFFER_WIDTH = 0x8D42
const GL_RENDERBUFFER_HEIGHT = 0x8D43
const GL_RENDERBUFFER_INTERNAL_FORMAT = 0x8D44
const GL_RENDERBUFFER_RED_SIZE = 0x8D50
const GL_RENDERBUFFER_GREEN_SIZE = 0x8D51
const GL_RENDERBUFFER_BLUE_SIZE = 0x8D52
const GL_RENDERBUFFER_ALPHA_SIZE = 0x8D53
const GL_RENDERBUFFER_DEPTH_SIZE = 0x8D54
const GL_RENDERBUFFER_STENCIL_SIZE = 0x8D55

const GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 0x8CD0
const GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 0x8CD1
const GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 0x8CD2
const GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 0x8CD3

const GL_COLOR_ATTACHMENT0 = 0x8CE0
const GL_DEPTH_ATTACHMENT = 0x8D00
const GL_STENCIL_ATTACHMENT = 0x8D20

const GL_NONE = 0

const GL_FRAMEBUFFER_COMPLETE = 0x8CD5
const GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 0x8CD6
const GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 0x8CD7
const GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS = 0x8CD9
const GL_FRAMEBUFFER_UNSUPPORTED = 0x8CDD

const GL_FRAMEBUFFER_BINDING = 0x8CA6
const GL_RENDERBUFFER_BINDING = 0x8CA7
const GL_MAX_RENDERBUFFER_SIZE = 0x84E8

const GL_INVALID_FRAMEBUFFER_OPERATION = 0x0506

/*-------------------------------------------------------------------------
 * GL core functions.
 *-----------------------------------------------------------------------*/

//GL_APICALL void         GL_APIENTRY glActiveTexture (GLenum texture);
//todo
//GL_APICALL void         GL_APIENTRY glAttachShader (GLuint program, GLuint shader);
//todo
//GL_APICALL void         GL_APIENTRY glBindAttribLocation (GLuint program, GLuint index, const GLchar* name);
//todo
//GL_APICALL void         GL_APIENTRY glBindBuffer (GLenum target, GLuint buffer);
//todo
//GL_APICALL void         GL_APIENTRY glBindFramebuffer (GLenum target, GLuint framebuffer);
//todo
//GL_APICALL void         GL_APIENTRY glBindRenderbuffer (GLenum target, GLuint renderbuffer);
//todo
//GL_APICALL void         GL_APIENTRY glBindTexture (GLenum target, GLuint texture);
//todo
//GL_APICALL void         GL_APIENTRY glBlendColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
//todo
//GL_APICALL void         GL_APIENTRY glBlendEquation ( GLenum mode );
//todo
//GL_APICALL void         GL_APIENTRY glBlendEquationSeparate (GLenum modeRGB, GLenum modeAlpha);
//todo
//GL_APICALL void         GL_APIENTRY glBlendFunc (GLenum sfactor, GLenum dfactor);
//todo
//GL_APICALL void         GL_APIENTRY glBlendFuncSeparate (GLenum srcRGB, GLenum dstRGB, GLenum srcAlpha, GLenum dstAlpha);
//todo
//GL_APICALL void         GL_APIENTRY glBufferData (GLenum target, GLsizeiptr size, const GLvoid* data, GLenum usage);
//todo
//GL_APICALL void         GL_APIENTRY glBufferSubData (GLenum target, GLintptr offset, GLsizeiptr size, const GLvoid* data);
//todo
//GL_APICALL GLenum       GL_APIENTRY glCheckFramebufferStatus (GLenum target);
//todo
//GL_APICALL void         GL_APIENTRY glClear (GLbitfield mask);
//todo
//GL_APICALL void         GL_APIENTRY glClearColor (GLclampf red, GLclampf green, GLclampf blue, GLclampf alpha);
//todo
//GL_APICALL void         GL_APIENTRY glClearDepthf (GLclampf depth);
//todo
//GL_APICALL void         GL_APIENTRY glClearStencil (GLint s);
//todo
//GL_APICALL void         GL_APIENTRY glColorMask (GLboolean red, GLboolean green, GLboolean blue, GLboolean alpha);
//todo
//GL_APICALL void         GL_APIENTRY glCompileShader (GLuint shader);
//todo
//GL_APICALL void         GL_APIENTRY glCompressedTexImage2D (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLint border, GLsizei imageSize, const GLvoid* data);
//todo
//GL_APICALL void         GL_APIENTRY glCompressedTexSubImage2D (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLsizei imageSize, const GLvoid* data);
//todo
//GL_APICALL void         GL_APIENTRY glCopyTexImage2D (GLenum target, GLint level, GLenum internalformat, GLint x, GLint y, GLsizei width, GLsizei height, GLint border);
//todo
//GL_APICALL void         GL_APIENTRY glCopyTexSubImage2D (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint x, GLint y, GLsizei width, GLsizei height);
//todo
//GL_APICALL GLuint       GL_APIENTRY glCreateProgram (void);
//todo
//GL_APICALL GLuint       GL_APIENTRY glCreateShader (GLenum type);
//todo
//GL_APICALL void         GL_APIENTRY glCullFace (GLenum mode);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteBuffers (GLsizei n, const GLuint* buffers);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteFramebuffers (GLsizei n, const GLuint* framebuffers);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteProgram (GLuint program);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteRenderbuffers (GLsizei n, const GLuint* renderbuffers);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteShader (GLuint shader);
//todo
//GL_APICALL void         GL_APIENTRY glDeleteTextures (GLsizei n, const GLuint* textures);
//todo
//GL_APICALL void         GL_APIENTRY glDepthFunc (GLenum func);
//todo
//GL_APICALL void         GL_APIENTRY glDepthMask (GLboolean flag);
//todo
//GL_APICALL void         GL_APIENTRY glDepthRangef (GLclampf zNear, GLclampf zFar);
//todo
//GL_APICALL void         GL_APIENTRY glDetachShader (GLuint program, GLuint shader);
//todo
//GL_APICALL void         GL_APIENTRY glDisable (GLenum cap);
//todo
//GL_APICALL void         GL_APIENTRY glDisableVertexAttribArray (GLuint index);
//todo
//GL_APICALL void         GL_APIENTRY glDrawArrays (GLenum mode, GLint first, GLsizei count);
//todo
//GL_APICALL void         GL_APIENTRY glDrawElements (GLenum mode, GLsizei count, GLenum type, const GLvoid* indices);
//todo
//GL_APICALL void         GL_APIENTRY glEnable (GLenum cap);
//todo
//GL_APICALL void         GL_APIENTRY glEnableVertexAttribArray (GLuint index);
//todo
//GL_APICALL void         GL_APIENTRY glFinish (void);
//todo
//GL_APICALL void         GL_APIENTRY glFlush (void);
//todo
//GL_APICALL void         GL_APIENTRY glFramebufferRenderbuffer (GLenum target, GLenum attachment, GLenum renderbuffertarget, GLuint renderbuffer);
//todo
//GL_APICALL void         GL_APIENTRY glFramebufferTexture2D (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level);
//todo
//GL_APICALL void         GL_APIENTRY glFrontFace (GLenum mode);
//todo
//GL_APICALL void         GL_APIENTRY glGenBuffers (GLsizei n, GLuint* buffers);
//todo
//GL_APICALL void         GL_APIENTRY glGenerateMipmap (GLenum target);
//todo
//GL_APICALL void         GL_APIENTRY glGenFramebuffers (GLsizei n, GLuint* framebuffers);
//todo
//GL_APICALL void         GL_APIENTRY glGenRenderbuffers (GLsizei n, GLuint* renderbuffers);
//todo
//GL_APICALL void         GL_APIENTRY glGenTextures (GLsizei n, GLuint* textures);
//todo
//GL_APICALL void         GL_APIENTRY glGetActiveAttrib (GLuint program, GLuint index, GLsizei bufsize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
//todo
//GL_APICALL void         GL_APIENTRY glGetActiveUniform (GLuint program, GLuint index, GLsizei bufsize, GLsizei* length, GLint* size, GLenum* type, GLchar* name);
//todo
//GL_APICALL void         GL_APIENTRY glGetAttachedShaders (GLuint program, GLsizei maxcount, GLsizei* count, GLuint* shaders);
//todo
//GL_APICALL GLint        GL_APIENTRY glGetAttribLocation (GLuint program, const GLchar* name);
//todo
//GL_APICALL void         GL_APIENTRY glGetBooleanv (GLenum pname, GLboolean* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetBufferParameteriv (GLenum target, GLenum pname, GLint* params);
//todo
//GL_APICALL GLenum       GL_APIENTRY glGetError (void);
//todo
//GL_APICALL void         GL_APIENTRY glGetFloatv (GLenum pname, GLfloat* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetFramebufferAttachmentParameteriv (GLenum target, GLenum attachment, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetIntegerv (GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetProgramiv (GLuint program, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetProgramInfoLog (GLuint program, GLsizei bufsize, GLsizei* length, GLchar* infolog);
//todo
//GL_APICALL void         GL_APIENTRY glGetRenderbufferParameteriv (GLenum target, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetShaderiv (GLuint shader, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetShaderInfoLog (GLuint shader, GLsizei bufsize, GLsizei* length, GLchar* infolog);
//todo
//GL_APICALL void         GL_APIENTRY glGetShaderPrecisionFormat (GLenum shadertype, GLenum precisiontype, GLint* range, GLint* precision);
//todo
//GL_APICALL void         GL_APIENTRY glGetShaderSource (GLuint shader, GLsizei bufsize, GLsizei* length, GLchar* source);
//todo
//GL_APICALL const GLubyte* GL_APIENTRY glGetString (GLenum name);
//todo
//GL_APICALL void         GL_APIENTRY glGetTexParameterfv (GLenum target, GLenum pname, GLfloat* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetTexParameteriv (GLenum target, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetUniformfv (GLuint program, GLint location, GLfloat* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetUniformiv (GLuint program, GLint location, GLint* params);
//todo
//GL_APICALL GLint        GL_APIENTRY glGetUniformLocation (GLuint program, const GLchar* name);
//todo
//GL_APICALL void         GL_APIENTRY glGetVertexAttribfv (GLuint index, GLenum pname, GLfloat* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetVertexAttribiv (GLuint index, GLenum pname, GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glGetVertexAttribPointerv (GLuint index, GLenum pname, GLvoid** pointer);
//todo
//GL_APICALL void         GL_APIENTRY glHint (GLenum target, GLenum mode);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsBuffer (GLuint buffer);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsEnabled (GLenum cap);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsFramebuffer (GLuint framebuffer);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsProgram (GLuint program);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsRenderbuffer (GLuint renderbuffer);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsShader (GLuint shader);
//todo
//GL_APICALL GLboolean    GL_APIENTRY glIsTexture (GLuint texture);
//todo
//GL_APICALL void         GL_APIENTRY glLineWidth (GLfloat width);
//todo
//GL_APICALL void         GL_APIENTRY glLinkProgram (GLuint program);
//todo
//GL_APICALL void         GL_APIENTRY glPixelStorei (GLenum pname, GLint param);
//todo
//GL_APICALL void         GL_APIENTRY glPolygonOffset (GLfloat factor, GLfloat units);
//todo
//GL_APICALL void         GL_APIENTRY glReadPixels (GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLvoid* pixels);
//todo
//GL_APICALL void         GL_APIENTRY glReleaseShaderCompiler (void);
//todo
//GL_APICALL void         GL_APIENTRY glRenderbufferStorage (GLenum target, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//GL_APICALL void         GL_APIENTRY glSampleCoverage (GLclampf value, GLboolean invert);
//todo
//GL_APICALL void         GL_APIENTRY glScissor (GLint x, GLint y, GLsizei width, GLsizei height);
//todo
//GL_APICALL void         GL_APIENTRY glShaderBinary (GLsizei n, const GLuint* shaders, GLenum binaryformat, const GLvoid* binary, GLsizei length);
//todo
//GL_APICALL void         GL_APIENTRY glShaderSource (GLuint shader, GLsizei count, const GLchar* const* string, const GLint* length);
//todo
//GL_APICALL void         GL_APIENTRY glStencilFunc (GLenum func, GLint ref, GLuint mask);
//todo
//GL_APICALL void         GL_APIENTRY glStencilFuncSeparate (GLenum face, GLenum func, GLint ref, GLuint mask);
//todo
//GL_APICALL void         GL_APIENTRY glStencilMask (GLuint mask);
//todo
//GL_APICALL void         GL_APIENTRY glStencilMaskSeparate (GLenum face, GLuint mask);
//todo
//GL_APICALL void         GL_APIENTRY glStencilOp (GLenum fail, GLenum zfail, GLenum zpass);
//todo
//GL_APICALL void         GL_APIENTRY glStencilOpSeparate (GLenum face, GLenum fail, GLenum zfail, GLenum zpass);
//todo
//GL_APICALL void         GL_APIENTRY glTexImage2D (GLenum target, GLint level, GLint internalformat, GLsizei width, GLsizei height, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
//todo
//GL_APICALL void         GL_APIENTRY glTexParameterf (GLenum target, GLenum pname, GLfloat param);
//todo
//GL_APICALL void         GL_APIENTRY glTexParameterfv (GLenum target, GLenum pname, const GLfloat* params);
//todo
//GL_APICALL void         GL_APIENTRY glTexParameteri (GLenum target, GLenum pname, GLint param);
//todo
//GL_APICALL void         GL_APIENTRY glTexParameteriv (GLenum target, GLenum pname, const GLint* params);
//todo
//GL_APICALL void         GL_APIENTRY glTexSubImage2D (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLsizei width, GLsizei height, GLenum format, GLenum type, const GLvoid* pixels);
//todo
//GL_APICALL void         GL_APIENTRY glUniform1f (GLint location, GLfloat x);
//todo
//GL_APICALL void         GL_APIENTRY glUniform1fv (GLint location, GLsizei count, const GLfloat* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform1i (GLint location, GLint x);
//todo
//GL_APICALL void         GL_APIENTRY glUniform1iv (GLint location, GLsizei count, const GLint* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform2f (GLint location, GLfloat x, GLfloat y);
//todo
//GL_APICALL void         GL_APIENTRY glUniform2fv (GLint location, GLsizei count, const GLfloat* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform2i (GLint location, GLint x, GLint y);
//todo
//GL_APICALL void         GL_APIENTRY glUniform2iv (GLint location, GLsizei count, const GLint* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform3f (GLint location, GLfloat x, GLfloat y, GLfloat z);
//todo
//GL_APICALL void         GL_APIENTRY glUniform3fv (GLint location, GLsizei count, const GLfloat* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform3i (GLint location, GLint x, GLint y, GLint z);
//todo
//GL_APICALL void         GL_APIENTRY glUniform3iv (GLint location, GLsizei count, const GLint* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform4f (GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
//todo
//GL_APICALL void         GL_APIENTRY glUniform4fv (GLint location, GLsizei count, const GLfloat* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniform4i (GLint location, GLint x, GLint y, GLint z, GLint w);
//todo
//GL_APICALL void         GL_APIENTRY glUniform4iv (GLint location, GLsizei count, const GLint* v);
//todo
//GL_APICALL void         GL_APIENTRY glUniformMatrix2fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
//todo
//GL_APICALL void         GL_APIENTRY glUniformMatrix3fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
//todo
//GL_APICALL void         GL_APIENTRY glUniformMatrix4fv (GLint location, GLsizei count, GLboolean transpose, const GLfloat* value);
//todo
//GL_APICALL void         GL_APIENTRY glUseProgram (GLuint program);
//todo
//GL_APICALL void         GL_APIENTRY glValidateProgram (GLuint program);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib1f (GLuint indx, GLfloat x);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib1fv (GLuint indx, const GLfloat* values);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib2f (GLuint indx, GLfloat x, GLfloat y);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib2fv (GLuint indx, const GLfloat* values);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib3f (GLuint indx, GLfloat x, GLfloat y, GLfloat z);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib3fv (GLuint indx, const GLfloat* values);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib4f (GLuint indx, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttrib4fv (GLuint indx, const GLfloat* values);
//todo
//GL_APICALL void         GL_APIENTRY glVertexAttribPointer (GLuint indx, GLint size, GLenum type, GLboolean normalized, GLsizei stride, const GLvoid* ptr);
//todo
//GL_APICALL void         GL_APIENTRY glViewport (GLint x, GLint y, GLsizei width, GLsizei height);
//todo
