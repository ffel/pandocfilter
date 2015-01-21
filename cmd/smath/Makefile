all:
	pandoc readme.md -o readme.json --filter pdtree 2> tree.txt
	pandoc readme.md -s -t latex -o readme.tex --filter ./frac
	# pandoc readme.md -o readme.pdf --filter ./frac


clean:
	rm *.json *.tex *.pdf tree.txt frac