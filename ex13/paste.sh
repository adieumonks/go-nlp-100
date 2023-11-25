wget https://nlp100.github.io/data/popular-names.txt
cut -f 1 popular-names.txt > col1.txt
cut -f 2 popular-names.txt > col2.txt
paste col1.txt col2.txt > marge.txt
cat marge.txt
rm popular-names.txt col1.txt col2.txt marge.txt
