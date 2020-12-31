---
title: "Reverse Engineering Source Code of the Biontech Pfizer Vaccine: Part 2"
date: 2020-12-31T12:22:03+01:00
draft: true
images:
 - bnt162b2.png
---
<meta name="twitter:card" content="summary_large_image">
<meta name="twitter:site" content="@powerdns_bert">
<meta name="twitter:creator" content="@powerdns_bert">
<meta name="twitter:title" content="Reverse Engineering the source code of the BioNTech/Pfizer SARS-CoV-2 Vaccine: Part 2">
<meta name="twitter:description" content="Welcome! In this post, we'll be taking a character-by-character look at the source code of the BioNTech/Pfizer SARS-CoV-2 mRNA vaccine.">
<meta name="twitter:image" content="https://berthub.eu/articles/bnt162b2.png">

All vaccine data on this page is sourced from this [World Health
Organization
document](https://mednet-communities.net/inn/db/media/docs/11889.doc).

This describes the RNA contents of the BNT162b2 SARS-CoV-2 vaccine.  We
should all be very grateful that BioNTech has shared this data with us.  And
of course we should also be grateful to the many many researchers that
worked for decades to bring the state of the art to the point that such a
vaccine could be developed.  It is marvelous.

Because it is so marvelous, I want to understand everything about the
vaccine. I wrote a page [Reverse Engineering the source code of the BioNTech/Pfizer SARS-CoV-2
Vaccine](https://berthub.eu/articles/posts/reverse-engineering-source-code-of-the-biontech-pfizer-vaccine/)
that describes in some detail what is in the mRNA of the vaccine. It helps
to read this page before continuing, I promise you it will be interesting.

The post left open some questions however, and this is where it gets
fascinating. 

The vaccine contains RNA code for a very *slightly* modified copy of the
SARS-CoV-2 S protein.

The RNA code itself however is *highly* modified from the viral original! And
from what we understand, these modifications make the vaccine **much much
more** effective. It would be a lot of fun to understand these
modifications.

Here is the beginning of the S protein in both the virus and the vaccine RNA
code. Exclamation marks denote differences.

```
Virus:   AUG UUU GUU UUU CUU GUU UUA UUG CCA CUA GUC UCU AGU CAG UGU GUU
Vaccine: AUG UUC GUG UUC CUG GUG CUG CUG CCU CUG GUG UCC AGC CAG UGU GUU
               !   !   !   !   ! ! ! !     !   !   !   !   !            
```

RNA is a string (literally) of RNA characters, `A`, `C`, `G` and `U`. There is no
physical framing on there, but it makes sense to analyse it in groups of
three.

Each group (called a codon) maps to an amino acid (denoted by a capital
letter).  A string of amino acids is a protein.  Here is what that looks
like:

```
Virus:   AUG UUU GUU UUU CUU GUU UUA UUG CCA CUA GUC UCU AGU CAG UGU GUU
          M   F   V   F   L   V   L   L   P   L   V   S   S   Q   C   V
Vaccine: AUG UUC GUG UUC CUG GUG CUG CUG CCU CUG GUG UCC AGC CAG UGU GUU
               !   !   !   !   ! ! ! !     !   !   !   !   !            
```

Here we can see that while the codons are different, the amino acid version
is the same. There are 4*4*4 codons but only 20 amino acids. This means you
can typically change every codon into one of two others, and still code for
the same amino acid.

So in the second codon, `UUU` was changed to `UUC`. This is a net addition
of one 'C' to the vaccine. The third codon changed from `GUU` to `GUG`, which is
a net addition of one `G`.

**It is known that a higher fraction of `G` and `C` characters improves the
efficiency of an mRNA vaccine**.

Now, if that was all there was to it, this could be the end of this page.
"The algorithm is change codons so we get more G and C in there". But then
we meet the 9th codon which changes `CCA` to `CCU`.

Throughout the ~4000 characters of the vaccine, this happens many times.

Our challenge
-------------
The goal is: find an algorithm that modifies the 'wild type' RNA code into
the BNT162b2 one. Because everyone would like to understand how to turn
viral RNA into an effective vaccine. The algorithm does not need to
reproduce the _exact_ RNA code of course, but it would be super nice if it
came up with something very similar, while also being brief.

To help you, I have provided the data in a number of forms.

