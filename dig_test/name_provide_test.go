package dig_test

import (
	"context"
	"testing"

	"go.uber.org/dig"
	"go.uber.org/fx"
)

func TestName1(t *testing.T) {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	p1 := func() (*DSN, error) {
		return &DSN{Addr: "primary DSN"}, nil
	}
	if err := c.Provide(p1, dig.Name("primary")); err != nil {
		t.Fatal(err)
	}

	p2 := func() (*DSN, error) {
		return &DSN{Addr: "secondary DSN"}, nil
	}
	if err := c.Provide(p2, dig.Name("secondary")); err != nil {
		t.Fatal(err)
	}

	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}

	if err := c.Invoke(func(db DBInfo) {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
	}); err != nil {
		t.Fatal(err)
	}
}

func TestName2(t *testing.T) {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	type DSNRev struct {
		dig.Out
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	p1 := func() (DSNRev, error) {
		return DSNRev{PrimaryDSN: &DSN{Addr: "Primary DSN"},
			SecondaryDSN: &DSN{Addr: "Secondary DSN"}}, nil
	}

	if err := c.Provide(p1); err != nil {
		t.Fatal(err)
	}

	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	inv1 := func(db DBInfo) {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
	}

	if err := c.Invoke(inv1); err != nil {
		t.Fatal(err)
	}
}

func TestName3(t *testing.T) {
	type DSN struct {
		Addr string
	}
	c := dig.New()

	p1 := func() (*DSN, error) {
		return &DSN{Addr: "primary DSN"}, nil
	}
	if err := c.Provide(p1, dig.Name("primary")); err != nil {
		t.Fatal(err)
	}
	p2 := func() (*DSN, error) {
		return &DSN{Addr: "secondary DSN"}, nil
	}
	if err := c.Provide(p2, dig.Name("secondary")); err != nil {
		t.Fatal(err)
	}

	type DBInfo struct {
		dig.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}

	if err := c.Provide(func(db DBInfo) string {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
		return "aaa"
	}, dig.Name("str1")); err != nil {
		t.Fatal(err)
	}

	if err := c.Provide(func(db DBInfo) string {
		t.Log(db.PrimaryDSN)
		t.Log(db.SecondaryDSN)
		return "bbb"
	}, dig.Name("str2")); err != nil {
		t.Fatal(err)
	}

	type StrInfo struct {
		dig.In
		Str1 string `name:"str1"`
		Str2 string `name:"str2"`
	}
	if err := c.Invoke(func(s StrInfo) {
		t.Log(s.Str1, "   ", s.Str2)
	}); err != nil {
		t.Fatal(err)
	}
}

func TestName4(t *testing.T) {
	type DSN struct {
		Addr string
	}

	p1 := func() (*DSN, error) {
		return &DSN{Addr: "primary DSN"}, nil
	}
	p2 := func() (*DSN, error) {
		return &DSN{Addr: "secondary DSN"}, nil
	}

	type DBInfo struct {
		fx.In
		PrimaryDSN   *DSN `name:"primary"`
		SecondaryDSN *DSN `name:"secondary"`
	}
	type StrInfo struct {
		fx.In
		Str1 string `name:"str1"`
		Str2 string `name:"str2"`
	}

	app := fx.New(fx.Provide(fx.Annotated{
		Name:   "primary",
		Target: p1,
	}), fx.Provide(fx.Annotated{
		Name:   "secondary",
		Target: p2,
	}), fx.Provide(fx.Annotated{
		Name: "str1",
		Target: func(db DBInfo) string {
			t.Log(db.PrimaryDSN)
			t.Log(db.SecondaryDSN)
			return "aaa"
		},
	}), fx.Provide(fx.Annotated{
		Name: "str2",
		Target: func(db DBInfo) string {
			t.Log(db.PrimaryDSN)
			t.Log(db.SecondaryDSN)
			return "bbb"
		},
	}), fx.Invoke(func(s StrInfo) {
		t.Log(s.Str1, "   ", s.Str2)
	}))
	app.Start(context.Background())
}
