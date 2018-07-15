AUTOMAKE_OPTIONS = foreign
SUBDIRS = lib/sqlite

AM_CPPFLAGS = -I$(srcdir)/lib/sqlite

bin_PROGRAMS = main
main_SOURCES = src/main.cpp

main_LDADD = lib/sqlite/libsqlite3.la
