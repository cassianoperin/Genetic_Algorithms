package main

import (
  "os"
  "time"
  "strings"
  "sort"
  "fmt"
  "math/rand"
  "strconv"
  "regexp"
  // Graphics
  "github.com/faiface/pixel"
  "github.com/faiface/pixel/imdraw"
  "github.com/faiface/pixel/pixelgl"
  "golang.org/x/image/colornames"
	"github.com/faiface/pixel/text"
  "golang.org/x/image/font/basicfont"
)

var (
  // Main variables
  population_size int = 10
  gene_number int = 10
  k = 2 // Tournament size (number of participants)
  crossover_rate float64 = 0.7
  mutation_rate float64 = 0.0005  // 0,5% (I'm analyzing each gene so the mutation rate should be really small)
  generations int = 10
  elitism_percentual int = 10  // 10% of population size

  // Other variables
  population []string
  elitism_individuals int = (elitism_percentual * population_size) / 100
  debug bool = false

  // Graphics
  sizeX float64 = 1024
  sizeY float64 = 768
  graph_fitness []int
  // graph []int
  pixelY_size float64 = (sizeY -100) / float64(gene_number)
  pixelX_size float64 = (sizeX -100) / float64(generations)

  // Counters
  mutation_count, mutation_ind_count int
)

func graphics() {

  imd := imdraw.New(nil)

	cfg := pixelgl.WindowConfig{
		Title:  "Genetic Algorithms - One Max Problem",
		Bounds: pixel.R(0, 0, sizeX, sizeY),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

  basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)


  // ------------------- Draw Cartesian Plane ------------------- //
  imd.Color = colornames.Gray
  // X
  imd.Push(pixel.V(50, 0), pixel.V(50, sizeY - 50)) // Initial X,Y -> Final X,Y
  imd.Line(2)
  // X Arrow
  imd.Push(pixel.V(sizeX-50, 40))
  imd.Push(pixel.V(sizeX-50, 60))
  imd.Push(pixel.V(sizeX-30, 50))
	imd.Polygon(0)
  // Y
  imd.Push(pixel.V(0, 50), pixel.V(sizeX - 50, 50))
  imd.Line(2)
  // Y Arrow
  imd.Push(pixel.V(40, sizeY-50))
  imd.Push(pixel.V(60, sizeY-50))
  imd.Push(pixel.V(50, sizeY-30))
	imd.Polygon(0)

  // Plane Cartesian Zero
	txtCartesianZero := text.New(pixel.V(40, 25), basicAtlas)
  txtCartesianZero.Color = colornames.Black
	fmt.Fprintf(txtCartesianZero, "0")
  // X Generations Label
  txtCartesianGen := text.New(pixel.V(sizeX -950, 25), basicAtlas)
  txtCartesianGen.Color = colornames.Black
	fmt.Fprintf(txtCartesianGen, "Generations")
  // Y Fitness Label
  txtCartesianFit := text.New(pixel.V(1, 75), basicAtlas)
  txtCartesianFit.Color = colornames.Black
  fmt.Fprintf(txtCartesianFit, "Fitness")

  // ------------- Draw Max Fitness and Generations ------------- //

  // 100% Fitness Line
  imd.Color = colornames.Lightgray
  imd.Push(pixel.V(50, pixelY_size*float64(gene_number)), pixel.V(sizeX - 50 - pixelX_size, pixelY_size*float64(gene_number) ) )
  imd.Line(1)

  // Draw Generation Result Line
  imd.Push(pixel.V(50+(pixelX_size*float64(len(graph_fitness)-1)), 50), pixel.V(50+(pixelX_size*float64(len(graph_fitness)-1)), pixelY_size*float64(gene_number) ) )
  imd.Line(1)

  // Y 100% Label
  txtMaxY := text.New(pixel.V(20, pixelY_size*float64(gene_number)-5), basicAtlas)
  txtMaxY.Color = colornames.Black
	fmt.Fprintf(txtMaxY, strconv.Itoa(gene_number))

  // X 100% Label
	txtMaxX := text.New(pixel.V(sizeX-80, 25), basicAtlas)
  txtMaxX.Color = colornames.Black
	fmt.Fprintf(txtMaxX, strconv.Itoa(generations))

  // ------------- Draw Max Fitness and Generations ------------- //

  // Generation result Label
	txtResultGen := text.New(pixel.V(50+(pixelX_size*float64(len(graph_fitness)-1)), 25), basicAtlas)
  txtResultGen.Color = colornames.Blue
	fmt.Fprintf(txtResultGen, strconv.Itoa(len(graph_fitness)))

  // Fitness Result Label
  len_graph:= len(graph_fitness)
  txtResultFit := text.New(pixel.V(sizeX -45, pixelY_size*float64(graph_fitness[len_graph-1])-5), basicAtlas)
  txtResultFit.Color = colornames.Blue
	fmt.Fprintf(txtResultFit, strconv.Itoa(graph_fitness[len_graph-1]))

  // Y Start Label
  txtStartY := text.New(pixel.V(20, pixelY_size*float64(graph_fitness[0])-5), basicAtlas)
  txtStartY.Color = colornames.Blue
  text :=  strconv.Itoa(graph_fitness[0])
  // basicTxt5.Dot.X = basicTxt5.BoundsOf(text).W()
  fmt.Fprintf(txtStartY, text)

  // -------------------- Draw Fitness Graph -------------------- //
  x := float64(50)
  y := float64(graph_fitness[0]) * pixelY_size
  for i := 1 ; i < len(graph_fitness) ; i++ {
    if i % 2 == 0{
      imd.Color = colornames.Blue
    } else {
      imd.Color = colornames.Lightblue
    }
    // Initial X,Y -> Final X,Y
    imd.Push(pixel.V(x, y), pixel.V(x+pixelX_size, float64(graph_fitness[i]) * pixelY_size) )
    x+=pixelX_size
    y = float64(graph_fitness[i]) * pixelY_size
  }
	imd.Line(2)


  // ---------------------- Render Graphics --------------------- //
	for !win.Closed() {
		win.Clear(colornames.White)
    // Draw Objects
		imd.Draw(win)
    // Cartesian Plane text
    txtCartesianZero.Draw(win, pixel.IM.Scaled(txtCartesianZero.Orig, 1))
    txtCartesianGen.Draw(win, pixel.IM.Scaled(txtCartesianGen.Orig, 1))
    txtCartesianFit.Draw(win, pixel.IM.Scaled(txtCartesianFit.Orig, 1))
    // Draw Max Fitness and Generations
    txtMaxY.Draw(win, pixel.IM.Scaled(txtMaxY.Orig, 1))
    if len(graph_fitness) != generations {
      txtMaxX.Draw(win, pixel.IM.Scaled(txtMaxX.Orig, 1))
    }
    // Draw result information on graphic
    txtResultGen.Draw(win, pixel.IM.Scaled(txtResultGen.Orig, 1))
    txtResultFit.Draw(win, pixel.IM.Scaled(txtResultFit.Orig, 1))
    txtStartY.Draw(win, pixel.IM.Scaled(txtStartY.Orig, 1))

    // Update screen
		win.Update()
	}
}



