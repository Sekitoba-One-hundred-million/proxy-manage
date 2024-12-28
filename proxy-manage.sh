## !/bin/bash

main () {
  local PROXY_ID='i-0db77541139990ecd'
  local DNS_NAME="$(aws ec2 describe-instances | jq -r ".Reservations[].Instances[] | select(.InstanceId==\"${PROXY_ID}\") | .PublicDnsName")"

  if [ ! -z "${DNS_NAME}" ]; then
    aws ec2 stop-instances --instance-ids "${PROXY_ID}" >> /dev/null

    while : ;
    do
      sleep 5
      local state="$(aws ec2 describe-instances | jq -r ".Reservations[].Instances[] | select(.InstanceId==\"${PROXY_ID}\") | .State.Name")"

      if [ "${state}" == 'stopped' ]; then
        break
      fi
    done
  fi

  DNS_NAME=""
  aws ec2 start-instances --instance-ids "${PROXY_ID}" >> /dev/null

  while : ;
  do
    sleep 2
    DNS_NAME="$(aws ec2 describe-instances | jq -r '.Reservations[].Instances[].PublicDnsName')"

    if [ -z "${DNS_NAME}" ]; then
      continue
    fi

    local status="$(curl -k https://${DNS_NAME}/get -H "Host: httpbin.org" -o /dev/null -w '%{http_code}\n' -s)"
    
    if [ "${status}" == 200 ]; then
      break
    fi
  done

  echo "${DNS_NAME}"
}

main "$@"
