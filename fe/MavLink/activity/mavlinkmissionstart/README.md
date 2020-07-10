# 	Mavlink MissionStart - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkmissionstart
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkmissionstart
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
      "name": "param1",
      "type": "string"
    },
    {
      "name": "param2",
      "type": "string"
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
| param1       | Set the first mission item to run. min:0 |
| param2       | Set the last mission item to run. min:0 |

## Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
