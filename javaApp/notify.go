package javaApp

func (jni *JNI) Notify() error {
	jni.notifyCh <- struct{}{}
	return jni.err
}

func (jni *JNI) notify() error {
	if err := jni.app.CallMethod(jni.env, "showNotify", nil); err != nil {
		return err
	}
	return nil
}
