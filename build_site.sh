#! /bin/bash

rm -rf public
cd site
hugo && mv public ../public
