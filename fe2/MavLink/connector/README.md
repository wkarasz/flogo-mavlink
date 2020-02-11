<!--
title: Mavlink Connection
weight: 4622
-->
# Mavlink Connection
This connection allows you to configure properties necessary to establish a connection with a Mavlink instance. A Mavlink Connection is necessary to work with the activities and trigger under flogo-mavlink contribution.

## Installation

### Flogo CLI
```bash
flogo install github.com/wkarasz/flogo-mavlink/oss/connection
```

## Configuration

### Settings:
| Name             | Type       | Description
| :---             | :---       | :---    
| name             | string     | A name for the connection  - ***REQUIRED***
| description      | string     | A short description for the connection
| port		   | string     | Mavlink UDP service port - ***REQUIRED***


## Example
A sample Mavlink connection JSON

```json
{
 "settings": {
 "name": "mavlink.skyviper",
 "description": "",
 "port": "14550"
}
}
```
