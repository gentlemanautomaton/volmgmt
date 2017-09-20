package ioctltype

// Device IO Control Code Types
const (
	DeviceBeep              = iota + 1 // 1 FILE_DEVICE_BEEP
	DeviceCDRom                        // 2 FILE_DEVICE_CD_ROM
	DeviceCDRomFileSystem              // 3 FILE_DEVICE_CD_ROM_FILE_SYSTEM
	DeviceController                   // 4 FILE_DEVICE_CONTROLLER
	DeviceDataLink                     // 5 FILE_DEVICE_DATALINK
	DeviceDFS                          // 6 FILE_DEVICE_DFS
	DeviceDisk                         // 7 FILE_DEVICE_DISK
	DeviceDiskFileSystem               // 8 FILE_DEVICE_DISK_FILE_SYSTEM
	DeviceFileSystem                   // 9 FILE_DEVICE_FILE_SYSTEM
	DeviceInportPort                   // 10 FILE_DEVICE_INPORT_PORT
	DeviceKeyboard                     // 11 FILE_DEVICE_KEYBOARD
	DeviceMailSlot                     // 12 FILE_DEVICE_MAILSLOT
	DeviceMidiIn                       // 13 FILE_DEVICE_MIDI_IN
	DeviceMidiOut                      // 14 FILE_DEVICE_MIDI_OUT
	DeviceMouse                        // 15 FILE_DEVICE_MOUSE
	DeviceMultiUNCProvider             // 16 FILE_DEVICE_MULTI_UNC_PROVIDER
	DeviceNamedPipe                    // 17 FILE_DEVICE_NAMED_PIPE
	DeviceNetwork                      // 18 FILE_DEVICE_NETWORK
	DeviceNetworkBrowser               // 19 FILE_DEVICE_NETWORK_BROWSER
	DeviceNetworkFileSystem            // 20 FILE_DEVICE_NETWORK_FILE_SYSTEM
	DeviceNull                         // 21 FILE_DEVICE_NULL
	DeviceParallelPort                 // 22 FILE_DEVICE_PARALLEL_PORT
	DevicePhysicalNetCard              // 23 FILE_DEVICE_PHYSICAL_NETCARD
	DevicePrinter                      // 24 FILE_DEVICE_PRINTER
	DeviceScanner                      // 25 FILE_DEVICE_SCANNER
	DeviceSerialMousePort              // 26 FILE_DEVICE_SERIAL_MOUSE_PORT
	DeviceSerialPort                   // 27 FILE_DEVICE_SERIAL_PORT
	DeviceScreen                       // 28 FILE_DEVICE_SCREEN
	DeviceSound                        // 29 FILE_DEVICE_SOUND
	DeviceStreams                      // 30 FILE_DEVICE_STREAMS
	DeviceTape                         // 31 FILE_DEVICE_TAPE
	DeviceTapeFileSystem               // 32 FILE_DEVICE_TAPE_FILE_SYSTEM
	DeviceTransport                    // 33 FILE_DEVICE_TRANSPORT
	DeviceUnknown                      // 34 FILE_DEVICE_UNKNOWN
	DeviceVideo                        // 35 FILE_DEVICE_VIDEO
	DeviceVirtualDisk                  // 36 FILE_DEVICE_VIRTUAL_DISK
	DeviceWaveIn                       // 37 FILE_DEVICE_WAVE_IN
	DeviceWaveOut                      // 38 FILE_DEVICE_WAVE_OUT
	Device8042Port                     // 39 FILE_DEVICE_8042_PORT
	DeviceNetworkRedirector            // 40 FILE_DEVICE_NETWORK_REDIRECTOR
	DeviceBattery                      // 41 FILE_DEVICE_BATTERY
	DeviceBusExtender                  // 42 FILE_DEVICE_BUS_EXTENDER
	DeviceModem                        // 43 FILE_DEVICE_MODEM
	DeviceVDM                          // 44 FILE_DEVICE_VDM
	DeviceMassStorage                  // 45 FILE_DEVICE_MASS_STORAGE
	DeviceSMB                          // 46 FILE_DEVICE_SMB
	DeviceKS                           // 47 FILE_DEVICE_KS
	DeviceChanger                      // 48 FILE_DEVICE_CHANGER
	DeviceSmartCard                    // 49 FILE_DEVICE_SMARTCARD
	DeviceACPI                         // 50 FILE_DEVICE_ACPI
	DeviceDVD                          // 51 FILE_DEVICE_DVD
	DeviceFullscreenVideo              // 52 FILE_DEVICE_FULLSCREEN_VIDEO
	DeviceDFSFileSystem                // 53 FILE_DEVICE_DFS_FILE_SYSTEM
	DeviceDFSVolume                    // 54 FILE_DEVICE_DFS_VOLUME
	DeviceSerEnum                      // 55 FILE_DEVICE_SERENUM
	DeviceTermSrv                      // 56 FILE_DEVICE_TERMSRV
	DeviceKSEC                         // 57 FILE_DEVICE_KSEC
)
