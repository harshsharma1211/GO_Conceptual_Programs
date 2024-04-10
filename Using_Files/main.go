package main

func main() {


	Content :=" This needs to be done at all cost"

	file,err := os.Create("E:\\myLcogofile.txt")

	if err !=nil{

		panic(err)
	}
     length,err : io.WriteString(file,Content)

	 if err !=nil{

		panic(err)
	}

	fmt.Println("length is :",length)

	defer file.Close()
	readFile("E:\\myLcogofile.txt")
	

}