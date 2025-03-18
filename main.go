package main

import (
	"fmt"
	"fyne-test/javaApp"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type Chat struct {
	Name string
	Id   int
}

func init() {
	log.SetFlags(log.Lshortfile | log.Ltime)

	go func() {
		conn, err := net.DialTimeout("tcp", "192.168.193.97:9988", time.Second*3)
		if err != nil {
			log.Println("Connect to logging server failed:", err)
		} else {
			log.SetOutput(io.MultiWriter(conn, os.Stdout))
		}
	}()

}

// go env -w CGO_CFLAGS="-O2 -g -IC:\java\jdk-23\include -IC:\java\jdk-23\include\win32 -fdeclspec"
// fyne package -os android -appID com.messenger.app -javaSource D:\Users\akuzm\GolandProjects\fyne-test\javaApp\java

func main() {
	a := app.New()

	jni, err := javaApp.RunOnFyne()
	if err != nil {
		log.Fatal(err)
	}
	defer jni.Close()

	go func() {
		contactsPermIsGrant, err := jni.RequestContactsPermission()
		if err != nil {
			log.Println("get contacts permission failed:", err)
		}
		log.Println("contacts perm isGrant: ", contactsPermIsGrant)
		time.Sleep(time.Second)
		if err := jni.Notify(); err != nil {
			log.Println("notify failed:", err)
		}
	}()

	resp := []Chat{
		{
			Id:   1,
			Name: "Chat 1",
		},
		{
			Id:   2,
			Name: "Chat 2",
		},
		{
			Id:   3,
			Name: "Chat 3",
		},
	}

	w := a.NewWindow("test")

	var buttons []fyne.CanvasObject
	for _, chat := range resp {
		button := widget.NewButton(chat.Name, func() {
			w.SetContent(widget.NewLabel(fmt.Sprint("Это чат с Id=", chat.Id)))
		})

		buttons = append(buttons, button)
	}

	buttons = append(buttons, widget.NewButton("Load contacts", func() {
		log.Println("clicked load contacts")
		contacts, err := jni.LoadContacts()
		if err != nil {
			log.Println(err)
			return
		}

		contactContainer := container.NewVBox()

		for _, contact := range contacts {
			contactContainer.Add(widget.NewLabel(contact.Name + " " + contact.Phone))
		}

		w.SetContent(contactContainer)
	}))

	c := container.NewVBox(buttons...)

	w.SetContent(c)
	w.ShowAndRun()
}
