# Barcode

Simple go command to read a barcode from a USB keyboard wedge barcode scanner, print to std out and exit

```
GOARCH=arm GOARM=7 go build
./barcode /dev/input/by-id/usb-USB_Adapter_USB_Device-event-kbd
```
