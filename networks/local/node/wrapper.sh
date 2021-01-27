#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-fastock-chain-daemon.log}

##
## Run binary with all parameters
##
export OKEXCHAINDHOME="/fastock-chain-daemon/node${ID}/fastock-chain-daemon"

if [ -d "$(dirname "${OKEXCHAINDHOME}"/"${LOG}")" ]; then
  fastock-chain-daemon --chain-id okexchain-1 --home "${OKEXCHAINDHOME}" "$@" | tee "${OKExCHAINDHOME}/${LOG}"
else
  fastock-chain-daemon --chain-id okexchain-1 --home "${OKEXCHAINDHOME}" "$@"
fi

