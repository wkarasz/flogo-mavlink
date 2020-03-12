package mavlinkarm

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/coerce"

	//"context"
	"fmt"
        "time"
 
	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

var logArm = log.ChildLogger(log.RootLogger(), "mavlink-arm")
var activityMd = activity.ToMetadata(&Input{}, &Output{}, &Settings{})

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logArm.Errorf("Mavlink Arm/Disarm Activity init error : %s ", err.Error())
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
	logArm.Debugf("cleaning up Mavlink Arm/Disarm activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	cmd := context.GetInput("COMPONENT_MODE").(string)
	logArm.Debugf("Arm/Disarm request [%s]", cmd)
	log.RootLogger().Infof("Arm/Disarm request [%s]", cmd)

	commandrequest := ardupilotmega.MAV_CMD_COMPONENT_ARM_DISARM		
	
	node := *a.n
	var Param1 float32 = 0
        switch cmd {
		case "ARM":
                	Param1 = 1              // PARAM1, see MAV_CMD enum
		case "DISARM":
                	Param1 = 0              // PARAM1, see MAV_CMD enum
		default:
			logArm.Errorf("Unknown mode [%s]",cmd)
			return false, nil
	}
	
	node.WriteMessageAll(&ardupilotmega.MessageCommandLong{
        	TargetSystem: 0,        // System ID
                TargetComponent: 1,     // Component ID
                Command: commandrequest,
                Confirmation: 0,
                Param1: Param1,              // PARAM1, see MAV_CMD enum
	})


	// initialize timeout variables
        start := time.Now()
        counter := 1
        //resultpending := true

        for evt := range node.Events() {
                if frm, ok := evt.(*gomavlib.EventFrame); ok {

                        // if frm is the Message Command Ack && MAV_CMD_COMPONENT_ARM_DISARM
                        if msg, ok := frm.Message().(*ardupilotmega.MessageCommandAck); ok && msg.Command == ardupilotmega.MAV_CMD_COMPONENT_ARM_DISARM {
                                logArm.Debugf("MAV_CMD_COMPONENT_ARM_DISARM\n")
                                logArm.Debugf("RESULT: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
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
                        	})
				start = time.Now() // reset timeout stopwatch
				counter += 1
				logArm.Debugf("  Count: %d\n", counter)
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
