package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/**
 *  \brief SDL_sensor.h
 *
 *  In order to use these functions, SDL_Init() must have been called
 *  with the ::SDL_INIT_SENSOR flag.  This causes SDL to scan the system
 *  for sensors, and load appropriate drivers.
 */
type SDL_Sensor struct {
}

/**
 * This is a unique ID for a sensor for the time it is connected to the system,
 * and is never reused for the lifetime of the application.
 *
 * The ID value starts at 0 and increments from there. The value -1 is an invalid ID.
 */
type SDL_SensorID = sdlcommon.FSint32

/* The different sensors defined by SDL
 *
 * Additional sensors may be available, using platform dependent semantics.
 *
 * Hare are the additional Android sensors:
 * https://developer.android.com/reference/android/hardware/SensorEvent.html#values
 */
type SDL_SensorType = int32

const (
	SDL_SENSOR_INVALID = iota - 1 /**< Returned for an invalid sensor */
	SDL_SENSOR_UNKNOWN            /**< Unknown sensor type */
	SDL_SENSOR_ACCEL              /**< Accelerometer */
	SDL_SENSOR_GYRO               /**< Gyroscope */
)

/**
 * Accelerometer sensor
 *
 * The accelerometer returns the current acceleration in SI meters per
 * second squared. This measurement includes the force of gravity, so
 * a device at rest will have an value of SDL_STANDARD_GRAVITY away
 * from the center of the earth.
 *
 * values[0]: Acceleration on the x axis
 * values[1]: Acceleration on the y axis
 * values[2]: Acceleration on the z axis
 *
 * For phones held in portrait mode and game controllers held in front of you,
 * the axes are defined as follows:
 * -X ... +X : left ... right
 * -Y ... +Y : bottom ... top
 * -Z ... +Z : farther ... closer
 *
 * The axis data is not changed when the phone is rotated.
 *
 * \sa SDL_GetDisplayOrientation()
 */
const SDL_STANDARD_GRAVITY = 9.80665

/**
 * Gyroscope sensor
 *
 * The gyroscope returns the current rate of rotation in radians per second.
 * The rotation is positive in the counter-clockwise direction. That is,
 * an observer looking from a positive location on one of the axes would
 * see positive rotation on that axis when it appeared to be rotating
 * counter-clockwise.
 *
 * values[0]: Angular speed around the x axis (pitch)
 * values[1]: Angular speed around the y axis (yaw)
 * values[2]: Angular speed around the z axis (roll)
 *
 * For phones held in portrait mode and game controllers held in front of you,
 * the axes are defined as follows:
 * -X ... +X : left ... right
 * -Y ... +Y : bottom ... top
 * -Z ... +Z : farther ... closer
 *
 * The axis data is not changed when the phone or controller is rotated.
 *
 * \sa SDL_GetDisplayOrientation()
 */

/* Function prototypes */

/**
 * Locking for multi-threaded access to the sensor API
 *
 * If you are using the sensor API or handling events from multiple threads
 * you should use these locking functions to protect access to the sensors.
 *
 * In particular, you are guaranteed that the sensor list won't change, so the
 * API functions that take a sensor index will be valid, and sensor events
 * will not be delivered.
 */
//extern DECLSPEC void SDLCALL SDL_LockSensors(void);
func SDL_LockSensors() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LockSensors").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_UnlockSensors(void);
func SDL_UnlockSensors() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockSensors").Call()
	if t == 0 {

	}
	return
}

/**
 * Count the number of sensors attached to the system right now.
 *
 * \returns the number of sensors detected.
 */
//extern DECLSPEC int SDLCALL SDL_NumSensors(void);
func SDL_NumSensors() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_NumSensors").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the implementation dependent name of a sensor.
 *
 * \param device_index The sensor to obtain name from
 * \returns the sensor name, or NULL if `device_index` is out of range.
 */
//extern DECLSPEC const char *SDLCALL SDL_SensorGetDeviceName(int device_index);
func SDL_SensorGetDeviceName(device_index sdlcommon.FInt) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetDeviceName").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the type of a sensor.
 *
 * \param device_index The sensor to get the type from
 * \returns the SDL_SensorType, or `SDL_SENSOR_INVALID` if `device_index` is
 *          out of range.
 */
//extern DECLSPEC SDL_SensorType SDLCALL SDL_SensorGetDeviceType(int device_index);
func SDL_SensorGetDeviceType(device_index sdlcommon.FInt) (res SDL_SensorType) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetDeviceType").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = SDL_SensorType(t)
	return
}

/**
 * Get the platform dependent type of a sensor.
 *
 * \param device_index The sensor to check
 * \returns the sensor platform dependent type, or -1 if `device_index` is out
 *          of range.
 */
