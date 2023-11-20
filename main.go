package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(message string) string {
	fmt.Scanf(message)
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	banner()
	mypassword := "secret"
	fmt.Println("MD5 Hashed value: ", getMD5Hash(mypassword))

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

	fmt.Println("Enter text':")
}
