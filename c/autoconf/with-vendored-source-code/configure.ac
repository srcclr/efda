#                                               -*- Autoconf -*-
# Process this file with autoconf to produce a configure script.

AC_PREREQ([2.69])
AC_INIT([main], [0.01a], [REPORT-ADDRESS])
LT_INIT(static)
AM_INIT_AUTOMAKE([subdir-objects])
#AC_CONFIG_SRCDIR([src/main.cpp])

# Checks for programs.
AC_PROG_CXX
AC_PROG_CC

AC_MSG_CHECKING(whether -lpthread is needed for pthreads)
AC_CHECK_FUNC(pthread_create,
    [
        AC_MSG_RESULT(no)
        AC_SUBST([CabalNeedLibpthread],[False])
        need_lpthread=0
    ],
    [
        AC_CHECK_LIB(pthread, pthread_create,
            [
                AC_MSG_RESULT(yes)
                AC_SUBST([CabalNeedLibpthread],[True])
                need_lpthread=1
            ],
            [
                AC_SUBST([CabalNeedLibpthread],[False])
                AC_MSG_RESULT([no pthreads support found.])
                need_lpthread=0
            ])
    ])
AC_DEFINE_UNQUOTED([NEED_PTHREAD_LIB], [$need_lpthread],
    [Define 1 if we need to link code using pthreads with -lpthread])

dnl ** pthread_setname_np is a recent addition to glibc, and OS X has
dnl    a different single-argument version.
AC_CHECK_LIB(pthread, pthread_setname_np)
AC_MSG_CHECKING(for pthread_setname_np)
AC_TRY_LINK(
[
#define _GNU_SOURCE
#include <pthread.h>
],
[pthread_setname_np(pthread_self(), "name");],
    AC_MSG_RESULT(yes)
    AC_DEFINE([HAVE_PTHREAD_SETNAME_NP], [1],
      [Define to 1 if you have the glibc version of pthread_setname_np]),
    AC_MSG_RESULT(no)
)

 AC_CHECK_LIB(pthread, pthread_create,
            [
                AC_MSG_RESULT(yes)
                AC_SUBST([CabalNeedLibpthread],[True])
                need_lpthread=1
            ],
            [
                AC_SUBST([CabalNeedLibpthread],[False])
                AC_MSG_RESULT([no pthreads support found.])
                need_lpthread=0
            ])

AC_CHECK_LIB([dl], [dlopen])


AC_CONFIG_SUBDIRS(lib/sqlite)
AC_OUTPUT(Makefile src/Makefile)

#AC_CONFIG_FILES([Makefile])
