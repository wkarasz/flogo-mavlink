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
	Data interface{} `md:"SET_MODE,required"` // The JSON Object that will serve as the input data
}

// Output structure
type Output struct {
	Output map[string]interface{} `md:"output"` //The Output of the trigger
}

//FromMap method
func (i *Input) FromMap(values map[string]interface{}) error {
	i.Data, _ = values["SET_MODE"]
	return nil
}

//ToMap method
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"SET_MODE": i.Data,
	}
}

// ToMap method for Output
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"output": o.Output,
	}
}

// FromMap method for Output
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error

	o.Output, err = coerce.ToObject(values["output"])
	if err != nil {
		return err
	}

	return nil
}
