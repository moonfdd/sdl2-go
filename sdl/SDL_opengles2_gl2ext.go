package sdl

import "sdl2-go/common"

//#ifndef __gl2ext_h_
//#define __gl2ext_h_
//
///* $Revision: 22801 $ on $Date:: 2013-08-21 03:20:48 -0700 #$ */
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
//#ifndef GL_APIENTRYP
//#   define GL_APIENTRYP GL_APIENTRY*
//#endif
//
///* New types shared by several extensions */
//
//#ifndef __gl3_h_
///* These are defined with respect to <inttypes.h> in the
// * Apple extension spec, but they are also used by non-APPLE
// * extensions, and in the Khronos header we use the Khronos
// * portable types in khrplatform.h, which must be defined.
// */
//typedef khronos_int64_t GLint64;
//typedef khronos_uint64_t GLuint64;
//typedef struct __GLsync *GLsync;
//#endif
//
//
///*------------------------------------------------------------------------*
// * OES extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_OES_compressed_ETC1_RGB8_texture */
//#ifndef GL_OES_compressed_ETC1_RGB8_texture
const GL_ETC1_RGB8_OES = 0x8D64

//#endif
//
///* GL_OES_compressed_paletted_texture */
//#ifndef GL_OES_compressed_paletted_texture
const GL_PALETTE4_RGB8_OES = 0x8B90
const GL_PALETTE4_RGBA8_OES = 0x8B91
const GL_PALETTE4_R5_G6_B5_OES = 0x8B92
const GL_PALETTE4_RGBA4_OES = 0x8B93
const GL_PALETTE4_RGB5_A1_OES = 0x8B94
const GL_PALETTE8_RGB8_OES = 0x8B95
const GL_PALETTE8_RGBA8_OES = 0x8B96
const GL_PALETTE8_R5_G6_B5_OES = 0x8B97
const GL_PALETTE8_RGBA4_OES = 0x8B98
const GL_PALETTE8_RGB5_A1_OES = 0x8B99

//#endif
//
///* GL_OES_depth24 */
//#ifndef GL_OES_depth24
const GL_DEPTH_COMPONENT24_OES = 0x81A6

//#endif

///* GL_OES_depth32 */
//#ifndef GL_OES_depth32
//const GL_DEPTH_COMPONENT32_OES              =                  0x81A7
//#endif
//
///* GL_OES_depth_texture */
///* No new tokens introduced by this extension. */
//
///* GL_OES_EGL_image */
//#ifndef GL_OES_EGL_image
//typedef void* GLeglImageOES;
type GLeglImageOES = common.FVoidP

//#endif

///* GL_OES_EGL_image_external */
//#ifndef GL_OES_EGL_image_external
///* GLeglImageOES defined in GL_OES_EGL_image already. */
const GL_TEXTURE_EXTERNAL_OES = 0x8D65
const GL_SAMPLER_EXTERNAL_OES = 0x8D66
const GL_TEXTURE_BINDING_EXTERNAL_OES = 0x8D67
const GL_REQUIRED_TEXTURE_IMAGE_UNITS_OES = 0x8D68

//#endif
//
///* GL_OES_element_index_uint */
//#ifndef GL_OES_element_index_uint
//const GL_UNSIGNED_INT                            =             0x1405
//#endif
//
///* GL_OES_get_program_binary */
//#ifndef GL_OES_get_program_binary
const GL_PROGRAM_BINARY_LENGTH_OES = 0x8741
const GL_NUM_PROGRAM_BINARY_FORMATS_OES = 0x87FE
const GL_PROGRAM_BINARY_FORMATS_OES = 0x87FF

//#endif
//
///* GL_OES_mapbuffer */
//#ifndef GL_OES_mapbuffer
const GL_WRITE_ONLY_OES = 0x88B9
const GL_BUFFER_ACCESS_OES = 0x88BB
const GL_BUFFER_MAPPED_OES = 0x88BC
const GL_BUFFER_MAP_POINTER_OES = 0x88BD

//#endif
//
///* GL_OES_packed_depth_stencil */
//#ifndef GL_OES_packed_depth_stencil
//const GL_DEPTH_STENCIL_OES             =                       0x84F9
//const GL_UNSIGNED_INT_24_8_OES         =                       0x84FA
//const GL_DEPTH24_STENCIL8_OES           =                      0x88F0
//#endif
//
///* GL_OES_required_internalformat */
//#ifndef GL_OES_required_internalformat
const GL_ALPHA8_OES = 0x803C
const GL_DEPTH_COMPONENT16_OES = 0x81A5

/* reuse GL_DEPTH_COMPONENT24_OES */
/* reuse GL_DEPTH24_STENCIL8_OES */
/* reuse GL_DEPTH_COMPONENT32_OES */
const GL_LUMINANCE4_ALPHA4_OES = 0x8043
const GL_LUMINANCE8_ALPHA8_OES = 0x8045
const GL_LUMINANCE8_OES = 0x8040
const GL_RGBA4_OES = 0x8056
const GL_RGB5_A1_OES = 0x8057
const GL_RGB565_OES = 0x8D62

///* reuse GL_RGB8_OES */
///* reuse GL_RGBA8_OES */
///* reuse GL_RGB10_EXT */
///* reuse GL_RGB10_A2_EXT */
//#endif
//
///* GL_OES_rgb8_rgba8 */
//#ifndef GL_OES_rgb8_rgba8
const GL_RGB8_OES = 0x8051
const GL_RGBA8_OES = 0x8058

//#endif

///* GL_OES_standard_derivatives */
//#ifndef GL_OES_standard_derivatives
const GL_FRAGMENT_SHADER_DERIVATIVE_HINT_OES = 0x8B8B

//#endif
//
///* GL_OES_stencil1 */
//#ifndef GL_OES_stencil1
//const GL_STENCIL_INDEX1_OES                                   0x8D46
//#endif
//
///* GL_OES_stencil4 */
//#ifndef GL_OES_stencil4
const GL_STENCIL_INDEX4_OES = 0x8D47

//#endif
//
//#ifndef GL_OES_surfaceless_context
const GL_FRAMEBUFFER_UNDEFINED_OES = 0x8219

//#endif
//
///* GL_OES_texture_3D */
//#ifndef GL_OES_texture_3D
const GL_TEXTURE_WRAP_R_OES = 0x8072
const GL_TEXTURE_3D_OES = 0x806F
const GL_TEXTURE_BINDING_3D_OES = 0x806A
const GL_MAX_3D_TEXTURE_SIZE_OES = 0x8073
const GL_SAMPLER_3D_OES = 0x8B5F
const GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_3D_ZOFFSET_OES = 0x8CD4

//#endif
//
///* GL_OES_texture_float */
///* No new tokens introduced by this extension. */
//
///* GL_OES_texture_float_linear */
///* No new tokens introduced by this extension. */
//
///* GL_OES_texture_half_float */
//#ifndef GL_OES_texture_half_float
const GL_HALF_FLOAT_OES = 0x8D61

//#endif
//
///* GL_OES_texture_half_float_linear */
///* No new tokens introduced by this extension. */
//
///* GL_OES_texture_npot */
///* No new tokens introduced by this extension. */
//
///* GL_OES_vertex_array_object */
//#ifndef GL_OES_vertex_array_object
const GL_VERTEX_ARRAY_BINDING_OES = 0x85B5

//#endif
//
///* GL_OES_vertex_half_float */
///* GL_HALF_FLOAT_OES defined in GL_OES_texture_half_float already. */
//
///* GL_OES_vertex_type_10_10_10_2 */
//#ifndef GL_OES_vertex_type_10_10_10_2
const GL_UNSIGNED_INT_10_10_10_2_OES = 0x8DF6
const GL_INT_10_10_10_2_OES = 0x8DF7

//#endif
//
///*------------------------------------------------------------------------*
// * KHR extension tokens
// *------------------------------------------------------------------------*/
//
//#ifndef GL_KHR_debug
//typedef void (GL_APIENTRYP GLDEBUGPROCKHR)(GLenum source,GLenum type,GLuint id,GLenum severity,GLsizei length,const GLchar *message,const void *userParam);
//todo
const GL_DEBUG_OUTPUT_SYNCHRONOUS_KHR = 0x8242
const GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH_KHR = 0x8243
const GL_DEBUG_CALLBACK_FUNCTION_KHR = 0x8244
const GL_DEBUG_CALLBACK_USER_PARAM_KHR = 0x8245
const GL_DEBUG_SOURCE_API_KHR = 0x8246
const GL_DEBUG_SOURCE_WINDOW_SYSTEM_KHR = 0x8247
const GL_DEBUG_SOURCE_SHADER_COMPILER_KHR = 0x8248
const GL_DEBUG_SOURCE_THIRD_PARTY_KHR = 0x8249
const GL_DEBUG_SOURCE_APPLICATION_KHR = 0x824A
const GL_DEBUG_SOURCE_OTHER_KHR = 0x824B
const GL_DEBUG_TYPE_ERROR_KHR = 0x824C
const GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR_KHR = 0x824D
const GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR_KHR = 0x824E
const GL_DEBUG_TYPE_PORTABILITY_KHR = 0x824F
const GL_DEBUG_TYPE_PERFORMANCE_KHR = 0x8250
const GL_DEBUG_TYPE_OTHER_KHR = 0x8251
const GL_DEBUG_TYPE_MARKER_KHR = 0x8268
const GL_DEBUG_TYPE_PUSH_GROUP_KHR = 0x8269
const GL_DEBUG_TYPE_POP_GROUP_KHR = 0x826A
const GL_DEBUG_SEVERITY_NOTIFICATION_KHR = 0x826B
const GL_MAX_DEBUG_GROUP_STACK_DEPTH_KHR = 0x826C
const GL_DEBUG_GROUP_STACK_DEPTH_KHR = 0x826D
const GL_BUFFER_KHR = 0x82E0
const GL_SHADER_KHR = 0x82E1
const GL_PROGRAM_KHR = 0x82E2
const GL_QUERY_KHR = 0x82E3

/* PROGRAM_PIPELINE only in GL */
const GL_SAMPLER_KHR = 0x82E6

/* DISPLAY_LIST only in GL */
const GL_MAX_LABEL_LENGTH_KHR = 0x82E8
const GL_MAX_DEBUG_MESSAGE_LENGTH_KHR = 0x9143
const GL_MAX_DEBUG_LOGGED_MESSAGES_KHR = 0x9144
const GL_DEBUG_LOGGED_MESSAGES_KHR = 0x9145
const GL_DEBUG_SEVERITY_HIGH_KHR = 0x9146
const GL_DEBUG_SEVERITY_MEDIUM_KHR = 0x9147
const GL_DEBUG_SEVERITY_LOW_KHR = 0x9148
const GL_DEBUG_OUTPUT_KHR = 0x92E0
const GL_CONTEXT_FLAG_DEBUG_BIT_KHR = 0x00000002
const GL_STACK_OVERFLOW_KHR = 0x0503
const GL_STACK_UNDERFLOW_KHR = 0x0504

//#endif
//
//#ifndef GL_KHR_texture_compression_astc_ldr
const GL_COMPRESSED_RGBA_ASTC_4x4_KHR = 0x93B0
const GL_COMPRESSED_RGBA_ASTC_5x4_KHR = 0x93B1
const GL_COMPRESSED_RGBA_ASTC_5x5_KHR = 0x93B2
const GL_COMPRESSED_RGBA_ASTC_6x5_KHR = 0x93B3
const GL_COMPRESSED_RGBA_ASTC_6x6_KHR = 0x93B4
const GL_COMPRESSED_RGBA_ASTC_8x5_KHR = 0x93B5
const GL_COMPRESSED_RGBA_ASTC_8x6_KHR = 0x93B6
const GL_COMPRESSED_RGBA_ASTC_8x8_KHR = 0x93B7
const GL_COMPRESSED_RGBA_ASTC_10x5_KHR = 0x93B8
const GL_COMPRESSED_RGBA_ASTC_10x6_KHR = 0x93B9
const GL_COMPRESSED_RGBA_ASTC_10x8_KHR = 0x93BA
const GL_COMPRESSED_RGBA_ASTC_10x10_KHR = 0x93BB
const GL_COMPRESSED_RGBA_ASTC_12x10_KHR = 0x93BC
const GL_COMPRESSED_RGBA_ASTC_12x12_KHR = 0x93BD
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR = 0x93D0
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR = 0x93D1
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR = 0x93D2
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR = 0x93D3
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR = 0x93D4
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR = 0x93D5
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR = 0x93D6
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR = 0x93D7
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR = 0x93D8
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR = 0x93D9
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR = 0x93DA
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR = 0x93DB
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR = 0x93DC
const GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR = 0x93DD

//#endif
//
///*------------------------------------------------------------------------*
// * AMD extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_AMD_compressed_3DC_texture */
//#ifndef GL_AMD_compressed_3DC_texture
const GL_3DC_X_AMD = 0x87F9
const GL_3DC_XY_AMD = 0x87FA

//#endif
//
///* GL_AMD_compressed_ATC_texture */
//#ifndef GL_AMD_compressed_ATC_texture
const GL_ATC_RGB_AMD = 0x8C92
const GL_ATC_RGBA_EXPLICIT_ALPHA_AMD = 0x8C93
const GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = 0x87EE

//#endif
//
///* GL_AMD_performance_monitor */
//#ifndef GL_AMD_performance_monitor
const GL_COUNTER_TYPE_AMD = 0x8BC0
const GL_COUNTER_RANGE_AMD = 0x8BC1
const GL_UNSIGNED_INT64_AMD = 0x8BC2
const GL_PERCENTAGE_AMD = 0x8BC3
const GL_PERFMON_RESULT_AVAILABLE_AMD = 0x8BC4
const GL_PERFMON_RESULT_SIZE_AMD = 0x8BC5
const GL_PERFMON_RESULT_AMD = 0x8BC6

