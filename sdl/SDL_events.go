package sdl

import (
	"unsafe"

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
 *  \file SDL_events.h
 *
 *  Include file for SDL event handling.
 */

// #ifndef SDL_events_h_
// #define SDL_events_h_

// #include <SDL3/SDL_audio.h>
// #include <SDL3/SDL_error.h>
// #include <SDL3/SDL_gamepad.h>
// #include <SDL3/SDL_joystick.h>
// #include <SDL3/SDL_keyboard.h>
// #include <SDL3/SDL_mouse.h>
// #include <SDL3/SDL_quit.h>
// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_touch.h>
// #include <SDL3/SDL_video.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/* General keyboard/mouse state definitions */
const SDL_RELEASED = 0
const SDL_PRESSED = 1

/**
 * The types of events that can be delivered.
 */
type SDL_EventType int32

const (
	SDL_FIRSTEVENT = 0 /**< Unused (do not remove) */
)
const (
	/* Application events */
	SDL_QUIT = 0x100 + iota /**< User-requested quit */

	/* These application events have special meaning on iOS, see README-ios.md for details */
	SDL_APP_TERMINATING /**< The application is being terminated by the OS
	  Called on iOS in applicationWillTerminate()
	  Called on Android in onDestroy()
	*/
	SDL_APP_LOWMEMORY /**< The application is low on memory, free memory if possible.
	  Called on iOS in applicationDidReceiveMemoryWarning()
	  Called on Android in onLowMemory()
	*/
	SDL_APP_WILLENTERBACKGROUND /**< The application is about to enter the background
	  Called on iOS in applicationWillResignActive()
	  Called on Android in onPause()
	*/
	SDL_APP_DIDENTERBACKGROUND /**< The application did enter the background and may not get CPU for some time
	  Called on iOS in applicationDidEnterBackground()
	  Called on Android in onPause()
	*/
	SDL_APP_WILLENTERFOREGROUND /**< The application is about to enter the foreground
	  Called on iOS in applicationWillEnterForeground()
	  Called on Android in onResume()
	*/
	SDL_APP_DIDENTERFOREGROUND /**< The application is now interactive
	  Called on iOS in applicationDidBecomeActive()
	  Called on Android in onResume()
	*/

	SDL_LOCALECHANGED /**< The user's locale preferences have changed. */

	/* Display events */
	SDL_DISPLAYEVENT = 0x150 /**< Display state change */
)
const (
	/* Window events */
	SDL_WINDOWEVENT = 0x200 + iota /**< Window state change */
	SDL_SYSWMEVENT                 /**< System specific event */
)
const (
	/* Keyboard events */
	SDL_KEYDOWN       = 0x300 + iota /**< Key pressed */
	SDL_KEYUP                        /**< Key released */
	SDL_TEXTEDITING                  /**< Keyboard text editing (composition) */
	SDL_TEXTINPUT                    /**< Keyboard text input */
	SDL_KEYMAPCHANGED                /**< Keymap changed due to a system event such as an
	  input language or keyboard layout change.
	*/
)
const (
	/* Mouse events */
	SDL_MOUSEMOTION     = 0x400 + iota /**< Mouse moved */
	SDL_MOUSEBUTTONDOWN                /**< Mouse button pressed */
	SDL_MOUSEBUTTONUP                  /**< Mouse button released */
	SDL_MOUSEWHEEL                     /**< Mouse wheel motion */
)
const (
	/* Joystick events */
	SDL_JOYAXISMOTION    = 0x600 + iota /**< Joystick axis motion */
	SDL_JOYBALLMOTION                   /**< Joystick trackball motion */
	SDL_JOYHATMOTION                    /**< Joystick hat position change */
	SDL_JOYBUTTONDOWN                   /**< Joystick button pressed */
	SDL_JOYBUTTONUP                     /**< Joystick button released */
	SDL_JOYDEVICEADDED                  /**< A new joystick has been inserted into the system */
	SDL_JOYDEVICEREMOVED                /**< An opened joystick has been removed */
)
const (
	/* Game controller events */
	SDL_CONTROLLERAXISMOTION     = 0x650 + iota /**< Game controller axis motion */
	SDL_CONTROLLERBUTTONDOWN                    /**< Game controller button pressed */
	SDL_CONTROLLERBUTTONUP                      /**< Game controller button released */
	SDL_CONTROLLERDEVICEADDED                   /**< A new Game controller has been inserted into the system */
	SDL_CONTROLLERDEVICEREMOVED                 /**< An opened Game controller has been removed */
	SDL_CONTROLLERDEVICEREMAPPED                /**< The controller mapping was updated */
	SDL_CONTROLLERTOUCHPADDOWN                  /**< Game controller touchpad was touched */
	SDL_CONTROLLERTOUCHPADMOTION                /**< Game controller touchpad finger was moved */
	SDL_CONTROLLERTOUCHPADUP                    /**< Game controller touchpad finger was lifted */
	SDL_CONTROLLERSENSORUPDATE                  /**< Game controller sensor was updated */
)
const (
	/* Touch events */
	SDL_FINGERDOWN = 0x700 + iota
	SDL_FINGERUP
	SDL_FINGERMOTION
)
const (
	/* Gesture events */
	SDL_DOLLARGESTURE = 0x800
	SDL_DOLLARRECORD
	SDL_MULTIGESTURE

	/* Clipboard events */
	SDL_CLIPBOARDUPDATE = 0x900 /**< The clipboard changed */
)
const (
	/* Drag and drop events */
	SDL_DROPFILE     = 0x1000 + iota /**< The system requests a file open */
	SDL_DROPTEXT                     /**< text/plain drag-and-drop event */
	SDL_DROPBEGIN                    /**< A new set of drops is beginning (NULL filename) */
	SDL_DROPCOMPLETE                 /**< Current set of drops is now complete (NULL filename) */
)
const (
	/* Audio hotplug events */
	SDL_AUDIODEVICEADDED   = 0x1100 + iota /**< A new audio device is available */
	SDL_AUDIODEVICEREMOVED                 /**< An audio device has been removed. */

	/* Sensor events */
	SDL_SENSORUPDATE = 0x1200 /**< A sensor was updated */
)
const (
	/* Render events */
	SDL_RENDER_TARGETS_RESET = 0x2000 + iota /**< The render targets have been reset and their contents need to be updated */
	SDL_RENDER_DEVICE_RESET                  /**< The device has been reset and all textures need to be recreated */

	/** Events ::SDL_USEREVENT through ::SDL_LASTEVENT are for your use,
	 *  and should be allocated with SDL_RegisterEvents()
	 */
	SDL_USEREVENT = 0x8000

	/**
	 *  This last event is only for bounding internal arrays
	 */
	SDL_LASTEVENT = 0xFFFF
)

/**
 *  \brief Fields shared by every event
 */
type SDL_CommonEvent struct {
	Type      sdlcommon.FUint32T
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS()  */
}

/**
 *  \brief Display state change event data (event.display.*)
 */
type SDL_DisplayEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_DISPLAYEVENT_* */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	DisplayID SDL_DisplayID      /**< The associated display */
	Data1     sdlcommon.FSint32  /**< event dependent data */
}

