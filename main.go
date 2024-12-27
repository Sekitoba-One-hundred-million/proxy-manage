package main

import (
	"fmt"
	"os"
	"time"

	"github.com/kansei/proxy-manage/lib"
)

func main() {
  dnsFilePath := "/Volumes/Gilgamesh/proxy/domain"
  command := "proxy-manage.sh"
  
  if lib.IsFile( dnsFilePath ) {
    os.Remove( dnsFilePath )
  }

  for {
    if ! lib.IsFile( dnsFilePath ) {
      commandOut, err := lib.DoCommand( command )

      if err != nil {
        fmt.Println( err )
        os.Exit( 1 )
      }
      
      err = lib.WriteFile( commandOut, dnsFilePath )
      fmt.Println( commandOut )
    
      if err != nil {
        fmt.Println( err )
        os.Exit( 1 )
      }
    }

    time.Sleep( 1 )
  }
}