//#endif
//
///* GL_AMD_program_binary_Z400 */
//#ifndef GL_AMD_program_binary_Z400
const GL_Z400_BINARY_AMD = 0x8740

//#endif
//
///*------------------------------------------------------------------------*
// * ANGLE extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_ANGLE_depth_texture */
//#ifndef GL_ANGLE_depth_texture
const GL_DEPTH_COMPONENT = 0x1902
const GL_DEPTH_STENCIL_OES = 0x84F9
const GL_UNSIGNED_SHORT = 0x1403
const GL_UNSIGNED_INT = 0x1405
const GL_UNSIGNED_INT_24_8_OES = 0x84FA
const GL_DEPTH_COMPONENT16 = 0x81A5
const GL_DEPTH_COMPONENT32_OES = 0x81A7
const GL_DEPTH24_STENCIL8_OES = 0x88F0

//#endif
//
///* GL_ANGLE_framebuffer_blit */
//#ifndef GL_ANGLE_framebuffer_blit
const GL_READ_FRAMEBUFFER_ANGLE = 0x8CA8
const GL_DRAW_FRAMEBUFFER_ANGLE = 0x8CA9
const GL_DRAW_FRAMEBUFFER_BINDING_ANGLE = 0x8CA6
const GL_READ_FRAMEBUFFER_BINDING_ANGLE = 0x8CAA

//#endif
//
///* GL_ANGLE_framebuffer_multisample */
//#ifndef GL_ANGLE_framebuffer_multisample
const GL_RENDERBUFFER_SAMPLES_ANGLE = 0x8CAB
const GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_ANGLE = 0x8D56
const GL_MAX_SAMPLES_ANGLE = 0x8D57

//#endif
//
///* GL_ANGLE_instanced_arrays */
//#ifndef GL_ANGLE_instanced_arrays
const GL_VERTEX_ATTRIB_ARRAY_DIVISOR_ANGLE = 0x88FE

//#endif
//
///* GL_ANGLE_pack_reverse_row_order */
//#ifndef GL_ANGLE_pack_reverse_row_order
const GL_PACK_REVERSE_ROW_ORDER_ANGLE = 0x93A4

//#endif
//
///* GL_ANGLE_program_binary */
//#ifndef GL_ANGLE_program_binary
const GL_PROGRAM_BINARY_ANGLE = 0x93A6

//#endif
//
///* GL_ANGLE_texture_compression_dxt3 */
//#ifndef GL_ANGLE_texture_compression_dxt3
const GL_COMPRESSED_RGBA_S3TC_DXT3_ANGLE = 0x83F2

//#endif
//
///* GL_ANGLE_texture_compression_dxt5 */
//#ifndef GL_ANGLE_texture_compression_dxt5
const GL_COMPRESSED_RGBA_S3TC_DXT5_ANGLE = 0x83F3

//#endif
//
///* GL_ANGLE_texture_usage */
//#ifndef GL_ANGLE_texture_usage
const GL_TEXTURE_USAGE_ANGLE = 0x93A2
const GL_FRAMEBUFFER_ATTACHMENT_ANGLE = 0x93A3

//#endif
//
///* GL_ANGLE_translated_shader_source */
//#ifndef GL_ANGLE_translated_shader_source
const GL_TRANSLATED_SHADER_SOURCE_LENGTH_ANGLE = 0x93A0

//#endif
//
///*------------------------------------------------------------------------*
// * APPLE extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_APPLE_copy_texture_levels */
///* No new tokens introduced by this extension. */
//
///* GL_APPLE_framebuffer_multisample */
//#ifndef GL_APPLE_framebuffer_multisample
const GL_RENDERBUFFER_SAMPLES_APPLE = 0x8CAB
const GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_APPLE = 0x8D56
const GL_MAX_SAMPLES_APPLE = 0x8D57
const GL_READ_FRAMEBUFFER_APPLE = 0x8CA8
const GL_DRAW_FRAMEBUFFER_APPLE = 0x8CA9
const GL_DRAW_FRAMEBUFFER_BINDING_APPLE = 0x8CA6
const GL_READ_FRAMEBUFFER_BINDING_APPLE = 0x8CAA

//#endif
//
///* GL_APPLE_rgb_422 */
//#ifndef GL_APPLE_rgb_422
const GL_RGB_422_APPLE = 0x8A1F
const GL_UNSIGNED_SHORT_8_8_APPLE = 0x85BA
const GL_UNSIGNED_SHORT_8_8_REV_APPLE = 0x85BB

//#endif
//
///* GL_APPLE_sync */
//#ifndef GL_APPLE_sync

const GL_SYNC_OBJECT_APPLE = 0x8A53
const GL_MAX_SERVER_WAIT_TIMEOUT_APPLE = 0x9111
const GL_OBJECT_TYPE_APPLE = 0x9112
const GL_SYNC_CONDITION_APPLE = 0x9113
const GL_SYNC_STATUS_APPLE = 0x9114
const GL_SYNC_FLAGS_APPLE = 0x9115
const GL_SYNC_FENCE_APPLE = 0x9116
const GL_SYNC_GPU_COMMANDS_COMPLETE_APPLE = 0x9117
const GL_UNSIGNALED_APPLE = 0x9118
const GL_SIGNALED_APPLE = 0x9119
const GL_ALREADY_SIGNALED_APPLE = 0x911A
const GL_TIMEOUT_EXPIRED_APPLE = 0x911B
const GL_CONDITION_SATISFIED_APPLE = 0x911C
const GL_WAIT_FAILED_APPLE = 0x911D
const GL_SYNC_FLUSH_COMMANDS_BIT_APPLE = 0x00000001
const GL_TIMEOUT_IGNORED_APPLE = 0xFFFFFFFFFFFFFFFF

//#endif
//
///* GL_APPLE_texture_format_BGRA8888 */
//#ifndef GL_APPLE_texture_format_BGRA8888
//const GL_BGRA_EXT                            =                 0x80E1
//#endif
//
///* GL_APPLE_texture_max_level */
//#ifndef GL_APPLE_texture_max_level
const GL_TEXTURE_MAX_LEVEL_APPLE = 0x813D

//#endif
//
///*------------------------------------------------------------------------*
// * ARM extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_ARM_mali_program_binary */
//#ifndef GL_ARM_mali_program_binary
const GL_MALI_PROGRAM_BINARY_ARM = 0x8F61

////#endif
//
///* GL_ARM_mali_shader_binary */
//#ifndef GL_ARM_mali_shader_binary
const GL_MALI_SHADER_BINARY_ARM = 0x8F60

//#endif
//
///* GL_ARM_rgba8 */
///* No new tokens introduced by this extension. */
//
///*------------------------------------------------------------------------*
// * EXT extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_EXT_blend_minmax */
//#ifndef GL_EXT_blend_minmax
const GL_MIN_EXT = 0x8007
const GL_MAX_EXT = 0x8008

//#endif
//
///* GL_EXT_color_buffer_half_float */
//#ifndef GL_EXT_color_buffer_half_float
const GL_RGBA16F_EXT = 0x881A
const GL_RGB16F_EXT = 0x881B

//const GL_RG16F_EXT                                   =         0x822F
//const GL_R16F_EXT                                      =       0x822D
const GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE_EXT = 0x8211
const GL_UNSIGNED_NORMALIZED_EXT = 0x8C17

//#endif
//
///* GL_EXT_debug_label */
//#ifndef GL_EXT_debug_label
const GL_PROGRAM_PIPELINE_OBJECT_EXT = 0x8A4F
const GL_PROGRAM_OBJECT_EXT = 0x8B40
const GL_SHADER_OBJECT_EXT = 0x8B48
const GL_BUFFER_OBJECT_EXT = 0x9151
const GL_QUERY_OBJECT_EXT = 0x9153
const GL_VERTEX_ARRAY_OBJECT_EXT = 0x9154

//#endif
//
///* GL_EXT_debug_marker */
///* No new tokens introduced by this extension. */
//
///* GL_EXT_discard_framebuffer */
//#ifndef GL_EXT_discard_framebuffer
const GL_COLOR_EXT = 0x1800
const GL_DEPTH_EXT = 0x1801
const GL_STENCIL_EXT = 0x1802

//#endif
//
//#ifndef GL_EXT_disjoint_timer_query
const GL_QUERY_COUNTER_BITS_EXT = 0x8864
const GL_CURRENT_QUERY_EXT = 0x8865
const GL_QUERY_RESULT_EXT = 0x8866
const GL_QUERY_RESULT_AVAILABLE_EXT = 0x8867
const GL_TIME_ELAPSED_EXT = 0x88BF
const GL_TIMESTAMP_EXT = 0x8E28
const GL_GPU_DISJOINT_EXT = 0x8FBB

//#endif
//
//#ifndef GL_EXT_draw_buffers
//const GL_EXT_draw_buffers = 1
const GL_MAX_COLOR_ATTACHMENTS_EXT = 0x8CDF
const GL_MAX_DRAW_BUFFERS_EXT = 0x8824
const GL_DRAW_BUFFER0_EXT = 0x8825
const GL_DRAW_BUFFER1_EXT = 0x8826
const GL_DRAW_BUFFER2_EXT = 0x8827
const GL_DRAW_BUFFER3_EXT = 0x8828
const GL_DRAW_BUFFER4_EXT = 0x8829
const GL_DRAW_BUFFER5_EXT = 0x882A
const GL_DRAW_BUFFER6_EXT = 0x882B
const GL_DRAW_BUFFER7_EXT = 0x882C
const GL_DRAW_BUFFER8_EXT = 0x882D
const GL_DRAW_BUFFER9_EXT = 0x882E
const GL_DRAW_BUFFER10_EXT = 0x882F
const GL_DRAW_BUFFER11_EXT = 0x8830
const GL_DRAW_BUFFER12_EXT = 0x8831
const GL_DRAW_BUFFER13_EXT = 0x8832
const GL_DRAW_BUFFER14_EXT = 0x8833
const GL_DRAW_BUFFER15_EXT = 0x8834
const GL_COLOR_ATTACHMENT0_EXT = 0x8CE0
const GL_COLOR_ATTACHMENT1_EXT = 0x8CE1
const GL_COLOR_ATTACHMENT2_EXT = 0x8CE2
const GL_COLOR_ATTACHMENT3_EXT = 0x8CE3
const GL_COLOR_ATTACHMENT4_EXT = 0x8CE4
const GL_COLOR_ATTACHMENT5_EXT = 0x8CE5
const GL_COLOR_ATTACHMENT6_EXT = 0x8CE6
const GL_COLOR_ATTACHMENT7_EXT = 0x8CE7
const GL_COLOR_ATTACHMENT8_EXT = 0x8CE8
const GL_COLOR_ATTACHMENT9_EXT = 0x8CE9
const GL_COLOR_ATTACHMENT10_EXT = 0x8CEA
const GL_COLOR_ATTACHMENT11_EXT = 0x8CEB
const GL_COLOR_ATTACHMENT12_EXT = 0x8CEC
const GL_COLOR_ATTACHMENT13_EXT = 0x8CED
const GL_COLOR_ATTACHMENT14_EXT = 0x8CEE
const GL_COLOR_ATTACHMENT15_EXT = 0x8CEF

//#endif
//
///* GL_EXT_map_buffer_range */
//#ifndef GL_EXT_map_buffer_range
const GL_MAP_READ_BIT_EXT = 0x0001
const GL_MAP_WRITE_BIT_EXT = 0x0002
const GL_MAP_INVALIDATE_RANGE_BIT_EXT = 0x0004
const GL_MAP_INVALIDATE_BUFFER_BIT_EXT = 0x0008
const GL_MAP_FLUSH_EXPLICIT_BIT_EXT = 0x0010
const GL_MAP_UNSYNCHRONIZED_BIT_EXT = 0x0020

//#endif
//
///* GL_EXT_multisampled_render_to_texture */
//#ifndef GL_EXT_multisampled_render_to_texture
const GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_SAMPLES_EXT = 0x8D6C

/* reuse values from GL_EXT_framebuffer_multisample (desktop extension) */
const GL_RENDERBUFFER_SAMPLES_EXT = 0x8CAB
const GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_EXT = 0x8D56
const GL_MAX_SAMPLES_EXT = 0x8D57

//#endif
//
///* GL_EXT_multiview_draw_buffers */
//#ifndef GL_EXT_multiview_draw_buffers
const GL_COLOR_ATTACHMENT_EXT = 0x90F0
const GL_MULTIVIEW_EXT = 0x90F1
const GL_DRAW_BUFFER_EXT = 0x0C01
const GL_READ_BUFFER_EXT = 0x0C02
const GL_MAX_MULTIVIEW_BUFFERS_EXT = 0x90F2

//#endif
//
///* GL_EXT_multi_draw_arrays */
///* No new tokens introduced by this extension. */
//
///* GL_EXT_occlusion_query_boolean */
//#ifndef GL_EXT_occlusion_query_boolean
const GL_ANY_SAMPLES_PASSED_EXT = 0x8C2F
const GL_ANY_SAMPLES_PASSED_CONSERVATIVE_EXT = 0x8D6A

//const GL_CURRENT_QUERY_EXT                        =            0x8865
//const GL_QUERY_RESULT_EXT                         =            0x8866
//const GL_QUERY_RESULT_AVAILABLE_EXT                =           0x8867
//#endif
//
///* GL_EXT_read_format_bgra */
//#ifndef GL_EXT_read_format_bgra
//const GL_BGRA_EXT                                 =            0x80E1
const GL_UNSIGNED_SHORT_4_4_4_4_REV_EXT = 0x8365
const GL_UNSIGNED_SHORT_1_5_5_5_REV_EXT = 0x8366

