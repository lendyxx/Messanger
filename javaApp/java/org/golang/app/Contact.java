package org.golang.app;

public class Contact {
    public String Name;
    public String Phone;

    public Contact(String name, String phone) {
        this.Name = name;
        this.Phone = phone;
    }

    public byte[] NameBytes() {
        return this.Name.getBytes();
    }

    public byte[] PhoneBytes() {
        return this.Phone.getBytes();
    }
}