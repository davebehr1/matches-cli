# spanassessment

[![Test](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml/badge.svg)](https://github.com/davebehr1/spanassessment/actions/workflows/test.yml)


```
This simple cli is used to load match results via a file or stdin to generate a rank table.
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
        * --f / the file with match results / 
          * if you run the cli without this command you will be prompted for input
        * --o / the file the rank table will be written to / 
          * if you run the cli without this flag output will be written to stdout
        * --help / help
     
     *   Example:
           *  go run main.go grt --f=matches.txt
           *  go run main.go grt --f=matches.txt --o=ranktable.txt
           *  go run main.go grt --o=ranktable.txt
           *  go run main.go grt
              * you will get a prompt ```Enter Match Result:``` enter the match result in the format ```cheetahs 1, bulls 2```
              * to conclude entering in match results and see the final rank table type in ```done``` at the next prompt
     * Example Result:
     ```
     1. Lions, 11 pts 
     2. Tarantulas, 6 pts 
     3. FC Awesome, 1 pts 
     4. Snakes, 1 pts 
     5. Cheetahs, 0 pts 
     6. Grouches, 0 pts 
     7. Pumas, 0 pts 
     ```