/**
 *  \brief Window state change event data (event.window.*)
 */

type SDL_WindowEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_WINDOWEVENT */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID       /**< The associated window */
	Data1     sdlcommon.FSint32  /**< event dependent data */
	Data2     sdlcommon.FSint32  /**< event dependent data */
}

/**
 *  \brief Keyboard button event structure (event.key.*)
 */
type SDL_KeyboardEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_KEYDOWN or ::SDL_KEYUP */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID       /**< The window with keyboard focus, if any */
	State     sdlcommon.FUint8T  /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Repeat    sdlcommon.FUint8T  /**< Non-zero if this is a key repeat */
	Padding2  sdlcommon.FUint8T
	Padding3  sdlcommon.FUint8T
	Keysym    SDL_Keysym /**< The key that was pressed or released */
}

const SDL_TEXTEDITINGEVENT_TEXT_SIZE = (32)

/**
 *  \brief Keyboard text editing event structure (event.edit.*)
 */
type SDL_TextEditingEvent struct {
	Type      sdlcommon.FUint32T                              /**< ::SDL_TEXTEDITING */
	Timestamp sdlcommon.FUint64T                              /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID                                    /**< The window with keyboard focus, if any */
	Text      [SDL_TEXTEDITINGEVENT_TEXT_SIZE]sdlcommon.FChar /**< The editing text */
	Start     sdlcommon.FSint32                               /**< The start cursor of selected editing text */
	Length    sdlcommon.FSint32                               /**< The length of selected editing text */
}

/**
 *  \brief Extended keyboard text editing event structure (event.editExt.*) when text would be
 *  truncated if stored in the text buffer SDL_TextEditingEvent
 */
type SDL_TextEditingExtEvent struct {
	Type      sdlcommon.FUint32T      /**< ::SDL_EVENT_TEXT_EDITING_EXT */
	Timestamp sdlcommon.FUint64T      /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID            /**< The window with keyboard focus, if any */
	Text      *sdlcommon.FCharPStruct /**< The editing text, which should be freed with SDL_free(), and will not be NULL */
	Start     sdlcommon.FSint32       /**< The start cursor of selected editing text */
	Length    sdlcommon.FSint32       /**< The length of selected editing text */
}

const SDL_TEXTINPUTEVENT_TEXT_SIZE = (32)

/**
 *  \brief Keyboard text input event structure (event.text.*)
 */
type SDL_TextInputEvent struct {
	Type      sdlcommon.FUint32T                            /**< ::SDL_EVENT_TEXT_INPUT */
	Timestamp sdlcommon.FUint64T                            /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID                                  /**< The window with keyboard focus, if any */
	text      [SDL_TEXTINPUTEVENT_TEXT_SIZE]sdlcommon.FChar /**< The input text */
}

/**
 *  \brief Mouse motion event structure (event.motion.*)
 */
type SDL_MouseMotionEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_MOUSEMOTION */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID       /**< The window with mouse focus, if any */
	Which     SDL_MouseID        /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	State     sdlcommon.FUint32T /**< The current button state */
	X         sdlcommon.FFloat   /**< X coordinate, relative to window */
	Y         sdlcommon.FFloat   /**< Y coordinate, relative to window */
	Xrel      sdlcommon.FFloat   /**< The relative motion in the X direction */
	Trel      sdlcommon.FFloat   /**< The relative motion in the Y direction */
}

/**
 *  \brief Mouse button event structure (event.button.*)
 */
