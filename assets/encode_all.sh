#!/bin/sh

for file in $(find /media/video_tmp -name '*.ts'); do
  kinsro encode $file
done