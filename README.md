# Genetic Algorithms

A genetic algorithm is a search heuristic that is inspired by Charles Darwin's theory of natural evolution. This algorithm reflects the process of natural selection where the fittest individuals are selected for reproduction in order to produce offspring of the next generation, using these steps:

1 - Generate the **Initial Population**

2 - **Evaluation** of the current population (fitness) accordingly to what is expected from these individuals

3 - **Define the parents** based on Evaluation (based on roulette, rank or tournament)

4 - **Elitism** passing a percentage of extraordinary individuals to the next generation

5 - **Generate the children** (cut-point or crossover)

6 - **Mutation**

7 - Find the **Best Individual**

# One Max Problem
The objective here is to generate an individual (composite of 0's or 1's) with all bits 1 evolving the initial population.

**Pop. of 100 reached 3000 1's** | **Pop. of 250 reached 5000 1's**
:-------------------------:|:-------------------------:
<img width="430" alt="horizontal" src="https://github.com/cassianoperin/Genetic_Algorithms/blob/main/Images/Onemax_3000.png">  |  <img width="430" alt="vertical" src="https://github.com/cassianoperin/Genetic_Algorithms/blob/main/Images/Onemax_5000.png">


# Crack Passwords
The objective is to discover a password. The unique rule is to inform the program the number of characters of the password (gene_number).

**Pop. of 200 - 50 characters pwd** | **Pop. of 500 - 100 characters pwd**
:-------------------------:|:-------------------------:
<img width="430" alt="horizontal" src="https://github.com/cassianoperin/Genetic_Algorithms/blob/main/Images/Password_50.png">  |  <img width="430" alt="vertical" src="https://github.com/cassianoperin/Genetic_Algorithms/blob/main/Images/Password_100.png">

# Next steps

- Make parameters configurable via ini files

## Usage:

Just run the executable files.

###  Define the population size, number of genes, number of generations to evolve:

(Hardcoded as variables at this moment on each file)

- Number of generations (Generations)
- Population size (Population_size)
- Number of genes (Gene_number)
- Number of participants of tournament for parents selection (K)
- Crossover rate (Crossover_rate): How many individuals from the current generation will pass to the next one
- Mutation rate (Mutation_rate): Recommended to be really low to avoid to start depending on randomically changes to reach the objective
- Elitism percentual (Elitism_percentual)

## Compile

### 1. Mac

- 32 bits:

`env GOOS="darwin" GOARCH="386" go build -ldflags="-s -w" one_max.go`

`env GOOS="darwin" GOARCH="386" go build -ldflags="-s -w" crack_password.go`

- 64 bits:

`env GOOS="darwin" GOARCH="amd64" go build -ldflags="-s -w" one_max.go`

`env GOOS="darwin" GOARCH="amd64" go build -ldflags="-s -w" crack_password.go`

#### Compress binaries

`brew install upx`

`upx <binary_file>`

### 2. Linux

Instructions to build using Ubuntu.

#### Install requisites:

`sudo apt install pkg-config libgl1-mesa-dev licxi-dev libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev mesa-utils build-essential xorg-dev upx`

#### Build:

- 32 bits:

`env GOOS="linux" GOARCH="386" go build -ldflags="-s -w" one_max.go`

`env GOOS="linux" GOARCH="386" go build -ldflags="-s -w" crack_password.go`

- 64 bits:

`env GOOS="linux" GOARCH="amd64" go build -ldflags="-s -w" one_max.go`

`env GOOS="linux" GOARCH="amd64" go build -ldflags="-s -w" crack_password.go`

#### Compress binaries:

`upx <binary_file>

### 3. Windows

GO allows to create a Windows executable file using a MacOS:

#### Install mingw-w64 (support the GCC compiler on Windows systems):

`brew install mingw-w64`

#### Compile:

- 32 bits:

`env GOOS="windows" GOARCH="386" CGO_ENABLED="1" CC="i686-w64-mingw32-gcc" go build -ldflags="-s -w" one_max.go`

`env GOOS="windows" GOARCH="386" CGO_ENABLED="1" CC="i686-w64-mingw32-gcc" go build -ldflags="-s -w" crack_password.go`

- 64 bits:

`env GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build -ldflags="-s -w" one_max.go`

`env GOOS="windows" GOARCH="amd64" CGO_ENABLED="1" CC="x86_64-w64-mingw32-gcc" go build -ldflags="-s -w" crack_password.go`

* If you receive the message when running the executable, you need to ensure that the video drivers supports OpenGL (or the virtual driver in the case of virtualization).

* If you receive this message : "APIUnavailable: WGL: The driver does not appear to support OpenGL", please update your graphics driver os just copy the Mesa3D library from https://fdossena.com/?p=mesa/index.frag  (opengl32.dll) to the executable folder.

#### Compress binaries

`brew install upx`

`upx <binary_file>`
