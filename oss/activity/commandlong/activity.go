package mavlinkcommandlong

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/coerce"

	//"context"
	"time"
        
	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

var logCommandLong = log.ChildLogger(log.RootLogger(), "mavlink-commandlong")
var activityMd = activity.ToMetadata(&Input{}, &Output{}, &Settings{})

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logCommandLong.Errorf("Mavlink Set Mode Activity init error : %s ", err.Error())
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
	logCommandLong.Debugf("cleaning up Mavlink CommandLong activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	cmd := context.GetInput("command").(string)
	param1 := context.GetInput("param1").(float32)
	param2 := context.GetInput("param2").(float32)
	param3 := context.GetInput("param3").(float32)
	param4 := context.GetInput("param4").(float32)
	param5 := context.GetInput("param5").(float32)
	param6 := context.GetInput("param6").(float32)
	param7 := context.GetInput("param7").(float32)
	logCommandLong.Debugf("Command [%s]", cmd)
	log.RootLogger().Infof("Command [%s]", cmd)
		
	n := *a.n
	
	n.WriteMessageAll(&ardupilotmega.MessageCommandLong{
		TargetSystem: 0, // System ID
		TargetComponent: 0, // Component ID
		Command: ardupilotmega.FromString(cmd),
		Confirmation: 0,
		Param1: float32(param1),
		Param2: float32(param2),
		Param3: float32(param3),
		Param4: float32(param4),
		Param5: float32(param5),
		Param6: float32(param6),
		Param7: float32(param7),
	})

	// initialize timeout variables
	start := time.Now()
	counter := 1
	for evt := range n.Events() {
		if frm, ok := evt.(*gomavlib.EventFrame); ok {
			if msg, ok := frm.Message().(*ardupilotmega.MessageCommandAck); ok && msg.Command == ardupilotmega.FromString(cmd) {
				if msg.Result == 0 {
				// Command was successful
					context.SetOutput("result", msg.Result)
					return true, nil
				} else if msg.Result != 5 {
					logCommandLong.Errorf("Operation failed: MAV_RESULT=%d\n",msg.Result)
					context.SetOutput("result", msg.Result)
					return true, nil		
				}
			}

			if time.Since(start) > 1.5e9 && counter <= 5 {
				// resend message
				n.WriteMessageAll(&ardupilotmega.MessageCommandLong{
       		         		TargetSystem: 0, // System ID
                			TargetComponent: 0, // Component ID
                			Command: ardupilotmega.FromString(cmd),
                			Confirmation: uint8(counter),
                			Param1: float32(param1),
                			Param2: float32(param2),
                			Param3: float32(param3),
                			Param4: float32(param4),
                			Param5: float32(param5),
                			Param6: float32(param6),
               		 		Param7: float32(param7),
        			})

				start = time.Now() // reset timeout stopwatch
				counter += 1
			} else if counter > 5 {
				logCommandLong.Errorf("Operation timed out\n")
				context.SetOutput("result", "Operation timed out")
				return true, nil
			}
		}
	}
	return true, nil
}
