# 	Mavlink Waypoints - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkuploadmission
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkuploadmission
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
      "name": "MISSION",
      "type": "string",
      "required": true
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
| MISSION       | The mission commands and waypoints.  Follows syntax of `[{mission_element},{mission_element},{mission_element}]` where mission_element schema is defined below. |

**mission_element schema**
```json
{
	"Index": 0,
	"CurrentWP": 1,
	"CoordFrame": 0,
	"Command": 16,
	"Param1": 0,
	"Param2": 0,
	"Param3": 0,
	"Param4": 0,
	"X": 41.668469,
	"Y": -72.646413,
	"Z": 13.439352,
	"Autocontinue": 1
}
```
e.g. 
```json
[{"Index": 0,"CurrentWP": 1,"CoordFrame": 0,"Command": 16,"Param1": 0,"Param2": 0,"Param3": 0,"Param4": 0,"X": 41.668469,"Y": -72.646413,"Z": 13.439352,"Autocontinue": 1},{"Index": 1,"CurrentWP": 0,"CoordFrame": 3,"Command": 22,"Param1": 0.0,"Param2": 0.0,"Param3": 0.0,"Param4": 0.0,"X": 0.0,"Y": 0.0,"Z": 10.0,"Autocontinue": 1},{"Index": 2,"CurrentWP": 0,"CoordFrame": 0,"Command": 16,"Param1": 0.0,"Param2": 0.0,"Param3": 0.0,"Param4": 0.0,"X": 41.66856880,"Y": -72.64635710,"Z": 10.0,"Autocontinue": 1},{"Index": 3,"CurrentWP": 0,"CoordFrame": 0,"Command": 16,"Param1": 0.0,"Param2": 0.0,"Param3": 0.0,"Param4": 0.0,"X": 41.66856880,"Y": -72.64635710,"Z": 10.0,"Autocontinue": 1},{"Index": 4,"CurrentWP": 0,"CoordFrame": 3,"Command": 20,"Param1": 0.0,"Param2": 0.0,"Param3": 0.0,"Param4": 0.0,"X": 0.0,"Y": 0.0,"Z": 0.0,"Autocontinue": 1}]
```


## Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
