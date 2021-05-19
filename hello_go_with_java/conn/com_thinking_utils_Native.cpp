//
// Created by yong on 5/19/21.
//

#include <iostream>
#include "com_thinking_utils_Native.h"

using namespace std;

JNIEXPORT void JNICALL Java_com_thinking_utils_Native_SayHello
        (JNIEnv *, jclass) {
    cout << "hello fuck" << endl;
}