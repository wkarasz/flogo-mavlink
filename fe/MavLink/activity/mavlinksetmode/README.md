# 	Mavlink SetMode - Activity

Communicate with Ardupilot system using Mavlink link.
This activity sets the mode of the system using the SetMode message type and standard values for BaseMode.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinksetmode
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinksetmode
```

## Schema
Inputs and Outputs:
```json
{
  "settings": [
    {
      "name": "connection",
      "type": "connection",
      "required": true,
      "display": {
        "name": "MavLink Connection",
        "description": "Select your MavLink Connection",
        "type": "connection",
        "selection": "single"
      },
      "allowed": []
    }
  ],
  "inputs": [
    {
      "name": "SET_MODE",
      "type": "string",
      "display": {
        "name":"Mode",
        "description": "Set the mode of the UAV",
        "type":"dropdown",
        "selection":"single",
        "mappable": true
      },
      "allowed": [
        "MAV_MODE_PREFLIGHT",
        "MAV_MODE_STABILIZE_DISARMED",
        "MAV_MODE_STABILIZE_ARMED", 
        "MAV_MODE_MANUAL_DISARMED",
        "MAV_MODE_MANUAL_ARMED",
        "MAV_MODE_GUIDED_DISARMED",
        "MAV_MODE_GUIDED_ARMED",
        "MAV_MODE_AUTO_DISARMED",
        "MAV_MODE_AUTO_ARMED",
        "MAV_MODE_TEST_DISARMED",
        "MAV_MODE_TEST_ARMED"
      ],
      "value": "MAV_MODE_MANUAL_ARMED"
    }
  ],
  "outputs": [
    {
      "name": "result",
      "type": "string"
    }
  ]
}
```
## Inputs
| Input            | Description    |
|:-----------------|:---------------|
| SET_MODE       | The ENUM of MAV_MODE |

## Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
