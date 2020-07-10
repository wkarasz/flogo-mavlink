package mavlinkmissionstart

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/coerce"

	//"context"
	"fmt"
        "time"
	"strconv"
 
	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

var logMav = log.ChildLogger(log.RootLogger(), "mavlink-missionstart")
var activityMd = activity.ToMetadata(&Input{}, &Output{}, &Settings{})

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logMav.Errorf("Mavlink Mission Start Activity init error : %s ", err.Error())
	}
}

// New functioncommon
func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}
	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}
	if settings.Connection != "" {
		log.RootLogger().Infof("***Connection [%s]***",settings.Connection)
		mConn, toConnerr := coerce.ToConnection(settings.Connection)
		if toConnerr != nil {
			return nil, toConnerr
		}
		node := mConn.GetConnection().(*gomavlib.Node)
		return &Activity{settings: settings, n: node}, nil
	}
	return nil, nil
}


// Activity is a stub for your Activity implementation
type Activity struct {
	settings  *Settings
	n	  *gomavlib.Node
}

// Metadata implements activity.Activity.Metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

/*
// Cleanup method
func (a *Activity) Cleanup() error {
	logMav.Debugf("cleaning up Mavlink Mission Start activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	param1 := context.GetInput("param1").(string)
	param2 := context.GetInput("param2").(string)
	logMav.Debugf("Mission Start param1: [%s] param2: [%s]", param1, param2)
	log.RootLogger().Infof("Mission Start param1: [%s] param2: [%s]", param1, param2)

	commandrequest := ardupilotmega.MAV_CMD_MISSION_START		
	
	node := *a.n
	var Param1 float32 = 0
	var Param2 float32 = 0

	Param1_intermediate, err := strconv.ParseFloat(param1, 32)
	if err != nil {
	        logMav.Errorf("[Param1 convert error] [%s], %s\n", param1, err.Error())
    	}
	Param2_intermediate, err := strconv.ParseFloat(param2, 32)
	if err != nil {
	        logMav.Errorf("[Param2 convert error] [%s], %s\n", param2, err.Error())
	}	
	Param1 = float32(Param1_intermediate)
	Param2 = float32(Param2_intermediate)

	node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
        	TargetSystem: 0,        // System ID
                TargetComponent: 1,     // Component ID
                Command: commandrequest,
                Confirmation: 0,
                Param1: Param1,              // PARAM1, see MAV_CMD enum
		Param2: Param2,
	})


	// initialize timeout variables
        start := time.Now()
        counter := 1
        //resultpending := true

        for evt := range node.Events() {
                if frm, ok := evt.(*gomavlib.EventFrame); ok {

                        // if frm is the Message Command Ack && MAV_CMD_MISSION_START
                        if msg, ok := frm.Message().(*ardupilotmega.MessageCommandAck); ok && msg.Command == ardupilotmega.MAV_CMD_MISSION_START {
                                logMav.Debugf("MAV_CMD_MISSION_START\n")
                                logMav.Debugf("RESULT: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
				err = context.SetOutput("result","Operation Executed")
				if err != nil {
					return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
				}
				return true, nil
                        }
                 }

		// check if time.Since(start) > timeout, elapsed time (time.Since reports nanonseconds)
		if time.Since(start) > 1.5e9 {
			if counter > 5 {
				// Error out
				fmt.Errorf("Error - Autopilot did not reply in time\n")
				return true, nil
			} else {
				// Resend command
				node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
                          	      TargetSystem: 0,        // System ID
                                	TargetComponent: 1,     // Component ID
                                	Command: commandrequest,
                                	Confirmation: 0,
                                	Param1: Param1,              // PARAM1, see MAV_CMD enum
					Param2: Param2,
                        	})
				start = time.Now() // reset timeout stopwatch
				counter += 1
				logMav.Debugf("  Count: %d\n", counter)
			}
				
		}

/*                              // Check the result codes of the message, msg.Result
                                // 0 - accepted
                                // 1 - temporarily rejected
                                // 2 - denied
                                // 3 - unsupported
                                // 4 - failed
                                // 5 - result in progress
*/

        }

	//Set result in the output
	//actData := make(map[string]interface{})
	//actData["result"] = "Operation Executed"
	//output := &Output{}
	//output.Result = "Operate executed"
	//err = context.SetOutputObject(output)
	err = context.SetOutput("result","Operation Executed")
	if err != nil {
		return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
	}
	return true, nil
}