//#endif
//
///* GL_EXT_robustness */
//#ifndef GL_EXT_robustness
///* reuse GL_NO_ERROR */
const GL_GUILTY_CONTEXT_RESET_EXT = 0x8253
const GL_INNOCENT_CONTEXT_RESET_EXT = 0x8254
const GL_UNKNOWN_CONTEXT_RESET_EXT = 0x8255
const GL_CONTEXT_ROBUST_ACCESS_EXT = 0x90F3
const GL_RESET_NOTIFICATION_STRATEGY_EXT = 0x8256
const GL_LOSE_CONTEXT_ON_RESET_EXT = 0x8252
const GL_NO_RESET_NOTIFICATION_EXT = 0x8261

//#endif
//
///* GL_EXT_separate_shader_objects */
//#ifndef GL_EXT_separate_shader_objects
const GL_VERTEX_SHADER_BIT_EXT = 0x00000001
const GL_FRAGMENT_SHADER_BIT_EXT = 0x00000002
const GL_ALL_SHADER_BITS_EXT = 0xFFFFFFFF
const GL_PROGRAM_SEPARABLE_EXT = 0x8258
const GL_ACTIVE_PROGRAM_EXT = 0x8259
const GL_PROGRAM_PIPELINE_BINDING_EXT = 0x825A

//#endif
//
///* GL_EXT_shader_framebuffer_fetch */
//#ifndef GL_EXT_shader_framebuffer_fetch
const GL_FRAGMENT_SHADER_DISCARDS_SAMPLES_EXT = 0x8A52

//#endif
//
///* GL_EXT_shader_texture_lod */
///* No new tokens introduced by this extension. */
//
///* GL_EXT_shadow_samplers */
//#ifndef GL_EXT_shadow_samplers
const GL_TEXTURE_COMPARE_MODE_EXT = 0x884C
const GL_TEXTURE_COMPARE_FUNC_EXT = 0x884D
const GL_COMPARE_REF_TO_TEXTURE_EXT = 0x884E
const GL_SAMPLER_2D_SHADOW_EXT = 0x8B62

//#endif
//
///* GL_EXT_sRGB */
//#ifndef GL_EXT_sRGB
const GL_SRGB_EXT = 0x8C40
const GL_SRGB_ALPHA_EXT = 0x8C42
const GL_SRGB8_ALPHA8_EXT = 0x8C43
const GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING_EXT = 0x8210

//#endif
//
///* GL_EXT_sRGB_write_control */
//#ifndef GL_EXT_sRGB_write_control
const GL_EXT_sRGB_write_control = 1
const GL_FRAMEBUFFER_SRGB_EXT = 0x8DB9

//#endif
//
///* GL_EXT_texture_compression_dxt1 */
//#ifndef GL_EXT_texture_compression_dxt1
const GL_COMPRESSED_RGB_S3TC_DXT1_EXT = 0x83F0
const GL_COMPRESSED_RGBA_S3TC_DXT1_EXT = 0x83F1

//#endif
//
///* GL_EXT_texture_filter_anisotropic */
//#ifndef GL_EXT_texture_filter_anisotropic
const GL_TEXTURE_MAX_ANISOTROPY_EXT = 0x84FE
const GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = 0x84FF

//#endif
//
///* GL_EXT_texture_format_BGRA8888 */
//#ifndef GL_EXT_texture_format_BGRA8888
const GL_BGRA_EXT = 0x80E1

//#endif
//
///* GL_EXT_texture_rg */
//#ifndef GL_EXT_texture_rg
const GL_RED_EXT = 0x1903
const GL_RG_EXT = 0x8227

//const GL_R8_EXT                      =                         0x8229
//const GL_RG8_EXT                      =                        0x822B
//#endif
//
///* GL_EXT_texture_sRGB_decode */
//#ifndef GL_EXT_texture_sRGB_decode
const GL_EXT_texture_sRGB_decode = 1
const GL_TEXTURE_SRGB_DECODE_EXT = 0x8A48
const GL_DECODE_EXT = 0x8A49
const GL_SKIP_DECODE_EXT = 0x8A4A

//#endif
//
///* GL_EXT_texture_storage */
//#ifndef GL_EXT_texture_storage
const GL_TEXTURE_IMMUTABLE_FORMAT_EXT = 0x912F
const GL_ALPHA8_EXT = 0x803C
const GL_LUMINANCE8_EXT = 0x8040
const GL_LUMINANCE8_ALPHA8_EXT = 0x8045
const GL_RGBA32F_EXT = 0x8814
const GL_RGB32F_EXT = 0x8815
const GL_ALPHA32F_EXT = 0x8816
const GL_LUMINANCE32F_EXT = 0x8818
const GL_LUMINANCE_ALPHA32F_EXT = 0x8819

/* reuse GL_RGBA16F_EXT */
/* reuse GL_RGB16F_EXT */
const GL_ALPHA16F_EXT = 0x881C
const GL_LUMINANCE16F_EXT = 0x881E
const GL_LUMINANCE_ALPHA16F_EXT = 0x881F
const GL_RGB10_A2_EXT = 0x8059
const GL_RGB10_EXT = 0x8052
const GL_BGRA8_EXT = 0x93A1
const GL_R8_EXT = 0x8229
const GL_RG8_EXT = 0x822B
const GL_R32F_EXT = 0x822E
const GL_RG32F_EXT = 0x8230
const GL_R16F_EXT = 0x822D
const GL_RG16F_EXT = 0x822F

//#endif
//
///* GL_EXT_texture_type_2_10_10_10_REV */
//#ifndef GL_EXT_texture_type_2_10_10_10_REV
const GL_UNSIGNED_INT_2_10_10_10_REV_EXT = 0x8368

//#endif
//
///* GL_EXT_unpack_subimage */
//#ifndef GL_EXT_unpack_subimage
const GL_UNPACK_ROW_LENGTH_EXT = 0x0CF2
const GL_UNPACK_SKIP_ROWS_EXT = 0x0CF3
const GL_UNPACK_SKIP_PIXELS_EXT = 0x0CF4

//#endif
//
///*------------------------------------------------------------------------*
// * DMP extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_DMP_shader_binary */
//#ifndef GL_DMP_shader_binary
const GL_SHADER_BINARY_DMP = 0x9250

//#endif
//
///*------------------------------------------------------------------------*
// * FJ extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_FJ_shader_binary_GCCSO */
//#ifndef GL_FJ_shader_binary_GCCSO
const GL_GCCSO_SHADER_BINARY_FJ = 0x9260

//#endif
//
///*------------------------------------------------------------------------*
// * IMG extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_IMG_program_binary */
//#ifndef GL_IMG_program_binary
const GL_SGX_PROGRAM_BINARY_IMG = 0x9130

//#endif
//
///* GL_IMG_read_format */
//#ifndef GL_IMG_read_format
const GL_BGRA_IMG = 0x80E1
const GL_UNSIGNED_SHORT_4_4_4_4_REV_IMG = 0x8365

//#endif
//
///* GL_IMG_shader_binary */
//#ifndef GL_IMG_shader_binary
const GL_SGX_BINARY_IMG = 0x8C0A

//#endif
//
///* GL_IMG_texture_compression_pvrtc */
//#ifndef GL_IMG_texture_compression_pvrtc
const GL_COMPRESSED_RGB_PVRTC_4BPPV1_IMG = 0x8C00
const GL_COMPRESSED_RGB_PVRTC_2BPPV1_IMG = 0x8C01
const GL_COMPRESSED_RGBA_PVRTC_4BPPV1_IMG = 0x8C02
const GL_COMPRESSED_RGBA_PVRTC_2BPPV1_IMG = 0x8C03

//#endif
//
///* GL_IMG_texture_compression_pvrtc2 */
//#ifndef GL_IMG_texture_compression_pvrtc2
const GL_COMPRESSED_RGBA_PVRTC_2BPPV2_IMG = 0x9137
const GL_COMPRESSED_RGBA_PVRTC_4BPPV2_IMG = 0x9138

//#endif
//
///* GL_IMG_multisampled_render_to_texture */
//#ifndef GL_IMG_multisampled_render_to_texture
const GL_RENDERBUFFER_SAMPLES_IMG = 0x9133
const GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_IMG = 0x9134
const GL_MAX_SAMPLES_IMG = 0x9135
const GL_TEXTURE_SAMPLES_IMG = 0x9136

//#endif
//
///*------------------------------------------------------------------------*
// * NV extension tokens
// *------------------------------------------------------------------------*/
//
///* GL_NV_coverage_sample */
//#ifndef GL_NV_coverage_sample
const GL_COVERAGE_COMPONENT_NV = 0x8ED0
const GL_COVERAGE_COMPONENT4_NV = 0x8ED1
const GL_COVERAGE_ATTACHMENT_NV = 0x8ED2
const GL_COVERAGE_BUFFERS_NV = 0x8ED3
const GL_COVERAGE_SAMPLES_NV = 0x8ED4
const GL_COVERAGE_ALL_FRAGMENTS_NV = 0x8ED5
const GL_COVERAGE_EDGE_FRAGMENTS_NV = 0x8ED6
const GL_COVERAGE_AUTOMATIC_NV = 0x8ED7
const GL_COVERAGE_BUFFER_BIT_NV = 0x00008000

//#endif
//
///* GL_NV_depth_nonlinear */
//#ifndef GL_NV_depth_nonlinear
const GL_DEPTH_COMPONENT16_NONLINEAR_NV = 0x8E2C

//#endif

///* GL_NV_draw_buffers */
//#ifndef GL_NV_draw_buffers
const GL_MAX_DRAW_BUFFERS_NV = 0x8824
const GL_DRAW_BUFFER0_NV = 0x8825
const GL_DRAW_BUFFER1_NV = 0x8826
const GL_DRAW_BUFFER2_NV = 0x8827
const GL_DRAW_BUFFER3_NV = 0x8828
const GL_DRAW_BUFFER4_NV = 0x8829
const GL_DRAW_BUFFER5_NV = 0x882A
const GL_DRAW_BUFFER6_NV = 0x882B
const GL_DRAW_BUFFER7_NV = 0x882C
const GL_DRAW_BUFFER8_NV = 0x882D
const GL_DRAW_BUFFER9_NV = 0x882E
const GL_DRAW_BUFFER10_NV = 0x882F
const GL_DRAW_BUFFER11_NV = 0x8830
const GL_DRAW_BUFFER12_NV = 0x8831
const GL_DRAW_BUFFER13_NV = 0x8832
const GL_DRAW_BUFFER14_NV = 0x8833
const GL_DRAW_BUFFER15_NV = 0x8834
const GL_COLOR_ATTACHMENT0_NV = 0x8CE0
const GL_COLOR_ATTACHMENT1_NV = 0x8CE1
const GL_COLOR_ATTACHMENT2_NV = 0x8CE2
const GL_COLOR_ATTACHMENT3_NV = 0x8CE3
const GL_COLOR_ATTACHMENT4_NV = 0x8CE4
const GL_COLOR_ATTACHMENT5_NV = 0x8CE5
const GL_COLOR_ATTACHMENT6_NV = 0x8CE6
const GL_COLOR_ATTACHMENT7_NV = 0x8CE7
const GL_COLOR_ATTACHMENT8_NV = 0x8CE8
const GL_COLOR_ATTACHMENT9_NV = 0x8CE9
const GL_COLOR_ATTACHMENT10_NV = 0x8CEA
const GL_COLOR_ATTACHMENT11_NV = 0x8CEB
const GL_COLOR_ATTACHMENT12_NV = 0x8CEC
const GL_COLOR_ATTACHMENT13_NV = 0x8CED
const GL_COLOR_ATTACHMENT14_NV = 0x8CEE
const GL_COLOR_ATTACHMENT15_NV = 0x8CEF

//#endif
//
///* GL_NV_draw_instanced */
///* No new tokens introduced by this extension. */
//
///* GL_NV_fbo_color_attachments */
//#ifndef GL_NV_fbo_color_attachments
const GL_MAX_COLOR_ATTACHMENTS_NV = 0x8CDF

///* GL_COLOR_ATTACHMENT{0-15}_NV defined in GL_NV_draw_buffers already. */
//#endif
//
///* GL_NV_fence */
//#ifndef GL_NV_fence
const GL_ALL_COMPLETED_NV = 0x84F2
const GL_FENCE_STATUS_NV = 0x84F3
const GL_FENCE_CONDITION_NV = 0x84F4

//#endif
//
///* GL_NV_framebuffer_blit */
//#ifndef GL_NV_framebuffer_blit
const GL_READ_FRAMEBUFFER_NV = 0x8CA8
const GL_DRAW_FRAMEBUFFER_NV = 0x8CA9
const GL_DRAW_FRAMEBUFFER_BINDING_NV = 0x8CA6
const GL_READ_FRAMEBUFFER_BINDING_NV = 0x8CAA

//#endif
//
///* GL_NV_framebuffer_multisample */
//#ifndef GL_NV_framebuffer_multisample
const GL_RENDERBUFFER_SAMPLES_NV = 0x8CAB
const GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE_NV = 0x8D56
const GL_MAX_SAMPLES_NV = 0x8D57

//#endif
//
///* GL_NV_generate_mipmap_sRGB */
///* No new tokens introduced by this extension. */
//
///* GL_NV_instanced_arrays */
//#ifndef GL_NV_instanced_arrays
const GL_VERTEX_ATTRIB_ARRAY_DIVISOR_NV = 0x88FE

