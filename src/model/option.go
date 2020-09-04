package model

import (
	"github.com/Pengxn/go-xn/src/util/log"
)

// Option model, optional informations
type Option struct {
	ID    uint64 `json:"option_id,omitempty" xorm:"bigint(20) notnull autoincr pk 'option_id'"`
	Name  string `json:"option_name" xorm:"varvhar(255) notnull unique 'option_name'"`
	Value string `json:"option_value" xorm:"longtext notnull 'option_value'"`
}

// GetAllOptions returns all options.
func GetAllOptions() []Option {
	db := orm.NewSession()
	defer db.Close()

	var options []Option

	err := db.Omit("option_id").Find(&options)
	if err != nil {
		log.Errorf("GetAllOptions throw error: %s", err)
	}

	return options
}

// GetOptionByName returns an Option by 'option_name' if it exist.
// Not including 'option_id' field.
func GetOptionByName(optionName string) (bool, *Option) {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name: optionName,
	}

	has, err := db.Omit("option_id").Get(option)
	if err != nil {
		log.Errorf("GetOptionByName throw error: %s", err)
	}

	return has, option
}

// AddToOption adds option to DB `option` table, return boolean value.
// If result returned is `true`, insert option data to DB successful.
func AddToOption(option *Option) bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(option)
	if err != nil {
		log.Errorf("AddToOption throw error: %s", err)
	}

	isSuccess := false

	if affected > 0 {
		isSuccess = true
	}

	return isSuccess
}

// DeleteOptionByName deletes an Option by 'option_name'.
// If result returned is `true`, delete option data successful.
func DeleteOptionByName(optionName string) bool {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name: optionName,
	}

	affected, err := db.Delete(option)
	if err != nil {
		log.Errorf("DeleteOptionByName throw error: %s", err)
	}

	isSuccess := false

	if affected > 0 {
		isSuccess = true
	}

	return isSuccess
}

// UpdateOptionByName updates an option by 'option_name'.
// If result returned is `true`, update option data successful.
func UpdateOptionByName(option *Option) bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Where("option_name = ?", option.Name).
		Cols("option_value").
		Update(option)
	if err != nil {
		log.Errorf("UpdateOptionByName throw error: %s", err)
	}

	isSuccess := false

	if affected > 0 {
		isSuccess = true
	}

	return isSuccess
}

// OptionExist if the option_name of option exist
// If `true`, this option already exists.
func OptionExist(optionName string) bool {
	db := orm.NewSession()
	defer db.Close()

	has, err := db.Exist(&Option{
		Name: optionName,
	})
	if err != nil {
		log.Errorf("OptionExist throw error: %s", err)
	}

	return has
}
