# prtcls
Port close is a simple go app for MacOS that closes the app using a specified port.
This is good when you have orphaned processes that are blocking ports that you need to use on your system.

## Installation

You can use go to install prtcls using the following command

``` 
go install github.com/pmatt1988/prtcls
```
You may need to ensure that your go bin folder has been added to the path.
In your terminal initialization file (.bash_profile, .bashrc or .zshrc.)
add the following to your path.

```
export PATH=$PATH:$(go env GOPATH)/bin
```
Make sure to re-source your current session with source or simply open a new terminal.




## Usage
```
sudo prtcls <port>
```

