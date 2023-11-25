wget https://nlp100.github.io/data/popular-names.txt
cut -f 1 popular-names.txt | sort | uniq -c | sort -k 1 -r
rm popular-names.txt
