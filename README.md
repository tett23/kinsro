# kinsro

## Build

```
$ VINDEX_PATH=path_to_vindex_file make build
```

## Test

```
$ make test
```

## Setup Raspberry Pi

```
$ cp ~/.ssh/id_rsa.pub authorized_keys
$ sh ./scripts/setup_raspberry_pi.sh
$ cd ./infra
$ TARGET_IP={ip address} TARGET_HOST={(encode|media)N.kinsro.local} ansible-playbook -i production --limit initial_setup tasks/add_static_ip_address.yml
$ echo '{(encode|media)N} ansible_user=pi' >> production
```

### Setup encode servers

```
$ ansible-playbook -i production --limit encode tasks/kernel_update.yml
$ ansible-playbook -i production encode.yml
```

### Setup media servers

```
sudo echo 'arm_64bit=1' >> /boot/config.txt
SKIP_WARNING=1 sudo rpi-update
sudo shutdown -r now
```

sudo apt install -y xfsprogs
sudo fdisk /dev/sda
sudo parted -s /dev/sda mklabel gpt mkpart primary 0% 100%
sudo mkfs -t xfs /dev/sda
sudo mkdir -p /mnt/video1
sudo mount -a

### Setup video tmp

```
sudo fdisk /dev/sda # delete all partitions and save
sudo parted -s /dev/sda mklabel gpt mkpart primary 0% 100%
sudo mkfs -t ext4 /dev/sda
sudo mkdir -p /media/video_tmp
# edit fstab
sudo mount -a
sudo chown pi:pi /media/video_tmp
```
