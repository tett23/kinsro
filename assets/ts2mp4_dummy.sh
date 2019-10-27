#!/bin sh

ts2mp4() {
  local TS=$1
  local BASE=$(basename $TS .ts)
  local PRESET_PATH=$3
  local WIDTH=$4
  local HEIGHT=$5

  [ "${BASE}.ts" = "$(basename $TS)" ] || exit 1
  local DIR=$(dirname $TS)
  local OUTPUT="$DIR/$BASE.mp4"

  until [ ! `pgrep ffmpeg` ]
  do
    sleep 60
  done

  touch $OUTPUT

  exit
}

ts2mp4 $1 $2 $3 $4 $5
