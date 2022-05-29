package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

//typedef void (SDLCALL * SDL_WindowsMessageHook)(void *userdata, void *hWnd, unsigned int message, Uint64 wParam, Sint64 lParam);
type SDL_WindowsMessageHook = func(userdata sdlcommon.FVoidP, hWnd sdlcommon.FVoidP, message sdlcommon.FUint, wParam sdlcommon.FUint64T, lParam sdlcommon.FSint64) uintptr

/**
 * Set a callback for every Windows message, run before TranslateMessage().
 *
 * \param callback The SDL_WindowsMessageHook function to call.
 * \param userdata a pointer to pass to every iteration of `callback`
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowsMessageHook(SDL_WindowsMessageHook callback, void *userdata);
func SDL_SetWindowsMessageHook(callback SDL_WindowsMessageHook, userdata sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowsMessageHook").Call(
		sdlcommon.NewCallback(callback),
		userdata,
	)
	if t == 0 {

	}
	return
}

/**
 * Get the D3D9 adapter index that matches the specified display index.
 *
 * The returned adapter index can be passed to `IDirect3D9::CreateDevice` and
 * controls on which monitor a full screen application will appear.
 *
 * \param displayIndex the display index for which to get the D3D9 adapter
 *                     index
 * \returns the D3D9 adapter index on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.1.
 */
//extern DECLSPEC int SDLCALL SDL_Direct3D9GetAdapterIndex( int displayIndex );
func SDL_Direct3D9GetAdapterIndex(displayIndex sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Direct3D9GetAdapterIndex").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

type IDirect3DDevice9 struct {
}

/**
 * Get the D3D9 device associated with a renderer.
 *
 * Once you are done using the device, you should release it to avoid a
 * resource leak.
 *
 * \param renderer the renderer from which to get the associated D3D device
 * \returns the D3D9 device associated with given renderer or NULL if it is
 *          not a D3D9 renderer; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.1.
 */
//extern DECLSPEC IDirect3DDevice9* SDLCALL SDL_RenderGetD3D9Device(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderGetD3D9Device(displayIndex sdlcommon.FInt) (res *IDirect3DDevice9) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetD3D9Device").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = (*IDirect3DDevice9)(unsafe.Pointer(t))
	return
}

//typedef struct ID3D11Device ID3D11Device;
type ID3D11Device struct {
}

/**
 * Get the D3D11 device associated with a renderer.
 *
 * Once you are done using the device, you should release it to avoid a
 * resource leak.
 *
 * \param renderer the renderer from which to get the associated D3D11 device
 * \returns the D3D11 device associated with given renderer or NULL if it is
 *          not a D3D11 renderer; call SDL_GetError() for more information.
 */
//extern DECLSPEC ID3D11Device* SDLCALL SDL_RenderGetD3D11Device(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderGetD3D11Device(displayIndex sdlcommon.FInt) (res *ID3D11Device) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetD3D11Device").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = (*ID3D11Device)(unsafe.Pointer(t))
	return
}

/**
 * Get the DXGI Adapter and Output indices for the specified display index.
 *
 * The DXGI Adapter and Output indices can be passed to `EnumAdapters` and
 * `EnumOutputs` respectively to get the objects required to create a DX10 or
 * DX11 device and swap chain.
 *
 * Before SDL 2.0.4 this function did not return a value. Since SDL 2.0.4 it
 * returns an SDL_bool.
 *
 * \param displayIndex the display index for which to get both indices
 * \param adapterIndex a pointer to be filled in with the adapter index
 * \param outputIndex a pointer to be filled in with the output index
 * \returns SDL_TRUE on success or SDL_FALSE on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 2.0.2.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_DXGIGetOutputInfo( int displayIndex, int *adapterIndex, int *outputIndex );
func SDL_DXGIGetOutputInfo(displayIndex sdlcommon.FInt, adapterIndex *sdlcommon.FInt, outputIndex *sdlcommon.FInt) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DXGIGetOutputInfo").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(adapterIndex)),
		uintptr(unsafe.Pointer(outputIndex)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/* Platform specific functions for Linux */
//#ifdef __LINUX__

/**
 * Sets the UNIX nice value for a thread.
 *
 * This uses setpriority() if possible, and RealtimeKit if available.
 *
 * \param threadID the Unix thread ID to change priority of.
 * \param priority The new, Unix-specific, priority value.
 * \returns 0 on success, or -1 on error.
 */
