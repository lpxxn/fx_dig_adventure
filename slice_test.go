package fx_dig_adventure

import (
	"os"
	"testing"

	"go.uber.org/dig"
)

func TestGroup1(t *testing.T) {
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
	type inParams struct {
		dig.In

		StudentList []*Student `group:"stu"`
	}
	Info := func(params inParams) error {
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}
		return nil
	}
	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
}

func TestGroup2(t *testing.T) {
	type Student struct {
		Name string
		Age  int
	}
	type Rep struct {
		dig.Out
		StudentList []*Student `group:"stu,flatten"`
	}
	NewUser := func(name string, age int) func() Rep {
		return func() Rep {
			r := Rep{}
			r.StudentList = append(r.StudentList, &Student{
				Name: name,
				Age:  age,
			})
			return r
		}
	}

	container := dig.New()
	if err := container.Provide(NewUser("tom", 3)); err != nil {
		t.Fatal(err)
	}
	if err := container.Provide(NewUser("jerry", 1)); err != nil {
		t.Fatal(err)
	}
	type InParams struct {
		dig.In

		StudentList []*Student `group:"stu"`
	}
	Info := func(params InParams) error {
		for _, u := range params.StudentList {
			t.Log(u.Name, u.Age)
		}
		return nil
	}
	if err := container.Invoke(Info); err != nil {
		t.Fatal(err)
	}
	dig.Visualize(container, os.Stdout)
}
