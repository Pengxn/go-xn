package model

// Option model, optional informations
type Option struct {
	ID    uint64 `json:"option_id,omitempty" xorm:"bigint(20) notnull autoincr pk 'option_id'"`
	Name  string `json:"option_name" xorm:"varchar(255) notnull unique 'option_name'"`
	Value string `json:"option_value" xorm:"longtext notnull 'option_value'"`
}

// GetAllOptions returns all options.
func GetAllOptions() ([]Option, error) {
	db := orm.NewSession()
	defer db.Close()

	var options []Option
	err := db.Omit("option_id").Find(&options)

	return options, err
}

// GetOptionByName returns an Option by 'option_name' if it exist.
// Only include 'option_value' field.
func GetOptionByName(optionName string) (bool, Option, error) {
	db := orm.NewSession()
	defer db.Close()

	var option Option
	has, err := db.Select("option_value").
		Where("option_name = ?", optionName).
		Get(&option)

	return has, option, err
}

// AddOption adds option to DB `option` table, return boolean value.
// If result returned is `true`, insert option data to DB successful.
func AddOption(option *Option) (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(option)

	return affected > 0, err
}

// DeleteOptionByName deletes an Option by 'option_name'.
// If result returned is `true`, delete option data successful.
func DeleteOptionByName(optionName string) (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Delete(&Option{
		Name: optionName,
	})

	return affected > 0, err
}

// UpdateOptionByName updates an option by 'option_name'.
// If result returned is `true`, update option data successful.
func UpdateOptionByName(option Option) (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Where("option_name = ?", option.Name).
		Cols("option_value").
		Update(option)

	return affected > 0, err
}

// OptionExist if the option_name of option exist
// If `true`, this option already exists.
func OptionExist(optionName string) (bool, error) {
	db := orm.NewSession()
	defer db.Close()

	return db.Exist(&Option{
		Name: optionName,
	})
}
