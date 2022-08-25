package rname

import (
  "strings"
  "strconv"
  "fmt"
  "os"
  "path/filepath"
  "regexp"
)

// filesystem utility

func GetFilepathList(filter string) (list []string, err error) {

  entries, err := filepath.Glob(filter)
  if err != nil {
    return nil, err
  }

  for _, entry := range entries {
    isdir, err := isDir(entry)
    if err != nil {
      return nil, err
    }
    if ! isdir {
      list = append(list, entry)
    }
  }

  // both absolute and relative paths are included here
  return list, nil
}


func isDir(p string) (bool, error) {
  f, err := os.Stat(p)
  if err != nil {
    return false, err
  }
  return f.Mode().IsDir(), nil
}


func splitFilepath(fpath string) (dir, name, ext string) {

  dir   = filepath.Dir(fpath)
  base := filepath.Base(fpath)

  dotindex := strings.LastIndex(base, ".")

  // file has no extention
  if dotindex == -1 {
    name = base
    ext  = ""
    return
  }

  name = base[0:dotindex]
  ext  = base[dotindex:len(base)]
  return
}



// subcommand interface

type Command interface {
  Rewrite(string) string
}


// subcommand PREPEND

type PrependCommand struct {
  Width int
}

func (p *PrependCommand) Rewrite(fpath string) (newfpath string) {
  dir, name, ext := splitFilepath(fpath)
  return filepath.Join(dir, prependZeros(name, p.Width) + ext)
}

func prependZeros(s string, w int) string {
  r := strings.LastIndexAny(s, "0123456789")
  l := r

  if r == -1 { // number not found
    return s
  }

  for t := r-1; t >= 0; t -= 1 {
    if ! strings.ContainsAny(string(s[t]), "0123456789") { break }
    l = t
  }

  if r+1-l >= w { // number length is not enough
    return s
  }

  number, _ := strconv.Atoi(s[l:r+1])
  format := fmt.Sprintf("%%0%dd", w)  // => "%0?d"
  filled := fmt.Sprintf(format, number)

  return s[0:l] + filled + s[r+1:len(s)]
}


// subcommand SERIAL

type SerialCommand struct {
  Width   int
  Current int
}

func (s *SerialCommand) Rewrite(fpath string) (newfpath string) {
  dir, name, ext := splitFilepath(fpath)
  format := fmt.Sprintf("%%0%dd", s.Width)
  name    = fmt.Sprintf(format, s.Current)
  s.Current += 1
  return filepath.Join(dir, name + ext)
}


// subcommand FILLIN

type FillinCommand struct {
  Padding string
}

func (f *FillinCommand) Rewrite(fpath string) (newfpath string) {
  dir, name, ext := splitFilepath(fpath)
  name = strings.Replace(name, " ", f.Padding, -1)
  return filepath.Join(dir, name + ext)
}


// subcommand ERASE

type EraseCommand struct {
  Target string
}

func (e *EraseCommand) Rewrite(fpath string) (newfpath string) {
  dir, name, ext := splitFilepath(fpath)
  name = strings.Replace(name, e.Target, "", -1)
  return filepath.Join(dir, name + ext)
}


// subcommand REGEX

type RegexCommand struct {
  Pattern string
  Re      *regexp.Regexp
  Replace string
}

func (r *RegexCommand) Rewrite(fpath string) (newfpath string) {
  dir, name, ext := splitFilepath(fpath)
  name = r.Re.ReplaceAllString(name, r.Replace)
  return filepath.Join(dir, name + ext)
}


// regexp compile

func CompileStringToRegexp(s string) (*regexp.Regexp, error) {
  return regexp.Compile(s)
}


