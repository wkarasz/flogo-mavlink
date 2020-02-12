package mavlinkdatastream

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	
	_ "github.com/wkarasz/flogo-mavlink/fe2/MavLink/connector/mavlink"
	//"github.com/project-flogo/core/trigger"
	//"github.com/project-flogo/core/data/mapper"
	//"github.com/project-flogo/core/data/resolve"
	//"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/log"
	//"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

//var triggerMetadata *trigger.Metadata
var jsonTestMetadata = getTestJsonMetadata()

const settingsjson = `{
	"settings": {
		"connection": {
			"name": "myConn",
			"description": "Local Mavlink Connection",
			"settings": {
				"name": "myConn",
				"description": "Local Mavlink Connection",
				"port": "14550"
			}
		}
	}
}`

func getTestJsonMetadata() string {
	jsonMetadataBytes, err := ioutil.ReadFile("trigger.json")
	if err != nil {
		panic("No Json Metadata found for trigger.json path")
	}
	return string(jsonMetadataBytes)
}

func TestCreate(t *testing.T) {
	// New factory
	log.RootLogger().Info("***TEST : Executing Zero start***")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson), &m)
	
	log.RootLogger().Infof("Input Settings are : %v", m["settings"])
	assert.Nil(t, err1)

	//mf := mapper.NewFactory(resolve.GetBasicResolver())

	//support.RegisterAlias("connection", "connection", "github.com/wkarasz/flogo-mavlink/fe/connection")

	//iCtx := test.NewActivityInitContext(m["settings"],mf)
	//act, err := New(iCtx)
	//assert.Nil(t, err)
}

/*
func Test_One(t *testing.T) {
	log.RootLogger().Info("****TEST : Executing One start****")
	m := make(map[string]interface{})
	err1 := json.Unmarshal([]byte(settingsjson), &m)

	log.RootLogger().Infof("Input Settings are : %v", m["settings"])
	assert.Nil(t, err1)

	mf := mapper.NewFactory(resolve.GetBasicResolver())

	support.RegisterAlias("connection", "connection", "github.com/ayh20/flogo-components/activity/mqtt/connection")

	//fmt.Println("=======Settings========", m["settings"])
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err := New(iCtx)
	assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("message", "My test message1")

	_, err = act.Eval(tc)

	// Getting outputs
	testOutput := tc.GetOutput("result")
	jsonOutput, _ := json.Marshal(testOutput)
	log.RootLogger().Infof("jsonOutput is : %s", string(jsonOutput))
	log.RootLogger().Info("****TEST : Executing Find One ends****")
	assert.Nil(t, err)
}
*/

