# rname

rename utility for windows command prompt.


## Usage

```
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

  FILE_PATTERN is shell file name syntax like '*.jpg'.
  when FILE_PATTERN is omitted, '*' is used.
```


## Example

in Windows comannd prompt.

```
E:\temp>dir /b
(demo) abc.mp3
(demo) xyz.mp3
日本語-1.txt
日本語-10.txt
日本語-100.txt

E:\temp>rname test prepend
   (demo) abc.mp3
=> (demo) abc.mp3
   (demo) xyz.mp3
=> (demo) xyz.mp3
   日本語-1.txt
=> 日本語-00001.txt
   日本語-10.txt
=> 日本語-00010.txt
   日本語-100.txt
=> 日本語-00100.txt

E:\temp>rname test fillin *.mp3
   (demo) abc.mp3
=> (demo)_abc.mp3
   (demo) xyz.mp3
=> (demo)_xyz.mp3

E:\temp>rname test serial *.txt
   日本語-1.txt
=> 00001.txt
   日本語-10.txt
=> 00002.txt
   日本語-100.txt
=> 00003.txt

E:\temp>rname test regex -p="^(.)(.)" -r="$2$1" *.txt
   日本語-1.txt
=> 本日語-1.txt
   日本語-10.txt
=> 本日語-10.txt
   日本語-100.txt
=> 本日語-100.txt

E:\temp>
```


## Installation

windows binary is [here](https://github.com/inazak/rname/releases)


