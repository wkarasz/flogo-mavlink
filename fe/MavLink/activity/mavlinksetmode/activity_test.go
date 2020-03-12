package mavlinksetmode

import (
	"io/ioutil"
	"testing"
	"encoding/json"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/mapper"
	"github.com/project-flogo/core/data/resolve"
	"github.com/project-flogo/core/support"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/support/test"
	_ "github.com/wkarasz/flogo-mavlink/fe/MavLink/connector/mavlink"
	"github.com/stretchr/testify/assert"
)

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
	},
	"input" : {
		"SET_MODE": "MAV_MODE_PREFLIGHT"
	}
}`

var activityMetadata *activity.Metadata

func getTestJsonMetadata() string {
        jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
        if err != nil {
                panic("No Json Metadata found for activity.json path")
        }
        return string(jsonMetadataBytes)
}

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil{
			panic("No Json Metadata found for activity.json path")
		}

		activityMetadata = activity.ToMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {
	// New factory
	log.RootLogger().Info("***TEST : Executing Zero start***")
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(settingsjson), &m)

	log.RootLogger().Infof("Connection Input Settings are : %v", m["settings"])
	log.RootLogger().Infof("Activity Input Settings are : %v", m["input"])

	assert.Nil(t, err)

	mf := mapper.NewFactory(resolve.GetBasicResolver())
	support.RegisterAlias("connection", "connection", "github.com/wkarasz/flogo-mavlink/fe2/MavLink/connector/mavlink")
	log.RootLogger().Infof("====Settings=====\n%s", m["settings"])
	iCtx := test.NewActivityInitContext(m["settings"], mf)
	act, err1 := New(iCtx)
	assert.Nil(t, err1)

	tc := test.NewActivityContext(act.Metadata())
	// Setting inputs
	tc.SetInput("SET_MODE","MAV_MODE_PREFLIGHT")
	_, err = act.Eval(tc)
	
	//act := NewActivity(getActivityMetadata())

	//if act == nil {
	//	t.Error("Activity Not Created")
	//	t.Fail()
	//	return
	//}
}

/*
func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("panic during execution: %v", r)
		}
	}()

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())

	//setup attrs
	tc.SetInput("SET_MODE", "MAV_MODE_TEST_DISARMED")
	act.Eval(tc)

	//check result attr
	result := tc.GetOutput("result")
	assert.Equal(t, result, "OBDII by elm329@gmail.com")
}
*/
