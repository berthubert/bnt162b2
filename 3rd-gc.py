import csv


def read_csv(filename):
    records = []
    with open(filename, 'rt') as fp:
        reader = csv.reader(fp, delimiter=',')
        for ii, row in enumerate(reader):
            if ii > 0:
                records.append(row)

    return records


codons = read_csv('codon-table-grouped.csv')
#print(codons)

c2s= {}
for c in codons:
    c2s[c[1]] = c[0]
print(c2s)


virvac = read_csv("side-by-side.csv")
#print(virvac)

matches = 0

for element in virvac:

    vir = element[1]
    vac = element[2]

    print(f'{vir} v {vac}, amino: {c2s[vir]} == {c2s[vac]}.')

    our = vir

    if vir[2] == 'G' or vir[2] == 'C':
        print('codon ended on G or C already, not doing anything')

    else:
        prop = vir[:2]+"G"
        print(f'Attempting G substution, new candidate {prop}')

        if c2s[vir] == c2s[prop]:
            print('amino acid still the same, done!')
            our = prop
        else:
            print(f'Oops, maino acid changed. Trying C, new candidate {prop}')
            prop = vir[:2] + "C"

            if c2s[vir] == c2s[prop]:
                print('Amino acid still the same, done!')
                our = prop

    if vac == our:
        print('Matched the vaccine!')
        matches +=1 
    else:
        print('No Match.')

print(100*matches / len(virvac))
