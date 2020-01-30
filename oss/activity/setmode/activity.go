package mavlinksetmode

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/coerce"

	//"context"
        
	"github.com/wkarasz/gomavlib"
	"github.com/wkarasz/gomavlib/dialects/ardupilotmega"
)

var logSetMode = log.ChildLogger(log.RootLogger(), "mavlink-setmode")
var activityMd = activity.ToMetadata(&Input{}, &Output{}, &Settings{})

func init() {
	err := activity.Register(&Activity{}, New)
	if err != nil {
		logSetMode.Errorf("Mavlink Set Mode Activity init error : %s ", err.Error())
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
		log.RootLogger().Infof("Connection [%s]",settings.Connection)
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
	logSetMode.Debugf("cleaning up Mavlink SetMode activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	cmd := context.GetInput("SET_MODE").(string)
	logSetMode.Debugf("Set Mode [%s]", cmd)
	
	n := *a.n

        switch cmd {
		case "MAV_MODE_PREFLIGHT":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
				TargetSystem: 0,
                		BaseMode: ardupilotmega.MAV_MODE_PREFLIGHT,
			})
		case "MAV_MODE_STABILIZE_DISARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_STABILIZE_DISARMED,
                        })
        	case "MAV_MODE_STABILIZE_ARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_STABILIZE_ARMED,
                        })
		case "MAV_MODE_MANUAL_DISARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_MANUAL_DISARMED,
                        })
		case "MAV_MODE_MANUAL_ARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_MANUAL_ARMED,
                        })
		case "MAV_MODE_GUIDED_DISARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_GUIDED_DISARMED,
                        })
		case "MAV_MODE_GUIDED_ARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_GUIDED_ARMED,
                        })
		case "MAV_MODE_AUTO_DISARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_AUTO_DISARMED,
                        })
		case "MAV_MODE_AUTO_ARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_AUTO_ARMED,
                        })
		case "MAV_MODE_TEST_DISARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_TEST_DISARMED,
                        })
		case "MAV_MODE_TEST_ARMED":
			n.WriteMessageAll(&ardupilotmega.MessageSetMode{
                                TargetSystem: 0,
                                BaseMode: ardupilotmega.MAV_MODE_TEST_ARMED,
                        })
		default:
			logSetMode.Errorf("Unknown mode [%s]",cmd)
			return
	}
	return true, nil
}