//extern DECLSPEC int SDLCALL SDL_SensorGetDeviceNonPortableType(int device_index);
func SDL_SensorGetDeviceNonPortableType(device_index sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetDeviceNonPortableType").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the instance ID of a sensor.
 *
 * \param device_index The sensor to get instance id from
 * \returns the sensor instance ID, or -1 if `device_index` is out of range.
 */
//extern DECLSPEC SDL_SensorID SDLCALL SDL_SensorGetDeviceInstanceID(int device_index);
func SDL_SensorGetDeviceInstanceID(device_index sdlcommon.FInt) (res SDL_SensorID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetDeviceInstanceID").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = SDL_SensorID(t)
	return
}

/**
 * Open a sensor for use.
 *
 * \param device_index The sensor to open
 * \returns an SDL_Sensor sensor object, or NULL if an error occurred.
 */
//extern DECLSPEC SDL_Sensor *SDLCALL SDL_SensorOpen(int device_index);
func SDL_SensorOpen(device_index sdlcommon.FInt) (res *SDL_Sensor) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorOpen").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = (*SDL_Sensor)(unsafe.Pointer(t))
	return
}

/**
 * Return the SDL_Sensor associated with an instance id.
 *
 * \param instance_id The sensor from instance id
 * \returns an SDL_Sensor object.
 */
//extern DECLSPEC SDL_Sensor *SDLCALL SDL_SensorFromInstanceID(SDL_SensorID instance_id);
func SDL_SensorFromInstanceID(instance_id SDL_SensorID) (res *SDL_Sensor) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorFromInstanceID").Call(
		uintptr(instance_id),
	)
	if t == 0 {

	}
	res = (*SDL_Sensor)(unsafe.Pointer(t))
	return
}

/**
 * Get the implementation dependent name of a sensor
 *
 * \param sensor The SDL_Sensor object
 * \returns the sensor name, or NULL if `sensor` is NULL.
 */
//extern DECLSPEC const char *SDLCALL SDL_SensorGetName(SDL_Sensor *sensor);
func (sensor *SDL_Sensor) SDL_SensorGetName() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetName").Call(
		uintptr(unsafe.Pointer(sensor)),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the type of a sensor.
 *
 * \param sensor The SDL_Sensor object to inspect
 * \returns the SDL_SensorType type, or `SDL_SENSOR_INVALID` if `sensor` is
 *          NULL.
 */
//extern DECLSPEC SDL_SensorType SDLCALL SDL_SensorGetType(SDL_Sensor *sensor);
func (sensor *SDL_Sensor) SDL_SensorGetType() (res SDL_SensorType) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetType").Call(
		uintptr(unsafe.Pointer(sensor)),
	)
	if t == 0 {

	}
	res = SDL_SensorType(t)
	return
}

/**
 * Get the platform dependent type of a sensor.
 *
 * \param sensor The SDL_Sensor object to inspect
 * \returns the sensor platform dependent type, or -1 if `sensor` is NULL.
 */
//extern DECLSPEC int SDLCALL SDL_SensorGetNonPortableType(SDL_Sensor *sensor);
func (sensor *SDL_Sensor) SDL_SensorGetNonPortableType() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetNonPortableType").Call(
		uintptr(unsafe.Pointer(sensor)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the instance ID of a sensor.
 *
 * \param sensor The SDL_Sensor object to inspect
 * \returns the sensor instance ID, or -1 if `sensor` is NULL.
 */
//extern DECLSPEC SDL_SensorID SDLCALL SDL_SensorGetInstanceID(SDL_Sensor *sensor);
func (sensor *SDL_Sensor) SDL_SensorGetInstanceID() (res SDL_SensorID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetInstanceID").Call(
		uintptr(unsafe.Pointer(sensor)),
	)
	if t == 0 {

	}
	res = SDL_SensorID(t)
	return
}

/**
 * Get the current state of an opened sensor.
 *
 * The number of values and interpretation of the data is sensor dependent.
 *
 * \param sensor The SDL_Sensor object to query
 * \param data A pointer filled with the current sensor state
 * \param num_values The number of values to write to data
 * \returns 0 or -1 if an error occurred.
 */
//extern DECLSPEC int SDLCALL SDL_SensorGetData(SDL_Sensor * sensor, float *data, int num_values);
func (sensor *SDL_Sensor) SDL_SensorGetData(data *sdlcommon.FFloat, num_values sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorGetData").Call(
		uintptr(unsafe.Pointer(sensor)),
		uintptr(unsafe.Pointer(data)),
		uintptr(num_values),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Close a sensor previously opened with SDL_SensorOpen().
 *
 * \param sensor The SDL_Sensor object to close
 */
//extern DECLSPEC void SDLCALL SDL_SensorClose(SDL_Sensor * sensor);
func (sensor *SDL_Sensor) SDL_SensorClose() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorClose").Call(
		uintptr(unsafe.Pointer(sensor)),
	)
	if t == 0 {

	}
	return
}

/**
 * Update the current state of the open sensors.
 *
 * This is called automatically by the event loop if sensor events are
 * enabled.
 *
 * This needs to be called from the thread that initialized the sensor
 * subsystem.
 */
//extern DECLSPEC void SDLCALL SDL_SensorUpdate(void);
func SDL_SensorUpdate() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SensorUpdate").Call()
	if t == 0 {

	}
	return
}
