# go(lang) notes

Notes from my personal go-ing journey.

## features
- binary executable
- minimalist
- automatic garbage collection
- built in formatting engine
- built in testing, benchmarking
- advanced concurrency techniques
- low boilerplate code
- networking api for distributed, cloud applications

## go CLI commands

`go mod init <module_name>` - init a go module. will create go.mod file

`GOOS="linux" go build` - build binary for a platform

## more notes

- Is Go object oriented?
- go code: modules -> packages
- package - made up of all the files in directory
- each module has a go.mod file 
- go mod init <module-name/url> 
- using a local module (not-published) - 
`go mod edit -replace <module-path>=<local-path>`
example - 
`go mod edit -replace example.com/greetings=../greetings`
- 