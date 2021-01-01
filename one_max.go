package main

import (
  "os"
  "time"
  "strings"
  "fmt"
  "math/rand"
  "strconv"
  "regexp"
)

var (
  population_size int = 30
  gene_number int = 50
  population []string
  k = 3 // Tournament size (number of participants)
  crossover_rate float64 = 0.7
  mutation_rate float64 = 0.005
  generations int = 30
)


// ------------------- Validate Parameters -------------------- //
func validate_parameters() {
  // Minimal Population Size size accepted is 2
  if population_size % 2 == 1 {
    fmt.Printf("\nPopulation size should be ODD numbers. Exiting\n")
    os.Exit(2)
  }

  // Population Size should be positive
  if population_size <= 0 {
    fmt.Printf("\nPopulation size should be Positive. Exiting\n")
    os.Exit(2)
  }
}


// ------------------- Generate Individuals ------------------- //
func generate_individuals(gene_number int) string {
  var individual string = ""

  // Initialize rand source
  rand.Seed(time.Now().UnixNano())

  for i := 0 ; i < gene_number ; i++ {
    individual += strconv.Itoa(rand.Intn(2))
  }

  return individual
}


// ------------------- Generate Population -------------------- //
func generate_population(population_size int, gene_number int) {

  fmt.Printf("Initializing the Population:\n")

  // Generate each individual for population
  for i := 0 ; i < population_size ; i++ {
    population = append( population, generate_individuals(gene_number) )
  }

}


// ------------------- Generate Evaluation -------------------- //
func fitness(individual string) int {
  ones := regexp.MustCompile("1")
  matches := ones.FindAllStringIndex(individual, -1)

  return len(matches)
}


// ---------------------- Define Parents ---------------------- //
func define_parents(pop []string, pop_size int, k int) []string {
  var parents []string

  // Quantity of tournaments is equal to the size of population
  for tournament := 0 ; tournament < pop_size ; tournament ++ {
    var (
      competitors []string
      score []int
    )

    // Each tournament, K competitors
    for i := 0 ; i < k ; i++ {
      competitors = append( competitors, pop[rand.Intn(pop_size)] )
    }

    // Calculate the score of K competitors
    for i := 0 ; i < k ; i++ {
      score = append( score, fitness(competitors[i]) )
    }

    bigger := score[0]
    winner := competitors[0]

    for i := 0 ; i < k ; i++ {
      if score[i] > bigger {
        bigger = score[i]
        winner = competitors[i]
      }
    }

    parents = append(parents, winner)

    fmt.Printf("Tournament: %d\t Competitors: %s\t Scores: %d\t Winner: %s (%d)\n", tournament, competitors, score, winner, bigger)

  }

  return parents

}


// -------------------- Generate Children --------------------- //
func generate_children(parents []string, pop_size int) []string {
  var (
    father1, father2, child1, child2 string
    pop_new []string
  )

  fmt.Printf("\nSelected parents:\n")

  for i := 0 ; i < pop_size / 2 ; i++ {
    // Define the couples
    randomIndex := rand.Intn(len(parents))
    father1 = parents[randomIndex]

    randomIndex = rand.Intn(len(parents))
    father2 = parents[randomIndex]

    fmt.Printf("%d) %s with %s\n", i, father1, father2)

    // Define if will have crossover (the parents will be copied to next generation)
    if rand.Float64() < crossover_rate {

      // Define the cut-point
      cut_point := rand.Intn(gene_number -1) + 1
      fmt.Printf("\tCut-point: %d\n",cut_point)

      // Split father's values
      // Father1
      father1_split := strings.Split(father1,"")
      father1_split_p1 := father1_split[0:cut_point]
      father1_split_p2 := father1_split[cut_point:]
      // Father2
      father2_split := strings.Split(father2,"")
      father2_split_p1 := father2_split[0:cut_point]
      father2_split_p2 := father2_split[cut_point:]

      // Child1
      child1_p1 := strings.Join(father1_split_p1,"")
      child1_p2 := strings.Join(father2_split_p2,"")
      child1 = child1_p1 + child1_p2
      fmt.Printf("\tChild1: %s + %s: %s\n", child1_p1, child1_p2, child1)

      // Child2
      child2_p1 := strings.Join(father2_split_p1,"")
      child2_p2 := strings.Join(father1_split_p2,"")
      child2 = child2_p1 + child2_p2
      fmt.Printf("\tChild2: %s + %s: %s\n", child2_p1, child2_p2, child2)

      // Put the childs in the new generation
      pop_new = append(pop_new, child1)
      pop_new = append(pop_new, child2)

    } else {
      fmt.Printf("\tCrossover:\n")
      pop_new = append(pop_new, father1)
      pop_new = append(pop_new, father2)
      fmt.Printf("\tChild1 (Father1): %s\n", father1)
      fmt.Printf("\tChild2 (Father2): %s\n", father2)
    }

  }

  return pop_new
}



