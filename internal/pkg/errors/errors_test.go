package errors

import (
	"errors"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestError(t *testing.T) {
	convey.Convey("have no cause", t, func() {
		err := New(1, "winter is comming")
		t.Log(err)

		convey.So(Type(err), convey.ShouldEqual, 1)
	})
	convey.Convey("have cause", t, func() {
		err := Wrap(1, New(2, "Hear me roar"), "Lannisters Always pay their debts.")
		t.Log(err)

		convey.So(Type(err), convey.ShouldEqual, 1)
	})
	convey.Convey("withmsg", t, func() {
		err := WithMsg(New(2, "Valar Morghulis"), "Valar Dohaeris")
		t.Log(err)

		convey.So(Type(err), convey.ShouldEqual, 2)
	})
	convey.Convey("unknow err", t, func() {
		err := errors.New("Ours Is the Fury")
		t.Log(err)

		convey.So(Type(err), convey.ShouldEqual, UnknownErr)
	})

	convey.Convey("nil err", t, func() {
		convey.So(Type(nil), convey.ShouldEqual, NoErr)
	})
}
