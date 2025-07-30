date = $(shell git log | grep -m 1 Date | sed -r 's/Date: +[A-Z][a-z]+ ([A-Z][a-z]+) ([0-9]+) [^ ]+ ([0-9]+) .+/\2_\1_\3/')
version = $(shell git describe)


all: clioDoc.pdf
clioDoc.pdf: clioDoc.tex clio.tex
	echo $(date) | tr '_' ' ' > date.txt
	echo $(version) | tr '-' ' ' | awk '{printf "%s", $$1; if ($$2) printf "-%s", $$2; printf "\n"}' > version.txt
	latex clioDoc
	latex clioDoc
	latex clioDoc
	dvipdf -dALLOWPSTRANSPARENCY clioDoc
clio.tex: ../clio.org
	bash ../scripts/org2nw ../clio.org | awk -f ../scripts/preWeave.awk | noweave -n -x > clio.tex
publish: clioDoc.pdf
	if [ -d ~/WinHome/ownCloud/ ]; then \
		cp clioDoc.pdf ~/WinHome/ownCloud/docs/; \
	fi
clean:
	rm -f clio.tex *.pdf *.aux *.bbl *.blg *.dvi *.log *.toc 
