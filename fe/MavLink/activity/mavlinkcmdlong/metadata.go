package mavlinkcmdlong

import (
	"github.com/project-flogo/core/data/coerce"
)

// Settings structure
type Settings struct {
	Connection string `md:"connection,required"` // The Mavlink connection
}

// HandlerSettings structure
type HandlerSettings struct {
//	Database     string `md:"databaseName,required"` // MongoDB Database name
}

//Input structure
type Input struct {
	Connection interface{} `md:"connection,required"`
	MavCommand interface{} `md:"mavcommand"`
	Param1 interface{} `md:"param1"` // The JSON Object that will serve as the input data
	Param2 interface{} `md:"param2"` //
	Param3 interface{} `md:"param3"`
	Param4 interface{} `md:"param4"`
	Param5 interface{} `md:"param5"`
	Param6 interface{} `md:"param6"`
	Param7 interface{} `md:"param7"`
}

//FromMap method
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Connection, _ = values["connection"]
	i.MavCommand, _ = values["mavcommand"]
	i.Param1, _ = values["param1"]
	i.Param2, _ = values["param2"]
	i.Param3, _ = values["param3"]
	i.Param4, _ = values["param4"]
	i.Param5, _ = values["param5"]
	i.Param6, _ = values["param6"]
	i.Param7, _ = values["param7"]
	return nil
}

//ToMap method
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"connection": i.Connection,
		"mavcommand": i.MavCommand,
		"param1": i.Param1,
		"param2": i.Param2,
		"param3": i.Param3,
		"param4": i.Param4,
		"param5": i.Param5,
		"param6": i.Param6,
		"param7": i.Param7,
	}
}

// Output structure
type Output struct {
	//Output map[string]interface{} `md:"output"` //The Output of the trigger
	Result	string `md:"result"`
}

// ToMap method for Output
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"result": o.Result,
	}
}

// FromMap method for Output
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Result, err = coerce.ToString(values["result"])
	if err != nil {
		return err
	}

	return nil
}
