import (
    "appengine/aetest"
    "testing"
)

func TestPost(t *testing.T) {
	t.Parallel()
	
    c, err := aetest.NewContext(nil)
    if err != nil {
        t.Fatal(err)
    }
    defer c.Close()
}