//extern DECLSPEC int SDLCALL SDL_LinuxSetThreadPriority(Sint64 threadID, int priority);
func SDL_LinuxSetThreadPriority(threadID sdlcommon.FSint64, priority sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LinuxSetThreadPriority").Call(
		uintptr(threadID),
		uintptr(priority),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//#endif /* __LINUX__ */

/* Platform specific functions for iOS */
//#ifdef __IPHONEOS__

//#define SDL_iOSSetAnimationCallback(window, interval, callback, callbackParam) SDL_iPhoneSetAnimationCallback(window, interval, callback, callbackParam)
//extern DECLSPEC int SDLCALL SDL_iPhoneSetAnimationCallback(SDL_Window * window, int interval, void (*callback)(void*), void *callbackParam);
func (window *SDL_Window) SDL_iPhoneSetAnimationCallback(interval sdlcommon.FInt, callback *func(p sdlcommon.FVoidP), callbackParam sdlcommon.FVoidP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LinuxSetThreadPriority").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(interval),
		uintptr(unsafe.Pointer(callback)),
		callbackParam,
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//#define SDL_iOSSetEventPump(enabled) SDL_iPhoneSetEventPump(enabled)
//extern DECLSPEC void SDLCALL SDL_iPhoneSetEventPump(SDL_bool enabled);
func SDL_iPhoneSetEventPump(enabled bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_iPhoneSetEventPump").Call(
		uintptr(sdlcommon.CBool(enabled)),
	)
	if t == 0 {

	}
	return
}

//#endif /* __IPHONEOS__ */

/* Platform specific functions for Android */
//#ifdef __ANDROID__

/**
 * Get the Android Java Native Interface Environment of the current thread.
 *
 * This is the JNIEnv one needs to access the Java virtual machine from native
 * code, and is needed for many Android APIs to be usable from C.
 *
 * The prototype of the function in SDL's code actually declare a void* return
 * type, even if the implementation returns a pointer to a JNIEnv. The
 * rationale being that the SDL headers can avoid including jni.h.
 *
 * \returns a pointer to Java native interface object (JNIEnv) to which the
 *          current thread is attached, or 0 on error.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AndroidGetActivity
 */
//extern DECLSPEC void * SDLCALL SDL_AndroidGetJNIEnv(void);
func SDL_AndroidGetJNIEnv() (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidGetJNIEnv").Call()
	if t == 0 {

	}
	res = t
	return
}

/**
 * Retrieve the Java instance of the Android activity class.
 *
 * The prototype of the function in SDL's code actually declares a void*
 * return type, even if the implementation returns a jobject. The rationale
 * being that the SDL headers can avoid including jni.h.
 *
 * The jobject returned by the function is a local reference and must be
 * released by the caller. See the PushLocalFrame() and PopLocalFrame() or
 * DeleteLocalRef() functions of the Java native interface:
 *
 * https://docs.oracle.com/javase/1.5.0/docs/guide/jni/spec/functions.html
 *
 * \returns the jobject representing the instance of the Activity class of the
 *          Android application, or NULL on error.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AndroidGetJNIEnv
 */
//extern DECLSPEC void * SDLCALL SDL_AndroidGetActivity(void);
func SDL_AndroidGetActivity() (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidGetActivity").Call()
	if t == 0 {

	}
	res = t
	return
}

/**
 * Query Android API level of the current device.
 *
 * - API level 30: Android 11
 * - API level 29: Android 10
 * - API level 28: Android 9
 * - API level 27: Android 8.1
 * - API level 26: Android 8.0
 * - API level 25: Android 7.1
 * - API level 24: Android 7.0
 * - API level 23: Android 6.0
 * - API level 22: Android 5.1
 * - API level 21: Android 5.0
 * - API level 20: Android 4.4W
 * - API level 19: Android 4.4
 * - API level 18: Android 4.3
 * - API level 17: Android 4.2
 * - API level 16: Android 4.1
 * - API level 15: Android 4.0.3
 * - API level 14: Android 4.0
 * - API level 13: Android 3.2
 * - API level 12: Android 3.1
 * - API level 11: Android 3.0
 * - API level 10: Android 2.3.3
 *
 * \returns the Android API level.
 */
//extern DECLSPEC int SDLCALL SDL_GetAndroidSDKVersion(void);
func SDL_GetAndroidSDKVersion() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAndroidSDKVersion").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Query if the application is running on Android TV.
 *
 * \returns SDL_TRUE if this is Android TV, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsAndroidTV(void);
func SDL_IsAndroidTV() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsAndroidTV").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Query if the application is running on a Chromebook.
 *
 * \returns SDL_TRUE if this is a Chromebook, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsChromebook(void);
func SDL_IsChromebook() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsChromebook").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Query if the application is running on a Samsung DeX docking station.
 *
 * \returns SDL_TRUE if this is a DeX docking station, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsDeXMode(void);
func SDL_IsDeXMode() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsDeXMode").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Trigger the Android system back button behavior.
 */
//extern DECLSPEC void SDLCALL SDL_AndroidBackButton(void);
func SDL_AndroidBackButton() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidBackButton").Call()
	if t == 0 {

	}
	return
}

