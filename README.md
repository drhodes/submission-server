A program used for submitting student answers to the remote answer server. 

todo document this after it is working.

There are some environmental variables the server requires, it will
check to make sure they exist and fall over if they don't.

Designed to live in a kubernetes cluster, this server takes answer
submissions from jupyter notebooks and sends requests to another
server that can 


```bash
# ------------------------------------------------------
export STAFF_SUBMITTER_USERID="staff"       
export STAFF_SUBMITTER_PASSWD="[enter a good password]" 

# -------------------------------------------------------

export ANSWER_SERVER="[answerserver:port]"
export ANSWER_SERVER_USERID="[this can be anything]"
export ANSWER_SERVER_PASSWD="[enter a good password]" 
```

There environment variables are used elsewhere.

`ANSWER_SERVER`, `ANSWER_SERVER_USERID` and `ANSWER_SERVER_PASSWD` are
shared with the [RemoXBlock](https://github.com/drhodes), which uses needs them to fetch answers
from the answer server.

`STAFF_SUBMITTER_USERID` and `STAFF_SUBMITTER_PASSWD` are located in the
kubernetes singleuser section of the jupyterhub config.yaml file.
