all: clioDoc.pdf
clioDoc.pdf: clioDoc.tex clio.tex
	latex clioDoc
	latex clioDoc
	latex clioDoc
	dvipdf clioDoc
clio.tex: ../clio.org
	bash ../scripts/org2nw ../clio.org | awk -f ../scripts/preWeave.awk | noweave -n -x > clio.tex
clean:
	rm -f clio.tex *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc 