/**
  See the official Android developer guide for more information:
  http://developer.android.com/guide/topics/data/data-storage.html
*/
const SDL_ANDROID_EXTERNAL_STORAGE_READ = 0x01
const SDL_ANDROID_EXTERNAL_STORAGE_WRITE = 0x02

/**
 * Get the path used for internal storage for this application.
 *
 * This path is unique to your application and cannot be written to by other
 * applications.
 *
 * Your internal storage path is typically:
 * `/data/data/your.app.package/files`.
 *
 * \returns the path used for internal storage or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AndroidGetExternalStorageState
 */
//extern DECLSPEC const char * SDLCALL SDL_AndroidGetInternalStoragePath(void);
func SDL_AndroidGetInternalStoragePath() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidGetInternalStoragePath").Call()
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the current state of external storage.
 *
 * The current state of external storage, a bitmask of these values:
 * `SDL_ANDROID_EXTERNAL_STORAGE_READ`, `SDL_ANDROID_EXTERNAL_STORAGE_WRITE`.
 *
 * If external storage is currently unavailable, this will return 0.
 *
 * \returns the current state of external storage on success or 0 on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AndroidGetExternalStoragePath
 */
//extern DECLSPEC int SDLCALL SDL_AndroidGetExternalStorageState(void);
func SDL_AndroidGetExternalStorageState() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidGetExternalStorageState").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the path used for external storage for this application.
 *
 * This path is unique to your application, but is public and can be written
 * to by other applications.
 *
 * Your external storage path is typically:
 * `/storage/sdcard0/Android/data/your.app.package/files`.
 *
 * \returns the path used for external storage for this application on success
 *          or NULL on failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AndroidGetExternalStorageState
 */
//extern DECLSPEC const char * SDLCALL SDL_AndroidGetExternalStoragePath(void);
func SDL_AndroidGetExternalStoragePath() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidGetExternalStoragePath").Call()
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Request permissions at runtime.
 *
 * This blocks the calling thread until the permission is granted or denied.
 *
 * \param permission The permission to request.
 * \returns SDL_TRUE if the permission was granted, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_AndroidRequestPermission(const char *permission);
func SDL_AndroidRequestPermission(permission sdlcommon.FConstCharP) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidRequestPermission").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(permission))),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Shows an Android toast notification.
 *
 * Toasts are a sort of lightweight notification that are unique to Android.
 *
 * https://developer.android.com/guide/topics/ui/notifiers/toasts
 *
 * Shows toast in UI thread.
 *
 * For the `gravity` parameter, choose a value from here, or -1 if you don't
 * have a preference:
 *
 * https://developer.android.com/reference/android/view/Gravity
 *
 * \param message text message to be shown
 * \param duration 0=short, 1=long
 * \param gravity where the notification should appear on the screen.
 * \param xoffset set this parameter only when gravity >=0
 * \param yoffset set this parameter only when gravity >=0
 * \returns 0 if success, -1 if any error occurs.
 */
