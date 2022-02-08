#/usr/bin/env bash


# ------------------------------------------------------------------
# python

SUBMITTER_PY=/home/jovyan/.ipython/submitter.py
SUBMITTER_PY_URL="https://raw.githubusercontent.com/drhodes/submitter-py/main/submitter.py"

# put the python submitter script on the path
wget $SUBMITTER_PY_URL -O $SUBMITTER_PY


# ------------------------------------------------------------------
# javascript that enables the submit button

SUBMITTER_JS_URL="https://raw.githubusercontent.com/drhodes/submitter-py/main/js/custom.js"
SUBMITTER_JS_DIR=/home/jovyan/.jupyter/custom/
SUBMITTER_JS=$SUBMITTER_JS_DIR/custom.js

mkdir -p $SUBMITTER_JS_DIR

# put the python submitter script on the path
wget $SUBMITTER_JS_URL -O $SUBMITTER_JS