type SDL_MouseButtonEvent struct {
	Type      sdlcommon.FUint32T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Timestamp sdlcommon.FUint64T /**< In milliseconds, populated using SDL_GetTicks() */
	WindowID  SDL_WindowID       /**< The window with mouse focus, if any */
	Which     SDL_MouseID        /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	Button    sdlcommon.FUint8T  /**< The mouse button index */
	State     sdlcommon.FUint8T  /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Clicks    sdlcommon.FUint8T  /**< 1 for single-click, 2 for double-click, etc. */
	Padding1  sdlcommon.FUint8T
	X         sdlcommon.FFloat /**< X coordinate, relative to window */
	Y         sdlcommon.FFloat /**< Y coordinate, relative to window */
}

/**
 *  \brief Mouse wheel event structure (event.wheel.*)
 */
type SDL_MouseWheelEvent struct {
	Type      sdlcommon.FUint32T     /**< ::SDL_MOUSEWHEEL */
	Timestamp sdlcommon.FUint64T     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  sdlcommon.SDL_WindowID /**< The window with mouse focus, if any */
	Which     sdlcommon.SDL_MouseID  /**< The mouse instance id, or SDL_TOUCH_MOUSEID */
	X         sdlcommon.FFloat       /**< The amount scrolled horizontally, positive to the right and negative to the left */
	Y         sdlcommon.FFloat       /**< The amount scrolled vertically, positive away from the user and negative toward the user */
	Direction sdlcommon.FUint32T     /**< Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back */
	MouseX    sdlcommon.FFloat       /**< X coordinate, relative to window */
	MouseY    sdlcommon.FFloat       /**< Y coordinate, relative to window */
}

/**
 *  \brief Joystick axis motion event structure (event.jaxis.*)
 */
type SDL_JoyAxisEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_JOYAXISMOTION */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	Axis      sdlcommon.FUint8T  /**< The joystick axis index */
	Padding1  sdlcommon.FUint8T
	Padding2  sdlcommon.FUint8T
	Padding3  sdlcommon.FUint8T
	Value     sdlcommon.FSint16 /**< The axis value (range: -32768 to 32767) */
	Padding4  sdlcommon.FUint16T
}

/**
 *  \brief Joystick hat position change event structure (event.jhat.*)
 */
type SDL_JoyHatEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_JOYHATMOTION */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	Hat       sdlcommon.FUint8T  /**< The joystick hat index */
	Value     sdlcommon.FUint8T  /**< The hat position value.
	*   \sa ::SDL_HAT_LEFTUP ::SDL_HAT_UP ::SDL_HAT_RIGHTUP
	*   \sa ::SDL_HAT_LEFT ::SDL_HAT_CENTERED ::SDL_HAT_RIGHT
	*   \sa ::SDL_HAT_LEFTDOWN ::SDL_HAT_DOWN ::SDL_HAT_RIGHTDOWN
	*
	*   Note that zero means the POV is centered.
	 */
	Padding1 sdlcommon.FUint8T
	Padding2 sdlcommon.FUint8T
}

/**
 *  \brief Joystick button event structure (event.jbutton.*)
 */
type SDL_JoyButtonEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_JOYBUTTONDOWN or ::SDL_JOYBUTTONUP */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	Button    sdlcommon.FUint8T  /**< The joystick button index */
	State     sdlcommon.FUint8T  /**< ::SDL_PRESSED or ::SDL_RELEASED */
	Padding1  sdlcommon.FUint8T
	Padding2  sdlcommon.FUint8T
}

/**
 *  \brief Joystick device event structure (event.jdevice.*)
 */
type SDL_JoyDeviceEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_JOYDEVICEADDED or ::SDL_JOYDEVICEREMOVED */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick device index for the ADDED event, instance id for the REMOVED event */
}

/**
 *  \brief Joysick battery level change event structure (event.jbattery.*)
 */
type SDL_JoyBatteryEvent struct {
	Type      sdlcommon.FUint32T     /**< ::SDL_EVENT_JOYSTICK_BATTERY_UPDATED */
	Timestamp sdlcommon.FUint64T     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID         /**< The joystick instance id */
	level     SDL_JoystickPowerLevel /**< The joystick battery level */
}

/**
/**
 *  \brief Gamepad axis motion event structure (event.caxis.*)
*/
type SDL_GamepadAxisEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_GAMEPAD_AXIS_MOTION */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	Axis      sdlcommon.FUint8T  /**< The gamepad axis (SDL_GamepadAxis) */
	Padding1  sdlcommon.FUint8T
	Padding2  sdlcommon.FUint8T
	Padding3  sdlcommon.FUint8T
	value     sdlcommon.FSint16 /**< The axis value (range: -32768 to 32767) */
	padding4  sdlcommon.FUint16T
}

/**
 *  \brief Gamepad button event structure (event.cbutton.*)
 */
type SDL_GamepadButtonEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_GAMEPAD_BUTTON_DOWN or ::SDL_EVENT_GAMEPAD_BUTTON_UP */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	button    sdlcommon.FUint8T  /**< The gamepad axis (SDL_GamepadAxis) */
	state     sdlcommon.FUint8T
	Padding1  sdlcommon.FUint8T
	Padding2  sdlcommon.FUint8T
}

/**
 *  \brief Gamepad device event structure (event.cdevice.*)
 */
type SDL_GamepadDeviceEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_GAMEPAD_ADDED, ::SDL_EVENT_GAMEPAD_REMOVED, or ::SDL_EVENT_GAMEPAD_REMAPPED */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
}

/**
 *  \brief Gamepad touchpad event structure (event.ctouchpad.*)
 */
