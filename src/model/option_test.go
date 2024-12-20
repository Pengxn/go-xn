package model

import (
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/Pengxn/go-xn/src/util/log"
)

func init() {
	CleanTable(testORM, &Option{})
	db := testORM.NewSession()
	defer db.Close()
	num, err := db.Insert(
		&Option{Name: "site_url", Value: "https://xn--02f.com", ID: 1},
		&Option{Name: "will_be_deleted", Value: "data_will_be_deleted"},
		&Option{Name: "will_be_update", Value: "data_will_be_update"},
		&Option{Name: "test_exist", Value: "1"},
	)
	if err != nil {
		log.Infof("Init data to option table error: %+v, num: %d", err, num)
	}
}

func TestGetAllOptions(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("Test GetAllOptions.", t, func() {
		options, err := GetAllOptions()

		So(err, ShouldBeNil)
		So(options, ShouldResemble, []Option{
			{Name: "site_url", Value: "https://xn--02f.com"},
			{Name: "will_be_deleted", Value: "data_will_be_deleted"},
			{Name: "will_be_update", Value: "data_will_be_update"},
			{Name: "test_exist", Value: "1"},
		})
	})
}

func TestGetOptionByName(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("When option record is exists.", t, func() {
		has, option, err := GetOptionByName("site_url")

		So(err, ShouldBeNil)
		So(has, ShouldEqual, true)
		So(option, ShouldResemble, Option{Value: "https://xn--02f.com"})
	})

	Convey("When option record does not exist.", t, func() {
		has, option, err := GetOptionByName("not_exist")

		So(err, ShouldBeNil)
		So(has, ShouldEqual, false)
		So(option, ShouldResemble, Option{})
	})
}

func TestAddToOption(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("Add option record successfully.", t, func() {
		success, err := AddOption(&Option{Name: "add_record", Value: "add_data"})
		So(err, ShouldBeNil)
		So(success, ShouldEqual, true)

		db := testORM.NewSession()
		defer db.Close()

		option := Option{Name: "add_record"}
		_, err = db.Omit("option_id").Get(&option)
		So(option, ShouldResemble, Option{Name: "add_record", Value: "add_data"})
		So(err, ShouldBeNil)
	})
}

func TestDeleteOptionByName(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("Delete option record successfully.", t, func() {
		success, err := DeleteOptionByName("will_be_deleted")
		So(err, ShouldBeNil)
		So(success, ShouldEqual, true)

		db := testORM.NewSession()
		defer db.Close()

		option := []Option{}
		err = db.Where("option_name = ?", "will_be_deleted").Find(&option)
		So(len(option), ShouldEqual, 0)
		So(err, ShouldBeNil)
	})
}

func TestUpdateOptionByName(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("Update option record successfully.", t, func() {
		success, err := UpdateOptionByName(Option{Name: "will_be_update", Value: "update_data"})
		So(err, ShouldBeNil)
		So(success, ShouldEqual, true)

		db := testORM.NewSession()
		defer db.Close()

		option := Option{Name: "will_be_update"}
		_, err = db.Omit("option_id").Get(&option)
		So(option, ShouldResemble, Option{Name: "will_be_update", Value: "update_data"})
		So(err, ShouldBeNil)
	})
}

func TestOptionExist(t *testing.T) {
	p := ApplyGlobalVar(&orm, testORM)
	defer p.Reset()

	Convey("When option record is exists.", t, func() {
		success, err := OptionExist("test_exist")
		So(err, ShouldBeNil)
		So(success, ShouldEqual, true)
	})

	Convey("When option record does not exist.", t, func() {
		success, err := OptionExist("not_exist")
		So(err, ShouldBeNil)
		So(success, ShouldEqual, false)
	})
}
