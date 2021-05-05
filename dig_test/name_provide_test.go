package dig_test

import (
	"testing"

	"go.uber.org/dig"
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
