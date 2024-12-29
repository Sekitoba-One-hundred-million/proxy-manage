## !/bin/bash

main () {
  local PROXY_ID='i-0db77541139990ecd'
  aws ec2 stop-instances --instance-ids "${PROXY_ID}" >> /dev/null
}

main "$@"
