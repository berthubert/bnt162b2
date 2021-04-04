The big BNT162b2 archive
------------------------
All vaccine data here is sourced from this [World Health
Organization
document](https://mednet-communities.net/inn/db/media/docs/11889.doc).

This describes the RNA contents of the BNT162b2 SARS-CoV-2 vaccine.  We
should all be very grateful that BioNTech has shared this data with us.  And
of course we should also be grateful to the many many researchers that
worked for decades to bring the state of the art to the point that such a
vaccine could be developed.  It is marvelous.

This GitHub repository is a companion to [Reverse Engineering the source code of the BioNTech/Pfizer SARS-CoV-2
Vaccine](https://berthub.eu/articles/posts/reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/)
and [part
2](https://berthub.eu/articles/posts/part-2-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/).

In part 2 we find the challenge: Can we find an algorithm that turns the
viral RNA into the vaccine RNA?

If so that would help explain how the vaccine is designed. It would also be
useful for other researchers to turn viral RNA into RNA that gets converted
into proteins efficiently.  

Details are in [part 2](https://berthub.eu/articles/posts/part-2-reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/).

Data files
----------

 * [ncov-s.fasta](ncov-s.fasta): the unprocessed RNA of the virus S protein
 * [vaccine-s.fasta](vaccine-s.fasta): the unprocessed RNA of the vaccine S protein
 * [side-by-side.csv](side-by-side.csv): the two files aligned, with virual and vaccine codons side by side 



