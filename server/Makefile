include $(GOROOT)/src/Make.inc

DIRS=\
pkg/location\
pkg/people\
pkg/rooms\
pkg/packets\
cmd/opensky-server

all: install

clean.dirs: $(addsuffix .clean, $(DIRS))

install.dirs: $(addsuffix .install, $(DIRS))

%.clean:
	+cd $* && gomake clean
	
%.install:
	+cd $* && gomake install
	
clean: clean.dirs

install: install.dirs
