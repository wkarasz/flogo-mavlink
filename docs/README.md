# 	Flogo-Mavlink Project

Welcome to the Flogo-Mavlink Project.<br>
<br>
The goal of this project is to enable developers to access all the features of the Mavlink protocol from within the Flogo ecosystem ([Project Flogo](https://flogo.io) and [TIBCO Flogo Enterprise](https://www.tibco.com/products/tibco-flogo)).<br>
<br>
Quite simply, the project is a Flogo extension, creating a new palette of activities and trigger for interfacing with a Mavlink supported vehicle.  The capabilities of the extension are ever-growing, currently including:
- Arm/disarm
- Upload mission
- Start mission
- Set mode
- Issue any generic MAV_CMD type
- Listen to all events emitted by vehicle
<br>
<b>Demo Video</b>
<iframe width="560" height="315" src="https://www.youtube.com/embed/iFnUDRtWmmQ" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<br>
## Getting Started  
  
### Using Flogo Enterprise  
- Upload the extension to Flogo Enterprise using the provided zip file **[MavLink.zip](https://raw.githubusercontent.com/wkarasz/flogo-mavlink/master/fe/MavLink.zip)**<br>
<img src="https://raw.githubusercontent.com/wkarasz/flogo-mavlink/master/img/upload_zip_extension.jpg"/><br>
<br>
- Create a new application
  - Add the Mavlink activity or trigger
  - Configure the Mavlink connector
- Build/compile target binary
- Run the binary
