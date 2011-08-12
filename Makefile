include $(GOROOT)/src/Make.inc

TARG    = perlin
GOFILES = perlin.go
CLEANFILES=testnoise

include $(GOROOT)/src/Make.pkg

testnoise: testnoise.go install
	$(GC) testnoise.go
	$(LD) -o $@ testnoise.$(O)
