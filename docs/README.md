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
<br>
<b>Demo Video</b>
<iframe width="560" height="315" src="https://www.youtube.com/embed/iFnUDRtWmmQ" frameborder="0" allow="accelerometer; autoplay; encrypted-media; gyroscope; picture-in-picture" allowfullscreen></iframe>
<br>
## Getting Started  
<br>
<b>Implementation Concept</b><br>
The flogo-mavlink extension works by starting up a UDP listener service and receiving UDP packets that are broadcast by the drone; thus the ground control system and drone must be on the same network.<br>
<img src="https://raw.githubusercontent.com/wkarasz/flogo-mavlink/master/img/simple_diagram.png"/><br>
<br>
### Using Flogo Enterprise  
- Upload the extension to Flogo Enterprise using the provided zip file **[MavLink.zip](https://raw.githubusercontent.com/wkarasz/flogo-mavlink/master/fe/MavLink.zip)**<br>
<img src="https://raw.githubusercontent.com/wkarasz/flogo-mavlink/master/img/upload_zip_extension.jpg"/><br>
<br>
- Create a new application
  - Add the Mavlink activity or trigger
  - Configure the Mavlink connector
    - Set the listener port; default 14550
- Build/compile target binary
- Run the binary
  
  
### Labs ###  

**Coming Soon** While you can try out the flogo-mavlink extension with a real drone, it can be cheaper and easier to simulate the drone.  Ardupilot has an excellent SITL (software in the loop) option to substitute a virtual drone for a real one and it works well with this extension, so no longer do you have to worry about crashing or gaining flight airspace clearance. This lab will show you how to get started using this option.  
  
**Coming Soon** Uploading a mission

**Coming Soon** Listening for mission events  
  