//#endif
//
///* GL_NV_read_buffer */
//#ifndef GL_NV_read_buffer
const GL_READ_BUFFER_NV = 0x0C02

//#endif

/* GL_NV_read_buffer_front */
/* No new tokens introduced by this extension. */

/* GL_NV_read_depth */
/* No new tokens introduced by this extension. */

/* GL_NV_read_depth_stencil */
/* No new tokens introduced by this extension. */

/* GL_NV_read_stencil */
/* No new tokens introduced by this extension. */

/* GL_NV_shadow_samplers_array */
//#ifndef GL_NV_shadow_samplers_array
const GL_SAMPLER_2D_ARRAY_SHADOW_NV = 0x8DC4

//#endif

/* GL_NV_shadow_samplers_cube */
//#ifndef GL_NV_shadow_samplers_cube
const GL_SAMPLER_CUBE_SHADOW_NV = 0x8DC5

//#endif

/* GL_NV_sRGB_formats */
//#ifndef GL_NV_sRGB_formats
const GL_SLUMINANCE_NV = 0x8C46
const GL_SLUMINANCE_ALPHA_NV = 0x8C44
const GL_SRGB8_NV = 0x8C41
const GL_SLUMINANCE8_NV = 0x8C47
const GL_SLUMINANCE8_ALPHA8_NV = 0x8C45
const GL_COMPRESSED_SRGB_S3TC_DXT1_NV = 0x8C4C
const GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT1_NV = 0x8C4D
const GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT3_NV = 0x8C4E
const GL_COMPRESSED_SRGB_ALPHA_S3TC_DXT5_NV = 0x8C4F
const GL_ETC1_SRGB8_NV = 0x88EE

//#endif

/* GL_NV_texture_border_clamp */
//#ifndef GL_NV_texture_border_clamp
const GL_TEXTURE_BORDER_COLOR_NV = 0x1004
const GL_CLAMP_TO_BORDER_NV = 0x812D

//#endif

/* GL_NV_texture_compression_s3tc_update */
/* No new tokens introduced by this extension. */

/* GL_NV_texture_npot_2D_mipmap */
/* No new tokens introduced by this extension. */

/*------------------------------------------------------------------------*
 * QCOM extension tokens
 *------------------------------------------------------------------------*/

/* GL_QCOM_alpha_test */
//#ifndef GL_QCOM_alpha_test
const GL_ALPHA_TEST_QCOM = 0x0BC0
const GL_ALPHA_TEST_FUNC_QCOM = 0x0BC1
const GL_ALPHA_TEST_REF_QCOM = 0x0BC2

//#endif
//
///* GL_QCOM_binning_control */
//#ifndef GL_QCOM_binning_control
const GL_BINNING_CONTROL_HINT_QCOM = 0x8FB0
const GL_CPU_OPTIMIZED_QCOM = 0x8FB1
const GL_GPU_OPTIMIZED_QCOM = 0x8FB2
const GL_RENDER_DIRECT_TO_FRAMEBUFFER_QCOM = 0x8FB3

//#endif
//
///* GL_QCOM_driver_control */
///* No new tokens introduced by this extension. */
//
///* GL_QCOM_extended_get */
//#ifndef GL_QCOM_extended_get
const GL_TEXTURE_WIDTH_QCOM = 0x8BD2
const GL_TEXTURE_HEIGHT_QCOM = 0x8BD3
const GL_TEXTURE_DEPTH_QCOM = 0x8BD4
const GL_TEXTURE_INTERNAL_FORMAT_QCOM = 0x8BD5
const GL_TEXTURE_FORMAT_QCOM = 0x8BD6
const GL_TEXTURE_TYPE_QCOM = 0x8BD7
const GL_TEXTURE_IMAGE_VALID_QCOM = 0x8BD8
const GL_TEXTURE_NUM_LEVELS_QCOM = 0x8BD9
const GL_TEXTURE_TARGET_QCOM = 0x8BDA
const GL_TEXTURE_OBJECT_VALID_QCOM = 0x8BDB
const GL_STATE_RESTORE = 0x8BDC

//#endif
//
///* GL_QCOM_extended_get2 */
///* No new tokens introduced by this extension. */
//
///* GL_QCOM_perfmon_global_mode */
//#ifndef GL_QCOM_perfmon_global_mode
const GL_PERFMON_GLOBAL_MODE_QCOM = 0x8FA0

//#endif
//
///* GL_QCOM_writeonly_rendering */
//#ifndef GL_QCOM_writeonly_rendering
const GL_WRITEONLY_RENDERING_QCOM = 0x8823

//#endif
//
///* GL_QCOM_tiled_rendering */
//#ifndef GL_QCOM_tiled_rendering
const GL_COLOR_BUFFER_BIT0_QCOM = 0x00000001
const GL_COLOR_BUFFER_BIT1_QCOM = 0x00000002
const GL_COLOR_BUFFER_BIT2_QCOM = 0x00000004
const GL_COLOR_BUFFER_BIT3_QCOM = 0x00000008
const GL_COLOR_BUFFER_BIT4_QCOM = 0x00000010
const GL_COLOR_BUFFER_BIT5_QCOM = 0x00000020
const GL_COLOR_BUFFER_BIT6_QCOM = 0x00000040
const GL_COLOR_BUFFER_BIT7_QCOM = 0x00000080
const GL_DEPTH_BUFFER_BIT0_QCOM = 0x00000100
const GL_DEPTH_BUFFER_BIT1_QCOM = 0x00000200
const GL_DEPTH_BUFFER_BIT2_QCOM = 0x00000400
const GL_DEPTH_BUFFER_BIT3_QCOM = 0x00000800
const GL_DEPTH_BUFFER_BIT4_QCOM = 0x00001000
const GL_DEPTH_BUFFER_BIT5_QCOM = 0x00002000
const GL_DEPTH_BUFFER_BIT6_QCOM = 0x00004000
const GL_DEPTH_BUFFER_BIT7_QCOM = 0x00008000
const GL_STENCIL_BUFFER_BIT0_QCOM = 0x00010000
const GL_STENCIL_BUFFER_BIT1_QCOM = 0x00020000
const GL_STENCIL_BUFFER_BIT2_QCOM = 0x00040000
const GL_STENCIL_BUFFER_BIT3_QCOM = 0x00080000
const GL_STENCIL_BUFFER_BIT4_QCOM = 0x00100000
const GL_STENCIL_BUFFER_BIT5_QCOM = 0x00200000
const GL_STENCIL_BUFFER_BIT6_QCOM = 0x00400000
const GL_STENCIL_BUFFER_BIT7_QCOM = 0x00800000
const GL_MULTISAMPLE_BUFFER_BIT0_QCOM = 0x01000000
const GL_MULTISAMPLE_BUFFER_BIT1_QCOM = 0x02000000
const GL_MULTISAMPLE_BUFFER_BIT2_QCOM = 0x04000000
const GL_MULTISAMPLE_BUFFER_BIT3_QCOM = 0x08000000
const GL_MULTISAMPLE_BUFFER_BIT4_QCOM = 0x10000000
const GL_MULTISAMPLE_BUFFER_BIT5_QCOM = 0x20000000
const GL_MULTISAMPLE_BUFFER_BIT6_QCOM = 0x40000000
const GL_MULTISAMPLE_BUFFER_BIT7_QCOM = 0x80000000

//#endif

/*------------------------------------------------------------------------*
 * VIV extension tokens
 *------------------------------------------------------------------------*/

/* GL_VIV_shader_binary */
//#ifndef GL_VIV_shader_binary
const GL_SHADER_BINARY_VIV = 0x8FC4

//#endif
//
///*------------------------------------------------------------------------*
// * End of extension tokens, start of corresponding extension functions
// *------------------------------------------------------------------------*/
//
///*------------------------------------------------------------------------*
// * OES extension functions
// *------------------------------------------------------------------------*/
//
///* GL_OES_compressed_ETC1_RGB8_texture */
//#ifndef GL_OES_compressed_ETC1_RGB8_texture
const GL_OES_compressed_ETC1_RGB8_texture = 1

//#endif
//
///* GL_OES_compressed_paletted_texture */
//#ifndef GL_OES_compressed_paletted_texture
const GL_OES_compressed_paletted_texture = 1

//#endif
//
///* GL_OES_depth24 */
//#ifndef GL_OES_depth24
const GL_OES_depth24 = 1

//#endif
//
///* GL_OES_depth32 */
//#ifndef GL_OES_depth32
const GL_OES_depth32 = 1

//#endif
//
///* GL_OES_depth_texture */
//#ifndef GL_OES_depth_texture
const GL_OES_depth_texture = 1

//#endif
//
///* GL_OES_EGL_image */
//#ifndef GL_OES_EGL_image
const GL_OES_EGL_image = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glEGLImageTargetTexture2DOES (GLenum target, GLeglImageOES image);
//GL_APICALL void GL_APIENTRY glEGLImageTargetRenderbufferStorageOES (GLenum target, GLeglImageOES image);
//#endif
//typedef void (GL_APIENTRYP PFNGLEGLIMAGETARGETTEXTURE2DOESPROC) (GLenum target, GLeglImageOES image);
//todo
//typedef void (GL_APIENTRYP PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOESPROC) (GLenum target, GLeglImageOES image);
//todo
//#endif

/* GL_OES_EGL_image_external */
//#ifndef GL_OES_EGL_image_external
const GL_OES_EGL_image_external = 1

/* glEGLImageTargetTexture2DOES defined in GL_OES_EGL_image already. */
//#endif
//
///* GL_OES_element_index_uint */
//#ifndef GL_OES_element_index_uint
const GL_OES_element_index_uint = 1

//#endif
//
///* GL_OES_fbo_render_mipmap */
//#ifndef GL_OES_fbo_render_mipmap
const GL_OES_fbo_render_mipmap = 1

//#endif
//
///* GL_OES_fragment_precision_high */
//#ifndef GL_OES_fragment_precision_high
const GL_OES_fragment_precision_high = 1

//#endif
//
///* GL_OES_get_program_binary */
//#ifndef GL_OES_get_program_binary
const GL_OES_get_program_binary = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glGetProgramBinaryOES (GLuint program, GLsizei bufSize, GLsizei *length, GLenum *binaryFormat, GLvoid *binary);
//GL_APICALL void GL_APIENTRY glProgramBinaryOES (GLuint program, GLenum binaryFormat, const GLvoid *binary, GLint length);
//#endif
//typedef void (GL_APIENTRYP PFNGLGETPROGRAMBINARYOESPROC) (GLuint program, GLsizei bufSize, GLsizei *length, GLenum *binaryFormat, GLvoid *binary);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMBINARYOESPROC) (GLuint program, GLenum binaryFormat, const GLvoid *binary, GLint length);
//todo
//#endif

/* GL_OES_mapbuffer */
//#ifndef GL_OES_mapbuffer
const GL_OES_mapbuffer = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void* GL_APIENTRY glMapBufferOES (GLenum target, GLenum access);
//todo
func glMapBufferOES() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("glMapBufferOES").Call()
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

//GL_APICALL GLboolean GL_APIENTRY glUnmapBufferOES (GLenum target);
//todo
func glUnmapBufferOES() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("glUnmapBufferOES").Call()
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

//GL_APICALL void GL_APIENTRY glGetBufferPointervOES (GLenum target, GLenum pname, GLvoid **params);
//todo
func glGetBufferPointervOES() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("glGetBufferPointervOES").Call()
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

//#endif
//typedef void* (GL_APIENTRYP PFNGLMAPBUFFEROESPROC) (GLenum target, GLenum access);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLUNMAPBUFFEROESPROC) (GLenum target);
//todo
//typedef void (GL_APIENTRYP PFNGLGETBUFFERPOINTERVOESPROC) (GLenum target, GLenum pname, GLvoid **params);
//todo
//#endif

/* GL_OES_packed_depth_stencil */
//#ifndef GL_OES_packed_depth_stencil
const GL_OES_packed_depth_stencil = 1

//#endif
//
///* GL_OES_required_internalformat */
//#ifndef GL_OES_required_internalformat
const GL_OES_required_internalformat = 1

//#endif
//
///* GL_OES_rgb8_rgba8 */
//#ifndef GL_OES_rgb8_rgba8
const GL_OES_rgb8_rgba8 = 1

//#endif
//
///* GL_OES_standard_derivatives */
//#ifndef GL_OES_standard_derivatives
const GL_OES_standard_derivatives = 1

//#endif
//
///* GL_OES_stencil1 */
//#ifndef GL_OES_stencil1
const GL_OES_stencil1 = 1

//#endif
//
///* GL_OES_stencil4 */
//#ifndef GL_OES_stencil4
const GL_OES_stencil4 = 1

//#endif
//
//#ifndef GL_OES_surfaceless_context
const GL_OES_surfaceless_context = 1

//#endif

