package main

import (
  "fmt"
  "flag"
  "os"
  "github.com/inazak/rname"
)

var usage =`
Usage:

  rname [test] prepend [-width=5] [FILE_PATTERN]
    prepend zeros for number
    ex) abc-1.jpg  =>  abc-00001.jpg

  rname [test] serial  [-width=5] [-start=1] [FILE_PATTERN]
    replace filename to serial number
    ex) abc.jpg  =>  00001.jpg

  rname [test] fillin [-padding='_'] [FILE_PATTERN]
    fill padding in place of space
    ex) abc def.jpg => abc_def.jpg

  rname [test] erase -target='?' [FILE_PATTERN]
    erase string matched target
    ex) erase -t="-demo" : abc-demo.jpg => abc.jpg

  rname [test] regex -pattern='?' [-replace=''] [FILE_PATTERN]
    substitute regex-pattern to replacement text    
    when replace text is omitted, erase matched.
    ex) regex -p="^(.)(.)" -r="$2$1" : abc.jpg => bac.jpg

  wildcard can be used for FILE_PATTERN, like '*.jpg'.
  when FILE_PATTERN is omitted, '*' is used.
`

func main() {

  prepend       := flag.NewFlagSet("prepend", flag.ExitOnError)
  prependWidth  := prepend.Int("width", 5, "width of prepending zeros")
  prependW      := prepend.Int("w",     0, "width of prepending zeros")

  serial        := flag.NewFlagSet("serial", flag.ExitOnError)
  serialWidth   := serial.Int("width", 5, "width of prepending zeros")
  serialW       := serial.Int("w",     0, "width of prepending zeros")
  serialStart   := serial.Int("start", 1, "strat number")
  serialS       := serial.Int("s",    -1, "start number")

  fillin        := flag.NewFlagSet("fillin", flag.ExitOnError)
  fillinPadding := fillin.String("padding", "_", "padding in place of space")
  fillinP       := fillin.String("p",       "",  "padding in place of space")

  erase         := flag.NewFlagSet("erase", flag.ExitOnError)
  eraseTarget   := erase.String("target", "", "string to erase")
  eraseT        := erase.String("t",      "", "string to erase")

  regex         := flag.NewFlagSet("regex", flag.ExitOnError)
  regexPattern  := regex.String("pattern", "", "regex pattern for search")
  regexP        := regex.String("p",       "", "regex pattern for search")
  regexReplace  := regex.String("replace", "", "replacement text")
  regexR        := regex.String("r",       "", "replacement text")

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

  case "fillin":
    fillin.Parse(args[1:])

  case "erase":
    erase.Parse(args[1:])

  case "regex":
    regex.Parse(args[1:])


  default:
    fmt.Printf("%q is not valid subcommand.\n%s", args[0], usage)
    os.Exit(1)
  }


  var comm rname.Command
  filter := "*"

  // subcommand PREPEND
  if prepend.Parsed() {
    if *prependW != 0 { *prependWidth = *prependW }
    comm = &rname.PrependCommand{ Width: *prependWidth }
    if len(prepend.Args()) == 1 {
      filter = prepend.Args()[0]
    }
  }
  // subcommand SERIAL
  if serial.Parsed() {
    if *serialW != 0  { *serialWidth = *serialW }
    if *serialS != -1 { *serialStart = *serialS }
    comm = &rname.SerialCommand{ Width: *serialWidth, Current: *serialStart }
    if len(serial.Args()) == 1 {
      filter = serial.Args()[0]
    }
  }
  // subcommand FILLIN
  if fillin.Parsed() {
    if *fillinP != "" { *fillinPadding = *fillinP }
    comm = &rname.FillinCommand{ Padding: *fillinPadding }
    if len(fillin.Args()) == 1 {
      filter = fillin.Args()[0]
    }
  }
  // subcommand ERASE
  if erase.Parsed() {
    if *eraseT != "" { *eraseTarget = *eraseT }
    // exit when erase target not found
    if *eraseTarget == "" {
      fmt.Printf("%s", usage)
      os.Exit(1)
    }
    comm = &rname.EraseCommand{ Target: *eraseTarget }
    if len(erase.Args()) == 1 {
      filter = erase.Args()[0]
    }
  }
  // subcommand REGEX
  if regex.Parsed() {
    if *regexP != "" { *regexPattern = *regexP }
    if *regexR != "" { *regexReplace = *regexR }
    // exit when regex pattern not found
    if *regexPattern == "" {
      fmt.Printf("%s", usage)
      os.Exit(1)
    }
    re, err := rname.CompileStringToRegexp(*regexPattern)
    // exit when compile fail
    if err != nil {
      fmt.Printf("Error: %v", err)
      os.Exit(1)
    }
    comm = &rname.RegexCommand{ Pattern: *regexPattern, Re: re, Replace: *regexReplace }
    if len(regex.Args()) == 1 {
      filter = regex.Args()[0]
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

    newfpath := comm.Rewrite(fpath)

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


