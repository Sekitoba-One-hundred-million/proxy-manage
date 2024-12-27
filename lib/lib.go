package lib

import (
	"io/ioutil"
	"os"
  "os/exec"
	"strings"
)

func IsFile( filePath string ) bool {
  info, err := os.Stat( filePath )
  
  if err != nil {
    return false
  }
  
  return !info.IsDir()
}

func DoCommand( command string ) ( string, error ) {
  out, err := exec.Command( "bash", command ).Output()

  if err != nil {
    return "", err
  }
  
  strOut := string( out )
  strOut = strings.TrimSuffix( strOut, "\n" )
  return strOut, nil
}

func ReadFile( filePath string ) ( string, error ) {
  fileContent := ""
  bytes, err := ioutil.ReadFile( filePath )
  
	if err != nil {
    return fileContent, err
	}

  fileContent = string( bytes )
  fileContent = strings.TrimSuffix( fileContent, "\n" )
  return fileContent, nil
}

func WriteFile( data string, filePath string ) error {
  f, err := os.Create( filePath )

  if err != nil {
    return err
  }

  _, err = f.Write( []byte( data ) )
  f.Close()
  return err
}
