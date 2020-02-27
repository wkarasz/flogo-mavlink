package mavlinksetmode

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
	Data interface{} `md:"SET_MODE,required"` // The JSON Object that will serve as the input data
}

//FromMap method
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Connection, _ = values["connection"]
	i.Data, _ = values["SET_MODE"]
	return nil
}

//ToMap method
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"connection": i.Connection,
		"SET_MODE": i.Data,
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
