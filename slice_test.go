package fx_dig_adventure

import (
	"testing"

	"go.uber.org/dig"
)

func TestSlice1(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	NewUser := func(name string, age int) func() *Student {
		return func() *Student {
			return &Student{name, age}
		}
	}
	container := dig.New()
	if err := container.Provide(NewUser("tom", 3), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1), dig.Group("stu")); err != nil {
		t.Fatal(err)
	}
	type UserParams struct {
		dig.In

		StudentList []*Student `group:"stu"`
	}
	Info := func(params UserParams) error {
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}
		return nil
	}
	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
}