type SDL_GamepadTouchpadEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_GAMEPAD_TOUCHPAD_DOWN or ::SDL_EVENT_GAMEPAD_TOUCHPAD_MOTION or ::SDL_EVENT_GAMEPAD_TOUCHPAD_UP */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The joystick instance id */
	Touchpad  sdlcommon.FSint32  /**< The index of the touchpad */
	Finger    sdlcommon.FSint32  /**< The index of the finger on the touchpad */
	x         sdlcommon.FFloat   /**< Normalized in the range 0...1 with 0 being on the left */
	y         sdlcommon.FFloat   /**< Normalized in the range 0...1 with 0 being at the top */
	pressure  sdlcommon.FFloat   /**< Normalized in the range 0...1 */
}

/**
 *  \brief Gamepad sensor event structure (event.csensor.*)
 */
type SDL_GamepadSensorEvent struct {
	Type             sdlcommon.FUint32T  /**< ::SDL_EVENT_GAMEPAD_SENSOR_UPDATE */
	Timestamp        sdlcommon.FUint64T  /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which            SDL_JoystickID      /**< The joystick instance id */
	Sensor           sdlcommon.FSint32   /**< The type of the sensor, one of the values of ::SDL_SensorType */
	Data             [3]sdlcommon.FFloat /**< Up to 3 values from the sensor, as defined in SDL_sensor.h */
	Sensor_timestamp sdlcommon.FUint64T  /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

/**
 *  \brief Audio device event structure (event.adevice.*)
 */
type SDL_AudioDeviceEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_AUDIO_DEVICE_ADDED, or ::SDL_EVENT_AUDIO_DEVICE_REMOVED */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which     SDL_JoystickID     /**< The audio device index for the ADDED event (valid until next SDL_GetNumAudioDevices() call), SDL_AudioDeviceID for the REMOVED event */
	Iscapture sdlcommon.FUint8T  /**< zero if an output device, non-zero if a capture device. */
	Padding1  sdlcommon.FUint8T
	Padding2  sdlcommon.FUint8T
	Padding3  sdlcommon.FUint8T
}

/**
 *  \brief Touch finger event structure (event.tfinger.*)
 */
type SDL_TouchFingerEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_FINGER_MOTION or ::SDL_EVENT_FINGER_DOWN or ::SDL_EVENT_FINGER_UP */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	TouchId   SDL_TouchID        /**< The touch device id */
	FingerId  SDL_FingerID
	X         sdlcommon.FFloat /**< Normalized in the range 0...1 */
	Y         sdlcommon.FFloat /**< Normalized in the range 0...1 */
	Dx        sdlcommon.FFloat /**< Normalized in the range -1...1 */
	Dy        sdlcommon.FFloat /**< Normalized in the range -1...1 */
	Pressure  sdlcommon.FFloat /**< Normalized in the range 0...1 */
	WindowID  SDL_WindowID     /**< The window underneath the finger, if any */
}

/**
 *  \brief An event used to request a file open by the system (event.drop.*)
 *         This event is enabled by default, you can disable it with SDL_SetEventEnabled().
 *  \note If this event is enabled, you must free the filename in the event.
 */
type SDL_DropEvent struct {
	Type      sdlcommon.FUint32T     /**< ::SDL_EVENT_DROP_BEGIN or ::SDL_EVENT_DROP_FILE or ::SDL_EVENT_DROP_TEXT or ::SDL_EVENT_DROP_COMPLETE */
	Timestamp sdlcommon.FUint64T     /**< In nanoseconds, populated using SDL_GetTicksNS() */
	File      sdlcommon.FCharPStruct /**< The file name, which should be freed with SDL_free(), is NULL on begin/complete */
	WindowID  SDL_WindowID           /**< The window that was dropped on, if any */
}

/**
 *  \brief Sensor event structure (event.sensor.*)
 */
type SDL_SensorEvent struct {
	Type            sdlcommon.FUint32T  /**< ::SDL_EVENT_SENSOR_UPDATE */
	Timestamp       sdlcommon.FUint64T  /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Which           SDL_SensorID        /**< The instance ID of the sensor */
	Data            [6]sdlcommon.FFloat /**< Up to 6 values from the sensor - additional values can be queried using SDL_GetSensorData() */
	SensorTimestamp sdlcommon.FUint64T  /**< The timestamp of the sensor reading in nanoseconds, not necessarily synchronized with the system clock */
}

/**
 *  \brief The "quit requested" event
 */
type SDL_QuitEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_QUIT */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
}

/**
 *  \brief OS Specific event
 */
type SDL_OSEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_QUIT */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
}

/**
 *  \brief A user-defined event type (event.user.*)
 */
type SDL_UserEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_USER through ::SDL_EVENT_LAST-1 */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	WindowID  SDL_WindowID       /**< The associated window if any */
	Code      sdlcommon.FSint32  /**< User defined event code */
	Data1     sdlcommon.FVoidP   /**< User defined data pointer */
	Data2     sdlcommon.FVoidP   /**< User defined data pointer */
}

// struct SDL_SysWMmsg;
// typedef struct SDL_SysWMmsg SDL_SysWMmsg;

/**
 *  \brief A video driver dependent system event (event.syswm.*)
 *         This event is disabled by default, you can enable it with SDL_SetEventEnabled()
 *
 *  \note If you want to use this event, you should include SDL_syswm.h.
 */
type SDL_SysWMEvent struct {
	Type      sdlcommon.FUint32T /**< ::SDL_EVENT_SYSWM */
	Timestamp sdlcommon.FUint64T /**< In nanoseconds, populated using SDL_GetTicksNS() */
	Msg       *SDL_SysWMmsg      /**< driver dependent data, defined in SDL_syswm.h */
}

