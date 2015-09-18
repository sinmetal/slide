import (
    "appengine/aetest"
    "testing"
)

func TestPost(t *testing.T) {
    c, err := aetest.NewContext(nil)
    if err != nil {
        t.Fatal(err)
    }
    defer c.Close()
}