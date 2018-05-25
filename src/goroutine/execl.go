package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	file, err := os.Create("/Users/zhangyalei/test.xls")
	if err != nil {
		fmt.Println(err)
		log.Fatal("The test.xml is not cunzai")
	}

	defer func(file *os.File) {
		file.Close()
	}(file)

	file.WriteString("\xEF\xBB\xBF")

	w := csv.NewWriter(file)
	w.Write([]string{"编号", "姓名", "年龄"})
	records := [][]string{}

	for i := 1; i < 21; i++ {
		num := strconv.FormatInt(int64(i), 10)
		age := strconv.FormatInt(int64(rand.Intn(100)), 10)
		str := []string{num, "name" + num, age}
		records = append(records, str)
	}
	w.Flush()
	w.WriteAll(records)
}
