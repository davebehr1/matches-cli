# spanassessment

[![Test](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml)


```
This simple cli is used to load match resuts via a file or stdin to generate a rank table.
Rules:
  * if a team has won a game they get 3 points
  * if they draw they get one point
  * if they lose they get no points
```

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
