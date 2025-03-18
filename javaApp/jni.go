package javaApp

import (
	"errors"
	"fmt"
	"fyne-test/gores"
	"github.com/timob/jnigi"
	"log"
	"sync"
)

type Contact struct {
	Name  string
	Phone string
}

type JNI struct {
	env    *jnigi.Env
	initWg sync.WaitGroup
	close  chan struct{}
	err    error

	app                         *jnigi.ObjectRef
	getContactsCh               chan struct{}
	contactsCh                  chan []Contact
	requestContactsPermissionCh chan struct{}
	contactsPermissionCh        chan bool
	notifyCh                    chan struct{}
}

func NewJNI() *JNI {
	jni := &JNI{
		close: make(chan struct{}),

		getContactsCh: make(chan struct{}),
		contactsCh:    make(chan []Contact),

		requestContactsPermissionCh: make(chan struct{}),
		contactsPermissionCh:        make(chan bool),

		notifyCh: make(chan struct{}),
	}
	jni.initWg.Add(1)
	return jni
}

func (jni *JNI) Run(env *jnigi.Env) error {
	defer func() {
		log.Println("EXIT FROM jni.Run")
		if err := recover(); err != nil {
			log.Println("Recovered: ", err)
		}
	}()

	jni.env = env

	if err := jni.initAppActivity(); err != nil {
		log.Println(err)
		return errors.New("init app activity failed: " + err.Error())
	}

	jni.initWg.Done()
	for {
		log.Println("start listen")
		select {
		case <-jni.close:
			return nil

		case <-jni.getContactsCh:
			var contacts []Contact
			contacts, jni.err = jni.loadContacts()
			jni.contactsCh <- contacts

		case <-jni.requestContactsPermissionCh:
			var isGrant bool
			isGrant, jni.err = jni.requestContactsPermission()
			jni.contactsPermissionCh <- isGrant

		case <-jni.notifyCh:
			jni.err = jni.notify()
		}
		jni.err = nil
	}
}

func (jni *JNI) Wait() {
	jni.initWg.Wait()
}

func (jni *JNI) Close() {
	jni.close <- struct{}{}
}

func (jni *JNI) initAppActivity() (err error) {
	if err := jni.env.CallStaticMethod("org/golang/app/GoNativeActivity", "waitInit", nil); err != nil {
		return fmt.Errorf("waitInit failed: %w", err)
	}

	jni.app = jnigi.NewObjectRef("org/golang/app/GoNativeActivity")

	if err := jni.env.GetStaticField("org/golang/app/GoNativeActivity", "goNativeActivity", jni.app); err != nil {
		return fmt.Errorf("get app activity failed: %w", err)
	}

	if jni.app.IsNil() {
		return errors.New("app object is nil")
	}

	if err := jni.env.CallStaticMethod("org/golang/app/GoNativeActivity", "setNotifyIcon", nil, gores.Ic_notification_png); err != nil {
		return fmt.Errorf("set notify icon failed: %w", err)
	}

	return nil
}
