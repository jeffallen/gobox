# Makefile generated by gb: http://go-gb.googlecode.com
# gb provides configuration-free building and distributing

include $(GOROOT)/src/Make.inc

TARG=gobox
GOFILES=\
	applets.go\
	gobox.go\
	main.go\

# gb: this is the local install
GBROOT=../..

# gb: compile/link against local install
GC+= -I $(GBROOT)/_obj
LD+= -L $(GBROOT)/_obj

# gb: default target is in GBROOT this way
command:

include $(GOROOT)/src/Make.cmd

# gb: copy to local install
$(GBROOT)/bin/$(TARG): $(TARG)
	mkdir -p $(dir $@); cp -f $< $@
command: $(GBROOT)/bin/$(TARG)

# gb: local dependencies
$(TARG): $(GBROOT)/_obj/applets/ps.a
$(TARG): $(GBROOT)/_obj/applets/grep.a
$(TARG): $(GBROOT)/_obj/applets/head.a
$(TARG): $(GBROOT)/_obj/applets/kill.a
$(TARG): $(GBROOT)/_obj/applets/mkdir.a
$(TARG): $(GBROOT)/_obj/applets/umount.a
$(TARG): $(GBROOT)/_obj/applets/mknod.a
$(TARG): $(GBROOT)/_obj/applets/rm.a
$(TARG): $(GBROOT)/_obj/applets/gzip.a
$(TARG): $(GBROOT)/_obj/applets/shell.a
$(TARG): $(GBROOT)/_obj/applets/echo.a
$(TARG): $(GBROOT)/_obj/applets/ls.a
$(TARG): $(GBROOT)/_obj/applets/httpd.a
$(TARG): $(GBROOT)/_obj/applets/telnetd.a
$(TARG): $(GBROOT)/_obj/applets/mount.a
$(TARG): $(GBROOT)/_obj/applets/wget.a
$(TARG): $(GBROOT)/_obj/common.a
$(TARG): $(GBROOT)/_obj/applets/cat.a
$(TARG): $(GBROOT)/_obj/applets/chroot.a
