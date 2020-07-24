# 	Mavlink CommandLong - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkcmdlong
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/fe/MavLink/activity/mavlinkcmdlong
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
      "name": "mavcommand",
      "type": "string"
    },
    {
      "name": "param1",
      "type": "string"
    },
    {
      "name": "param2",
      "type": "string"
    },
    {
      "name": "param3",
      "type": "string"
    },
    {
      "name": "param4",
      "type": "string"
    },
    {
      "name": "param5",
      "type": "string"
    },
    {
      "name": "param6",
      "type": "string"
    },
    {
      "name": "param7",
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
| mavcommand   | Mav Command |
| param1       | Parameter 1 (for the specific command). float32 |
| param2       | Parameter 2 (for the specific command). float32 |
| param3       | Parameter 3 (for the specific command). float32 |
| param4       | Parameter 4 (for the specific command). float32 |
| param5      | Parameter 5 (for the specific command). float32 |
| param6       | Parameter 6 (for the specific command). float32 |
| param7       | Parameter 7 (for the specific command). float32 |

## Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
