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

// GetOptionByName return Option (not including 'option_id') by 'option_name'
func GetOptionByName(optionName string) *Option {
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
