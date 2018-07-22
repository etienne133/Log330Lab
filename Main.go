package main

import(
	"os"
	"encoding/csv"
	"strconv"
	"math"
	"fmt"
	"strings"
	"errors"
)

type student struct{
	nom string
	effort [6]float64
	note float64
}

func main(){

	var input string


	fmt.Println("Press '1' to chose Tp1")
	fmt.Println("Press '2' to chose Tp2")
	fmt.Println("Press '3' to chose Tp3")
	fmt.Println("Press '4' to chose Tp3")
	fmt.Println("Press '5' to chose Tp3")


	fmt.Scanln(&input)
	input = strings.TrimRight(input, "\n")

	if input == "1"{
		tp1()
	}else if input == "2"{
		tp2()
	}else if input == "3"{
		tp3()
	}else if input == "4"{
		tp4()
	}else if input == "5"{
		tp5()
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

func tp3(){
	file,_ := os.Open(os.Args[1])
	reader := csv.NewReader(file)
	lines,_ := reader.ReadAll()

	reader.Comma = ','
	var values [][]float64
	for i := range lines {
		x,_ := strconv.ParseFloat(lines[i][0],64)
		y,_ := strconv.ParseFloat(lines[i][1],64)

		nextValue := [][]float64{{x, y}}
		values = append(values,nextValue...)
	}

	var averageX float64
	var averageY float64

	for i := 0; i < len(values); i++ {
		averageX += values[i][0]
		averageY += values[i][1]
	}
	n := float64(len(values))
	averageX = averageX /n
	averageY = averageY /n

	b1,_ :=  GenerateB1(values, averageX, averageY)
	b0,_ :=  GenerateB0(b1, averageX, averageY)

	fmt.Println("B1 :", b1)
	fmt.Println("B0 :", b0)

	fmt.Println("Enter a float value : ")

	var value float64
	fmt.Scanf("%f", &value)
	//y=b0+b1x

	var input string

	fmt.Println("Press x or y to generate in term of x or y ")
	fmt.Scan(&input)
	input = strings.TrimRight(input, "\n")

	var response float64

	if input == "x"{
		response = CalculateInTermsOfX(b0,b1,value)
		fmt.Println("response in term of x : ", response)

	}else if input == "y"{
		response,_ = CalculateInTermsOfY(b0,b1,value)
		fmt.Println("response in term of y : ", response)
	}


}

func tp4(){
	file,_ := os.Open(os.Args[1])
	reader := csv.NewReader(file)
	lines,_ := reader.ReadAll()

	var students []student
	for i := range lines {

		var student student


		student.nom =  lines[i][0]
		student.effort[0],_ = strconv.ParseFloat(lines[i][1],64)
		student.effort[1],_ = strconv.ParseFloat(lines[i][2],64)
		student.effort[2],_ = strconv.ParseFloat(lines[i][3],64)
		student.effort[3],_ = strconv.ParseFloat(lines[i][4],64)
		student.effort[4],_ = strconv.ParseFloat(lines[i][5],64)
		student.effort[5],_ = strconv.ParseFloat(lines[i][6],64)
		//student.effort[6],_ = strconv.ParseFloat(lines[i][6],64)

		student.note,_ = strconv.ParseFloat(lines[i][7],64)

		students = append(students,student)
	}



	for i:= 0; i< 6 ; i++ {
		data,_:=StudentToTab(i,students)
		correlation := Correlation(data)
		fmt.Println(checkRDegree(correlation))

	}


}

func tp5(){
	file,_ := os.Open(os.Args[1])
	reader := csv.NewReader(file)
	lines,_ := reader.ReadAll()

	reader.Comma = ','
	var values [][]float64

	for i := range lines {

		x,_ := strconv.ParseFloat(lines[i][0],64)
		y,_ := strconv.ParseFloat(lines[i][1],64)

		nextValue := [][]float64{{x, y}}
		values = append(values,nextValue...)
	}

	var averageX float64
	var averageY float64

	for i := 0; i < len(values); i++ {
		averageX += values[i][0]
		averageY += values[i][1]
	}
	n := float64(len(values))
	averageX = averageX /n
	averageY = averageY /n

	b1,_ :=  GenerateB1(values, averageX, averageY)
	b0,_ :=  GenerateB0(b1, averageX, averageY)

	fmt.Println("B1 :", b1)
	fmt.Println("B0 :", b0)

	varianceWithRegression,_ := VarianceWithRegression(values, b0,b1)
	fmt.Println("Variance :", varianceWithRegression)

	stdVar := stdVariation(varianceWithRegression)
	fmt.Println("Stdvar :", stdVar)

	Inter70,_ := Intervalle(1119, stdVar,1.108,values,averageX)

	fmt.Println("Intervalle 70 :", Inter70)

	Inter90,_ := Intervalle(1119, stdVar,1.860,values,averageX)

	fmt.Println("Intervalle 90 :", Inter90)

	fmt.Printf("Interval(Alpha=0.9) : [%.5f - %.5f]\n Interval(Alpha=0.7) : [%.5f - %.5f]",
		CalculateInTermsOfX(b0,b1,1119)-Inter90,
			CalculateInTermsOfX(b0,b1,1119)+Inter90,
				CalculateInTermsOfX(b0,b1,1119)-Inter70,
					CalculateInTermsOfX(b0,b1,1119)+Inter70)



}

func Intervalle(xk float64, stddev float64, AlphaByTwo float64, values [][]float64, average float64)(float64, error){

	n := len(values)
	if n == 0 {
		return -1000, errors.New("Error division by 0")
	}

	denominator := 0.0

	for i := 0; i < n; i++ {
		denominator += PowerN(values[i][0]-average,2)
	}

	if denominator == 0 {
		return -1000, errors.New("Error division by 0")
	}


	return 	stddev*AlphaByTwo*math.Sqrt(1.0+(1.0/float64(n))+(PowerN(xk-average,2)/(denominator))),nil

}

func StudentToTab(week int, students []student)([][]float64,error){
	var values [][]float64

	if(week <0){
		return nil, errors.New("Invalid week value")
	}

	for i:=0;i<len(students);i++{
		x:= students[i].effort[week]
		y:= students[i].note

		nextValue := [][]float64{{x, y}}
		values = append(values,nextValue...)

	}
	return values,nil
}

func GenerateB1(values [][]float64, averageX float64, averageY float64)(float64,error){

	var numerator = 0.0
	var denominator float64 =0.0


	fmt.Println("average x", averageX)
	fmt.Println("average y", averageY)


	for i := 0; i < len(values); i++ {
		numerator += (values[i][0]-averageX) *(values[i][1]-averageY)
		denominator +=(values[i][0]-averageX) *(values[i][0]-averageX)
	}

	if denominator ==0 {
		return -1000, errors.New("Error division by 0")
	}

	var b1 = numerator/ denominator

	return b1,nil
}

func GenerateB0(b1 float64, averageX float64, averageY float64)(float64,error){

	return averageY -b1*averageX,nil
}

func CalculateInTermsOfX(b0 float64, b1 float64, x float64)float64{
	return  b0+b1*x
}

func CalculateInTermsOfY(b0 float64, b1 float64, y float64)(float64,error){
	if b1 ==0 {
		return -1000, errors.New("Error division by 0")
	}
	return (y-b0)/b1,nil
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

func checkRDegree(r float64) (string,error) {
	absR := math.Abs(r)
	switch {
	case absR<0.2:
		return "Nulle à faible",nil
	case absR<0.4:
		return"Faible à moyenne.",nil
	case absR<0.6:
		return"Moyenne à forte",nil
	case absR<0.8:
		return"Forte à très forte",nil
	case absR<=1:
		return"Très forte à parfaite",nil
	default:
		return"Valeur invalide",errors.New("erreur r invalide")
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

func VarianceWithRegression(values [][]float64, b0 float64, b1 float64) (float64, error){
	var distance float64

	if len(values) == 1 {
		return -1000, errors.New("Error division by 0")

	}

	for i := 0; i < len(values); i++ {
		distance += PowerN(values[i][1]-b0-b1*values[i][0],2)
	}
	return distance/(float64(len(values))-1.0),nil
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

