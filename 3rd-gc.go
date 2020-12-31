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
		if(vir == vac) {
			matches++;
		} else {
			fmt.Printf("%s != %s, amino: %s == %s. ",
				vir, vac,
				c2s[vir], c2s[vac])

			vir = vir[:2]+"G"
			fmt.Printf("Attempting G substitution, new candidate '%s'. ", vir)
			if(c2s[vir] == c2s[vac]) {
				fmt.Printf("Protein still the same. ")
				if(vir == vac) {
					fmt.Printf("Match!")
					matches++
				}
				
			} else {
				fmt.Printf("Oops, amino acid changed. Trying C.")
				vir = vir[:2]+"C"
				if(c2s[vir] == c2s[vac]) {
					fmt.Printf("Protein still the same. ")
					if(vir == vac) {
						fmt.Printf("Match!")
						matches++
					}
					
				}
			}
			fmt.Printf("\n")
		}
	}
	fmt.Printf("%.1f%%\n", 100.0*matches/float64(len(virvac)))

}
