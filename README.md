# Genetic Algorithms

A genetic algorithm is a search heuristic that is inspired by Charles Darwin's theory of natural evolution. This algorithm reflects the process of natural selection where the fittest individuals are selected for reproduction in order to produce offspring of the next generation, using these steps:

1 - Generate the **Initial Population**

2 - **Evaluation** of the current population (fitness) accordingly to what is expected from these individuals

3 - **Define the parents** based on Evaluation (based on roulette, rank or tournament)

4 - **Elitism** passing a percentage of extraordinary individuals to the next generation

5 - **Generate the children** (cut-point or crossover)

6 - **Mutation**

7 - Find the **Best Individual**

TODO:
- Create debug option
- Add evolution over generations graph
- Add mutation (individuals) over generations graph
- Add mutation (genes) over generations graph
- Add crossover over generations graph
- Mutation is too high when we have too many genes, maybe calculate it automatically to match 0,05


# One Max Problem
The objective here is to generate an individual (composite of 0's or 1's) with all bits 1 evolving the initial population.

Usage:
- Define the population size, number of genes, number of generations to evolve.
- Set the parameters for the number of participants on tournament, the crossover (how many individuals from the current generation will pass to the next one) and mutation rate (recommended to be really slow to avoid to start depending on randomically changes to reach the objective).

Requisites:
`go get github.com/faiface/pixel`
`go get github.com/faiface/pixel/imdraw`
`go get github.com/faiface/pixel/pixelgl`
`go get golang.org/x/image/colornames`
`go get github.com/faiface/pixel/text`
`go get github.com/faiface/pixel`

Run:
`go run one_max.go`
