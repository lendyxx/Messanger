package javaApp

import (
	"errors"
	"fyne.io/fyne/v2/driver"
	"github.com/timob/jnigi"
	"log"
	"unsafe"
)

func RunOnFyne() (*JNI, error) {
	jni := NewJNI()

	go func() {
		defer func() {
			if rec := recover(); rec != nil {
				log.Fatal(rec)
			}
		}()

		if err := driver.RunNative(func(a any) (err error) {
			switch v := a.(type) {
			case *driver.AndroidContext:
				_, env := jnigi.UseJVM(unsafe.Pointer(v.VM), unsafe.Pointer(v.Env), unsafe.Pointer(v.Ctx))
				return jni.Run(env)

			default:
				return errors.New("JNI work on android only")
			}

		}); err != nil {
			log.Fatal(err)
		}
	}()

	jni.Wait()

	return jni, nil
}

func test() error {
	/*log.Printf("AndroidContext: %+v", v)

	_, env := jnigi.UseJVM(unsafe.Pointer(v.VM), unsafe.Pointer(v.Env), unsafe.Pointer(v.Ctx))

	my, err := env.NewObject("com/messenger/myclass/MyClass")
	if err != nil {
		return errors.New("create com/messenger/myclass/MyClass error: " + err.Error())
	}

	testA := jnigi.NewObjectRef("java/lang/String")
	if err := my.GetField(env, "TestA", testA); err != nil {
		return err
	}

	var testAStr []byte
	if err := testA.CallMethod(env, "getBytes", &testAStr); err != nil {
		return err
	}

	log.Printf("TestA: %s", string(testAStr))*/
	return nil
}
