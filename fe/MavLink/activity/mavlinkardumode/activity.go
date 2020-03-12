package mavlinkardumode

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

var logArm = log.ChildLogger(log.RootLogger(), "mavlink-ardu-mode")
var activityMd = activity.ToMetadata(&Input{}, &Output{}, &Settings{})

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logArm.Errorf("Mavlink Ardu Mode Activity init error : %s ", err.Error())
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
	logArm.Debugf("cleaning up Mavlink Ardu Mode activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	cmd := context.GetInput("COMPONENT_MODE").(string)
	logArm.Debugf("Ardu Mode request [%s]", cmd)
	log.RootLogger().Infof("Ardu Mode request [%s]", cmd)
	
	var BaseMode uint8 = 89
	var CustomMode uint32 = 3

	node := *a.n
        switch cmd {
		// see https://diydrones.com/forum/topics/set-flight-modes-with-mavlink
		case "AUTO":
                	BaseMode = 89              
			CustomMode = 3
		case "LOITER":
                	BaseMode = 89
			CustomMode = 5
		case "ALTHOLD":
			BaseMode = 81
			CustomMode = 2
		case "RTL":
			BaseMode = 89
			CustomMode = 6             
		default:
			logArm.Errorf("Unknown mode [%s]",cmd)
			return false, nil
	}
	
	node.WriteMessageAll(&ardupilotmega.MessageSetMode{
        	TargetSystem: 0,        // System ID
		BaseMode: ardupilotmega.MAV_MODE(BaseMode),
		CustomMode: CustomMode,
	})


	// initialize timeout variables
        start := time.Now()
        counter := 1
        //resultpending := true

        for evt := range node.Events() {
                if frm, ok := evt.(*gomavlib.EventFrame); ok {
			// if frm is MessageSetMode
                        if _, ok := frm.Message().(*ardupilotmega.MessageSetMode); ok {
                                fmt.Printf("[MessageSetMode]\n  RESULT: id=%d, %+v\n\n", frm.Message().GetId(), frm.Message())
                        }


			// check if time.Since(start) > timeout, elapsed time (time.Since reports nanonseconds)
			if time.Since(start) > 1.5e9 {
				if counter > 5 {
					// Error out
					fmt.Errorf("Error - Autopilot did not reply in time\n")
					return true, nil
				}

				// Resend command
				node.WriteMessageAll(&ardupilotmega.MessageSetMode{
                       	 		TargetSystem: 0,        // System ID
                       		 	BaseMode: ardupilotmega.MAV_MODE(BaseMode),
                       		 	CustomMode: CustomMode,             
                       		})
					
				start = time.Now() // reset timeout stopwatch
				counter += 1
				logArm.Debugf("  Count: %d\n", counter)
			}
				
		}


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
