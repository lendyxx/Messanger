package com.messenger.myclass;


public class MyClass {
    public String TestA;
    public String TestB;

    public MyClass() {
        this.TestA = "test a!";
        this.TestB = "test b!";
    }

    public void callTest() {
        System.out.println("call testVoid");
        testVoid(true);
        System.out.println("call testVoid - ok");
      }

    public native void testVoid(boolean isGrant);

}