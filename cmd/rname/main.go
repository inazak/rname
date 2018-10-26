package main

import (
  "fmt"
  "flag"
  "os"
  "github.com/inazak/rname"
)

var usage =`
Usage:

  rname [test] prepend [-w=4] [PATTERN]
    prepend zeros for number
    ex) abc-1.jpg  =>  abc-0001.jpg

  rname [test] serial  [-w=4] [PATTERN]
    replace filename to serial number
    ex) abc.jpg  =>  0001.jpg

  PATTERN is shell file name syntax like '*.jpg'.
  when PATTERN is omitted, '*' is used.
`

func main() {

  // subcommand flagset
  prepend      := flag.NewFlagSet("prepend", flag.ExitOnError)
  prependWidth := prepend.Int("width", 4, "width of prepending zeros")

  serial       := flag.NewFlagSet("serial", flag.ExitOnError)
  serialWidth  := serial.Int("width", 4, "width of prepending zeros")

  // exit when arguments not found
  if len(os.Args) == 1 {
    fmt.Printf("%s", usage)
    os.Exit(1)
  }

  var args []string
  var isTest bool

  if os.Args[1] == "test" {
    // exit when subcommand not found
    if len(os.Args) < 3 {
      fmt.Printf("%s", usage)
      os.Exit(1)
    }
    // set testflag
    isTest = true
    args = os.Args[2:]
  } else {
    args = os.Args[1:]
  }


  // arguments parse
  switch args[0] {

  case "prepend":
    prepend.Parse(args[1:])

  case "serial":
    serial.Parse(args[1:])

  default:
    fmt.Printf("%q is not valid subcommand.\n%s", args[0], usage)
    os.Exit(1)
  }


  var com rname.Command //replace function interface
  filter := "*"

  // subcommand PREPEND
  if prepend.Parsed() {
    com = &rname.PrependCommand{ Width: *prependWidth }
    if len(prepend.Args()) == 1 {
      filter = prepend.Args()[0]
    }
  }
  // subcommand SERIAL
  if serial.Parsed() {
    com = &rname.SerialCommand{ Width: *serialWidth, Current: 1 }
    if len(serial.Args()) == 1 {
      filter = serial.Args()[0]
    }
  }

  // get filepath list
  list, err := rname.GetFilepathList(filter)
  if err != nil {
    fmt.Printf("Error: %v\n", err)
    os.Exit(2)
  }

  // do action for each filepath
  for _, fpath := range list {
    // use command interface function
    newfpath := com.Replace(fpath)

    // only print in test mode
    if isTest {
      fmt.Printf("   %v\n=> %v\n", fpath, newfpath)

    // do rename
    } else {
      if fpath != newfpath {
        err := os.Rename(fpath, newfpath)
        if err != nil {
          fmt.Printf("Error: %v\n", err)
          //continue
        }
      }
    }
  }
}


/*
--- goal ---

rname serial [-w 4]
  rename all file to serial number

rname prepend [-w 4]
  prepend zeros for number

rname fillin [-d '_']
  fill '_' in place of space

rname erase REGEXP
  erace match char

rname rename REGEXP STRING
  rename
*/



