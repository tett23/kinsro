#!/bin/sh

ts2mp4() {
  local PI_HOME=/home/pi
  local TS=$1
  local OUT=$2
  local WIDTH=$3
  local HEIGHT=$4
  local BASE=$(basename $TS .ts)
  [ "${BASE}.ts" = "$(basename $TS)" ] || exit 1

  until [ ! `pgrep ffmpeg` ]
  do
    sleep 60
  done

  local X264_HIGH_HDTV=" \
  -fflags +discardcorrupt \
  -c:v h264_omx \
  -b:v 5000k -rc_init_occupancy 5000k -bufsize 5000k -maxrate 5000k -minrate 2000k \
  -vf yadif=1:-1 \
  -aspect 16:9 -s ${WIDTH}x${HEIGHT} \
  -r 30000/1001 \
  -fpre $PI_HOME/encode/ts2mp4.ffpreset \
  -c:a copy -bsf:a aac_adtstoasc \
    "


  # -vf yadif=1:-1 \
  # -rc_init_occupancy 2000k -bufsize 2000k -maxrate 5000k -minrate 2000k \
  # -fpre libx264-hq-ts.ffpreset \
  # -f mp4 \
  #	-b:v 5000k -rc_init_occupancy 2000k -bufsize 20000k -maxrate 12000k \
  # -b:v 20000k -rc_init_occupancy 2000k -bufsize 20000k -maxrate 25000k \
  # -c:v mpeg2_mmal \
  # -vf yadif=1:-1,scale=${WIDTH}:${HEIGHT} \
  # -aspect 16:9 \
  # -c:v h264_omx \
  # -b:v 20000k \
  # -c:a copy -bsf:a aac_adtstoasc \
  # -fpre libx264-hq-ts.ffpreset \
  # -aspect 16:9 -s ${WIDTH}x${HEIGHT} \
  # -c:v mpeg2_mmal \
  # -f mp4 \
  # -fflags +discardcorrupt \
  # -b:v 20000k -rc_init_occupancy 2000k -bufsize 20000k -maxrate 25000k \
  # -acodec libfaac -ac 2 -ar 48000 -ab 128k \
  # ffmpeg -fflags +discardcorrupt -c:v mpeg2_mmal -i isdbt.ts -c:a copy -bsf:a aac_adtstoasc -c:v h264_omx -b:v 5000k -y out.mp4

  #	X264_HIGH_HDTV=" \
  #	  -vaapi_device /dev/dri/renderD128 \
  #	  -hwaccel vaapi -hwaccel_output_format vaapi \
  #	  -c:v h264_vaapi \
  #	  -vf 'format=nv12|vaapi,hwupload,deinterlace_vaapi,scale_vaapi=w=1920:h=1080' \
  #	  -c:v h264_vaapi -profile 100 -level 40 -qp 23 -aspect 16:9 \
  #	  -c:a copy \
  #	  "

  # OUT=sed -e 's/ /\ /g' ${OUT}
  # echo "-y -i $TS ${X264_HIGH_HDTV} '${OUT}'"
  ffmpeg -y -i "$TS" $X264_HIGH_HDTV "$OUT"

  exit
}

ts2mp4 $1 $2 $3 $4