/**
 *  \brief General event structure
 */
type SDL_Event struct {
	Type sdlcommon.FUint32T /**< Event type, shared with all events */
	// SDL_CommonEvent common;                 /**< Common event data */
	// SDL_DisplayEvent display;               /**< Display event data */
	// SDL_WindowEvent window;                 /**< Window event data */
	// SDL_KeyboardEvent key;                  /**< Keyboard event data */
	// SDL_TextEditingEvent edit;              /**< Text editing event data */
	// SDL_TextEditingExtEvent editExt;        /**< Extended text editing event data */
	// SDL_TextInputEvent text;                /**< Text input event data */
	// SDL_MouseMotionEvent motion;            /**< Mouse motion event data */
	// SDL_MouseButtonEvent button;            /**< Mouse button event data */
	// SDL_MouseWheelEvent wheel;              /**< Mouse wheel event data */
	// SDL_JoyAxisEvent jaxis;                 /**< Joystick axis event data */
	// SDL_JoyHatEvent jhat;                   /**< Joystick hat event data */
	// SDL_JoyButtonEvent jbutton;             /**< Joystick button event data */
	// SDL_JoyDeviceEvent jdevice;             /**< Joystick device change event data */
	// SDL_JoyBatteryEvent jbattery;           /**< Joystick battery event data */
	// SDL_GamepadAxisEvent caxis;             /**< Gamepad axis event data */
	// SDL_GamepadButtonEvent cbutton;         /**< Gamepad button event data */
	// SDL_GamepadDeviceEvent cdevice;         /**< Gamepad device event data */
	// SDL_GamepadTouchpadEvent ctouchpad;     /**< Gamepad touchpad event data */
	// SDL_GamepadSensorEvent csensor;         /**< Gamepad sensor event data */
	// SDL_AudioDeviceEvent adevice;           /**< Audio device event data */
	// SDL_SensorEvent sensor;                 /**< Sensor event data */
	// SDL_QuitEvent quit;                     /**< Quit request event data */
	// SDL_UserEvent user;                     /**< Custom event data */
	// SDL_SysWMEvent syswm;                   /**< System dependent window event data */
	// SDL_TouchFingerEvent tfinger;           /**< Touch finger event data */
	// SDL_DropEvent drop;                     /**< Drag and drop event data */

	/* This is necessary for ABI compatibility between Visual C++ and GCC.
	   Visual C++ will respect the push pack pragma and use 52 bytes (size of
	   SDL_TextEditingEvent, the largest structure for 32-bit and 64-bit
	   architectures) for this union, and GCC will use the alignment of the
	   largest datatype within the union, which is 8 bytes on 64-bit
	   architectures.

	   So... we'll add padding to force the size to be 56 bytes for both.

	   On architectures where pointers are 16 bytes, this needs rounding up to
	   the next multiple of 16, 64, and on architectures where pointers are
	   even larger the size of SDL_UserEvent will dominate as being 3 pointers.
	*/
	padding [128 - 4]sdlcommon.FUint8T
}

/* Make sure we haven't broken binary compatibility */
// SDL_COMPILE_TIME_ASSERT(SDL_Event, sizeof(SDL_Event) == sizeof(((SDL_Event *)NULL)->padding));

/* Function prototypes */

/**
 * Pump the event loop, gathering events from the input devices.
 *
 * This function updates the event queue and internal input device state.
 *
 * **WARNING**: This should only be run in the thread that initialized the
 * video subsystem, and for extra safety, you should consider only doing those
 * things on the main thread in any case.
 *
 * SDL_PumpEvents() gathers all the pending input information from devices and
 * places it in the event queue. Without calls to SDL_PumpEvents() no events
 * would ever be placed on the queue. Often the need for calls to
 * SDL_PumpEvents() is hidden from the user since SDL_PollEvent() and
 * SDL_WaitEvent() implicitly call SDL_PumpEvents(). However, if you are not
 * polling or waiting for events (e.g. you are filtering them), then you must
 * call SDL_PumpEvents() to force an event queue update.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_WaitEvent
 */
// extern DECLSPEC void SDLCALL SDL_PumpEvents(void);
func SDL_PumpEvents() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_PumpEvents").Call()
}

/* @{ */
type SDL_eventaction int32

const (
	SDL_ADDEVENT = iota
	SDL_PEEKEVENT
	SDL_GETEVENT
)

/**
 * Check the event queue for messages and optionally return them.
 *
 * `action` may be any of the following:
 *
 * - `SDL_ADDEVENT`: up to `numevents` events will be added to the back of the
 *   event queue.
 * - `SDL_PEEKEVENT`: `numevents` events at the front of the event queue,
 *   within the specified minimum and maximum type, will be returned to the
 *   caller and will _not_ be removed from the queue.
 * - `SDL_GETEVENT`: up to `numevents` events at the front of the event queue,
 *   within the specified minimum and maximum type, will be returned to the
 *   caller and will be removed from the queue.
 *
 * You may have to call SDL_PumpEvents() before calling this function.
 * Otherwise, the events may not be ready to be filtered when you call
 * SDL_PeepEvents().
 *
 * This function is thread-safe.
 *
 * \param events destination buffer for the retrieved events
 * \param numevents if action is SDL_ADDEVENT, the number of events to add
 *                  back to the event queue; if action is SDL_PEEKEVENT or
 *                  SDL_GETEVENT, the maximum number of events to retrieve
 * \param action action to take; see [[#action|Remarks]] for details
 * \param minType minimum value of the event type to be considered;
 *                SDL_EVENT_FIRST is a safe choice
 * \param maxType maximum value of the event type to be considered;
 *                SDL_EVENT_LAST is a safe choice
 * \returns the number of events actually stored or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PumpEvents
 * \sa SDL_PushEvent
 */
