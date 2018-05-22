package main

import(
	"os"
	"encoding/csv"
	"strconv"
	"math"
	"fmt"
)

func main(){
	var vals []int
	file,_ := os.Open(os.Args[1]);
	reader := csv.NewReader(file);
	lines,_ := reader.ReadAll();
	for i := range lines {
		value,_ := strconv.Atoi(lines[i][0]);
		vals = append(vals, value);
	}

	mean := mean(vals);
	distance := distancePowerTwo(mean,vals);
	variance := variance(distance,vals);
	stdVariation := stdVariation(variance);

	fmt.Println(vals)
	fmt.Println("Moyenne : ", mean)
	fmt.Println("Distance :", distance)
	fmt.Println("Variance :", variance)
	fmt.Println("Écart type :", stdVariation)


}

func mean(values []int) float64{

	var mean float64;
	for _,value := range values {
		mean += float64(value);
	}
	mean = mean/float64(len(values))
	return mean;
}

func distancePowerTwo(mean float64, values []int) float64{

	var distance float64;
	for _,value := range values {
		distance += (float64(value)-mean)*(float64(value)-mean);
	}
	return distance;
}

func variance(distance float64, values []int) float64{

	var variance float64;

	variance = distance/(float64(len(values))-1.0); // erreur sur cette ligne. Elle a été oublier
	return variance;
}

func stdVariation(variance float64) float64{

	return math.Sqrt(variance);
}

