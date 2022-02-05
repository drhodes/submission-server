source secret-env.bash

U=$STAFF_SUBMITTER_USERID
P=$STAFF_SUBMITTER_PASSWD
ENDPOINT=localhost:3000/pod_starting

set -x

curl -u "$U:$P" -X POST $ENDPOINT \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d "edx-anon-id=jupyter-0123456789ABCDEF-ABC12"

U=student
P=student

ENDPOINT=localhost:3000/submit-answers
curl -u "$U:$P" -X POST $ENDPOINT \
     -H "Content-Type: application/x-www-form-urlencoded" \
     -d 'edx-anon-id=jupyter-0123456789ABCDEF-ABC12&labname=SimpleLab1&lab-answers={"q1":2, "q2":7}'