// ------------------- Validate Parameters -------------------- //
func validate_parameters(pop_size int, competitors int) {
  // Minimal Population Size size accepted is 2
  if pop_size % 2 == 1 {
    fmt.Printf("\nPopulation size should be ODD numbers. Exiting\n")
    os.Exit(0)
  }

  // Population Size should be positive
  if pop_size <= 0 {
    fmt.Printf("\nPopulation size should be Positive. Exiting\n")
    os.Exit(0)
  }

  // K (competitors) must be at least 2
  if competitors < 2 {
    fmt.Printf("\nNumber of competitors (k) must be at least 2. Exiting\n")
    os.Exit(0)
  }
}


// ------------------- Generate Individuals ------------------- //
func generate_individuals(gene_nr int) string {
  var individual string = ""

  // Initialize rand source
  rand.Seed(time.Now().UnixNano())

  for i := 0 ; i < gene_nr ; i++ {
    individual += strconv.Itoa(rand.Intn(2))
  }

  return individual
}


// --------- Generate the Evaluation of an Individual --------- //
func fitness_individual(individual string) int {
  ones := regexp.MustCompile("1")
  matches := ones.FindAllStringIndex(individual, -1)

  return len(matches)
}


// --------- Generate the Evaluation of a Population ---------- //
func fitness_population(pop []string) []int {

  var score []int

  for i := 0 ; i < len(pop) ; i++ {
    score = append( score, fitness_individual(pop[i]) )
  }

  return score
}


