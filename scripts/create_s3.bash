#!/usr/bin/env bash
#
# example:
#    ./create_s3.bash
#    ./create_s3.bash --prd

set -e
cd `dirname $0`/./

readonly TEMPALTE_FILE="./s3_template.yml" 
env="stg"

[[ "${1}" = "--prd" ]] && isPrd=true || isPrd=false
"${isPrd}" && echo "****** Create Bucket Production Mode!!! ******"
"${isPrd}" && env="prd"

stack_name="${env}-email-bounce-bucket"

aws cloudformation deploy \
    --stack-name ${stack_name} \
    --template-file "${TEMPALTE_FILE}" \
    --parameter-overrides \
      Env=${env} \
    --profile us-east-1-user \

aws cloudformation update-termination-protection \
    --enable-termination-protection \
    --stack-name ${stack_name} \
    --profile us-east-1-user \

exit 0