//extern DECLSPEC int SDLCALL SDL_AndroidShowToast(const char* message, int duration, int gravity, int xoffset, int yoffset);
func SDL_AndroidShowToast(message sdlcommon.FConstCharP, duration, gravity, xoffset, yoffset sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AndroidShowToast").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(message))),
		uintptr(duration),
		uintptr(gravity),
		uintptr(xoffset),
		uintptr(yoffset),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//#endif /* __ANDROID__ */

/* Platform specific functions for WinRT */
//#ifdef __WINRT__

/**
 *  \brief WinRT / Windows Phone path types
 */
type SDL_WinRT_Path = int32

const (
	/** \brief The installed app's root directory.
	  Files here are likely to be read-only. */
	SDL_WINRT_PATH_INSTALLED_LOCATION = 0

	/** \brief The app's local data store.  Files may be written here */
	SDL_WINRT_PATH_LOCAL_FOLDER

	/** \brief The app's roaming data store.  Unsupported on Windows Phone.
	  Files written here may be copied to other machines via a network
	  connection.
	*/
	SDL_WINRT_PATH_ROAMING_FOLDER

	/** \brief The app's temporary data store.  Unsupported on Windows Phone.
	  Files written here may be deleted at any time. */
	SDL_WINRT_PATH_TEMP_FOLDER
)

/**
 *  \brief WinRT Device Family
 */
type SDL_WinRT_DeviceFamily = int32

const (
	/** \brief Unknown family  */
	SDL_WINRT_DEVICEFAMILY_UNKNOWN = 0

	/** \brief Desktop family*/
	SDL_WINRT_DEVICEFAMILY_DESKTOP

	/** \brief Mobile family (for example smartphone) */
	SDL_WINRT_DEVICEFAMILY_MOBILE

	/** \brief XBox family */
	SDL_WINRT_DEVICEFAMILY_XBOX
)

/**
 * Retrieve a WinRT defined path on the local file system.
 *
 * Not all paths are available on all versions of Windows. This is especially
 * true on Windows Phone. Check the documentation for the given SDL_WinRT_Path
 * for more information on which path types are supported where.
 *
 * Documentation on most app-specific path types on WinRT can be found on
 * MSDN, at the URL:
 *
 * https://msdn.microsoft.com/en-us/library/windows/apps/hh464917.aspx
 *
 * \param pathType the type of path to retrieve, one of SDL_WinRT_Path
 * \returns a UCS-2 string (16-bit, wide-char) containing the path, or NULL if
 *          the path is not available for any reason; call SDL_GetError() for
 *          more information.
 *
 * \since This function is available since SDL 2.0.3.
 *
 * \sa SDL_WinRTGetFSPathUTF8
 */
//extern DECLSPEC const wchar_t * SDLCALL SDL_WinRTGetFSPathUNICODE(SDL_WinRT_Path pathType);
func SDL_WinRTGetFSPathUNICODE(pathType SDL_WinRT_Path) (res sdlcommon.FWcharT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WinRTGetFSPathUNICODE").Call(
		uintptr(pathType),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Retrieve a WinRT defined path on the local file system.
 *
 * Not all paths are available on all versions of Windows. This is especially
 * true on Windows Phone. Check the documentation for the given SDL_WinRT_Path
 * for more information on which path types are supported where.
 *
 * Documentation on most app-specific path types on WinRT can be found on
 * MSDN, at the URL:
 *
 * https://msdn.microsoft.com/en-us/library/windows/apps/hh464917.aspx
 *
 * \param pathType the type of path to retrieve, one of SDL_WinRT_Path
 * \returns a UTF-8 string (8-bit, multi-byte) containing the path, or NULL if
 *          the path is not available for any reason; call SDL_GetError() for
 *          more information.
 *
 * \since This function is available since SDL 2.0.3.
 *
 * \sa SDL_WinRTGetFSPathUNICODE
 */
//extern DECLSPEC const char * SDLCALL SDL_WinRTGetFSPathUTF8(SDL_WinRT_Path pathType);
func SDL_WinRTGetFSPathUTF8(pathType SDL_WinRT_Path) (res sdlcommon.FWcharT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WinRTGetFSPathUTF8").Call(
		uintptr(pathType),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Detects the device family of WinRT plattform at runtime.
 *
 * \returns a value from the SDL_WinRT_DeviceFamily enum.
 */
//extern DECLSPEC SDL_WinRT_DeviceFamily SDLCALL SDL_WinRTGetDeviceFamily();
func SDL_WinRTGetDeviceFamily() (res SDL_WinRT_DeviceFamily) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WinRTGetDeviceFamily").Call()
	if t == 0 {

	}
	res = SDL_WinRT_DeviceFamily(t)
	return
}

//#endif /* __WINRT__ */

/**
 * Query if the current device is a tablet.
 *
 * If SDL can't determine this, it will return SDL_FALSE.
 *
 * \returns SDL_TRUE if the device is a tablet, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsTablet(void);
func SDL_IsTablet() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsTablet").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/* Functions used by iOS application delegates to notify SDL about state changes */
//extern DECLSPEC void SDLCALL SDL_OnApplicationWillTerminate(void);
func SDL_OnApplicationWillTerminate() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationWillTerminate").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_OnApplicationDidReceiveMemoryWarning(void);
func SDL_OnApplicationDidReceiveMemoryWarning() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationDidReceiveMemoryWarning").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_OnApplicationWillResignActive(void);
func SDL_OnApplicationWillResignActive() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationWillResignActive").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_OnApplicationDidEnterBackground(void);
func SDL_OnApplicationDidEnterBackground() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationDidEnterBackground").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_OnApplicationWillEnterForeground(void);
func SDL_OnApplicationWillEnterForeground() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationWillEnterForeground").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_OnApplicationDidBecomeActive(void);
func SDL_OnApplicationDidBecomeActive() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationDidBecomeActive").Call()
	if t == 0 {

	}
	return
}

//#ifdef __IPHONEOS__
//extern DECLSPEC void SDLCALL SDL_OnApplicationDidChangeStatusBarOrientation(void);
func SDL_OnApplicationDidChangeStatusBarOrientation() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OnApplicationDidChangeStatusBarOrientation").Call()
	if t == 0 {

	}
	return
}

//#endif
