package mavlinkuploadmission

import (
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/data/coerce"

	//"context"
	"fmt"
	"math"
	"encoding/json"
	"time"
	"strings"
        
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
	logSetMode.Debugf("cleaning up Mavlink SetMode activity")
	ctx, cancel := ctx.WithTimeout(ctx.Background(), 30*time.Second)
	defer cancel()
}
*/

type Waypoint struct {
        Index uint16
        CurrentWP uint8
        CoordFrame uint8
        Command uint16
        Param1 float32
        Param2 float32
        Param3 float32
        Param4 float32
        X float32
        Y float32
        Z float32
        Autocontinue uint8
}

// Eval implements activity.Activity.Eval
func (a *Activity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	waypointsjson2 := context.GetInput("MISSION").(string)
	waypointsjson := strings.Replace(waypointsjson2, "\\", "", -1)
	logSetMode.Debugf("[MavLinkUploadMission] %+v", waypointsjson)
	//log.RootLogger().Infof("MavLinkUploadMission %+v", waypointsjson2)
		
	node := *a.n

        var stage1 = false // MissionClearAll sent
        var stage2 = false // MissionCount sent
        var stage3 = false // All MissionItem sent
        var stage4 = false // MissionAck received
        var stage5 = false // Mission verified
        var bItemInt = false // Flag to determine if protocol is using ItemInt specification
        fmt.Println(stage1, stage2, stage3,stage4,stage5)

        fmt.Printf("[Sending MessageMissionClearAll]\n")
        node.WriteMessageAll(&ardupilotmega.MessageMissionClearAll{
                TargetSystem: 0,
                TargetComponent: 0,
                MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,
        })
        stage1 = true

        var waypoints []Waypoint
        err = json.Unmarshal([]byte(waypointsjson), &waypoints)
        if err != nil {
                fmt.Println("error:", err)
        }
        totalmissionitems := len(waypoints)
        fmt.Printf("  Total mission waypoints: %d\n",totalmissionitems)

        fmt.Printf("[Sending MessageMissionCount]\n")
        node.WriteMessageAll(&ardupilotmega.MessageMissionCount{
                TargetSystem: 0,        // System ID
                TargetComponent: 0,     // Component ID
                Count: uint16(totalmissionitems),               // Number of mission items in the sequence
                MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION, // MAV_MISSION_TYPE
        })
        stage2 = true

        // initialize timeout variables
        start := time.Now()
        counter := 1
        missionitemcount := 0
        var requesteditem uint16 = 0

        for evt := range node.Events() {
                if frm, ok := evt.(*gomavlib.EventFrame); ok {
                        // if frm is MessageStatustext
                        if _, ok := frm.Message().(*ardupilotmega.MessageStatustext); ok {
                                fmt.Printf("RESULT: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
                        }

                        // if frm is MessageMissionRequest
                        if MsgMissionRequest, ok := frm.Message().(*ardupilotmega.MessageMissionRequest); ok && (stage1 && stage2 && !stage3) {
                                fmt.Printf("RESULT: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
                                fmt.Printf("Sequence request: %d\n",MsgMissionRequest.Seq)

                                requesteditem = MsgMissionRequest.Seq

                                start = time.Now() // reset timeout stopwatch
                                // Send MessageMissionItem
                                //node.WriteMessageAll(&ardupilotmega.MessageMissionItem{
                                node.WriteMessageTo(frm.Channel, &ardupilotmega.MessageMissionItem{
                                        TargetSystem: 0,        // uint8        // System ID
                                        TargetComponent: 0,     // uint8        // Component ID
                                        Seq: waypoints[requesteditem].Index,  // uint16       // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
                                        Frame: ardupilotmega.MAV_FRAME(waypoints[requesteditem].CoordFrame),                 // MAV_FRAME `mavenum:"uint8"`// The coordinate system of the waypoint.
                                        Command: ardupilotmega.MAV_CMD(waypoints[requesteditem].Command),                    // MAV_CMD `mavenum:"uint16"` // The scheduled action for the waypoint.
                                        Current: waypoints[requesteditem].CurrentWP,                 // uint8        // false:0, true:1
                                        Autocontinue: waypoints[requesteditem].Autocontinue,         // uint8        // Autocontinue to next waypoint
                                        Param1: waypoints[requesteditem].Param1,                     // float32      // PARAM1, see MAV_CMD enum
                                        Param2: waypoints[requesteditem].Param2,                     // float32      // PARAM2, see MAV_CMD enum
                                        Param3: waypoints[requesteditem].Param3,                     // float32      // PARAM3, see MAV_CMD enum
                                        Param4: waypoints[requesteditem].Param4,                     // float32      // PARAM4, see MAV_CMD enum
                                        X: float32(waypoints[requesteditem].X),                 // float32        // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
                                        Y: float32(waypoints[requesteditem].Y),                 // float32        // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
                                        Z: waypoints[requesteditem].Z,                               // float32      // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
                                        MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,            // MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`     // Mission type.
                                })
                                bItemInt = false
                                fmt.Printf("  Writing mission item [MessageMissionItem=%d]\n",requesteditem)
                                if requesteditem == uint16(totalmissionitems-1) {
                                        stage3 = true
                                }
                                missionitemcount += 1
                                counter = 1 // reset counter
                        } else if MsgMissionRequestInt, ok := frm.Message().(*ardupilotmega.MessageMissionRequestInt); ok && (stage1 && stage2 && !stage3) { // if frm is MessageMissionRequestInt send MessageMissionItemInt

                                fmt.Printf("RESULT: id=%d, %+v\n", frm.Message().GetId(), frm.Message())

                                requesteditem = MsgMissionRequestInt.Seq

                                start = time.Now() // reset timeout stopwatch
                                // Send MessageMissionItemInt
                                node.WriteMessageAll(&ardupilotmega.MessageMissionItemInt{
                                        TargetSystem: 0,        // uint8        // System ID
                                        TargetComponent: 0,     // uint8        // Component ID
                                        Seq: waypoints[requesteditem].Index,    // uint16       // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
                                        Frame: ardupilotmega.MAV_FRAME(waypoints[requesteditem].CoordFrame),                    // MAV_FRAME `mavenum:"uint8"`// The coordinate system of the waypoint.
                                        Command: ardupilotmega.MAV_CMD(waypoints[requesteditem].Command),                       // MAV_CMD `mavenum:"uint16"` // The scheduled action for the waypoint.
                                        Current: waypoints[requesteditem].CurrentWP,                    // uint8        // false:0, true:1
                                        Autocontinue: waypoints[requesteditem].Autocontinue,            // uint8        // Autocontinue to next waypoint
                                        Param1: waypoints[requesteditem].Param1,                        // float32      // PARAM1, see MAV_CMD enum
                                        Param2: waypoints[requesteditem].Param2,                        // float32      // PARAM2, see MAV_CMD enum
                                        Param3: waypoints[requesteditem].Param3,                        // float32      // PARAM3, see MAV_CMD enum
                                        Param4: waypoints[requesteditem].Param4,                        // float32      // PARAM4, see MAV_CMD enum
                                        X: int32(float64(waypoints[requesteditem].X)*math.Pow(10,7)),   // int32        // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
                                        Y: int32(float64(waypoints[requesteditem].Y)*math.Pow(10,7)),   // int32        // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
                                        Z: waypoints[requesteditem].Z,                          // float32      // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
                                        MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,            // MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`     // Mission type.
                                })
                                bItemInt = true
                                fmt.Printf("Writing mission item [MessageMissionItemInt=%d]\n",requesteditem)
                                if requesteditem == uint16(totalmissionitems-1) {
                                        stage3 = true
                                }
                                missionitemcount += 1
                                counter = 1 // reset counter
                        }

                        // if frm is MessageMissionAck
                        // verify result and exit
                        if MsgMissionAck, ok := frm.Message().(*ardupilotmega.MessageMissionAck); ok && (stage1 && stage2 && stage3 && !stage4) {
                                fmt.Printf("[MessageMissionAck]: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
                                fmt.Printf("[MessageMissionAck]: Total waypoints set=%d\n",missionitemcount)
                                fmt.Printf("  Last requested item: %d\n", requesteditem)
                                if MsgMissionAck.Type==0 {
                                        stage4 = true
					err = context.SetOutput("result","Operation Executed")
					if err != nil {
						return false, fmt.Errorf("error setting output for Activity [%s]: %s", context.Name(), err.Error())
					}
                                        return true, nil
                                }
                        } else if _, ok := frm.Message().(*ardupilotmega.MessageMissionAck); ok {
                                fmt.Printf("[MessageMissionAck]: id=%d, %+v\n", frm.Message().GetId(), frm.Message())
                                fmt.Printf("  Current requested item: %d\n", requesteditem)
                        }

                        // check if time.Since(start) > timeout, elapsed time (time.Since reports nanoseconds)
                        if time.Since(start) > 1.5e9 {
                                if counter > 5 {
                                        // Error out
                                        fmt.Errorf("Error - Autopilot did not reply in time\n")
                                        return true, nil

                                } else if (stage1 && stage2 && !stage3) && missionitemcount == 0 {
                                        // Re-send MessageMissionCount
                                        fmt.Printf("[Re-Sending MessageMissionCount]\n")
                                        node.WriteMessageAll(&ardupilotmega.MessageMissionCount{
                                                TargetSystem: 0,        // System ID
                                                TargetComponent: 0,     // Component ID
                                                Count: uint16(totalmissionitems),               // Number of mission items in the sequence
                                                MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION, // MAV_MISSION_TYPE
                                        })

                                } else if (stage1 && stage2 && !stage3) {
                                        // Re-send last MessageMissionRequest or MessageMissionRequestInt
                                        if bItemInt {
                                                node.WriteMessageAll(&ardupilotmega.MessageMissionItemInt{
                                                        TargetSystem: 0,        // uint8        // System ID
                                                        TargetComponent: 0,     // uint8        // Component ID
                                                        Seq: waypoints[requesteditem].Index,    // uint16       // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
                                                        Frame: ardupilotmega.MAV_FRAME(waypoints[requesteditem].CoordFrame),                    // MAV_FRAME `mavenum:"uint8"`// The coordinate system of the waypoint.
                                                        Command: ardupilotmega.MAV_CMD(waypoints[requesteditem].Command),                       // MAV_CMD `mavenum:"uint16"` // The scheduled action for the waypoint.
                                                        Current: waypoints[requesteditem].CurrentWP,                    // uint8        // false:0, true:1
                                                        Autocontinue: waypoints[requesteditem].Autocontinue,            // uint8        // Autocontinue to next waypoint
                                                        Param1: waypoints[requesteditem].Param1,                        // float32      // PARAM1, see MAV_CMD enum
                                                        Param2: waypoints[requesteditem].Param2,                        // float32      // PARAM2, see MAV_CMD enum
                                                        Param3: waypoints[requesteditem].Param3,                        // float32      // PARAM3, see MAV_CMD enum
                                                        Param4: waypoints[requesteditem].Param4,                        // float32      // PARAM4, see MAV_CMD enum
                                                        X: int32(float64(waypoints[requesteditem].X)*math.Pow(10,7)),   // int32        // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
                                                        Y: int32(float64(waypoints[requesteditem].Y)*math.Pow(10,7)),   // int32        // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
                                                        Z: waypoints[requesteditem].Z,                          // float32      // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
                                                        MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,            // MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`     // Mission type.
                                                })
                                        } else {
                                                node.WriteMessageTo(frm.Channel, &ardupilotmega.MessageMissionItem{
                                                        TargetSystem: 0,        // uint8        // System ID
                                                        TargetComponent: 0,     // uint8        // Component ID
                                                        Seq: waypoints[requesteditem].Index,  // uint16       // Waypoint ID (sequence number). Starts at zero. Increases monotonically for each waypoint, no gaps in the sequence (0,1,2,3,4).
                                                        Frame: ardupilotmega.MAV_FRAME(waypoints[requesteditem].CoordFrame),                 // MAV_FRAME `mavenum:"uint8"`// The coordinate system of the waypoint.
                                                        Command: ardupilotmega.MAV_CMD(waypoints[requesteditem].Command),                    // MAV_CMD `mavenum:"uint16"` // The scheduled action for the waypoint.
                                                        Current: waypoints[requesteditem].CurrentWP,                 // uint8        // false:0, true:1
                                                        Autocontinue: waypoints[requesteditem].Autocontinue,         // uint8        // Autocontinue to next waypoint
                                                        Param1: waypoints[requesteditem].Param1,                     // float32      // PARAM1, see MAV_CMD enum
                                                        Param2: waypoints[requesteditem].Param2,                     // float32      // PARAM2, see MAV_CMD enum
                                                        Param3: waypoints[requesteditem].Param3,                     // float32      // PARAM3, see MAV_CMD enum
                                                        Param4: waypoints[requesteditem].Param4,                     // float32      // PARAM4, see MAV_CMD enum
                                                        X: float32(waypoints[requesteditem].X),                 // float32        // PARAM5 / local: x position in meters * 1e4, global: latitude in degrees * 10^7
                                                        Y: float32(waypoints[requesteditem].Y),                 // float32        // PARAM6 / y position: local: x position in meters * 1e4, global: longitude in degrees *10^7
                                                        Z: waypoints[requesteditem].Z,                               // float32      // PARAM7 / z position: global: altitude in meters (relative or absolute, depending on frame.
                                                        MissionType: ardupilotmega.MAV_MISSION_TYPE_MISSION,            // MAV_MISSION_TYPE `mavenum:"uint8" mavext:"true"`     // Mission type.
                                                })
                                        }
                                }
                                start = time.Now()  // reset timeout stopwatch
                                counter += 1
                                fmt.Printf("  Count: %d\n", counter)
                        }
                }
        }
        return true, nil
}
