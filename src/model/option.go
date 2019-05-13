package model

// Option model, website some information
type Option struct {
	ID    uint64 `json:"option_id,omitempty" xorm:"bigint(20) notnull autoincr pk 'option_id'"`
	Name  string `json:"option_name" xorm:"varvhar(255) notnull 'option_name'"`
	Value string `json:"option_value" xorm:"longtext notnull 'option_value'"`
}

// GetAllOptions return all options
func GetAllOptions() []Option {
	db := orm.NewSession()
	defer db.Close()

	var options []Option

	err := db.Omit("option_id").Find(&options)

	if err != nil {
		panic(err)
	}

	return options
}

// GetOptionByName return an Option by 'option_name' if it exist
// Not including 'option_id'
func GetOptionByName(optionName string) (bool, *Option) {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name: optionName,
	}

	has, err := db.Omit("option_id").Get(option)

	if err != nil {
		panic(err)
	}

	return has, option
}

// AddToOption will add option to DB option table
func AddToOption(option *Option) int64 {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.Insert(option)

	if err != nil {
		panic(err)
	}

	return affected
}

// DeleteOptionByName will delete an Option by 'option_name'
func DeleteOptionByName(optionName string) int64 {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name: optionName,
	}

	success, err := db.Delete(option)

	if err != nil {
		panic(err)
	}

	return success
}

// UpdateOptionByName will update an option by 'option_name'
func UpdateOptionByName(optionName string, optionValue string) int64 {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name:  optionName,
		Value: optionValue,
	}

	success, err := db.Cols("option_value").Update(option)

	if err != nil {
		panic(err)
	}

	return success
}

// OptionExist if the option_name of article exist
func OptionExist(optionName string) bool {
	db := orm.NewSession()
	defer db.Close()

	has, err := db.Exist(&Option{
		Name: optionName,
	})

	if err != nil {
		panic(err)
	}

	return has
}