// extern DECLSPEC int SDLCALL SDL_PeepEvents(SDL_Event * events, int numevents,
//                                            SDL_eventaction action,
//                                            Uint32 minType, Uint32 maxType);
func (events *SDL_Event) SDL_PeepEvents(
	numevents sdlcommon.FInt,
	action SDL_eventaction,
	minType sdlcommon.FUint32T,
	maxType sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PeepEvents").Call(
		uintptr(unsafe.Pointer(events)),
		uintptr(numevents),
		uintptr(action),
		uintptr(minType),
		uintptr(maxType),
	)
	res = sdlcommon.FInt(t)
	return
}

/* @} */

/**
 * Check for the existence of a certain event type in the event queue.
 *
 * If you need to check for a range of event types, use SDL_HasEvents()
 * instead.
 *
 * \param type the type of event to be queried; see SDL_EventType for details
 * \returns SDL_TRUE if events matching `type` are present, or SDL_FALSE if
 *          events matching `type` are not present.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_HasEvents
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_HasEvent(Uint32 type);
func SDL_HasEvent(type0 sdlcommon.FUint32T) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasEvent").Call(
		uintptr(type0),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Check for the existence of certain event types in the event queue.
 *
 * If you need to check for a single event type, use SDL_HasEvent() instead.
 *
 * \param minType the low end of event type to be queried, inclusive; see
 *                SDL_EventType for details
 * \param maxType the high end of event type to be queried, inclusive; see
 *                SDL_EventType for details
 * \returns SDL_TRUE if events with type >= `minType` and <= `maxType` are
 *          present, or SDL_FALSE if not.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_HasEvents
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_HasEvents(Uint32 minType, Uint32 maxType);
func SDL_HasEvents(minType, maxType sdlcommon.FUint32T) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasEvents").Call(
		uintptr(minType),
		uintptr(maxType),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Clear events of a specific type from the event queue.
 *
 * This will unconditionally remove any events from the queue that match
 * `type`. If you need to remove a range of event types, use SDL_FlushEvents()
 * instead.
 *
 * It's also normal to just ignore events you don't care about in your event
 * loop without calling this function.
 *
 * This function only affects currently queued events. If you want to make
 * sure that all pending OS events are flushed, you can call SDL_PumpEvents()
 * on the main thread immediately before the flush call.
 *
 * \param type the type of event to be cleared; see SDL_EventType for details
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_FlushEvents
 */
// extern DECLSPEC void SDLCALL SDL_FlushEvent(Uint32 type);
func SDL_FlushEvent(type0 sdlcommon.FUint32T) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_FlushEvent").Call(
		uintptr(type0),
	)
	return
}

/**
 * Clear events of a range of types from the event queue.
 *
 * This will unconditionally remove any events from the queue that are in the
 * range of `minType` to `maxType`, inclusive. If you need to remove a single
 * event type, use SDL_FlushEvent() instead.
 *
 * It's also normal to just ignore events you don't care about in your event
 * loop without calling this function.
 *
 * This function only affects currently queued events. If you want to make
 * sure that all pending OS events are flushed, you can call SDL_PumpEvents()
 * on the main thread immediately before the flush call.
 *
 * \param minType the low end of event type to be cleared, inclusive; see
 *                SDL_EventType for details
 * \param maxType the high end of event type to be cleared, inclusive; see
 *                SDL_EventType for details
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_FlushEvent
 */
// extern DECLSPEC void SDLCALL SDL_FlushEvents(Uint32 minType, Uint32 maxType);
func SDL_FlushEvents(minType, maxType sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FlushEvents").Call(
		uintptr(minType),
		uintptr(maxType),
	)
	return
}

/**
 * Poll for currently pending events.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`. The 1 returned refers to
 * this event, immediately stored in the SDL Event structure -- not an event
 * to follow.
 *
 * If `event` is NULL, it simply returns 1 if there is an event in the queue,
 * but will not remove it from the queue.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that set the video mode.
 *
 * SDL_PollEvent() is the favored way of receiving system events since it can
 * be done from the main loop and does not suspend the main loop while waiting
 * on an event to be posted.
 *
 * The common practice is to fully process the event queue once every frame,
 * usually as a first step before updating the game's state:
 *
 * ```c
 * while (game_is_still_running) {
 *     SDL_Event event;
 *     while (SDL_PollEvent(&event)) {  // poll until all events are handled!
 *         // decide what to do with this event.
 *     }
 *
 *     // update game state, draw the current frame
 * }
 * ```
 *
 * \param event the SDL_Event structure to be filled with the next event from
 *              the queue, or NULL
 * \returns 1 if there is a pending event or 0 if there are none available.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetEventFilter
 * \sa SDL_PeepEvents
 * \sa SDL_PushEvent
 * \sa SDL_SetEventFilter
 * \sa SDL_WaitEvent
 * \sa SDL_WaitEventTimeout
 */
