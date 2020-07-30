package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var (
	filename string
	template int
)

func init() {
	flag.StringVar(&filename, "file", "", "filename")
	flag.IntVar(&template, "template", 0, "timestamp")
	flag.Parse()
}

func main() {
	for filename == "" {
		fmt.Printf("请输入文件名: ")
		fmt.Scanf("%s", &filename)
		if filename == "" {
			fmt.Println("\n文件名为空，请重新选择")
			continue
		}
		fmt.Println()
	}

	for template < 1 || template > 4 {
		fmt.Println("1. Cleveland")
		fmt.Println("2. Central")
		fmt.Println("3. edgecombe")
		fmt.Println("4. Mayland")
		fmt.Printf("请输入解析模板编号: ")
		fmt.Scanf("%d", &template)
		if template < 1 || template > 4 {
			fmt.Println("\n选择错误，请重新选择")
			continue
		}
		fmt.Println()
	}

	gen(filename, template)

}

func outputFilename(filename string) string {
	index := strings.LastIndex(filename, ".")
	if index == -1 {
		return filename + "_out.csv"
	}
	return filename[:index] + "_out.csv"
}

func gen(filename string, template int) {
	fmt.Printf("正在使用模板[%d]转换[%s]\n", template, filename)
	ifile, err := os.Open(filename)
	if err != nil {
		log.Printf("open file %s %v\n", filename, err)
		os.Exit(1)
	}
	defer ifile.Close()
	output := outputFilename(filename)
	ofile, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Printf("open file %s %v\n", output, err)
		os.Exit(1)
	}
	defer ofile.Close()

	br := bufio.NewReader(ifile)
	for {
		data, _, err := br.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("read line %v\n", err)
			break
		}

		line := strings.TrimSpace(string(data))

		result := ""
		switch template {
		case 1:
			result = genCleveland(line)
		case 2:
			result = genCentral(line)
		case 3:
			result = genEdgecombe(line)
		case 4:
			result = genMayland(line)
		}

		if result == "" {
			continue
		}

		_, err = ofile.WriteString(result + "\r\n")
		if err != nil {
			log.Printf("write err %v", err)
			break
		}
	}

}

func genCleveland(line string) string {
	if line == "" {
		return ""
	}

	result := ""
	if len(line) < 3 {
		result = line
	} else {
		result = line[len(line)-3:]
	}

	return fmt.Sprintf("%s,%s@my.clevelandcc.edu,1999CCC#%s", line, line, result)
}

func genCentral(line string) string {
	items := strings.SplitN(line, ",", 2)
	if len(items) != 2 {
		return ""
	}

	names := strings.SplitN(items[0], " ", 2)
	if len(names) != 2 {
		return ""
	}

	result := ""
	if len(names[0]) != 0 {
		result += names[0][0:1]
	}

	if len(names[1]) >= 4 {
		result += names[1][0:4]
	} else {
		result += names[1]
	}

	if len(items[1]) >= 3 {
		result += items[1][len(items[1])-3:]
	} else {
		result += items[1]
	}

	return fmt.Sprintf("%s,%s,%s@cougarmail.cccc.edu", line, result, result)
}

func genEdgecombe(line string) string {
	items := strings.SplitN(line, ",", 2)
	if len(items) != 2 {
		return ""
	}

	ssn := strings.Replace(items[1], "-", "", -1)
	result := ssn
	if len(result) > 6 {
		result = result[len(result)-6:]
	}
	return fmt.Sprintf("%s,%s@myedge.cc,%s", line, items[0], result)
}

func genMayland(line string) string {
	items := strings.SplitN(line, ",", 3)
	if len(items) != 3 {
		return ""
	}
	user := items[0]
	name := items[1]
	day, month, year, err := parseBirthday(items[2])
	if err != nil {
		return ""
	}

	result := ""
	names := strings.SplitN(name, " ", 2)
	if len(names) != 2 {
		return ""
	}

	if len(names[1]) > 2 {
		result += names[1][0:2]
	} else {
		result += names[1]
	}

	result += day + month + year

	return fmt.Sprintf("%s,%s@students.mayland.edu,%s", line, user, result)
}

func parseBirthday(str string) (day, month, year string, err error) {
	items := strings.SplitN(str, "/", 3)
	if len(items) != 3 {
		err = fmt.Errorf("items no enouth")
		return
	}
	day = items[0]
	month = items[1]
	year = items[2]
	if len(year) > 2 {
		year = year[len(year)-2:]
	}
	err = nil
	return
}