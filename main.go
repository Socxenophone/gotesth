package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	banner()
	var userInput string
	fmt.Print("Enter text: ")
	fmt.Scan(&userInput)
	fmt.Println("MD5 Hashed value:", getMD5Hash(userInput))
}

func banner() {
	fmt.Println("Hasher ver 1.0 : Hash Stuff")
	fmt.Printf(` 
	
'||'  '||'                '||                      
 ||    ||   ....    ....   || ..     ....  ... ..  
 ||''''||  '' .||  ||. '   ||' ||  .|...||  ||' '' 
 ||    ||  .|' ||  . '|..  ||  ||  ||       ||     
.||.  .||. '|..'|' |'..|' .||. ||.  '|...' .||.    
 
`)
}
