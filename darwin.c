// Copyright (c) 2016 Vinzenz Feenstra, Red Hat Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the name of Red Hat Inc. nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

#include <Carbon/Carbon.h>

typedef struct uniCharBuffer{
    UniChar buffer[4];
    UniCharCount Length;
} uniCharBuffer;

uniCharBuffer translate(uint32 keyCode) {
    TISInputSourceRef currentKeyboard = TISCopyCurrentKeyboardInputSource();
    CFDataRef uchr = (CFDataRef)TISGetInputSourceProperty(currentKeyboard,
                                                          kTISPropertyUnicodeKeyLayoutData);

    const UCKeyboardLayout *keyboardLayout = (const UCKeyboardLayout*)CFDataGetBytePtr(uchr);

    UInt32 mods = GetCurrentEventKeyModifiers();
    UInt32 modsNoCtrlAlt = mods & ~(optionKey|controlKey);
    uniCharBuffer buffer = {};

    if(keyboardLayout) {

        UInt32 deadKeyState = 0;
        UniCharCount maxStringLength = 4;

        OSStatus status = UCKeyTranslate(keyboardLayout,
                                         keyCode, kUCKeyActionDown, 0,
                                         LMGetKbdType(), 0,
                                         &deadKeyState,
                                         maxStringLength,
                                         &buffer.Length, buffer.buffer);

        if (status == noErr && buffer.Length == 0 && deadKeyState)  {

            status = UCKeyTranslate(keyboardLayout,
                                    kVK_Space, kUCKeyActionDown, 0,
                                    LMGetKbdType(), 0,
                                    &deadKeyState,
                                    maxStringLength,
                                    &buffer.Length, buffer.buffer);

        }

    }
    return buffer;
}