/* GL_OES_texture_3D */
//#ifndef GL_OES_texture_3D
const GL_OES_texture_3D = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glTexImage3DOES (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
//GL_APICALL void GL_APIENTRY glTexSubImage3DOES (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid* pixels);
//GL_APICALL void GL_APIENTRY glCopyTexSubImage3DOES (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glCompressedTexImage3DOES (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid* data);
//GL_APICALL void GL_APIENTRY glCompressedTexSubImage3DOES (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid* data);
//GL_APICALL void GL_APIENTRY glFramebufferTexture3DOES (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
//#endif
//typedef void (GL_APIENTRYP PFNGLTEXIMAGE3DOESPROC) (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLenum format, GLenum type, const GLvoid* pixels);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXSUBIMAGE3DOESPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, const GLvoid* pixels);
//todo
//typedef void (GL_APIENTRYP PFNGLCOPYTEXSUBIMAGE3DOESPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLint x, GLint y, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLCOMPRESSEDTEXIMAGE3DOESPROC) (GLenum target, GLint level, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth, GLint border, GLsizei imageSize, const GLvoid* data);
//todo
//typedef void (GL_APIENTRYP PFNGLCOMPRESSEDTEXSUBIMAGE3DOESPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLsizei imageSize, const GLvoid* data);
//todo
//typedef void (GL_APIENTRYP PFNGLFRAMEBUFFERTEXTURE3DOESPROC) (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLint zoffset);
//todo
//#endif

/* GL_OES_texture_float */
//#ifndef GL_OES_texture_float
const GL_OES_texture_float = 1

//#endif
//
///* GL_OES_texture_float_linear */
//#ifndef GL_OES_texture_float_linear
const GL_OES_texture_float_linear = 1

//#endif
//
///* GL_OES_texture_half_float */
//#ifndef GL_OES_texture_half_float
const GL_OES_texture_half_float = 1

//#endif
//
///* GL_OES_texture_half_float_linear */
//#ifndef GL_OES_texture_half_float_linear
const GL_OES_texture_half_float_linear = 1

//#endif
//
///* GL_OES_texture_npot */
//#ifndef GL_OES_texture_npot
const GL_OES_texture_npot = 1

//#endif
//
///* GL_OES_vertex_array_object */
//#ifndef GL_OES_vertex_array_object
const GL_OES_vertex_array_object = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glBindVertexArrayOES (GLuint array);
//GL_APICALL void GL_APIENTRY glDeleteVertexArraysOES (GLsizei n, const GLuint *arrays);
//GL_APICALL void GL_APIENTRY glGenVertexArraysOES (GLsizei n, GLuint *arrays);
//GL_APICALL GLboolean GL_APIENTRY glIsVertexArrayOES (GLuint array);
//#endif
//typedef void (GL_APIENTRYP PFNGLBINDVERTEXARRAYOESPROC) (GLuint array);
//todo
//typedef void (GL_APIENTRYP PFNGLDELETEVERTEXARRAYSOESPROC) (GLsizei n, const GLuint *arrays);
//todo
//typedef void (GL_APIENTRYP PFNGLGENVERTEXARRAYSOESPROC) (GLsizei n, GLuint *arrays);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLISVERTEXARRAYOESPROC) (GLuint array);
//todo
//#endif

/* GL_OES_vertex_half_float */
//#ifndef GL_OES_vertex_half_float
const GL_OES_vertex_half_float = 1

//#endif
//
///* GL_OES_vertex_type_10_10_10_2 */
//#ifndef GL_OES_vertex_type_10_10_10_2
const GL_OES_vertex_type_10_10_10_2 = 1

//#endif
//
///*------------------------------------------------------------------------*
// * KHR extension functions
// *------------------------------------------------------------------------*/
//
//#ifndef GL_KHR_debug
const GL_KHR_debug = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDebugMessageControlKHR (GLenum source, GLenum type, GLenum severity, GLsizei count, const GLuint *ids, GLboolean enabled);
//GL_APICALL void GL_APIENTRY glDebugMessageInsertKHR (GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar *buf);
//GL_APICALL void GL_APIENTRY glDebugMessageCallbackKHR (GLDEBUGPROCKHR callback, const void *userParam);
//GL_APICALL GLuint GL_APIENTRY glGetDebugMessageLogKHR (GLuint count, GLsizei bufsize, GLenum *sources, GLenum *types, GLuint *ids, GLenum *severities, GLsizei *lengths, GLchar *messageLog);
//GL_APICALL void GL_APIENTRY glPushDebugGroupKHR (GLenum source, GLuint id, GLsizei length, const GLchar *message);
//GL_APICALL void GL_APIENTRY glPopDebugGroupKHR (void);
//GL_APICALL void GL_APIENTRY glObjectLabelKHR (GLenum identifier, GLuint name, GLsizei length, const GLchar *label);
//GL_APICALL void GL_APIENTRY glGetObjectLabelKHR (GLenum identifier, GLuint name, GLsizei bufSize, GLsizei *length, GLchar *label);
//GL_APICALL void GL_APIENTRY glObjectPtrLabelKHR (const void *ptr, GLsizei length, const GLchar *label);
//GL_APICALL void GL_APIENTRY glGetObjectPtrLabelKHR (const void *ptr, GLsizei bufSize, GLsizei *length, GLchar *label);
//GL_APICALL void GL_APIENTRY glGetPointervKHR (GLenum pname, GLvoid **params);
//#endif
//typedef void (GL_APIENTRYP PFNGLDEBUGMESSAGECONTROLKHRPROC) (GLenum source, GLenum type, GLenum severity, GLsizei count, const GLuint *ids, GLboolean enabled);
//todo
//typedef void (GL_APIENTRYP PFNGLDEBUGMESSAGEINSERTKHRPROC) (GLenum source, GLenum type, GLuint id, GLenum severity, GLsizei length, const GLchar *buf);
//todo
//typedef void (GL_APIENTRYP PFNGLDEBUGMESSAGECALLBACKKHRPROC) (GLDEBUGPROCKHR callback, const void *userParam);
//todo
//typedef GLuint (GL_APIENTRYP PFNGLGETDEBUGMESSAGELOGKHRPROC) (GLuint count, GLsizei bufsize, GLenum *sources, GLenum *types, GLuint *ids, GLenum *severities, GLsizei *lengths, GLchar *messageLog);
//todo
//typedef void (GL_APIENTRYP PFNGLPUSHDEBUGGROUPKHRPROC) (GLenum source, GLuint id, GLsizei length, const GLchar *message);
//todo
//typedef void (GL_APIENTRYP PFNGLPOPDEBUGGROUPKHRPROC) (void);
//todo
//typedef void (GL_APIENTRYP PFNGLOBJECTLABELKHRPROC) (GLenum identifier, GLuint name, GLsizei length, const GLchar *label);
//todo
//typedef void (GL_APIENTRYP PFNGLGETOBJECTLABELKHRPROC) (GLenum identifier, GLuint name, GLsizei bufSize, GLsizei *length, GLchar *label);
//todo
//typedef void (GL_APIENTRYP PFNGLOBJECTPTRLABELKHRPROC) (const void *ptr, GLsizei length, const GLchar *label);
//todo
//typedef void (GL_APIENTRYP PFNGLGETOBJECTPTRLABELKHRPROC) (const void *ptr, GLsizei bufSize, GLsizei *length, GLchar *label);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPOINTERVKHRPROC) (GLenum pname, GLvoid **params);
//#endif
//
//#ifndef GL_KHR_texture_compression_astc_ldr
const GL_KHR_texture_compression_astc_ldr = 1

//#endif
//
//
///*------------------------------------------------------------------------*
// * AMD extension functions
// *------------------------------------------------------------------------*/
//
///* GL_AMD_compressed_3DC_texture */
//#ifndef GL_AMD_compressed_3DC_texture
const GL_AMD_compressed_3DC_texture = 1

//#endif
//
///* GL_AMD_compressed_ATC_texture */
//#ifndef GL_AMD_compressed_ATC_texture
const GL_AMD_compressed_ATC_texture = 1

//#endif
//
///* AMD_performance_monitor */
//#ifndef GL_AMD_performance_monitor
const GL_AMD_performance_monitor = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glGetPerfMonitorGroupsAMD (GLint *numGroups, GLsizei groupsSize, GLuint *groups);
//GL_APICALL void GL_APIENTRY glGetPerfMonitorCountersAMD (GLuint group, GLint *numCounters, GLint *maxActiveCounters, GLsizei counterSize, GLuint *counters);
//GL_APICALL void GL_APIENTRY glGetPerfMonitorGroupStringAMD (GLuint group, GLsizei bufSize, GLsizei *length, GLchar *groupString);
//GL_APICALL void GL_APIENTRY glGetPerfMonitorCounterStringAMD (GLuint group, GLuint counter, GLsizei bufSize, GLsizei *length, GLchar *counterString);
//GL_APICALL void GL_APIENTRY glGetPerfMonitorCounterInfoAMD (GLuint group, GLuint counter, GLenum pname, GLvoid *data);
//GL_APICALL void GL_APIENTRY glGenPerfMonitorsAMD (GLsizei n, GLuint *monitors);
//GL_APICALL void GL_APIENTRY glDeletePerfMonitorsAMD (GLsizei n, GLuint *monitors);
//GL_APICALL void GL_APIENTRY glSelectPerfMonitorCountersAMD (GLuint monitor, GLboolean enable, GLuint group, GLint numCounters, GLuint *countersList);
//GL_APICALL void GL_APIENTRY glBeginPerfMonitorAMD (GLuint monitor);
//GL_APICALL void GL_APIENTRY glEndPerfMonitorAMD (GLuint monitor);
//GL_APICALL void GL_APIENTRY glGetPerfMonitorCounterDataAMD (GLuint monitor, GLenum pname, GLsizei dataSize, GLuint *data, GLint *bytesWritten);
//#endif
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORGROUPSAMDPROC) (GLint *numGroups, GLsizei groupsSize, GLuint *groups);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORCOUNTERSAMDPROC) (GLuint group, GLint *numCounters, GLint *maxActiveCounters, GLsizei counterSize, GLuint *counters);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORGROUPSTRINGAMDPROC) (GLuint group, GLsizei bufSize, GLsizei *length, GLchar *groupString);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORCOUNTERSTRINGAMDPROC) (GLuint group, GLuint counter, GLsizei bufSize, GLsizei *length, GLchar *counterString);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORCOUNTERINFOAMDPROC) (GLuint group, GLuint counter, GLenum pname, GLvoid *data);
//todo
//typedef void (GL_APIENTRYP PFNGLGENPERFMONITORSAMDPROC) (GLsizei n, GLuint *monitors);
//todo
//typedef void (GL_APIENTRYP PFNGLDELETEPERFMONITORSAMDPROC) (GLsizei n, GLuint *monitors);
//todo
//typedef void (GL_APIENTRYP PFNGLSELECTPERFMONITORCOUNTERSAMDPROC) (GLuint monitor, GLboolean enable, GLuint group, GLint numCounters, GLuint *countersList);
//todo
//typedef void (GL_APIENTRYP PFNGLBEGINPERFMONITORAMDPROC) (GLuint monitor);
//todo
//typedef void (GL_APIENTRYP PFNGLENDPERFMONITORAMDPROC) (GLuint monitor);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPERFMONITORCOUNTERDATAAMDPROC) (GLuint monitor, GLenum pname, GLsizei dataSize, GLuint *data, GLint *bytesWritten);
//todo
//#endif

/* GL_AMD_program_binary_Z400 */
//#ifndef GL_AMD_program_binary_Z400
const GL_AMD_program_binary_Z400 = 1

//#endif
//
///*------------------------------------------------------------------------*
// * ANGLE extension functions
// *------------------------------------------------------------------------*/
//
///* GL_ANGLE_depth_texture */
//#ifndef GL_ANGLE_depth_texture
const GL_ANGLE_depth_texture = 1