// extern DECLSPEC int SDLCALL SDL_PollEvent(SDL_Event * event);
func (event *SDL_Event) SDL_PollEvent() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PollEvent").Call(
		uintptr(unsafe.Pointer(event)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Wait indefinitely for the next available event.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that initialized the video subsystem.
 *
 * \param event the SDL_Event structure to be filled in with the next event
 *              from the queue, or NULL
 * \returns 1 on success or 0 if there was an error while waiting for events;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PumpEvents
 * \sa SDL_WaitEventTimeout
 */
// extern DECLSPEC int SDLCALL SDL_WaitEvent(SDL_Event *event);
func (event *SDL_Event) SDL_WaitEvent() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WaitEvent").Call(
		uintptr(unsafe.Pointer(event)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Wait until the specified timeout (in milliseconds) for the next available
 * event.
 *
 * If `event` is not NULL, the next event is removed from the queue and stored
 * in the SDL_Event structure pointed to by `event`.
 *
 * As this function may implicitly call SDL_PumpEvents(), you can only call
 * this function in the thread that initialized the video subsystem.
 *
 * The timeout is not guaranteed, the actual wait time could be longer due to
 * system scheduling.
 *
 * \param event the SDL_Event structure to be filled in with the next event
 *              from the queue, or NULL
 * \param timeoutMS the maximum number of milliseconds to wait for the next
 *                  available event
 * \returns 1 on success or 0 if there was an error while waiting for events;
 *          call SDL_GetError() for more information. This also returns 0 if
 *          the timeout elapsed without an event arriving.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PollEvent
 * \sa SDL_PumpEvents
 * \sa SDL_WaitEvent
 */
// extern DECLSPEC int SDLCALL SDL_WaitEventTimeout(SDL_Event *event, Sint32 timeoutMS);
func (event *SDL_Event) SDL_WaitEventTimeout(timeoutMS sdlcommon.FSint32) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WaitEventTimeout").Call(
		uintptr(unsafe.Pointer(event)),
		uintptr(timeoutMS),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Add an event to the event queue.
 *
 * The event queue can actually be used as a two way communication channel.
 * Not only can events be read from the queue, but the user can also push
 * their own events onto it. `event` is a pointer to the event structure you
 * wish to push onto the queue. The event is copied into the queue, and the
 * caller may dispose of the memory pointed to after SDL_PushEvent() returns.
 *
 * Note: Pushing device input events onto the queue doesn't modify the state
 * of the device within SDL.
 *
 * This function is thread-safe, and can be called from other threads safely.
 *
 * Note: Events pushed onto the queue with SDL_PushEvent() get passed through
 * the event filter but events added with SDL_PeepEvents() do not.
 *
 * For pushing application-specific events, please use SDL_RegisterEvents() to
 * get an event type that does not conflict with other code that also wants
 * its own custom event types.
 *
 * \param event the SDL_Event to be added to the queue
 * \returns 1 on success, 0 if the event was filtered, or a negative error
 *          code on failure; call SDL_GetError() for more information. A
 *          common reason for error is the event queue being full.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PeepEvents
 * \sa SDL_PollEvent
 * \sa SDL_RegisterEvents
 */
// extern DECLSPEC int SDLCALL SDL_PushEvent(SDL_Event * event);
func (event *SDL_Event) SDL_PushEvent() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PushEvent").Call(
		uintptr(unsafe.Pointer(event)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * A function pointer used for callbacks that watch the event queue.
 *
 * \param userdata what was passed as `userdata` to SDL_SetEventFilter()
 *        or SDL_AddEventWatch, etc
 * \param event the event that triggered the callback
 * \returns 1 to permit event to be added to the queue, and 0 to disallow
 *          it. When used with SDL_AddEventWatch, the return value is ignored.
 *
 * \sa SDL_SetEventFilter
 * \sa SDL_AddEventWatch
 */
// typedef int (SDLCALL * SDL_EventFilter) (void *userdata, SDL_Event * event);
type SDL_EventFilter = func(userdata sdlcommon.FVoidP, event *SDL_Event) uintptr

/**
 * Set up a filter to process all events before they change internal state and
 * are posted to the internal event queue.
 *
 * If the filter function returns 1 when called, then the event will be added
 * to the internal queue. If it returns 0, then the event will be dropped from
 * the queue, but the internal state will still be updated. This allows
 * selective filtering of dynamically arriving events.
 *
 * **WARNING**: Be very careful of what you do in the event filter function,
 * as it may run in a different thread!
 *
 * On platforms that support it, if the quit event is generated by an
 * interrupt signal (e.g. pressing Ctrl-C), it will be delivered to the
 * application at the next event poll.
 *
 * There is one caveat when dealing with the ::SDL_QuitEvent event type. The
 * event filter is only called when the window manager desires to close the
 * application window. If the event filter returns 1, then the window will be
 * closed, otherwise the window will remain open if possible.
 *
 * Note: Disabled events never make it to the event filter function; see
 * SDL_SetEventEnabled().
 *
 * Note: If you just want to inspect events without filtering, you should use
 * SDL_AddEventWatch() instead.
 *
 * Note: Events pushed onto the queue with SDL_PushEvent() get passed through
 * the event filter, but events pushed onto the queue with SDL_PeepEvents() do
 * not.
 *
 * \param filter An SDL_EventFilter function to call when an event happens
 * \param userdata a pointer that is passed to `filter`
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_AddEventWatch
 * \sa SDL_SetEventEnabled
 * \sa SDL_GetEventFilter
 * \sa SDL_PeepEvents
 * \sa SDL_PushEvent
 */
// extern DECLSPEC void SDLCALL SDL_SetEventFilter(SDL_EventFilter filter,
//                                                 void *userdata);
func SDL_SetEventFilter(filter SDL_EventFilter, userdata sdlcommon.FVoidP) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_SetEventFilter").Call(
		sdlcommon.NewCallback(filter),
		uintptr(unsafe.Pointer(userdata)),
	)
}

/**
 * Query the current event filter.
 *
 * This function can be used to "chain" filters, by saving the existing filter
 * before replacing it with a function that will call that saved filter.
 *
 * \param filter the current callback function will be stored here
 * \param userdata the pointer that is passed to the current event filter will
 *                 be stored here
 * \returns SDL_TRUE on success or SDL_FALSE if there is no event filter set.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetEventFilter
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetEventFilter(SDL_EventFilter * filter,
//                                                     void **userdata);
func SDL_GetEventFilter(filter SDL_EventFilter, userdata *sdlcommon.FVoidP) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetEventFilter").Call(
		sdlcommon.NewCallback(filter),
		uintptr(unsafe.Pointer(userdata)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Add a callback to be triggered when an event is added to the event queue.
 *
 * `filter` will be called when an event happens, and its return value is
 * ignored.
 *
 * **WARNING**: Be very careful of what you do in the event filter function,
 * as it may run in a different thread!
 *
 * If the quit event is generated by a signal (e.g. SIGINT), it will bypass
 * the internal queue and be delivered to the watch callback immediately, and
 * arrive at the next event poll.
 *
 * Note: the callback is called for events posted by the user through
 * SDL_PushEvent(), but not for disabled events, nor for events by a filter
 * callback set with SDL_SetEventFilter(), nor for events posted by the user
 * through SDL_PeepEvents().
 *
 * \param filter an SDL_EventFilter function to call when an event happens.
 * \param userdata a pointer that is passed to `filter`
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DelEventWatch
 * \sa SDL_SetEventFilter
 */
// extern DECLSPEC void SDLCALL SDL_AddEventWatch(SDL_EventFilter filter,
//                                                void *userdata);
func SDL_AddEventWatch(filter SDL_EventFilter, userdata sdlcommon.FVoidP) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_AddEventWatch").Call(
		sdlcommon.NewCallback(filter),
		userdata,
	)
}