// ------------------------- Elitism -------------------------- //
func elitism(pop []string, pop_score []int, pop_size int, elitism_number int) ([]string, []string) {
  var (
    elite, elite_score, tmp_slice []string
  )

  // Append score + individual in one slice
  for i := 0 ; i < pop_size ; i++ {
    tmp_slice = append( tmp_slice, strconv.Itoa(pop_score[i]) + "," + pop[i] )
  }

  // Sort slice
  sort.Strings(tmp_slice)

  // Insert individuals on Elite slice and score on elite_score
  for i := pop_size -1 ; i > ( pop_size -1 ) - elitism_number ; i-- {
    tmp_slice := strings.Split(tmp_slice[i],",")
    elite = append( elite, tmp_slice[1] )       // Individual
    elite_score = append( elite_score, tmp_slice[0] )   //Score
  }

  return elite, elite_score
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
      score = append( score, fitness_individual(competitors[i]) )
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

    if debug {
      fmt.Printf("\tTournament: %d\t Competitors: %s\t Scores: %d\t Winner: %s (%d)\n", tournament, competitors, score, winner, bigger)
    }

  }

  return parents

}


// -------------------- Generate Children --------------------- //
func generate_children(parents []string, pop_size int, elitism_number int, elite []string) ([]string, int) {
  var (
    father1, father2, child1, child2 string
    pop_new []string
    cross_count int = 0
  )

  if debug {
    fmt.Printf("\n\tSelected parents:\n")
  }

  for i := 0 ; i < pop_size / 2 ; i++ {
    // Define the couples
    randomIndex := rand.Intn(len(parents))
    father1 = parents[randomIndex]

    randomIndex = rand.Intn(len(parents))
    father2 = parents[randomIndex]

    if debug {
      fmt.Printf("\t%d) %s with %s\n", i, father1, father2)
    }

    // Define if will have crossover (the parents will be copied to next generation)
    if rand.Float64() < crossover_rate {

      // Define the cut-point
      cut_point := rand.Intn(gene_number -1) + 1
      if debug {
        fmt.Printf("\t\tCut-point: %d\n",cut_point)
      }

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
      if debug {
        fmt.Printf("\t\tChild1: %s + %s: %s\n", child1_p1, child1_p2, child1)
      }

      // Child2
      child2_p1 := strings.Join(father2_split_p1,"")
      child2_p2 := strings.Join(father1_split_p2,"")
      child2 = child2_p1 + child2_p2
      if debug {
        fmt.Printf("\t\tChild2: %s + %s: %s\n", child2_p1, child2_p2, child2)
      }

      // Put the childs in the new generation
      pop_new = append(pop_new, child1)
      pop_new = append(pop_new, child2)

    } else {
      if debug {
        fmt.Printf("\t\tCrossover:\n")
      }
      pop_new = append(pop_new, father1)
      pop_new = append(pop_new, father2)
      if debug {
        fmt.Printf("\t\tChild1 (Father1): %s\n", father1)
        fmt.Printf("\t\tChild2 (Father2): %s\n", father2)
      }
      cross_count++
    }

  }

  // Ensure place of elite members on next generation
  if elitism_number > 0 {
    if debug {
      fmt.Printf("\n\tElitism: Regular individual removal:\n")
    }

    // Remove randomically the number os elite elements
    for i := 0 ; i < elitism_number ; i++ {
      random := rand.Intn(len(pop_new))
      if debug {
        fmt.Printf("\t\tIndividual %d:\t%s removed randomically from new population\n", i, pop_new[random])
      }

      // Remove the element at index 'random' from pop_new
      pop_new[random] = pop_new[len(pop_new)-1] // Copy last element to index 'random'.
      pop_new[len(pop_new)-1] = ""   // Erase last element (write zero value).
      pop_new = pop_new[:len(pop_new)-1]   // Truncate slice.
    }

    // Insert Elite Members on next generation
    if debug {
      fmt.Printf("\n\tElitism: Elite individual insertion:\n")
    }
    for i := 0 ; i < elitism_number ; i++ {
      pop_new = append( pop_new, elite[i] )
      if debug {
        fmt.Printf("\t\tIndividual %d\t%s inserted to new population\n", i, elite[i])
      }
    }
  }

  return pop_new, cross_count
}