//#endif
//
///* GL_ANGLE_framebuffer_blit */
//#ifndef GL_ANGLE_framebuffer_blit
const GL_ANGLE_framebuffer_blit = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glBlitFramebufferANGLE (GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
//#endif
//typedef void (GL_APIENTRYP PFNGLBLITFRAMEBUFFERANGLEPROC) (GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
//todo
//#endif

/* GL_ANGLE_framebuffer_multisample */
//#ifndef GL_ANGLE_framebuffer_multisample
const GL_ANGLE_framebuffer_multisample = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glRenderbufferStorageMultisampleANGLE (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//#endif
//typedef void (GL_APIENTRYP PFNGLRENDERBUFFERSTORAGEMULTISAMPLEANGLEPROC) (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//#endif
//
//#ifndef GL_ANGLE_instanced_arrays
const GL_ANGLE_instanced_arrays = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDrawArraysInstancedANGLE (GLenum mode, GLint first, GLsizei count, GLsizei primcount);
//GL_APICALL void GL_APIENTRY glDrawElementsInstancedANGLE (GLenum mode, GLsizei count, GLenum type, const void *indices, GLsizei primcount);
//GL_APICALL void GL_APIENTRY glVertexAttribDivisorANGLE (GLuint index, GLuint divisor);
//#endif
//typedef void (GL_APIENTRYP PFNGLDRAWARRAYSINSTANCEDANGLEPROC) (GLenum mode, GLint first, GLsizei count, GLsizei primcount);
//todo
//typedef void (GL_APIENTRYP PFNGLDRAWELEMENTSINSTANCEDANGLEPROC) (GLenum mode, GLsizei count, GLenum type, const void *indices, GLsizei primcount);
//todo
//typedef void (GL_APIENTRYP PFNGLVERTEXATTRIBDIVISORANGLEPROC) (GLuint index, GLuint divisor);
//todo
//#endif

///* GL_ANGLE_pack_reverse_row_order */
//#ifndef GL_ANGLE_pack_reverse_row_order
const GL_ANGLE_pack_reverse_row_order = 1

//#endif
//
///* GL_ANGLE_program_binary */
//#ifndef GL_ANGLE_program_binary
const GL_ANGLE_program_binary = 1

//#endif
//
///* GL_ANGLE_texture_compression_dxt3 */
//#ifndef GL_ANGLE_texture_compression_dxt3
const GL_ANGLE_texture_compression_dxt3 = 1

//#endif
//
///* GL_ANGLE_texture_compression_dxt5 */
//#ifndef GL_ANGLE_texture_compression_dxt5
const GL_ANGLE_texture_compression_dxt5 = 1

//#endif
//
///* GL_ANGLE_texture_usage */
//#ifndef GL_ANGLE_texture_usage
const GL_ANGLE_texture_usage = 1

//#endif
//
//#ifndef GL_ANGLE_translated_shader_source
const GL_ANGLE_translated_shader_source = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glGetTranslatedShaderSourceANGLE (GLuint shader, GLsizei bufsize, GLsizei *length, GLchar *source);
//#endif
//typedef void (GL_APIENTRYP PFNGLGETTRANSLATEDSHADERSOURCEANGLEPROC) (GLuint shader, GLsizei bufsize, GLsizei *length, GLchar *source);
//todo
//#endif

/*------------------------------------------------------------------------*
 * APPLE extension functions
 *------------------------------------------------------------------------*/

/* GL_APPLE_copy_texture_levels */
//#ifndef GL_APPLE_copy_texture_levels
const GL_APPLE_copy_texture_levels = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glCopyTextureLevelsAPPLE (GLuint destinationTexture, GLuint sourceTexture, GLint sourceBaseLevel, GLsizei sourceLevelCount);
//#endif
//typedef void (GL_APIENTRYP PFNGLCOPYTEXTURELEVELSAPPLEPROC) (GLuint destinationTexture, GLuint sourceTexture, GLint sourceBaseLevel, GLsizei sourceLevelCount);
//todo
//#endif
//
///* GL_APPLE_framebuffer_multisample */
//#ifndef GL_APPLE_framebuffer_multisample
const GL_APPLE_framebuffer_multisample = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glRenderbufferStorageMultisampleAPPLE (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glResolveMultisampleFramebufferAPPLE (void);
//#endif /* GL_GLEXT_PROTOTYPES */
//typedef void (GL_APIENTRYP PFNGLRENDERBUFFERSTORAGEMULTISAMPLEAPPLEPROC) (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLRESOLVEMULTISAMPLEFRAMEBUFFERAPPLEPROC) (void);
//todo
//#endif
//
///* GL_APPLE_rgb_422 */
//#ifndef GL_APPLE_rgb_422
const GL_APPLE_rgb_422 = 1

//#endif

/* GL_APPLE_sync */
//#ifndef GL_APPLE_sync
const GL_APPLE_sync = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL GLsync GL_APIENTRY glFenceSyncAPPLE (GLenum condition, GLbitfield flags);
//GL_APICALL GLboolean GL_APIENTRY glIsSyncAPPLE (GLsync sync);
//GL_APICALL void GL_APIENTRY glDeleteSyncAPPLE (GLsync sync);
//GL_APICALL GLenum GL_APIENTRY glClientWaitSyncAPPLE (GLsync sync, GLbitfield flags, GLuint64 timeout);
//GL_APICALL void GL_APIENTRY glWaitSyncAPPLE (GLsync sync, GLbitfield flags, GLuint64 timeout);
//GL_APICALL void GL_APIENTRY glGetInteger64vAPPLE (GLenum pname, GLint64 *params);
//GL_APICALL void GL_APIENTRY glGetSyncivAPPLE (GLsync sync, GLenum pname, GLsizei bufSize, GLsizei *length, GLint *values);
//#endif
//typedef GLsync (GL_APIENTRYP PFNGLFENCESYNCAPPLEPROC) (GLenum condition, GLbitfield flags);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLISSYNCAPPLEPROC) (GLsync sync);
//todo
//typedef void (GL_APIENTRYP PFNGLDELETESYNCAPPLEPROC) (GLsync sync);
//todo
//typedef GLenum (GL_APIENTRYP PFNGLCLIENTWAITSYNCAPPLEPROC) (GLsync sync, GLbitfield flags, GLuint64 timeout);
//todo
//typedef void (GL_APIENTRYP PFNGLWAITSYNCAPPLEPROC) (GLsync sync, GLbitfield flags, GLuint64 timeout);
//todo
//typedef void (GL_APIENTRYP PFNGLGETINTEGER64VAPPLEPROC) (GLenum pname, GLint64 *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETSYNCIVAPPLEPROC) (GLsync sync, GLenum pname, GLsizei bufSize, GLsizei *length, GLint *values);
//todo
//#endif

/* GL_APPLE_texture_format_BGRA8888 */
//#ifndef GL_APPLE_texture_format_BGRA8888
const GL_APPLE_texture_format_BGRA8888 = 1

//#endif
//
///* GL_APPLE_texture_max_level */
//#ifndef GL_APPLE_texture_max_level
const GL_APPLE_texture_max_level = 1

//#endif
//
///*------------------------------------------------------------------------*
// * ARM extension functions
// *------------------------------------------------------------------------*/
//
///* GL_ARM_mali_program_binary */
//#ifndef GL_ARM_mali_program_binary
const GL_ARM_mali_program_binary = 1

//#endif
//
///* GL_ARM_mali_shader_binary */
//#ifndef GL_ARM_mali_shader_binary
const GL_ARM_mali_shader_binary = 1

//#endif
//
///* GL_ARM_rgba8 */
//#ifndef GL_ARM_rgba8
const GL_ARM_rgba8 = 1

//#endif
//
///*------------------------------------------------------------------------*
// * EXT extension functions
// *------------------------------------------------------------------------*/
//
///* GL_EXT_blend_minmax */
//#ifndef GL_EXT_blend_minmax
const GL_EXT_blend_minmax = 1

//#endif
//
///* GL_EXT_color_buffer_half_float */
//#ifndef GL_EXT_color_buffer_half_float
const GL_EXT_color_buffer_half_float = 1

//#endif
//
///* GL_EXT_debug_label */
//#ifndef GL_EXT_debug_label
const GL_EXT_debug_label = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glLabelObjectEXT (GLenum type, GLuint object, GLsizei length, const GLchar *label);
//GL_APICALL void GL_APIENTRY glGetObjectLabelEXT (GLenum type, GLuint object, GLsizei bufSize, GLsizei *length, GLchar *label);
//#endif
//typedef void (GL_APIENTRYP PFNGLLABELOBJECTEXTPROC) (GLenum type, GLuint object, GLsizei length, const GLchar *label);
//todo
//typedef void (GL_APIENTRYP PFNGLGETOBJECTLABELEXTPROC) (GLenum type, GLuint object, GLsizei bufSize, GLsizei *length, GLchar *label);
//todo
//#endif

/* GL_EXT_debug_marker */
//#ifndef GL_EXT_debug_marker
const GL_EXT_debug_marker = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glInsertEventMarkerEXT (GLsizei length, const GLchar *marker);
//GL_APICALL void GL_APIENTRY glPushGroupMarkerEXT (GLsizei length, const GLchar *marker);
//GL_APICALL void GL_APIENTRY glPopGroupMarkerEXT (void);
//#endif
//typedef void (GL_APIENTRYP PFNGLINSERTEVENTMARKEREXTPROC) (GLsizei length, const GLchar *marker);
//todo
//typedef void (GL_APIENTRYP PFNGLPUSHGROUPMARKEREXTPROC) (GLsizei length, const GLchar *marker);
//todo
//typedef void (GL_APIENTRYP PFNGLPOPGROUPMARKEREXTPROC) (void);
//todo
//#endif

/* GL_EXT_discard_framebuffer */
//#ifndef GL_EXT_discard_framebuffer
const GL_EXT_discard_framebuffer = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDiscardFramebufferEXT (GLenum target, GLsizei numAttachments, const GLenum *attachments);
//#endif
//typedef void (GL_APIENTRYP PFNGLDISCARDFRAMEBUFFEREXTPROC) (GLenum target, GLsizei numAttachments, const GLenum *attachments);
//todo
//#endif

//#ifndef GL_EXT_disjoint_timer_query
const GL_EXT_disjoint_timer_query = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glGenQueriesEXT (GLsizei n, GLuint *ids);
//GL_APICALL void GL_APIENTRY glDeleteQueriesEXT (GLsizei n, const GLuint *ids);
//GL_APICALL GLboolean GL_APIENTRY glIsQueryEXT (GLuint id);
//GL_APICALL void GL_APIENTRY glBeginQueryEXT (GLenum target, GLuint id);
//GL_APICALL void GL_APIENTRY glEndQueryEXT (GLenum target);
//GL_APICALL void GL_APIENTRY glQueryCounterEXT (GLuint id, GLenum target);
//GL_APICALL void GL_APIENTRY glGetQueryivEXT (GLenum target, GLenum pname, GLint *params);
//GL_APICALL void GL_APIENTRY glGetQueryObjectivEXT (GLuint id, GLenum pname, GLint *params);
//GL_APICALL void GL_APIENTRY glGetQueryObjectuivEXT (GLuint id, GLenum pname, GLuint *params);
//GL_APICALL void GL_APIENTRY glGetQueryObjecti64vEXT (GLuint id, GLenum pname, GLint64 *params);
//GL_APICALL void GL_APIENTRY glGetQueryObjectui64vEXT (GLuint id, GLenum pname, GLuint64 *params);
//#endif
//typedef void (GL_APIENTRYP PFNGLGENQUERIESEXTPROC) (GLsizei n, GLuint *ids);
//todo
//typedef void (GL_APIENTRYP PFNGLDELETEQUERIESEXTPROC) (GLsizei n, const GLuint *ids);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLISQUERYEXTPROC) (GLuint id);
//todo
//typedef void (GL_APIENTRYP PFNGLBEGINQUERYEXTPROC) (GLenum target, GLuint id);
//todo
//typedef void (GL_APIENTRYP PFNGLENDQUERYEXTPROC) (GLenum target);
//todo
//typedef void (GL_APIENTRYP PFNGLQUERYCOUNTEREXTPROC) (GLuint id, GLenum target);
//todo
//typedef void (GL_APIENTRYP PFNGLGETQUERYIVEXTPROC) (GLenum target, GLenum pname, GLint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETQUERYOBJECTIVEXTPROC) (GLuint id, GLenum pname, GLint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETQUERYOBJECTUIVEXTPROC) (GLuint id, GLenum pname, GLuint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETQUERYOBJECTI64VEXTPROC) (GLuint id, GLenum pname, GLint64 *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETQUERYOBJECTUI64VEXTPROC) (GLuint id, GLenum pname, GLuint64 *params);
//todo
//#endif /* GL_EXT_disjoint_timer_query */

//#ifndef GL_EXT_draw_buffers
const GL_EXT_draw_buffers = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDrawBuffersEXT (GLsizei n, const GLenum *bufs);
//#endif
//typedef void (GL_APIENTRYP PFNGLDRAWBUFFERSEXTPROC) (GLsizei n, const GLenum *bufs);
//todo
//#endif /* GL_EXT_draw_buffers */
//
///* GL_EXT_map_buffer_range */
//#ifndef GL_EXT_map_buffer_range
const GL_EXT_map_buffer_range = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void* GL_APIENTRY glMapBufferRangeEXT (GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
//GL_APICALL void GL_APIENTRY glFlushMappedBufferRangeEXT (GLenum target, GLintptr offset, GLsizeiptr length);
//#endif
//typedef void* (GL_APIENTRYP PFNGLMAPBUFFERRANGEEXTPROC) (GLenum target, GLintptr offset, GLsizeiptr length, GLbitfield access);
//todo
//typedef void (GL_APIENTRYP PFNGLFLUSHMAPPEDBUFFERRANGEEXTPROC) (GLenum target, GLintptr offset, GLsizeiptr length);
//todo
//#endif
//
///* GL_EXT_multisampled_render_to_texture */
//#ifndef GL_EXT_multisampled_render_to_texture
const GL_EXT_multisampled_render_to_texture = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glRenderbufferStorageMultisampleEXT (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glFramebufferTexture2DMultisampleEXT (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLsizei samples);
//#endif
//typedef void (GL_APIENTRYP PFNGLRENDERBUFFERSTORAGEMULTISAMPLEEXTPROC) (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEEXTPROC) (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLsizei samples);
//todo
//#endif

/* GL_EXT_multiview_draw_buffers */
//#ifndef GL_EXT_multiview_draw_buffers
const GL_EXT_multiview_draw_buffers = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glReadBufferIndexedEXT (GLenum src, GLint index);
//GL_APICALL void GL_APIENTRY glDrawBuffersIndexedEXT (GLint n, const GLenum *location, const GLint *indices);
//GL_APICALL void GL_APIENTRY glGetIntegeri_vEXT (GLenum target, GLuint index, GLint *data);
//#endif
//typedef void (GL_APIENTRYP PFNGLREADBUFFERINDEXEDEXTPROC) (GLenum src, GLint index);
//todo
//typedef void (GL_APIENTRYP PFNGLDRAWBUFFERSINDEXEDEXTPROC) (GLint n, const GLenum *location, const GLint *indices);
//todo
//typedef void (GL_APIENTRYP PFNGLGETINTEGERI_VEXTPROC) (GLenum target, GLuint index, GLint *data);
//todo
//#endif
//
//#ifndef GL_EXT_multi_draw_arrays
const GL_EXT_multi_draw_arrays = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glMultiDrawArraysEXT (GLenum mode, const GLint *first, const GLsizei *count, GLsizei primcount);
//GL_APICALL void GL_APIENTRY glMultiDrawElementsEXT (GLenum mode, const GLsizei *count, GLenum type, const GLvoid **indices, GLsizei primcount);
//#endif /* GL_GLEXT_PROTOTYPES */
//typedef void (GL_APIENTRYP PFNGLMULTIDRAWARRAYSEXTPROC) (GLenum mode, const GLint *first, const GLsizei *count, GLsizei primcount);
//todo
//typedef void (GL_APIENTRYP PFNGLMULTIDRAWELEMENTSEXTPROC) (GLenum mode, const GLsizei *count, GLenum type, const GLvoid **indices, GLsizei primcount);
//todo
//#endif
//
///* GL_EXT_occlusion_query_boolean */
//#ifndef GL_EXT_occlusion_query_boolean
const GL_EXT_occlusion_query_boolean = 1

/* All entry points also exist in GL_EXT_disjoint_timer_query */
//#endif
//
///* GL_EXT_read_format_bgra */
//#ifndef GL_EXT_read_format_bgra
const GL_EXT_read_format_bgra = 1

//#endif
//
///* GL_EXT_robustness */
//#ifndef GL_EXT_robustness
const GL_EXT_robustness = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL GLenum GL_APIENTRY glGetGraphicsResetStatusEXT (void);
//GL_APICALL void GL_APIENTRY glReadnPixelsEXT (GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLsizei bufSize, GLvoid *data);
//GL_APICALL void GL_APIENTRY glGetnUniformfvEXT (GLuint program, GLint location, GLsizei bufSize, GLfloat *params);
//GL_APICALL void GL_APIENTRY glGetnUniformivEXT (GLuint program, GLint location, GLsizei bufSize, GLint *params);
//#endif
//typedef GLenum (GL_APIENTRYP PFNGLGETGRAPHICSRESETSTATUSEXTPROC) (void);
//todo
//typedef void (GL_APIENTRYP PFNGLREADNPIXELSEXTPROC) (GLint x, GLint y, GLsizei width, GLsizei height, GLenum format, GLenum type, GLsizei bufSize, GLvoid *data);
//todo
//typedef void (GL_APIENTRYP PFNGLGETNUNIFORMFVEXTPROC) (GLuint program, GLint location, GLsizei bufSize, GLfloat *params);
//todo
//typedef void (GL_APIENTRYP PFNGLGETNUNIFORMIVEXTPROC) (GLuint program, GLint location, GLsizei bufSize, GLint *params);
//todo
//#endif

/* GL_EXT_separate_shader_objects */
//#ifndef GL_EXT_separate_shader_objects
const GL_EXT_separate_shader_objects = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glUseProgramStagesEXT (GLuint pipeline, GLbitfield stages, GLuint program);
//GL_APICALL void GL_APIENTRY glActiveShaderProgramEXT (GLuint pipeline, GLuint program);
//GL_APICALL GLuint GL_APIENTRY glCreateShaderProgramvEXT (GLenum type, GLsizei count, const GLchar **strings);
//GL_APICALL void GL_APIENTRY glBindProgramPipelineEXT (GLuint pipeline);
//GL_APICALL void GL_APIENTRY glDeleteProgramPipelinesEXT (GLsizei n, const GLuint *pipelines);
//GL_APICALL void GL_APIENTRY glGenProgramPipelinesEXT (GLsizei n, GLuint *pipelines);
//GL_APICALL GLboolean GL_APIENTRY glIsProgramPipelineEXT (GLuint pipeline);
//GL_APICALL void GL_APIENTRY glProgramParameteriEXT (GLuint program, GLenum pname, GLint value);
//GL_APICALL void GL_APIENTRY glGetProgramPipelineivEXT (GLuint pipeline, GLenum pname, GLint *params);
//GL_APICALL void GL_APIENTRY glProgramUniform1iEXT (GLuint program, GLint location, GLint x);
//GL_APICALL void GL_APIENTRY glProgramUniform2iEXT (GLuint program, GLint location, GLint x, GLint y);
//GL_APICALL void GL_APIENTRY glProgramUniform3iEXT (GLuint program, GLint location, GLint x, GLint y, GLint z);
//GL_APICALL void GL_APIENTRY glProgramUniform4iEXT (GLuint program, GLint location, GLint x, GLint y, GLint z, GLint w);
//GL_APICALL void GL_APIENTRY glProgramUniform1fEXT (GLuint program, GLint location, GLfloat x);
//GL_APICALL void GL_APIENTRY glProgramUniform2fEXT (GLuint program, GLint location, GLfloat x, GLfloat y);
//GL_APICALL void GL_APIENTRY glProgramUniform3fEXT (GLuint program, GLint location, GLfloat x, GLfloat y, GLfloat z);
//GL_APICALL void GL_APIENTRY glProgramUniform4fEXT (GLuint program, GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
//GL_APICALL void GL_APIENTRY glProgramUniform1ivEXT (GLuint program, GLint location, GLsizei count, const GLint *value);
//GL_APICALL void GL_APIENTRY glProgramUniform2ivEXT (GLuint program, GLint location, GLsizei count, const GLint *value);
//GL_APICALL void GL_APIENTRY glProgramUniform3ivEXT (GLuint program, GLint location, GLsizei count, const GLint *value);
//GL_APICALL void GL_APIENTRY glProgramUniform4ivEXT (GLuint program, GLint location, GLsizei count, const GLint *value);
//GL_APICALL void GL_APIENTRY glProgramUniform1fvEXT (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniform2fvEXT (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniform3fvEXT (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniform4fvEXT (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniformMatrix2fvEXT (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniformMatrix3fvEXT (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glProgramUniformMatrix4fvEXT (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//GL_APICALL void GL_APIENTRY glValidateProgramPipelineEXT (GLuint pipeline);
//GL_APICALL void GL_APIENTRY glGetProgramPipelineInfoLogEXT (GLuint pipeline, GLsizei bufSize, GLsizei *length, GLchar *infoLog);
//#endif
//typedef void (GL_APIENTRYP PFNGLUSEPROGRAMSTAGESEXTPROC) (GLuint pipeline, GLbitfield stages, GLuint program);
//todo
//typedef void (GL_APIENTRYP PFNGLACTIVESHADERPROGRAMEXTPROC) (GLuint pipeline, GLuint program);
//todo
//typedef GLuint (GL_APIENTRYP PFNGLCREATESHADERPROGRAMVEXTPROC) (GLenum type, GLsizei count, const GLchar **strings);
//todo
//typedef void (GL_APIENTRYP PFNGLBINDPROGRAMPIPELINEEXTPROC) (GLuint pipeline);
//todo
//typedef void (GL_APIENTRYP PFNGLDELETEPROGRAMPIPELINESEXTPROC) (GLsizei n, const GLuint *pipelines);
//todo
//typedef void (GL_APIENTRYP PFNGLGENPROGRAMPIPELINESEXTPROC) (GLsizei n, GLuint *pipelines);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLISPROGRAMPIPELINEEXTPROC) (GLuint pipeline);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMPARAMETERIEXTPROC) (GLuint program, GLenum pname, GLint value);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPROGRAMPIPELINEIVEXTPROC) (GLuint pipeline, GLenum pname, GLint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM1IEXTPROC) (GLuint program, GLint location, GLint x);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM2IEXTPROC) (GLuint program, GLint location, GLint x, GLint y);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM3IEXTPROC) (GLuint program, GLint location, GLint x, GLint y, GLint z);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM4IEXTPROC) (GLuint program, GLint location, GLint x, GLint y, GLint z, GLint w);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM1FEXTPROC) (GLuint program, GLint location, GLfloat x);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM2FEXTPROC) (GLuint program, GLint location, GLfloat x, GLfloat y);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM3FEXTPROC) (GLuint program, GLint location, GLfloat x, GLfloat y, GLfloat z);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM4FEXTPROC) (GLuint program, GLint location, GLfloat x, GLfloat y, GLfloat z, GLfloat w);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM1IVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLint *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM2IVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLint *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM3IVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLint *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM4IVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLint *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM1FVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM2FVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM3FVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORM4FVEXTPROC) (GLuint program, GLint location, GLsizei count, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORMMATRIX2FVEXTPROC) (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORMMATRIX3FVEXTPROC) (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLPROGRAMUNIFORMMATRIX4FVEXTPROC) (GLuint program, GLint location, GLsizei count, GLboolean transpose, const GLfloat *value);
//todo
//typedef void (GL_APIENTRYP PFNGLVALIDATEPROGRAMPIPELINEEXTPROC) (GLuint pipeline);
//todo
//typedef void (GL_APIENTRYP PFNGLGETPROGRAMPIPELINEINFOLOGEXTPROC) (GLuint pipeline, GLsizei bufSize, GLsizei *length, GLchar *infoLog);
//todo
//#endif
//
///* GL_EXT_shader_framebuffer_fetch */
//#ifndef GL_EXT_shader_framebuffer_fetch
const GL_EXT_shader_framebuffer_fetch = 1

//#endif
//
///* GL_EXT_shader_texture_lod */
//#ifndef GL_EXT_shader_texture_lod
const GL_EXT_shader_texture_lod = 1

//#endif
//
///* GL_EXT_shadow_samplers */
//#ifndef GL_EXT_shadow_samplers
const GL_EXT_shadow_samplers = 1

//#endif
//
///* GL_EXT_sRGB */
//#ifndef GL_EXT_sRGB
const GL_EXT_sRGB = 1

//#endif
//
///* GL_EXT_texture_compression_dxt1 */
//#ifndef GL_EXT_texture_compression_dxt1
const GL_EXT_texture_compression_dxt1 = 1

//#endif
//
///* GL_EXT_texture_filter_anisotropic */
//#ifndef GL_EXT_texture_filter_anisotropic
const GL_EXT_texture_filter_anisotropic = 1

//#endif
//
///* GL_EXT_texture_format_BGRA8888 */
//#ifndef GL_EXT_texture_format_BGRA8888
const GL_EXT_texture_format_BGRA8888 = 1

//#endif
//
///* GL_EXT_texture_rg */
//#ifndef GL_EXT_texture_rg
const GL_EXT_texture_rg = 1

//#endif
//
///* GL_EXT_texture_storage */
//#ifndef GL_EXT_texture_storage
const GL_EXT_texture_storage = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glTexStorage1DEXT (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
//GL_APICALL void GL_APIENTRY glTexStorage2DEXT (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glTexStorage3DEXT (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
//GL_APICALL void GL_APIENTRY glTextureStorage1DEXT (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
//GL_APICALL void GL_APIENTRY glTextureStorage2DEXT (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glTextureStorage3DEXT (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
//#endif
//typedef void (GL_APIENTRYP PFNGLTEXSTORAGE1DEXTPROC) (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXSTORAGE2DEXTPROC) (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXSTORAGE3DEXTPROC) (GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXTURESTORAGE1DEXTPROC) (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXTURESTORAGE2DEXTPROC) (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLTEXTURESTORAGE3DEXTPROC) (GLuint texture, GLenum target, GLsizei levels, GLenum internalformat, GLsizei width, GLsizei height, GLsizei depth);
//todo
//#endif

/* GL_EXT_texture_type_2_10_10_10_REV */
//#ifndef GL_EXT_texture_type_2_10_10_10_REV
const GL_EXT_texture_type_2_10_10_10_REV = 1

//#endif
//
///* GL_EXT_unpack_subimage */
//#ifndef GL_EXT_unpack_subimage
const GL_EXT_unpack_subimage = 1

//#endif
//
///*------------------------------------------------------------------------*
// * DMP extension functions
// *------------------------------------------------------------------------*/
//
///* GL_DMP_shader_binary */
//#ifndef GL_DMP_shader_binary
const GL_DMP_shader_binary = 1

//#endif
//
///*------------------------------------------------------------------------*
// * FJ extension functions
// *------------------------------------------------------------------------*/
//
///* GL_FJ_shader_binary_GCCSO */
//#ifndef GL_FJ_shader_binary_GCCSO
const GL_FJ_shader_binary_GCCSO = 1

//#endif
//
///*------------------------------------------------------------------------*
// * IMG extension functions
// *------------------------------------------------------------------------*/
//
///* GL_IMG_program_binary */
//#ifndef GL_IMG_program_binary
const GL_IMG_program_binary = 1

//#endif
//
///* GL_IMG_read_format */
//#ifndef GL_IMG_read_format
const GL_IMG_read_format = 1

//#endif
//
///* GL_IMG_shader_binary */
//#ifndef GL_IMG_shader_binary
const GL_IMG_shader_binary = 1

//#endif
//
///* GL_IMG_texture_compression_pvrtc */
//#ifndef GL_IMG_texture_compression_pvrtc
const GL_IMG_texture_compression_pvrtc = 1

//#endif
//
///* GL_IMG_texture_compression_pvrtc2 */
//#ifndef GL_IMG_texture_compression_pvrtc2
const GL_IMG_texture_compression_pvrtc2 = 1

//#endif
//
///* GL_IMG_multisampled_render_to_texture */
//#ifndef GL_IMG_multisampled_render_to_texture
const GL_IMG_multisampled_render_to_texture = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glRenderbufferStorageMultisampleIMG (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//GL_APICALL void GL_APIENTRY glFramebufferTexture2DMultisampleIMG (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLsizei samples);
//#endif
//typedef void (GL_APIENTRYP PFNGLRENDERBUFFERSTORAGEMULTISAMPLEIMGPROC) (GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//typedef void (GL_APIENTRYP PFNGLFRAMEBUFFERTEXTURE2DMULTISAMPLEIMGPROC) (GLenum target, GLenum attachment, GLenum textarget, GLuint texture, GLint level, GLsizei samples);
//todo
//#endif

/*------------------------------------------------------------------------*
 * NV extension functions
 *------------------------------------------------------------------------*/

/* GL_NV_coverage_sample */
//#ifndef GL_NV_coverage_sample
const GL_NV_coverage_sample = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glCoverageMaskNV (GLboolean mask);
//GL_APICALL void GL_APIENTRY glCoverageOperationNV (GLenum operation);
//#endif
//typedef void (GL_APIENTRYP PFNGLCOVERAGEMASKNVPROC) (GLboolean mask);
//todo
//typedef void (GL_APIENTRYP PFNGLCOVERAGEOPERATIONNVPROC) (GLenum operation);
//todo
//#endif

/* GL_NV_depth_nonlinear */
//#ifndef GL_NV_depth_nonlinear
const GL_NV_depth_nonlinear = 1

//#endif
//
///* GL_NV_draw_buffers */
//#ifndef GL_NV_draw_buffers
const GL_NV_draw_buffers = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDrawBuffersNV (GLsizei n, const GLenum *bufs);
//#endif
//typedef void (GL_APIENTRYP PFNGLDRAWBUFFERSNVPROC) (GLsizei n, const GLenum *bufs);
//todo
//#endif

/* GL_NV_draw_instanced */
//#ifndef GL_NV_draw_instanced
//const GL_NV_draw_instanced= 1
//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDrawArraysInstancedNV (GLenum mode, GLint first, GLsizei count, GLsizei primcount);
//GL_APICALL void GL_APIENTRY glDrawElementsInstancedNV (GLenum mode, GLsizei count, GLenum type, const GLvoid *indices, GLsizei primcount);
//#endif
//typedef void (GL_APIENTRYP PFNGLDRAWARRAYSINSTANCEDNVPROC) (GLenum mode, GLint first, GLsizei count, GLsizei primcount);
//todo
//typedef void (GL_APIENTRYP PFNGLDRAWELEMENTSINSTANCEDNVPROC) (GLenum mode, GLsizei count, GLenum type, const GLvoid *indices, GLsizei primcount);
//todo
//#endif

/* GL_NV_fbo_color_attachments */
//#ifndef GL_NV_fbo_color_attachments
const GL_NV_fbo_color_attachments = 1

//#endif
//
///* GL_NV_fence */
//#ifndef GL_NV_fence
const GL_NV_fence = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glDeleteFencesNV (GLsizei n, const GLuint *fences);
//GL_APICALL void GL_APIENTRY glGenFencesNV (GLsizei n, GLuint *fences);
//GL_APICALL GLboolean GL_APIENTRY glIsFenceNV (GLuint fence);
//GL_APICALL GLboolean GL_APIENTRY glTestFenceNV (GLuint fence);
//GL_APICALL void GL_APIENTRY glGetFenceivNV (GLuint fence, GLenum pname, GLint *params);
//GL_APICALL void GL_APIENTRY glFinishFenceNV (GLuint fence);
//GL_APICALL void GL_APIENTRY glSetFenceNV (GLuint fence, GLenum condition);
//#endif
//typedef void (GL_APIENTRYP PFNGLDELETEFENCESNVPROC) (GLsizei n, const GLuint *fences);
//todo
//typedef void (GL_APIENTRYP PFNGLGENFENCESNVPROC) (GLsizei n, GLuint *fences);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLISFENCENVPROC) (GLuint fence);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLTESTFENCENVPROC) (GLuint fence);
//todo
//typedef void (GL_APIENTRYP PFNGLGETFENCEIVNVPROC) (GLuint fence, GLenum pname, GLint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLFINISHFENCENVPROC) (GLuint fence);
//todo
//typedef void (GL_APIENTRYP PFNGLSETFENCENVPROC) (GLuint fence, GLenum condition);
//todo
//#endif

/* GL_NV_framebuffer_blit */
//#ifndef GL_NV_framebuffer_blit
const GL_NV_framebuffer_blit = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glBlitFramebufferNV (GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
//#endif
//typedef void (GL_APIENTRYP PFNGLBLITFRAMEBUFFERNVPROC) (GLint srcX0, GLint srcY0, GLint srcX1, GLint srcY1, GLint dstX0, GLint dstY0, GLint dstX1, GLint dstY1, GLbitfield mask, GLenum filter);
//todo
//#endif

/* GL_NV_framebuffer_multisample */
//#ifndef GL_NV_framebuffer_multisample
const GL_NV_framebuffer_multisample = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glRenderbufferStorageMultisampleNV ( GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//#endif
//typedef void (GL_APIENTRYP PFNGLRENDERBUFFERSTORAGEMULTISAMPLENVPROC) ( GLenum target, GLsizei samples, GLenum internalformat, GLsizei width, GLsizei height);
//todo
//#endif

/* GL_NV_generate_mipmap_sRGB */
//#ifndef GL_NV_generate_mipmap_sRGB
const GL_NV_generate_mipmap_sRGB = 1

//#endif
//
///* GL_NV_instanced_arrays */
//#ifndef GL_NV_instanced_arrays
const GL_NV_instanced_arrays = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glVertexAttribDivisorNV (GLuint index, GLuint divisor);
//#endif
//typedef void (GL_APIENTRYP PFNGLVERTEXATTRIBDIVISORNVPROC) (GLuint index, GLuint divisor);
//todo
//#endif

/* GL_NV_read_buffer */
//#ifndef GL_NV_read_buffer
const GL_NV_read_buffer = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glReadBufferNV (GLenum mode);
//#endif
//typedef void (GL_APIENTRYP PFNGLREADBUFFERNVPROC) (GLenum mode);
//todo
//#endif

/* GL_NV_read_buffer_front */
//#ifndef GL_NV_read_buffer_front
const GL_NV_read_buffer_front = 1

//#endif
//
///* GL_NV_read_depth */
//#ifndef GL_NV_read_depth
const GL_NV_read_depth = 1

//#endif
//
///* GL_NV_read_depth_stencil */
//#ifndef GL_NV_read_depth_stencil
const GL_NV_read_depth_stencil = 1

//#endif
//
///* GL_NV_read_stencil */
//#ifndef GL_NV_read_stencil
const GL_NV_read_stencil = 1

//#endif
//
///* GL_NV_shadow_samplers_array */
//#ifndef GL_NV_shadow_samplers_array
const GL_NV_shadow_samplers_array = 1

//#endif
//
///* GL_NV_shadow_samplers_cube */
//#ifndef GL_NV_shadow_samplers_cube
const GL_NV_shadow_samplers_cube = 1

//#endif
//
///* GL_NV_sRGB_formats */
//#ifndef GL_NV_sRGB_formats
const GL_NV_sRGB_formats = 1

//#endif
//
///* GL_NV_texture_border_clamp */
//#ifndef GL_NV_texture_border_clamp
const GL_NV_texture_border_clamp = 1

//#endif
//
///* GL_NV_texture_compression_s3tc_update */
//#ifndef GL_NV_texture_compression_s3tc_update
const GL_NV_texture_compression_s3tc_update = 1

//#endif
//
///* GL_NV_texture_npot_2D_mipmap */
//#ifndef GL_NV_texture_npot_2D_mipmap
const GL_NV_texture_npot_2D_mipmap = 1

//#endif
//
///*------------------------------------------------------------------------*
// * QCOM extension functions
// *------------------------------------------------------------------------*/
//
///* GL_QCOM_alpha_test */
//#ifndef GL_QCOM_alpha_test
const GL_QCOM_alpha_test = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glAlphaFuncQCOM (GLenum func, GLclampf ref);
//#endif
//typedef void (GL_APIENTRYP PFNGLALPHAFUNCQCOMPROC) (GLenum func, GLclampf ref);
//todo
//#endif

/* GL_QCOM_binning_control */
//#ifndef GL_QCOM_binning_control
const GL_QCOM_binning_control = 1

//#endif
//
///* GL_QCOM_driver_control */
//#ifndef GL_QCOM_driver_control
const GL_QCOM_driver_control = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glGetDriverControlsQCOM (GLint *num, GLsizei size, GLuint *driverControls);
//GL_APICALL void GL_APIENTRY glGetDriverControlStringQCOM (GLuint driverControl, GLsizei bufSize, GLsizei *length, GLchar *driverControlString);
//GL_APICALL void GL_APIENTRY glEnableDriverControlQCOM (GLuint driverControl);
//GL_APICALL void GL_APIENTRY glDisableDriverControlQCOM (GLuint driverControl);
//#endif
//typedef void (GL_APIENTRYP PFNGLGETDRIVERCONTROLSQCOMPROC) (GLint *num, GLsizei size, GLuint *driverControls);
//todo
//typedef void (GL_APIENTRYP PFNGLGETDRIVERCONTROLSTRINGQCOMPROC) (GLuint driverControl, GLsizei bufSize, GLsizei *length, GLchar *driverControlString);
//todo
//typedef void (GL_APIENTRYP PFNGLENABLEDRIVERCONTROLQCOMPROC) (GLuint driverControl);
//todo
//typedef void (GL_APIENTRYP PFNGLDISABLEDRIVERCONTROLQCOMPROC) (GLuint driverControl);
//todo
//#endif

/* GL_QCOM_extended_get */
//#ifndef GL_QCOM_extended_get
const GL_QCOM_extended_get = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glExtGetTexturesQCOM (GLuint *textures, GLint maxTextures, GLint *numTextures);
//GL_APICALL void GL_APIENTRY glExtGetBuffersQCOM (GLuint *buffers, GLint maxBuffers, GLint *numBuffers);
//GL_APICALL void GL_APIENTRY glExtGetRenderbuffersQCOM (GLuint *renderbuffers, GLint maxRenderbuffers, GLint *numRenderbuffers);
//GL_APICALL void GL_APIENTRY glExtGetFramebuffersQCOM (GLuint *framebuffers, GLint maxFramebuffers, GLint *numFramebuffers);
//GL_APICALL void GL_APIENTRY glExtGetTexLevelParameterivQCOM (GLuint texture, GLenum face, GLint level, GLenum pname, GLint *params);
//GL_APICALL void GL_APIENTRY glExtTexObjectStateOverrideiQCOM (GLenum target, GLenum pname, GLint param);
//GL_APICALL void GL_APIENTRY glExtGetTexSubImageQCOM (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid *texels);
//GL_APICALL void GL_APIENTRY glExtGetBufferPointervQCOM (GLenum target, GLvoid **params);
//#endif
//typedef void (GL_APIENTRYP PFNGLEXTGETTEXTURESQCOMPROC) (GLuint *textures, GLint maxTextures, GLint *numTextures);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETBUFFERSQCOMPROC) (GLuint *buffers, GLint maxBuffers, GLint *numBuffers);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETRENDERBUFFERSQCOMPROC) (GLuint *renderbuffers, GLint maxRenderbuffers, GLint *numRenderbuffers);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETFRAMEBUFFERSQCOMPROC) (GLuint *framebuffers, GLint maxFramebuffers, GLint *numFramebuffers);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETTEXLEVELPARAMETERIVQCOMPROC) (GLuint texture, GLenum face, GLint level, GLenum pname, GLint *params);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTTEXOBJECTSTATEOVERRIDEIQCOMPROC) (GLenum target, GLenum pname, GLint param);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETTEXSUBIMAGEQCOMPROC) (GLenum target, GLint level, GLint xoffset, GLint yoffset, GLint zoffset, GLsizei width, GLsizei height, GLsizei depth, GLenum format, GLenum type, GLvoid *texels);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETBUFFERPOINTERVQCOMPROC) (GLenum target, GLvoid **params);
//todo
//#endif

/* GL_QCOM_extended_get2 */
//#ifndef GL_QCOM_extended_get2
const GL_QCOM_extended_get2 = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glExtGetShadersQCOM (GLuint *shaders, GLint maxShaders, GLint *numShaders);
//GL_APICALL void GL_APIENTRY glExtGetProgramsQCOM (GLuint *programs, GLint maxPrograms, GLint *numPrograms);
//GL_APICALL GLboolean GL_APIENTRY glExtIsProgramBinaryQCOM (GLuint program);
//GL_APICALL void GL_APIENTRY glExtGetProgramBinarySourceQCOM (GLuint program, GLenum shadertype, GLchar *source, GLint *length);
//#endif
//typedef void (GL_APIENTRYP PFNGLEXTGETSHADERSQCOMPROC) (GLuint *shaders, GLint maxShaders, GLint *numShaders);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETPROGRAMSQCOMPROC) (GLuint *programs, GLint maxPrograms, GLint *numPrograms);
//todo
//typedef GLboolean (GL_APIENTRYP PFNGLEXTISPROGRAMBINARYQCOMPROC) (GLuint program);
//todo
//typedef void (GL_APIENTRYP PFNGLEXTGETPROGRAMBINARYSOURCEQCOMPROC) (GLuint program, GLenum shadertype, GLchar *source, GLint *length);
//todo
//#endif

/* GL_QCOM_perfmon_global_mode */
//#ifndef GL_QCOM_perfmon_global_mode
const GL_QCOM_perfmon_global_mode = 1

//#endif
//
///* GL_QCOM_writeonly_rendering */
//#ifndef GL_QCOM_writeonly_rendering
const GL_QCOM_writeonly_rendering = 1

//#endif
//
///* GL_QCOM_tiled_rendering */
//#ifndef GL_QCOM_tiled_rendering
const GL_QCOM_tiled_rendering = 1

//#ifdef GL_GLEXT_PROTOTYPES
//GL_APICALL void GL_APIENTRY glStartTilingQCOM (GLuint x, GLuint y, GLuint width, GLuint height, GLbitfield preserveMask);
//GL_APICALL void GL_APIENTRY glEndTilingQCOM (GLbitfield preserveMask);
//#endif
//typedef void (GL_APIENTRYP PFNGLSTARTTILINGQCOMPROC) (GLuint x, GLuint y, GLuint width, GLuint height, GLbitfield preserveMask);
//todo
//typedef void (GL_APIENTRYP PFNGLENDTILINGQCOMPROC) (GLbitfield preserveMask);
//todo
//#endif

/*------------------------------------------------------------------------*
 * VIV extension tokens
 *------------------------------------------------------------------------*/

/* GL_VIV_shader_binary */
//#ifndef GL_VIV_shader_binary
const GL_VIV_shader_binary = 1

//#endif
//
//#ifdef __cplusplus
//}
//#endif
//
//#endif /* __gl2ext_h_ */
