# 	Mavlink Command Long - Activity

Communicate with Ardupilot system using Mavlink link.

## Installation
Command for Flogo CLI:
```console
flogo install github.com/wkarasz/flogo-mavlink/oss/activity/commandlong
```

Link for Flogo Web UI:
```console
https://github.com/wkarasz/flogo-mavlink/oss/activity/commandlong
```

## Schema
Settings:
```json
{
  "connection": "string"
}
```

Inputs and Outputs:
```json
{
  "inputs": [
    {
      "name": "command",
      "type:" "string",
      "required": true
    },
    {
      "name": "paramN", // param1 | param2 | param3 | param4 | param5 | param6 | param7
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
| devicePath       | The path to the ELM device on the host system; e.g. /dev/ttyUSB0 |
| directCmd        | The raw command supported by the ELM chipset; e.g. AT@1<br>https://www.elmelectronics.com/wp-content/uploads/2017/01/AT_Command_Table.pdf|

# Outputs
| Output           | Description    |
|:-----------------|:---------------|
| result           | The result will contain a string response of the command or will contain an error message |
