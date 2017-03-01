#!/bin/bash
# Uploading to mobile
adb push events.sh /sdcard/tmp/
#adb shell chmod 755 /sdcard/tmp/events.sh
# Exec script
adb shell sh /sdcard/tmp/events.sh
