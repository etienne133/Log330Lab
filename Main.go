package main

import(
	"os"
	"encoding/csv"
	"strconv"
	"math"
	"fmt"
	"strings"
)

func main(){

	var input string


	fmt.Println("Press '1' to chose Tp1")
	fmt.Println("Press '2' to chose Tp2")

	fmt.Scanln(&input)
	input = strings.TrimRight(input, "\n")

	if input == "1"{
		tp1()
	}else if input == "2"{
		tp2()
	}




}
func tp1(){
	    var vals []int
		file,_ := os.Open(os.Args[1])
		reader := csv.NewReader(file)
		lines,_ := reader.ReadAll()
		for i := range lines {
			value,_ := strconv.Atoi(lines[i][0])
			vals = append(vals, value)
		}

		mean := mean(vals)
		distance := distancePowerTwo(mean,vals)
		variance := variance(distance,vals)
		stdVariation := stdVariation(variance)

		fmt.Println(vals)
		fmt.Println("Moyenne : ", mean)
		fmt.Println("Distance :", distance)
		fmt.Println("Variance :", variance)
		fmt.Println("Écart type :", stdVariation)


}

func tp2(){
	file,_ := os.Open(os.Args[1])
	reader := csv.NewReader(file)
	lines,_ := reader.ReadAll()
	//bufio.NewReader(os.Stdin).ReadBytes('1')

	reader.Comma = ','
	var values [][]float64
	for i := range lines {
		x,_ := strconv.ParseFloat(lines[i][0],64)
		y,_ := strconv.ParseFloat(lines[i][1],64)

		nextValue := [][]float64{{x, y}}
		values = append(values,nextValue...)
	}

	r :=  Correlation(values)
	fmt.Println("R :", r)
	fmt.Println(checkRDegree(r))
}

func PowerN(value float64, n int)float64{
	if(n == 0){
		return 1;
	}
	var result =value
	for i := 1; i < n; i++ {
		result *= value
	}
	return result
}

func Correlation(values [][]float64)float64{
	var sumX float64
	var sumXPower2 float64
	var sumYPower2 float64
	var sumY float64
	var sumXY float64
	n := float64(len(values))

	for i := 0; i < len(values); i++ {
		sumX += values[i][0]
		sumXPower2 += PowerN(values[i][0],2)
		sumY += values[i][1]
		sumYPower2 += PowerN(values[i][1],2)
		sumXY += values[i][0]*values[i][1]

	}

	numerator := sumXY-((sumX*sumY)/n)
	denominator := math.Sqrt((sumXPower2-(sumX*sumX)/n)*(sumYPower2-(sumY*sumY)/n))

	if denominator == 0{
		return 1000 // valeur invalide
	}
	return numerator/denominator
}

func checkRDegree(r float64) string {
	absR := math.Abs(r)
	switch {
	case absR<0.2:
		return "Nulle à faible"
	case absR<0.4:
		return"Faible à moyenne."
	case absR<0.6:
		return"Moyenne à forte"
	case absR<0.8:
		return"Forte à très forte"
	case absR<=1:
		return"Très forte à parfaite"
	default:
		return"Valeur invalide"
	}
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

	fmt.Println(distance)
	fmt.Println(float64(len(values))-1.0)
	variance = distance/(float64(len(values))-1.0); // erreur sur cette ligne. Elle a été oublier
	return variance;
}

func stdVariation(variance float64) float64{

	return math.Sqrt(variance);
}

