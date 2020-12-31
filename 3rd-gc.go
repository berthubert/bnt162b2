package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
)


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
		var prop string
		fmt.Printf("%s v %s, amino: %s == %s. ",
			vir, vac,
			c2s[vir], c2s[vac])

		// base case, don't do anything
		our = vir
		
		// don't do anything if codon ends on G or C already
		if(vir[2] == 'G' || vir[2] =='C') {
			fmt.Printf("Codon ended on G or C already, not doing anything.")
		} else {
			prop = vir[:2]+"G"
			fmt.Printf("Attempting G substitution, new candidate '%s'. ", prop)
			if(c2s[vir] == c2s[prop]) {
				fmt.Printf("Amino acid still the same, done!")
				our = prop
			} else {
				fmt.Printf("Oops, amino acid changed. Trying C, new candidate '%s'. ", prop)
				prop = vir[:2]+"C"
				if(c2s[vir] == c2s[prop]) {
					fmt.Printf("Amino acid still the same, done!")
					our=prop
				} 
				
			}
		
		}
		
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
