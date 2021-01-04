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
- Clean code
- Create debug option
- Check if mutation percentage should consider the number of individuals or genes


# One Max Problem
The objective here is to generate an individual (composite of 0's or 1's) with all bits 1 evolving the initial population.

Usage:
- Define the population size, number of genes, number of generations to evolve.
- Set the parameters for the number of participants on tournament, the crossover (how many individuals from the current generation will pass to the next one) and mutation rate (recommended to be really slow to avoid to start depending on randomically changes to reach the objective).

`go run one_max.go`
