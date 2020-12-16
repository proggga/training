
all: build

build:
	rm -rfv docs
	cp docs_raw/docs/home.md docs_raw/docs/index.md
	cd docs_raw && docnado --html ../docs
	rm docs_raw/docs/index.md


run:
	cd docs_raw && docnado