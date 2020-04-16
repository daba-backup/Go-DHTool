package go_dhtool_file

import(
	"bufio"
	"os"
)

func GetFileAllLines(filename string) ([]string,error){
	file,err:=os.Open(filename)
	if err!=nil{
		return nil,err
	}
	defer file.Close()

	lines:=make([]string,0)

	scanner:=bufio.NewScanner(file)
	for scanner.Scan(){
		line:=scanner.Text()
		lines=append(lines,line)
	}
	if err:=scanner.Err();err!=nil{
		return nil,err
	}

	return lines,nil
}
