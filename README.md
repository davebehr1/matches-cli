# spanassessment

[![Test](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml)



## Run tests:
  make test
  
## modes of execution:
  * ### docker:
    * make docker
    * docker run matchescli args...
 * ### executable:
    * go build -o matchescli
    * ./matchescli grt --f=matches.txt

## Cli commands:

*   ### generateranktable / grt:
     *  Flags:
        * --f
        * --help / help for command

     *   Example:
           *  Go run main.go grt --f=matches.txt’
           *  Go run main.go grt
