#!/bin/sh
# Time-stamp: <2023-03-10 16:04:07 krylon>

ldapsearch -x -b "ou=contacts,dc=krylon,dc=net" -s sub -a always -l 10 -H "ldap://achtfaden" -P 3 "(sn=Simpson)" sn givenName

