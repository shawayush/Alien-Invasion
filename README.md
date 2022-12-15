# Alien-Invasion

This contains a simulation for Alien-Invasion in golang.

Please use go 1.18+ version, you could use go installer from [here](https://go.dev/doc/install)

## Approcah to solve the problem:

The problem was broken into three parts:

### 1. Create a Map 

Parse the input file and create a map from parsing the input and using nodes and Links to connect the city with each other, A tree like Data Strucuture is formed.

### 2. Create Aliens

Parse the alien file and make alien with different attributes which could be used for simulation, while invadng cities and keep the track of the aliens. Also created another variable which would also be keeping the track of the simualtion going in a sublte way.

So in total, a object for cities, Aliens and simualtion is created.

### 3. Run simulation

Once everything is readfy for runnig the simulation. Which simulaties, aliens invading, city being destroyed and aliens being killed at the approach and also gives the O/P of how the simulation is running.


## Running the Simulation

Intially you could use `go run main.go -help`. This would Printout the following Output:

```
 -Alien string
        a file used to identify aliens (default "./test/aliens.txt")
  -aliens int
        number of aliens invading during an invasion (default 10)
  -city string
        a file used as city input to make simulation (default "./test/example4.txt")
  -iterations int
        number of iterations to simulate (default 10000)
```

So the simulation can run in different ways. You could directly start using `go run main.go` . This would further run the simulation with the default configuration.

You could specify the number of aliens to be invading like:
```
go run main.go -aliens 1
```
The above would only run the simualtion with the only 1 alien. You could n number to run the silunation. Further More the first 40 aliens are named, for more than 40 random intiger would be used. This intiger is created through time from the OS.

You could also use given the input for number of itterations to be used:
```
go run main.go -iterations 2
```
You could Input the city file by the following command :
```
go run main.go -city examplecli.txt
```

You could also use multiple inputs for controlling the simulation as:
```
go run main.go -city examplecli.txt -iterations 3 -aliens 30
```

Attached a [sample output](https://github.com/shawayush/Alien-Invasion/blob/main/sampleoutput.txt)
