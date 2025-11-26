## !/bin/bash

main () {
  local PROXY_ID='i-0619be4cc740fa700'
  aws ec2 stop-instances --instance-ids "${PROXY_ID}" >> /dev/null
}

main "$@"
