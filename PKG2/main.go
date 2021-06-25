package main

import (
	log "github.com/surajrathod2352/Practice/Go/PKG2/logger"  // "log" is the name thst you have given to the package
														       // other wise it will be as same as the the name you have goven in the file
															  // module name/direction to the forlder the packge is in
															  //if i use "." instead of name Then i dont have to initialize it
															  //for example Warn("text") will work if i start with "."
															  //"_" also stnds for somthing ot clear
	
)
func main(){
	log.Warn("This is the first test case for Warning_____")
	log.Log("This is the first test case for Logging______")
	log.Error("This is the first test case for Error_______")
}
