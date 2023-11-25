wget https://nlp100.github.io/data/popular-names.txt
sed -e "s/\t/ /g" popular-names.txt
rm popular-names.txt