/**
 * Remove an event watch callback added with SDL_AddEventWatch().
 *
 * This function takes the same input as SDL_AddEventWatch() to identify and
 * delete the corresponding callback.
 *
 * \param filter the function originally passed to SDL_AddEventWatch()
 * \param userdata the pointer originally passed to SDL_AddEventWatch()
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_AddEventWatch
 */
// extern DECLSPEC void SDLCALL SDL_DelEventWatch(SDL_EventFilter filter,
//                                                void *userdata);
func SDL_DelEventWatch(filter SDL_EventFilter, userdata sdlcommon.FVoidP) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DelEventWatch").Call(
		sdlcommon.NewCallback(filter),
		userdata,
	)
}

/**
 * Run a specific filter function on the current event queue, removing any
 * events for which the filter returns 0.
 *
 * See SDL_SetEventFilter() for more information. Unlike SDL_SetEventFilter(),
 * this function does not change the filter permanently, it only uses the
 * supplied filter until this function returns.
 *
 * \param filter the SDL_EventFilter function to call when an event happens
 * \param userdata a pointer that is passed to `filter`
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetEventFilter
 * \sa SDL_SetEventFilter
 */
// extern DECLSPEC void SDLCALL SDL_FilterEvents(SDL_EventFilter filter,
//                                               void *userdata);
func SDL_FilterEvents(filter SDL_EventFilter, userdata sdlcommon.FVoidP) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_FilterEvents").Call(
		sdlcommon.NewCallback(filter),
		userdata,
	)
}

/**
 * Set the state of processing events by type.
 *
 * \param type the type of event; see SDL_EventType for details
 * \param enabled whether to process the event or not
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_IsEventEnabled
 */
// extern DECLSPEC void SDLCALL SDL_SetEventEnabled(Uint32 type, SDL_bool enabled);
func SDL_SetEventEnabled(type0 sdlcommon.FUint32T, enabled bool) (res bool) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_SetEventEnabled").Call(
		uintptr(type0),
		sdlcommon.CBool(enabled),
	)
	return
}

/**
 * Query the state of processing events by type.
 *
 * \param type the type of event; see SDL_EventType for details
 * \returns SDL_TRUE if the event is being processed, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetEventEnabled
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_EventEnabled(Uint32 type);
func SDL_EventEnabled(type0 sdlcommon.FUint32T) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EventEnabled").Call(
		uintptr(type0),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Allocate a set of user-defined events, and return the beginning event
 * number for that set of events.
 *
 * Calling this function with `numevents` <= 0 is an error and will return
 * (Uint32)-1.
 *
 * Note, (Uint32)-1 means the maximum unsigned 32-bit integer value (or
 * 0xFFFFFFFF), but is clearer to write.
 *
 * \param numevents the number of events to be allocated
 * \returns the beginning event number, or (Uint32)-1 if there are not enough
 *          user-defined events left.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_PushEvent
 */
// extern DECLSPEC Uint32 SDLCALL SDL_RegisterEvents(int numevents);
func SDL_RegisterEvents(numevents sdlcommon.FInt) (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RegisterEvents").Call(
		uintptr(numevents),
	)
	res = sdlcommon.FUint32T(t)
	return
}

// /* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_events_h_ */
