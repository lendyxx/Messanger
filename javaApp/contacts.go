package javaApp

import (
	"errors"
	"github.com/timob/jnigi"
)

func (jni *JNI) LoadContacts() ([]Contact, error) {
	jni.getContactsCh <- struct{}{}
	return <-jni.contactsCh, jni.err
}

func (jni *JNI) RequestContactsPermission() (bool, error) {
	jni.requestContactsPermissionCh <- struct{}{}
	return <-jni.contactsPermissionCh, jni.err
}

func (jni *JNI) loadContacts() ([]Contact, error) {
	contacts := jnigi.NewObjectRef("java/util/List")
	if err := jni.app.CallMethod(jni.env, "GetContacts", contacts); err != nil {
		return nil, err
	}

	var size int
	if err := contacts.CallMethod(jni.env, "size", &size); err != nil {
		return nil, err
	}

	contactsList := make([]Contact, size)

	for i := 0; i < size; i++ {
		var contactName []byte
		var contactPhone []byte

		contactObj := jnigi.NewObjectRef("java/lang/Object")

		if err := contacts.CallMethod(jni.env, "get", contactObj, i); err != nil {
			return nil, err
		}

		if err := contactObj.CallMethod(jni.env, "NameBytes", &contactName); err != nil {
			return nil, err
		}
		if err := contactObj.CallMethod(jni.env, "PhoneBytes", &contactPhone); err != nil {
			return nil, err
		}

		contactsList[i] = Contact{
			Name:  string(contactName),
			Phone: string(contactPhone),
		}

	}

	return contactsList, nil
}

func (jni *JNI) requestContactsPermission() (bool, error) {
	var val bool
	if err := jni.app.CallMethod(jni.env, "RequestContactsPermission", &val); err != nil {
		return val, errors.New("request permission error:" + err.Error())
	}
	return val, nil
}
