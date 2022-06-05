package hello_world

import "testing"

func TestHelloWorld(t *testing.T){

    got := HelloWorld()
    want := "Hello world!"

    if got != want {
        t.Errorf("got %q, wanted %q", got, want)
    }
}