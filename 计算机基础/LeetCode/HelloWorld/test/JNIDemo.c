#include <jni.h>
#include "JNIDemo.h"


JNIEXPORT jstring JNICALL Java_JNIDemo_helloWorld (JNIEnv *env, jobject){
	return (*env)->NewStringUTF(env, "hello world!");
  
}
