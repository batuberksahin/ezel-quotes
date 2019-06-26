package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
	"bufio"
	"log"
	"strings"
	"regexp"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	valid "github.com/asaskevich/govalidator"
)

type Ezel struct {
	Quote []string `json:"quotes"`
}

func main() {
	for i := 1; i < 34; i++ {

		file, _ := os.Open(fmt.Sprintf("srt/Ezel %d.Bölüm HD.srt", i))
		defer file.Close()

		decoder := transform.NewReader(file, charmap.Windows1254.NewDecoder())

		scanner := bufio.NewScanner(decoder)

		var Quotes Ezel

		var text string
		var temp string
		var tempFlag bool
		var treeDot bool

		for scanner.Scan() {
			if len(string(scanner.Text())) > 3 {
				if !valid.IsInt(string(scanner.Text()[0])) && string(scanner.Text()[0]) != "[" {
					
					if !tempFlag {
						if string(scanner.Text())[len(string(scanner.Text()))-1:] == "!" || string(scanner.Text())[len(string(scanner.Text()))-1:] == "?" {
							text = scanner.Text()

							text = strings.Replace(text, "<i>", "", -1)
							text = strings.Replace(text, "</i>", "", -1)
							text = strings.Replace(text, "<- ", "", -1)
							text = strings.Replace(text, ">- ", "", -1)
							text = strings.Replace(text, "- ", "", -1)
							text = strings.Replace(text, "\"", "", -1)

							if WordCount(text) > 4 {
								Quotes.Quote = append(Quotes.Quote, text)
							}
						}else{
							if string(scanner.Text())[len(string(scanner.Text()))-3:] == "..." {
								text = scanner.Text()

								text = strings.Replace(text, "<i>", "", -1)
								text = strings.Replace(text, "</i>", "", -1)
								text = strings.Replace(text, "<- ", "", -1)
								text = strings.Replace(text, ">- ", "", -1)
								text = strings.Replace(text, "- ", "", -1)
								text = strings.Replace(text, "\"", "", -1)

								temp = text
								tempFlag = true
								treeDot = true
							}else if string(scanner.Text())[len(string(scanner.Text()))-1:] == "." {
								text = scanner.Text()

								text = strings.Replace(text, "<i>", "", -1)
								text = strings.Replace(text, "</i>", "", -1)
								text = strings.Replace(text, "<- ", "", -1)
								text = strings.Replace(text, ">- ", "", -1)
								text = strings.Replace(text, "- ", "", -1)
								text = strings.Replace(text, "\"", "", -1)
		
								if WordCount(text) > 4 {
									Quotes.Quote = append(Quotes.Quote, text)
								}
							}
						}
					}else{
						text = scanner.Text()

						text = strings.Replace(text, "<i>", "", -1)
						text = strings.Replace(text, "</i>", "", -1)
						text = strings.Replace(text, "<- ", "", -1)
						text = strings.Replace(text, ">- ", "", -1)
						text = strings.Replace(text, "- ", "", -1)
						text = strings.Replace(text, "\"", "", -1)

						if WordCount(text) > 2 {
							if treeDot {
								Quotes.Quote = append(Quotes.Quote, fmt.Sprintf("%s %s", temp[0:len(temp)-3], text[3:]))
							}else{
								Quotes.Quote = append(Quotes.Quote, fmt.Sprintf("%s %s", temp, text))
							}
						}

						treeDot = false
						tempFlag = false
					}

				}
			}
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		
		output, _ := json.MarshalIndent(Quotes, "", " ")
		_ = ioutil.WriteFile(fmt.Sprintf("Ezel-%d.json", i), output, 0644)

		fmt.Println("Yazdırma işlemi sona erdi.. - ", i)

	}
}

func WordCount(value string) int {
    re := regexp.MustCompile(`[\S]+`)

    results := re.FindAllString(value, -1)
    return len(results)
}