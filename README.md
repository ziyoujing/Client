> No longer maintained. I created this project in high school to mess with my teacher's computer, It was lots of fun but now I need to be a grown-up and respect others PCs. Thanks for the love this project received 

# Tiked
Multi-client remote administration tool targeting windows systems.

#### `Usage: [Command] [Target] [Arguments]`

# Commands
- **users***: Display list of users
- **help**: Displays information about the commands
- **off**: Powers the computer off (requires target, eg: off James)
- **lo**: Log out of account
- **kill***: Stops program
- **msg***: Sends a message to the victim (use - as spaces in argument)
- **yn***: Sends a yes or no message box if targeting a single user, receives response(use - as spaces in argument)
- **web***: Opens a web page or starts a program
- **ddos***: Starts an 8 threads DDOS HTTP GET flood attack
- **sdd**: Stops ddos attack
- **inf**: Infects usb drives with payload
- **pass**: Gets passwords from Google Chrome
- **passlist**: Prints passwords received
- **autoInf**: Continuously infects usb devices
- **stopAutoInf**: Stop autoInfecting usb devices
- **upgrade***: Downloads exe form direct download link
- **meterpreter***: Start meterpreter session to provided host
- **getav**: Returns AV procces name runnig
- **start-keylogger**: Start recording keys
- **keylog**: Send keylog
- **please***: Promt the user to start given command as Admin (Shows UAC)
- **uninstall**: Deletes files and clean registry

###### *(needs arguments)

# Target
+ Use **name** from users command to target **single target**
+ Use **asterisk** to target **all** computers

# Features
+ Multiple clients
+ Multi platform
+ No port fowarding
+ Near FUD
+ Mobile app

# TODO
- [ ] Encrypt communication
- [ ] Implement cap'n
- [ ] Decrypt stub and run from memory
- [ ] Save payload in registry and start it using an auxiliary program or a Startup key or with shellcode injected in explorer.exe or similar
