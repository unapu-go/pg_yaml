PROJ_NAME = pg_yaml
RM = rm -vrf
PLGO_DIR = ../../../gitlab.com/microo8/plgo/plgo

build:
	@ echo building...
	@ (cd $(PLGO_DIR) && go build) && $(PLGO_DIR)/plgo && echo done.

buildc:
	@ cd .docker && ./make.sh setup && ./make.sh compile

clean:
	@ [ -e build ] && echo cleaning... && $(RM) build && echo done.

install:
	@ cd build && sudo make install with_llvm=no

.PHONY: clean build
