package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "sort"
    "strings"
)

// modified from https://stackoverflow.com/questions/37562873/most-idiomatic-way-to-select-elements-from-an-array-in-golang
func filter(ss [][]string, test func([]string) bool) (ret [][]string) {
    for _, s := range ss {
        if test(s) {
            ret = append(ret, s)
        }
    }
    return
}

// from https://stackoverflow.com/questions/24999079/reading-csv-file-in-go
func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

// from https://gobyexample.com/sorting-by-functions
type byBestSynonym [][]string

func (s byBestSynonym) Len() int {
    return len(s)
}

func (s byBestSynonym) Swap(i, j int) {
    s[i][1], s[j][1] = s[j][1], s[i][1]
}

// sort by highest occurance of C and G
func (s byBestSynonym) Less(i, j int) bool {
    return strings.Count(s[i][1], "C") + strings.Count(s[i][1], "G") > strings.Count(s[j][1], "C") + strings.Count(s[j][1], "G")
}

func bestCodon(amino string, codons [][]string) string {
  synonym := func(s []string) bool { return s[0] == amino }
  synonyms := filter(codons, synonym)
  sort.Sort(byBestSynonym(synonyms))
  return synonyms[0][1]
}


func main() {
	// Read the codon/amino acid table
	codons := readCsvFile("codon-table-grouped.csv")[1:]

	c2s:= make(map[string]string)
	for _, element := range codons {
		c2s[element[1]]=element[0]
	}

	// read the codons
	virvac := readCsvFile("side-by-side.csv")[1:]

	matches := 0.0
	for _, element := range virvac {
		vir:=element[1]
		vac:=element[2]
		var our string
		fmt.Printf("%s v %s, amino: %s == %s. ",
			vir, vac,
			c2s[vir], c2s[vac])

		our = bestCodon(c2s[vir], codons)

		fmt.Printf(" ")
		if(vac == our) {
			fmt.Printf("Matched the vaccine!\n")
			matches++
		} else {
			fmt.Printf("No match.\n")
		}
	}
	fmt.Printf("%.1f%%\n", 100.0*matches/float64(len(virvac)))

}
