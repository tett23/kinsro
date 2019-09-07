FROM raspbian/stretch

WORKDIR /

RUN apt-get update -y -q
RUN apt-get upgrade -y 
RUN apt-get install -y \
    build-essential \
    libssl-dev \
    libreadline-dev \
    zlib1g-dev \
    curl
RUN apt-get install -y \
    openssh-server 

RUN mkdir /var/run/sshd
RUN echo 'PermitRootLogin yes' >> /etc/ssh/sshd_config
RUN useradd -m pi && echo "pi:pi" | chpasswd && gpasswd -a pi sudo
RUN mkdir -p /home/pi/.ssh && chown pi /home/pi/.ssh && chmod 700 /home/pi/.ssh
RUN echo 'pi ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
COPY ./authorized_keys /home/pi/.ssh/authorized_keys
RUN chown pi /home/pi/.ssh/authorized_keys && chmod 600 /home/pi/.ssh/authorized_keys