// ------------------------- Mutation ------------------------- //
func generate_mutation(new_pop []string, pop_size int, gene_nr int, mutation_rate float64) ([]string, int, int) {

  var (
    new_pop_mutated []string
    count_genes int = 0
    count_individuals int = 0
  )

  // For all individuals in population
  for i := 0 ; i < pop_size ; i ++ {

    var(
      individual string = ""
      individual_mutated_flag bool
    )

    individual = new_pop[i]

    // For each gene, check for mutations
    for gene := 0 ; gene < gene_nr ; gene ++ {

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

        if debug {
          fmt.Printf("\tIndividual #%d (%s) mutated on gene %d. New Individual: %s \n", i, new_pop[i], gene, individual)
        }

        count_genes ++  // Generation genes mutated count
        individual_mutated_flag = true

      }

    }

    // Generation individuals mutated count
    if individual_mutated_flag {
      count_individuals ++
      individual_mutated_flag = false
    }

    // Add mutated individuals to a new generation
    new_pop_mutated = append(new_pop_mutated, individual)
  }

  return new_pop_mutated, count_genes, count_individuals
}


// --------------------- Best Individual ---------------------- //
func best_individual() (string, int) {
  var score []int

  // Calculate the score of the latest population
  score = fitness_population(population)

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



// ------------------------- MAIN FUNCTION ------------------------- //
func main() {


  // --------------------- Validate parameters --------------------- //
  validate_parameters(population_size, k)


  // ----------------- 0 - Generate the population ----------------- //
  // Generate each individual for population
  for i := 0 ; i < population_size ; i++ {
    population = append( population, generate_individuals(gene_number) )
  }


  // ----------------------- Generations Loop ---------------------- //
  for i := 0 ; i < generations ; i ++ {

    if debug {
      fmt.Printf("\n// ---------------------------------- GENERATION: %d ---------------------------------- //\n\n", i)
    }

    // ----------------------- 1 - Evaluation ------------------------ //
    if debug {
      fmt.Printf("1 - Evaluation:\n\n")
    }
    population_score := fitness_population(population)

    // Show the evaluation of each individual
    if debug {
      for i := 0 ; i < population_size ; i ++ {
        fmt.Printf("\tIndividual %d:\t%s\tEvaluation %d\n", i, population[i], population_score[i])
      }
    }


    // ---------------------- 2 - Define Parents --------------------- //
    if debug {
      fmt.Printf("\n2 - Define Parents:\n\n")
    }

    parents := define_parents(population, population_size, k)

    if debug {
      fmt.Printf("\n\tParents: %s\n\n",parents)
    }


    // ------------------------- 3 - Elitism ------------------------- //
    elite, elite_score := elitism(population, population_score, population_size, elitism_individuals)
    if debug {
      fmt.Printf("\n3 - Elitism:\n\n\tNumber of elite members: %d\n\n", elitism_individuals)
      for i := 0 ; i < elitism_individuals ; i ++ {
        fmt.Printf("\tIndividual %d:\t%s set for elite with score: %s\n", i, elite[i], elite_score[i] )
      }
    }


    // -------------------- 4 - Generate Children -------------------- //
    new_population, crossover_count := generate_children(parents, population_size, elitism_individuals, elite)
    if debug {
      fmt.Printf("\n4 - Generate Chindren:\n\n\tNew population: %s\n", new_population)
    }


    // ------------------------ 5 - Mutation ------------------------- //
    new_population, mutation_count, mutation_ind_count = generate_mutation(new_population, population_size, gene_number, mutation_rate)
    if debug {
      fmt.Printf("\n5 - Mutation:\n\tMutated Generation: %s\n\n", new_population)
    }


    // ---- 6 - Replace population vector with new population one ---- //
    population = nil    // Clean ond population
    for i:= 0 ; i < len(new_population) ; i++ {
      population = append(population, new_population[i])
    }

    // -------------------- 7 - Best individual ---------------------- //
    best, score := best_individual()
    fmt.Printf("\nGENERATION: %d\n", i)
    fmt.Printf("Mutated individuals: %d\t\tMutated Genes: %d\n", mutation_ind_count, mutation_count)
    fmt.Printf("Crossovers: %d\n", crossover_count)
    fmt.Printf("Best Individual: %s\n", best)
    fmt.Printf("Fitness: %d\n\n", score)

    // Fill graphic position vector
    graph_fitness = append(graph_fitness, score)

    // Check if the objective is reached by some individual of this generation
    if score == gene_number {
      fmt.Printf("\nObjective reached! Fitness = %d on Generation: %d\n", gene_number, i)
      break
    }

  }

  // Generate Graphics
  pixelgl.Run(graphics)

}
