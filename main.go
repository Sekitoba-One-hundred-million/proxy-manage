package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/fsnotify/fsnotify"
	"github.com/kansei/sekitoba-proxy-manage/lib"
)

const dnsFilePath = "/Volumes/Gilgamesh/proxy/domain"

func stopProxy() {
  log.Println( "Start Stop" )
  command := "./proxy-stop.sh"
  _, err := lib.DoCommand( command )

  if err != nil {
    log.Println( err )
    os.Exit( 1 )
  }

  log.Println( "Finish Stop" )
}

func restartProxy() {
  log.Println( "Start Restart" )
  command := "./proxy-restart.sh"

  if lib.IsFile( dnsFilePath ) {
    os.Remove( dnsFilePath )
  }

  commandOut, err := lib.DoCommand( command )

  if err != nil {
    log.Println( err )
    os.Exit( 1 )
  }

  err = lib.WriteFile( commandOut, dnsFilePath )

  if err != nil {
    log.Println( err )
    os.Exit( 1 )
  }

  log.Println( "Finish Restart" )
}

func manageProxy(gen chan int) {
  watcher, err := fsnotify.NewWatcher()
  defer watcher.Close()
  
  if err != nil {
    log.Fatal( err )
    return
  }

  err = watcher.Add( "/Volumes/Gilgamesh/proxy/" )

  if err != nil {
    log.Fatal( err )
    return
  }
  
  for {
    select {
    case _ = <-gen:
      return
    default:
      select {
      case event, _ := <-watcher.Events:
        if( event.Op&fsnotify.Remove == fsnotify.Remove &&
          event.Name == dnsFilePath ) {
          restartProxy()
        }
      default:
        if( ! lib.IsFile( dnsFilePath ) ) {
          restartProxy()
        } else {
          continue
        }
      }
    }
  }
}

func getKillSignal(gen chan int) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
  gen<-1
  stopProxy()
}

func main() {
  restartProxy()
  gen := make(chan int)
  go manageProxy(gen)
	getKillSignal(gen)
	log.Println("正常終了")
}
