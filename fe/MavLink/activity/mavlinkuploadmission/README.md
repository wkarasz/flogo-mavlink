# 	Mavlink Waypoints - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinksetwaypoints
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinksetwaypoints
```

## Schema
Inputs and Outputs:
```json
{
  "inputs": [
    "name": "mode",
    "type": "string",
    "allowed": [
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
| devicePath       | The path to the ELM device on the host system; e.g. /dev/ttyUSB0 |
| directCmd        | The raw command supported by the ELM chipset; e.g. AT@1<br>https://www.elmelectronics.com/wp-content/uploads/2017/01/AT_Command_Table.pdf|

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
