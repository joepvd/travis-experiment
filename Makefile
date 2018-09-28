tar:
UNAME := $(shell uname -s | tr '[A-Z]' '[a-z]')
ifeq ($(UNAME),darwin)
  TAR := gtar
else
  ifeq ($(UNAME),linux)
    TAR := tar
  else
    $(error Operating system $(UNAME) is not yet supported)
  endif
endif
ifeq (,$(shell which $(TAR)))
  $(error Please ensure GNU tar is installed and is available as $(TAR))
endif
TAR := LC_ALL=C $(TAR)

.PHONY: hello
hello: tar
	@echo $(TAR)
