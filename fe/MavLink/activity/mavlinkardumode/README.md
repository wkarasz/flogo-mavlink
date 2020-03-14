# 	Mavlink Ardu Mode - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkardumode
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkardumode
```

## Schema
Inputs and Outputs:
```json
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
      "name": "COMPONENT_MODE",
      "type": "string",
      "display": {
        "name":"Component Mode",
        "description": "Set the mode of the component",
        "type":"dropdown",
        "selection":"single",
        "mappable": true
      },
      "allowed": [
        "AUTO",
        "LOITER",
        "RTL",
        "ALTHOLD"
      ],
      "value": "AUTO"
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
| COMPONENT_MODE       | The ENUM for MAV_MODE; e.g. AUTO, LOITER, etc. |

## Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
