package model

// Option model, website some information
type Option struct {
	ID    uint64 `json:"option_id,omitempty" xorm:"bigint(20) notnull autoincr pk 'option_id'"`
	Name  string `json:"option_name" xorm:"varvhar(255) notnull 'option_name'"`
	Value string `json:"option_value" xorm:"longtext notnull 'option_value'"`
}

// FindAllOptions return all options
func FindAllOptions() []Option {
	db := orm.NewSession()
	defer db.Close()

	var options []Option

	err := db.Omit("option_id").Find(&options)

	if err != nil {
		panic(err)
	}

	return options
}

// OptionExist if the option_name of article exist
func OptionExist(optionName string) bool {
	db := orm.NewSession()
	defer db.Close()

	has, _ := db.Exist(&Option{
		Name: optionName,
	})

	return has
}

// OptionByName return Option (not including 'option_id') by 'option_name'
func OptionByName(optionName string) *Option {
	db := orm.NewSession()
	defer db.Close()

	option := &Option{
		Name: optionName,
	}

	err := db.Find(option)

	if err != nil {
		panic(err)
	}

	return option
}
