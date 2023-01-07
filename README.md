# MMake
### A much simpler version of make written in Go

## Usage
- Compile the program
```mmake compile [<mmakefile>]```
- Compile & run the program
```mmake run [<makefile>]```
- Clean the project
```mmake clean [<makefile>]```

## The MMakefile
The MMakefile is a YAML file. It should:
- be in the project root
- have the filename ```MMakefile```

The only required property is ```name```. This is the structure of the file: <br/>
```yaml
name: [Project Name; will be used as the output name]
cc: [Compiler; g++ is the default one]
ccFlags: [Extra flags passed to the compiler]
outputFormat: [The output name; %s will be replaced with the name; Example: %s.exe; Leave blank to auto detect]
outputDirectory: [The output directory; bin by default]
include: [Include locations; default are include and headers]
  - include
  - headers
sourceOptions:
  files: [Glob patterns for source files; *.c *.cc *.cpp *.c++ are the default files]
  directoreis: [Directory glob patterns; keep in mind that if you want to include a directory and it's subdirectories, you'll need to write both dir/ and dir/**; default is source/ src/ sources/ source/** src/** sources/**]
```
