package fx_dig_adventure

import (
	"testing"

	"go.uber.org/dig"
)

func TestOption1(t *testing.T) {
	type Student struct {
		dig.Out
		Name string
		Age  *int `option:"true"`
	}

	c := dig.New()
	if err := c.Provide(func() Student {
		return Student{
			Name: "Tom",
		}
	}); err != nil {
		t.Fatal(err)
	}

	if err := c.Invoke(func(n string, age *int) {
		t.Logf("name: %s", n)
		if age == nil {
			t.Log("age is nil")
		} else {
			t.Logf("age: %d", age)
		}
	}); err != nil {
		t.Fatal(err)
	}
}