// ------------------------- Mutation ------------------------- //

func generate_mutation(new_pop []string, pop_size int, gene_number int, mutation_rate float64) []string {

  var new_pop_mutated []string

  // For all individuals in population
  for i := 0 ; i < pop_size ; i ++ {

    var individual string = ""
    individual = new_pop[i]

    // For each gene, check for mutations
    for gene := 0 ; gene < gene_number ; gene ++ {

      // Check if there is a mutation
      if mutation_rate >= rand.Float64() {

        individual_split := strings.Split(individual,"")

        // Invert the mutated gene
        if individual_split[gene] == "0" {
          individual_split[gene] = "1"

        } else {
          individual_split[gene] = "0"
        }

        // Update the mutated individual
        individual = strings.Join(individual_split,"")

        fmt.Printf("Individual #%d (%s) mutated on gene %d. New Individual: %s \n", i, new_pop[i], gene, individual)

      }

    }

    // Add mutated individuals to a new generation
    new_pop_mutated = append(new_pop_mutated, individual)

  }

  return new_pop_mutated

}


// --------------------- Best Individual ---------------------- //
func best_individual() (string, int) {
  var score []int

  // Calculate the score of the latest population
  for i := 0 ; i < len(population) ; i++ {
    score = append( score, fitness(population[i]) )
  }

  bigger := score[0]
  winner := population[0]

  for i := 0 ; i < len(score) ; i++ {
    if score[i] > bigger {
      bigger = score[i]
      winner = population[i]
    }
  }

  return winner, bigger
}



func main() {

  // Validate parameters
  validate_parameters()

  // 0 - Generate the population
  generate_population(population_size, gene_number)
  fmt.Printf("%s\n\n", population)

  for i := 0 ; i < generations ; i++ {

    fmt.Printf("\n// ---------------------------------- GENERATION: %d ---------------------------------- //\n\n", i)


    // 1 - Evaluation
    fmt.Printf("1 - Evaluation:\n")
    for i := 0 ; i < population_size ; i ++ {
      fmt.Printf("\nIndividual: %s\tEvaluation %d\n", population[i], fitness(population[i]))
    }

    // 2 - Define Parents
    fmt.Printf("\n2 - Define Parents:\n")
    parents := define_parents(population, population_size, k)
    fmt.Printf("\nParents:\n%s\n\n",parents)

    // 3 - Generate Children
    fmt.Printf("\n3 - Generate Chindren:\n")
    new_population := generate_children(parents, population_size)
    fmt.Printf("\nNew population:\n%s\n", new_population)

    // 4 - Mutation
    fmt.Printf("\n4 - Mutation:\n")
    new_population = generate_mutation(new_population, population_size, gene_number, mutation_rate)
    fmt.Printf("Mutated Generation: %s\n\n", new_population)

    // 5 - Replace population vector with new population one
    population = nil    // Clean ond population
    for i:= 0 ; i < len(new_population) ; i++ {
      population = append(population, new_population[i])
    }
  }

  // 6 - Best individual
  best, score := best_individual()
  fmt.Printf("\nBest Individual: %s with score %d\n\n", best, score)
}
