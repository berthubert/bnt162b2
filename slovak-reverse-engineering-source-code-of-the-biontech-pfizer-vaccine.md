---
title: "Ako to funguje: zdrojový kód BioNTech/Pfizer vakcíny SARS-CoV-2"
date: 2020-12-25T20:12:20+01:00
draft: false
images:
 - bnt162b2.png
---
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:site" content="@powerdns_bert">
<meta name="twitter:creator" content="@powerdns_bert">
<meta name="twitter:title" content="Ako to funguje: zdrojový kód BioNTech/Pfizer vakcíny SARS-CoV-2">
<meta name="twitter:description" content="Vitajte! V tomto článku sa pozrieme znak po znaku na zdrojový kód
BioNTech/Pfizer mRNA vakcíny SARS-CoV-2.">
<meta name="twitter:image" content="https://berthub.eu/articles/bnt162b2.png">

**Translations**: [ελληνικά](https://berthub.eu/articles/posts/greek-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/)
/ [中文](https://mp.weixin.qq.com/s/b0Mw8uKLYuXHJ5Bj3t2Dwg)
/ [Deutsch](https://berthub.eu/articles/posts/german-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/)
/ [Español](https://docs.google.com/document/d/1FxswEeem2kP1AUF979P7INBlkD1OzMzVHPEKeU9R2kw/edit)
/ [Français](https://renaudguerin.net/posts/explorons-le-code-source-du-vaccin-biontech-pfizer-sars-cov-2/)
/ [עִברִית](https://github.com/chilik/Hebrew-ReversingSARS-CoV-2mRNAVaccine/blob/main/%D7%94%D7%A0%D7%93%D7%A1%D7%94%20%D7%9C%D7%90%D7%97%D7%95%D7%A8%20%D7%A9%D7%9C%20%D7%A7%D7%95%D7%93%20%D7%94%D7%9E%D7%A7%D7%95%D7%A8%20%D7%A9%D7%9C%20%D7%94%D7%97%D7%99%D7%A1%D7%95%D7%9F%20BioNTech%20-%20Pfizer%20SARS-CoV-2.pdf)
/ [Hrvatski](https://docs.google.com/document/d/1BODRitAvGuDYGZCHU5LY-AkNhs9_1cVDubdRvz-cSPY/edit)
/ [Italiano](https://berthub.eu/articles/posts/italian-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/)
/ [नेपाली](https://onedrive.live.com/view.aspx?resid=9C571BA15BC4287D!15298&ithint=file%2cdocx&authkey=!ALATa2b8xetI7lQ)
/ [Polskie](https://randomseed.pl/rna/reverse-engineering-kodu-zrodlowego-szczepionki-biontech-pfizer-covid-sars-cov-2/)
/ [русский](https://localcrew.ru/reversepfizer)
/ [Português](https://docs.google.com/document/d/1pDo40DXcpXjzqAUfhFfup50-IQ2Qct-mhLnmRpjFZWM/edit).
/ [Slovensky](https://berthub.eu/articles/posts/slovak-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/).
/ [Markdown for translating](https://raw.githubusercontent.com/berthubert/bnt162b2/master/reverse-engineering-source-code-of-the-biontech-pfizer-vaccine.md)

Vitajte! V tomto článku sa pozrieme znak po znaku na zdrojový kód
BioNTech/Pfizer mRNA vakcíny SARS-CoV-2.

> *Chcem sa poďakovať veľkej skupine ľudí, ktorí si prečítali tento článok
> pred jeho publikáciou kvôli čitateľnosti a správnosti. Všetky chyby sú
> však moje, a rád by som o nich počul čo najskôr na
> bert@hubertnet.nl alebo [@PowerDNS_Bert](https://twitter.com/PowerDNS_Bert)*
>
> Chyby súvisiace so slovenským prekladom prosím hláste na ewalzel@gmail.com
> alebo [@eronisko](https://twitter.com/eronisko)

Tieto slová môžu znieť prekvapivo - vakcína je tekutina, ktorá sa
injekciou pichá do ruky.  Ako sa teda môžeme baviť o zdrojovom kóde?

Je to dobrá otázka.  Začnime preto malou časťou BioNTech/Pfizer vakcíny
[BNT162b2](https://en.wikipedia.org/wiki/Tozinameran), tiež známej ako Tozinameran
[tiež známej ako Comirnaty](https://twitter.com/PowerDNS_Bert/status/1342109138965422083).

<center>
{{< figure src="/articles/bnt162b2.png" caption="Prvých 500 znakov mRNA BNT162b2. Zdroj: [World Health Organization](https://mednet-communities.net/inn/db/media/docs/11889.doc)">}}
</center>

mRNA vakcína BNT162b má vo svojom srdci tento digitálny kód.  Má 4284 znakov, takže by sa
zmestil do zopár tweetov.  Na začiatku výroby vakcíny niekto nahral tento kód do DNA
tlačiarne (naozaj), ktorá potom premenila bajty z disku na molekuly DNA.

<center>
{{< figure src="/articles/bioxp-3200.jpg" caption=" DNA tlačiareň BioXp 3200 od firmy [Codex](https://codexdna.com/products/bioxp-system/)" >}}
</center>

Z takéhoto zariadenia vychádza v malých množstvách DNA, ktorá sa po
biologickom a chemickom spracovaní zmení na RNA (o tom trocha neskôr) v ampulke
vakcíny.  30mg dávka obsahuje práve tridsať miligramov RNA. Okrem toho sa tam
nachádza aj premyslený lipidový (tukový) baliaci systém, ktorý dostane mRNA
do našich buniek.

RNA je prechodná, "pracovná" verzia DNA. DNA je ako flash pamäť v biologickom
svete.  DNA je trvácna, vnútorne redundantná a veľmi spoľahlivá.  Ale podobne
ako počítače nespúšťajú zdrojový kód priamo z disku, predtým, než sa niečo
udeje, kód sa musí skopírovať na rýchlejší, všestrannejší a zároveň krehkejší
systém.

Pre počítače je týmto systémom RAM, pre biológiu je to RNA.  Tieto systémy
sú si až prekvapujúco podobné.  Narozdiel od flash disku, RAM degraduje
omnoho rýchlejšie, pokiaľ sa o ňu nestaráme.  Dôvod, prečo musí byť Pfizer/BioNTech
mRNA vakcína uložená v najmrazivejších mrazničkách je podobný: RNA je krehká.

Každý znak RNA váži rádovo 0.53&middot;10⁻²¹ gramov, čo znamená, že v každej
tridsaťmiligramovej dávke vakcíny sa nachádza 6&middot;10¹⁶ znakov.  Vyjadrené
v bajtoch, je to asi 25 petabajtov, aj ked je treba povedať, že tieto
pozostávajú z asi 2000 miliárd opakovaní tých istých 4284 znakov.  Skutočný
informačný obsah je o trocha viac než jeden kilobajt.
[Samotný SARS-CoV-2](https://www.ncbi.nlm.nih.gov/projects/sviewer/?id=NC_045512&tracks=[key:sequence_track,name:Sequence,display_name:Sequence,id:STD649220238,annots:Sequence,ShowLabel:false,ColorGaps:false,shown:true,order:1][key:gene_model_track,name:Genes,display_name:Genes,id:STD3194982005,annots:Unnamed,Options:ShowAllButGenes,CDSProductFeats:true,NtRuler:true,AaRuler:true,HighlightMode:2,ShowLabel:true,shown:true,order:9]&v=1:29903&c=null&select=null&slim=0) má asi 7,5 kilobajtov.

Naozaj krátky úvod
------------------------------
DNA je digitálny kód. Narozdiel od počítačov, ktore používajú 0 a 1, život
používa A, C, G a U/T (známe ako "nukleotidy", alebo "bázy").

V počítačoch ukladáme 0 a 1 ako prítomnosť alebo neprítomnosť elektrického náboja,
ako prúd, magnetickými prechodmi, ako napätie, ako moduláciu signálu, alebo ako
zmenu v odrazivosti.  Skrátka, nuly a jednotky nie sú akýmsi abstraktným
konceptom, ale existujú ako elektróny, či iné fyzické objekty.

V prírode sú A, C, G a U/T molekulami, ktoré sú uložené ako reťazce v DNA (alebo RNA).

V počítačoch 8 bitov spájame do bajtu a bajt je typickou jednotkou dát,
s ktorými potom pracujeme.

V prírode sa 3 nukleotidy spájajú do kodónu a kodón je typickou dátovou jednotkou.
Jeden kodón obsahuje 6 bitov informácií (2 bity na každý znak DNA, 3 znaky = 6 bitov.
To znamená 2⁶ = 64 rôznych hodnôt).

Zatiaľ celkom digitálne. Ak máte pochybnosti pozrite si digitálny kód priamo
[v dokumente od WHO](https://mednet-communities.net/inn/db/media/docs/11889.doc).

> *Ďalší materiál k tejto téme sa [nachádza aj
> tu](https://berthub.eu/articles/posts/what-is-life/) - tento článok ("What is life" -
> v angličtine - pozn. prekl.) vám  môže pomôcť lepšie pochopiť zvyšok tohto článku.
> Prípadne ak máte radi video, mám tu [pre vás dve hodiny](https://berthub.eu/dna).*

Takže čo vlastne ten kód robí?!
--------------------------
Základná myšlienka vakcíny je naučiť náš imunitný systém bojovať s patogénom bez
toho, aby sme sa ním nakazili.  V minulosti sa za týmto účelom vpichol oslabený
alebo utlmený (atenuovaný) vírus spolu s "adjuvans", ktoré pomáhajú vystrašiť
a tým naštartovať imunitný systém.  Toto bola rozhodne analógová technika, ktorá
si vyžadovala prípravu miliárd vajíčok (alebo hmyzu).  Tiež bola náročná na šťastie
a čas. Niekedy sa dokonca musel využiť iný (nesúvisiaci) vírus.

mRNA vakcína dokáže dosiahnuť rovnaký cieľ ("vytrénovať náš imunitný systém"), no
"laserovým" spôsobom.  A to myslím v oboch zmysloch slova: veľmi precízne a s
veľkou silou.

Funguje to takto. Injekcia obsahuje nestály genetický materiál, ktorý opisuje
slávny SARS-CoV-2 'Spike' proteín. Pomocou dômyselných chemických protriedkov
dokáže vakcína dopraviť tento gentický materiál do niektorých našich buniek.

Tieto bunky začnú pracovito vyrábať SARS-CoV-2 Spike proteíny v množstvách
dosť veľkých na to, aby naštartovali náš imunitný systém do akcie.  Keď
náš imunitný systém narazí na Spike proteíny a (hlavne) na signály, že
niektoré bunky boli napadnuté, vyvinie ráznu obrannú reakciu voči niekoľkým
aspektom Spike proteínov A aj ich výrobnému procesu.

A toto je to, vďaka čomu máme vakcínu s 95% účinnosťou.

Zdrojový kód!
----------------
[Začnime úplne od začiatku, napríklad tu](https://youtu.be/jp0opnxQ4rY?t=8).
Dokument od WHO obsahuje tento užitočný obrázok:

<center>
{{< figure src="/articles/vaccine-toc.png"  >}}
</center>

Toto je niečo ako obsah. Začneme "čiapkou" ("cap") ktorá je znázornená ako malá
šiltovka.

Tak, ako sú dôležité opcodes v počítačových súboroch, aj biologický
operačný systém potrebuje hlavičky, má linkery a veci ako volacie konvencie.

Kód vakcíny začína týmito dvoma nukleotidmi:

```
GA
```

Toto sa dá veľmi dobre prirovnať ku každému spustiteľnému DOS a Windows súboru,
[ktorý začína na MZ](https://en.wikipedia.org/wiki/DOS_MZ_executable), alebo
k UNIX skriptom, ktoré začínajú na
[`#!`](https://en.wikipedia.org/wiki/Shebang_(Unix).
V biológií ani v operačných systémoch sa tieto dva znaky nijak nespúšťajú. No je
ich treba, lebo bez nich by sa nič nezačalo diať.

mRNA "čiapka" [má mnoho funkcií](https://en.wikipedia.org/wiki/Five-prime_cap#Function). Napríklad označuje, že kód pochádza z bunkového
jadra. V našom prípade tomu tak samozrejme nie je, náš kód pochádza z vakcíny.
Ale toto bunke nepovieme. Čiapka napomáha nášmu kódu vyzerať bezpečne
a tým ho chráni pred predčasným zničením.

Prvé dva `GA` nukleotidy sú tiež chemicky mierne odlišné od zvyšku RNA.  Dá sa povedať,
že `GA` nesie mimopásmový (out-of-band) signál.

"5' neprekladaná oblasť"
------------------------------------
Trocha žargónu. Molekuly RNA sa dajú čítať len v jednom smere. Časť, kde
sa čítanie začína sa nazýva 5' alebo "five-prime". Čítanie sa končí na
konci 3' alebo three-prime.

Život sa skladá z proteínov (alebo z vecí vyrobených z proteínov).
A tieto proteíny sú opísané v RNA. Keď sa RNA premieňa na proteíny, hovoríme
tomu preklad, alebo translácia.

Tu máme 5' neprekladanú oblasť ("untranslated region" - "UTR"), teda časť kódu,
ktorá sa nebude nachádzať v proteíne.

```
GAAΨAAACΨAGΨAΨΨCΨΨCΨGGΨCCCCACAGACΨCAGAGAGAACCCGCCACC
```

A tu je prvé prekvapenie.  Bežné znaky RNA sú A, C, G a U.  U je tiež známe
ako "T" v DNA.  Ale tu máme Ψ, čo sa deje?

Toto je jedna z viacerých výnimočne dômyselných častí tejto vakcíny.
V našom tele beží veľmi účinný antivírusový systém ("prvý antivír").
Z tohto dôvodu sú bunky veľmi neochotné prijať cudziu RNA a snažia ju čím
skôr zničiť, než niečo spôsobí.

Toto je pre našu vakcínu trocha problematické - musí sa dostať cez náš
imunitný systém. Po mnohých rokoch experimentov sa zistilo, že keď sa U v
RNA vymení za mierne modifikovanú molekulu, náš imunitný systém o ňu
stratí záujem. Úplne vážne.

Takže vo vakcíne BioNTech/Pfizer sa každé U vymenilo za
1-metyl-3'-pseudouridylyl označované  Ψ.  Genialita tejto
zámeny je v tom, že aj keď Ψ upokojí náš imunitný systém, vo všetkých
dôležitých častiach bunky je prijaté ako normálne U.

V počítačovej bezpečnosti tiež poznáme tento trik - niekedy je možné
poslať mierne upravenú verziu správy ktorá zmätie firewally a bezpečnostné
riešenia, no je stále spracovaná backend servermi - a vtedy sme hacknutí.

Dnes využívame výsledky základného výskumu z minulosti.
[Objavitelia](https://twitter.com/PennMedicine/status/1341766354232365059)
techniky s Ψ museli bojovať o to, aby na [ich](https://www.statnews.com/2020/11/10/the-story-of-mrna-how-a-once-dismissed-idea-became-a-leading-technology-in-the-covid-vaccine-race/)
prácu mali dostatok prostriedkov a aby jej výsledky boli prijaté. Všetci by
sme mali byť za túto prácu vďační a som si istý, že [prídu aj Nobelove ceny](https://twitter.com/PowerDNS_Bert/status/1329861047168225281).

> Mnoho ľudí sa pýta, či by aj vírusy mohli použiť Ψ na to, aby obišli
> naše imunitné systémy. V skratke, toto je nanajvýš nepravdepodobné.
> Život jednoducho nemá mechanizmy na to, aby vyrobil nukleotidy
> 1-methyl-3'-pseudouridylylu. Vírusy sa pri svojej reprodukcii spoliehajú
> na štandardné biologické mechanizmy, ktoré túto schopnosť jednoducho nemajú.
> mRNA vakcíny v ľudskom tele zanikajú veľmi rýchlo a neexistuje možnosť, aby
> sa RNA s Ψ rozmnožovala naďalej. Dobré čítanie k téme: "[No, Really, mRNA Vaccines
> Are Not Going To Affect Your DNA](https://www.deplatformdisease.com/blog/no-really-mrna-vaccines-are-not-going-to-affect-your-dna)"

OK, späť k 5' UTR. Čo týchto 51 znakov vlastne robí. Ako to už v prírode býva,
takmer nič nemá len jednu konkrétnu fukciu.

Keď naše bunky potrebujú *preložiť* RNA na proteíny, tento proces prebieha
v prístroji, zvanom ribozóm.  Ribozóm je niečo ako 3D tlačiareň na proteíny.
Spracúva vlákno RNA a podľa neho vyrába reťazec aminokyselín, ktoré sa potom
zložia do proteínu.

<center>
<video controls width="90%" loop>
<source src="/articles/protein-short.mp4" type="video/mp4">
</video>
<br/>
Zdroj: [Wikipedia, používateľ Bensaccount](https://commons.wikimedia.org/wiki/File:Protein_translation.gif)
</center>


Toto vidíme hore.  Čierna páska na spodku je RNA. Páska, ktorá sa objaví
v zelenej časti je proteín, ktorý sa vytvára. Veci, ktoré prilietajú a
odlietajú sú aminokyseliny a adaptéry, vďaka ktorým pasujú do RNA.

Aby fungoval musí tento ribozóm musí fyzicky sedieť na vlákne RNA. Po tom,
čo sa usadí, môže vytvárať proteíny podľa RNA ktorú práve spracúva.
Z tohto sa môžeme dovtípiť, že časti na ktoré si sadne ako prvé, nemôže
spracovať hneď. Toto je len jedna z funkcií UTR: pristávacia plocha pre
ribozóm. UTR poskytuje navádzanie.

UTR tiež poskytuje metadáta: kedy má translácia začať? Koľko jej má byť?
Pre vakcínu použili "najokamžitejšiu" UTR akú našli, prevzatú z [génu
alfaglobínu](https://www.tandfonline.com/doi/full/10.1080/15476286.2018.1450054).
Tento gén je známy tým, že spoľahlivo vyrába veľké množstvo proteínu.
Vedcom sa v minulosti podarilo túto UTR ešte ďalej zoptimalizovať (podľa
dokumentu od WHO), takže toto nie je celkom UTR pre alfaglobín.
Je to ešte lepšie.

S-glykoproteínový signálny peptid
---------------------------------
Ako bolo spomenuté, cieľom vakcíny je prinútiť bunku, aby začala vyrábať
veľké množstvá Spike proteínu vírusu SARS-CoV-2. Doteraz sme sa stretli
prevažne s metadátami a "volacími konvenciami" v zdrojovom kóde vakcíny.
No teraz sa dostávame do samotného teritória vírusového proteínu.

Najprv sa však musíme dostať cez ďalšiu vrstvu metadát. Potom, čo ribozóm
(z vynikajúcej animácie hore) vytvorí proteín, tento proteín musí ešte niekam
ísť. Toto je zakódované v "S-glykoproteínovom signálovom peptide (rozšírenej
vedúcej sekvencii)"

Dá sa tomu chápať, že na začiatku proteínu je akýsi štítok s adresou, ktorý
je zakódovaný v proteíne. V tomto konkrétnom prípade signálny peptid vraví,
že proteín by mal z bunky vyjsť cez "endoplazmatické retikulum". Takéto parádne
výrazy nie sú ani v Star Treku!

"Signálny peptid" nie je veľmi dlhý, no keď sa pozrieme kód, nájdeme
medzi vírusom a RNA vakcíny rozdiely:

(Pre jednoduchšie porovnanie som vymenil zmenené Ψ za bežné RNA U)

```
           3   3   3   3   3   3   3   3   3   3   3   3   3   3   3   3
Vírus:   AUG UUU GUU UUU CUU GUU UUA UUG CCA CUA GUC UCU AGU CAG UGU GUU
Vakcína: AUG UUC GUG UUC CUG GUG CUG CUG CCU CUG GUG UCC AGC CAG UGU GUU
               !   !   !   !   ! ! ! !     !   !   !   !   !
```

Takže čo sa tu deje? Nespojil som RNA do trojíc náhodou. Tri znaky RNA tvoria
kodón. A každý kodón kóduje špecifickú aminokyselinu. Signálny peptid vo vakcíne
pozostáva z *presne* tých istých aminokyselín ako samotný vírus.

Ako je potom možné že RNA je iná?

Existuje 4³=64 rôznych kodónov, kedže existujú 4 rôzne RNA znaky a v kódone sú
po troch. Máme však len 20 rôznych aminokyselín. To znamená že niekoľko rôzných
kodónov kóduje tú istú aminokyselinu.

Život používa túto takmer univerzálnu tabuľku na mapovanie kodónov RNA na
aminokyseliny.

<center>
{{< figure src="/articles/rna-codon-table.png" caption="[Tabuľka kodónov RNA](https://en.wikipedia.org/wiki/DNA_and_RNA_codon_tables) (Wikipedia)" >}}
</center>

V tejto tabuľke vidíme, že modifikácie vo vakcíne (UUU -> UUC) sú všetky
*synonymické*. RNA kód vakcíny je iný, no vzniknú z neho rovnaké
aminokyseliny a proteíny.

Keď sa pozrieme bližšie, všimneme si, že väčšina zmien sa deje v tretej
pozicií v kodóne, označené "3" nad kodónom. A keď sa pozrieme do univerzálnej
tabuľky kodónov, uvidíme, že pre produkciu aminokyselín tretia pozícia v
kodóne nebýva rozhodujúca.

Takže tieto zmeny sú synonymické, ale načo tam potom sú? Keď sa pozrieme bližšie,
zistíme, že všetky zmeny *okrem jedinej* zvyšujú počet C a G.

Načo je to dobré? Ako sa spomína vyššie, náš imunitný systém nemá v obľube
"exogénnu" RNA, teda RNA nepochádzajúcu z bunky. Aby sa predišlo
takejto detekcii, už sme zamenili U v RNA za Ψ.

Ukázalo sa však, že RNA s [výšším počtom](https://www.nature.com/articles/nrd.2017.243)
G a C sa tiež [efektívnejšie premieňa na proteíny](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC1463026/).

A toto sa podarilo dosiahnuť zámenou čo najväčšieho počtu znakov za G alebo C.

> Celkom ma fascinuje, že *jedna* zmena neviedla k zvýšeniu počtu C alebo G,
> modifikácia CCA -> CCU. Ak niekto pozná dôvod prečo, dajte mi prosím vedieť!
> Rozumiem tomu, že niektoré kodóny sa v ľudskom tele objavujú častejšie, než
> iné, ale [tiež som čítal, že toto nevedie k veľkej zmene v rýchlosti prekladu](https://journals.plos.org/plosgenetics/article?id=10.1371/journal.pgen.1006024)

Samotný Spike proteín
------------------------
Ďalších 3777 znakov RNA vakcíny obsahuje podobné "kodónové optimalizácie"
na pridanie znakov C a G. Kvôli lepšej čitateľnosti ich sem neuvediem všetky,
ale rád by som sa zameral na jednu obzvlášť špeciálnu časť. Toto je časť,
vďaka ktorej to celé funguje; časť, ktorá nám pomôže vrátiť sa do normálnych
koľají.

```
                  *   *
          L   D   K   V   E   A   E   V   Q   I   D   R   L   I   T   G
Vírus:   CUU GAC AAA GUU GAG GCU GAA GUG CAA AUU GAU AGG UUG AUC ACA GGC
Vakcína: CUG GAC CCU CCU GAG GCC GAG GUG CAG AUC GAC AGA CUG AUC ACA GGC
          L   D   P   P   E   A   E   V   Q   I   D   R   L   I   T   G
           !     !!! !!        !   !       !   !   !   ! !
```

Vidíme tu známe synonymické zmeny v RNA. Napríklad, v prvom kodóne vidíme,
že CUU sa mení na CUG. Toto pridáva do vakcíny ďalšie "G" a my vieme, že
toto pomáha zlepšiť produkciu proteínov. CUU aj CUG oba kódujú aminokyselinu
"L", alebo leucín, takže v proteíne sa touto zmenou nič nemení.

Keď porovnáme celý Spike proteín vo vakcíne, všetky zmeny sú podobne
synonymické... okrem dvoch, na ktoré sa teraz pozeráme.

Tretí a štvrtý kodón vyššie predstavujú skutočné zmeny. Amynokyseliny K a V
sú obe vymenené za "P", prolín. Kvôli tejto zmene boli potrebné tri zmeny
("!!!") pre "K" a dve zmeny ("!!") pre "V".

**A tieto dve zmeny vystrelia účinnosť vakcíny do nových výšok**.

Ako je to možné? Ak sa pozriete na skutočnú časticu SARS-CoV-2, uvidíte Spike
proteíny ako skupinu výčnelkov:

<center>
{{< figure src="/articles/sars-em.jpg" caption="[Častice vírusu SARS](https://en.wikipedia.org/wiki/Severe_acute_respiratory_syndrome_coronavirus) (Wikipedia)" >}}
</center>

Výčnelky sú pripevnené k telu vírusu (k "nukleokapsidovému proteínu").
Ale pozor, naša vakcína vyrába iba výčnelky a tie nepripevňujeme k žiadnemu
telu vírusu.

Vo svojej nezmenenej forme sa samostatné Spike proteíny poskladajú do inej
štruktúry. Ak by sme dostali takúto vakcínu, naše telo by si vybudovalo
imunitu. Avšak, bola by to imunita voči iným, "poskladaným" spike proteínom.

A potom sa ukáže naozajstný SARS-CoV-2 s neposkladanými, pichľavými Spike
proteínmi. Vakcína by v tomto prípade nebola príliš účinná.

Čo s tým? V roku [2017 vedci popísali, ako dve zámeny za prolín na správnych
miestach](https://www.ncbi.nlm.nih.gov/pmc/articles/PMC5584442/) dokážu
SARS-CoV-1 a MERS S proteíny udržať v ich "neposkladanej" konfigurácii aj bez
toho, aby boli súčasťou celého vírusu. Funguje to vďaka tomu, že prolín je
veľmi stabilnou aminokyselinou. Slúži ako dlaha, vďaka ktorej môžeme proteín
zastabilizovať v stave, v ktorom ho potrebujeme ukázať imunitnému systému.

[Ľudia](https://twitter.com/goodwish916), ktorí [prišli](https://twitter.com/KizzyPhD)
na tento trik by si mali teraz neustále dávať high-fives a vyžarovať nafúkanosť.
[A bolo by to úplne v poriadku](https://twitter.com/McLellan_Lab/status/1291077489566142464).

> Novinka!  Kontaktovali ma ľudia z tímu [McLellan Lab](https://twitter.com/McLellan_Lab/status/1291077489566142464),
> jednej zo skupín, ktorá stojí za týmto prolínovým objavom. Dozvedel som
> sa, že high-fives sú nateraz pozastavené kvôli pandemickej situácií,
> no sú radi, že mohli prispieť k vývoju vakcín. Tiež chceli upozorniť na
> dôležitosť práce mnohých iných skupín, pracovníkov a dobrovoľníkov.

Koniec proteínu, ďalšie kroky
----------------------------------
Ak sa preskrolujeme cez zvyšok zdrojového kódu, môžeme si všimnúť drobných
zmien na konci Spike proteínu:

```
          V   L   K   G   V   K   L   H   Y   T   s
Vírus:   GUG CUC AAA GGA GUC AAA UUA CAU UAC ACA UAA
Vakcína: GUG CUG AAG GGC GUG AAA CUG CAC UAC ACA UGA UGA
          V   L   K   G   V   K   L   H   Y   T   s   s
               !   !   !   !     ! !   !          !
```

Na konci proteínu nájdeme "stop" kodón, označený malým "s". Toto je
spôsob ako slušne povedať, že proteín by mal končiť tu. Vírus používa
stop kodón UAA, vakcína zas, zrejme pre istotu, dva stop kodóny UGA.

3' neprekladaná oblasť
--------------------------
Podobne, ako keď ribozóm potreboval "naviesť" na 5'-konci, kde sme našli
"5' neprekladanú oblasť", takisto na konci oblasti kódujúcej proteín
nájdeme podobnú štruktúru, ktorú nazývame 3' UTR ("three prime untranslated region").

O 3' UTR sa dá napísať mnoho, no ja si pomôžem [citátom z Wikipedie](https://en.wikipedia.org/wiki/Three_prime_untranslated_region): "3' neprekladaná oblasť je dôležitá
pre génovú expresiu kvôli jej vplyvu na účinnosť lokalizácie, stability,
exportu a translácie mRNA .. **napriek ich nášmu súčasnému porozumeniu,
3' UTR sú pre nás stále relatívnou záhadou.**".

Čo vieme určite je, že niektoré 3' UTR sú veľmi úspešné pri napomáhaní
tvorby proteínov.  Podľa dokumentu od WHO, bola 3' UTR vo BioNTech/Pfizer
vakcíne vybraná z "mRNA amino-terminal enhancer of split (AES) a
mitochondriálnej kódovanej 12S ribozómovej RNA na dosiahnutie stability
RNA a vysokej expresie proteínu." K čomu len dodám: dobrá práca.

<center>
{{< figure src="/articles/vaccine.jpg" >}}
</center>


AAAAAAAAAAAAAAAAAAAAAA je koniec
----------------------------------------
Úplny koniec mRNA je polyadenylovaný. Čo je komplikovaný spôsob, ako
povedať, že končí kopou AAAAAAAAAAAAAAAAAAA. Zdá sa, že ešte aj mRNA
už mala roku 2020 dosť.

mRNA môže byť použitá viackrát, no pri každom použití stratí niekoľko A-čok
na konci. Keď sa A-čka minú, mRNA prestáva byť funkčná a zahodí sa. V tomto
zmysle slúži "poly-A" koniec ako ochrana pred degradáciou.

Niekoľko štúdií sa venovalo tomu, aký je optimálny počet A-čok na konci
mRNA vakcín. Vo voľne prístupnej literatúre som sa dozvedel, že toto číslo
môže byť až 120.

Vakcína BNT162b2 končí na:

```
                                     ****** ****
UAGCAAAAAA AAAAAAAAAA AAAAAAAAAA AAAAGCAUAU GACUAAAAAA AAAAAAAAAA
AAAAAAAAAA AAAAAAAAAA AAAAAAAAAA AAAAAAAAAA AAAAAAAAAA AAAA
```

Teda 30 A-čok, potom prepájací "linker" s desiatimi nukleotidmi (GCAUAUGACU) a
napokon ďalších 70 A-čok.

Predpokladám, že toto je výsledok ďalších optimalizácií na zvýšenie proteinovej
expresie.

Záverom
-----------
Poznáme teda presné zloženie mRNA vakcíny BNT162b2 a pri väčšine jej súčastí
aj rozumieme, na čo slúžia:

 * čiapka, ktorá sa stará o to, aby RNA vyzerala ako bežná mRNA
 * dobre známa a zoptimalizovaná 5' neprekladaná oblasť (UTR)
 * signálny peptid s optimalizovanými kodónmi, ktorý odošle Spike proteín
 na správne miesto (presne skoprírovaný z pôvodného vírusu)
 * verzia originálneho výčnelku s optimalizovanými kodónmi a s dvoma
 zámenami za "prolín", ktoré zaručia že proteín ostane v správnom tvare
 * dobre známa a zoptimalizovaná 3' neprekladaná oblasť
 * trocha záhadný poly-A koniec s nevysvetleným prepajacím "linkerom"

Kodónové optimalizácie pridávajú do mRNA kopu G-čok a C-čok. Použitie Ψ
(1-methyl-3'-pseudouridylylu) namiesto U zasa pomáha vyhnúť sa imunitnému
systému, takže mRNA vydrží dostatočne dlho na to, aby nám ho pomohla
vytrénovať.

Ďalšie čítane a pozeranie
-----------------------
V roku 2017 som mal dvojhodinovú prezentáciu o DNA, ktorú si môžete
[pozrieť tu](https://berthub.eu/dna). Ako tento článok, aj táto prezentácia
je zameraná na itčkárov.

Taktiež od roku 2001 udržiavam stránku "[DNA for
programmers](https://berthub.eu/amazing-dna)".

Tiež by sa vám mohol páčiť [tento úvod do nášho úžasného imunitného systému](https://berthub.eu/articles/posts/immune-system/).

A nakoniec, [tento zoznam mojich článkov](https://berthub.eu/articles) často spomína
DNA, SARS-CoV-2 a COVID.

<!--

Aminoacids
----------
Life is built out of or by proteins. Proteins are built out of 20 different
building blocks called amino acids. These amino acids have different chemical
and physical properties. Some attract each other, some are rigid, some have
a positive charge etc.

Once amino acids are chained together we call them a 'protein', and will
adopt a shape.  This is the famous 'folding'.  The shape and state of a
protein can depend on temperature, acidity, presence of (UV) light, magnetic
fields (!!) and many other things.

The famous "Spike" of SARS-CoV-2 is such a protein.

As noted, DNA is organized in 6-bit codons. Each codon uniquely encodes for
a single amino acid, using a nearly universal table:

<center>
{{< figure src="/articles/codon-table.png" caption="[The codon table](https://en.wikipedia.org/wiki/DNA_and_RNA_codon_tables) (Wikipedia)" >}}
</center>

From this table, we see that 'CCT' in DNA corresponds to the Proline
amino acid.


Ok but what is REALLY in the vaccine
------------------------------------
So, how does this work? DNA is like the 'flash drive' storage of biology.
DNA is very durable, internally redundant and very reliable. But much like
computers do not execute code directly from a flash drive, before something
happens, the code gets copied to a faster, more versatile yet far more
fragile system.

For computers, this is RAM, for biology it is RNA. The resemblence is
striking. RAM degrades very quickly unless lovingly tendered. The reason the
Pfizer/BioNTech mRNA vaccine must be stored in the deepest of deep freezers is the
same: RNA is a fragile flower.

We'll get to the difference later.

How the vaccine works
---------------------



For this reason also, there are no RNA printers. The DNA printer emits tiny
amounts of DNA, which are then injected into bacteria as little circular
loops of DNA called plasmids.



-->
