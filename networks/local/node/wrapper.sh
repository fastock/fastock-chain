#!/usr/bin/env sh

##
## Input parameters
##
ID=${ID:-0}
LOG=${LOG:-fastock-chain-daemon.log}

##
## Run binary with all parameters
##
export FASTOCKCHAINDAEMONHOME="/fastock-chain-daemon/node${ID}/fastock-chain-daemon"

if [ -d "$(dirname "${FASTOCKCHAINDAEMONHOME}"/"${LOG}")" ]; then
  fastock-chain-daemon --chain-id blockchain-1 --home "${FASTOCKCHAINDAEMONHOME}" "$@" | tee "${FASTOCKCHAINDAEMONHOME}/${LOG}"
else
  fastock-chain-daemon --chain-id blockchain-1 --home "${FASTOCKCHAINDAEMONHOME}" "$@"
fi

