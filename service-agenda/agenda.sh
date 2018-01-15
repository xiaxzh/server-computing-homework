#!/bin/sh
set -e

if [ "$1" = "agendad" ]
then
    exec agendad
else
    exec agenda "$@"
fi