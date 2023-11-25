wget https://nlp100.github.io/data/popular-names.txt
cut -f 1 popular-names.txt | sort | uniq
rm popular-names.txt
