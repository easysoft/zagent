#!/bin/bash
ck_ok(){
    if [ $? -ne 0 ];then
        echo -e "\033[31m $1 error. \033[0m"
        exit 1
    fi
}
command_exist(){
    if type $1 >/dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}
is_active(){
    check_command=`sudo systemctl is-active $1`
    if [ ${check_command} == 'active' ]; then
        return 0
    else
        return 1
    fi
}
create_dir(){
    if [ ! -d $1 ];then
        mkdir $1
    fi
    if [ ! -d $1 ];then
        echo "mkdir $1 error."
        exit 1
    fi
}
create_all_dir(){
    create_dir ${HOME}/zagent
    create_dir ${HOME}/zagent/kvm
    create_dir ${HOME}/zagent/token
    create_dir ${HOME}/zagent/websockify
    create_dir ${HOME}/zagent/nginx
    create_dir ${HOME}/zagent/nginx/conf.http.d
    create_dir ${HOME}/zagent/nginx/conf.stream.d
}

download_zagent()
{
    cd  ${HOME}/zagent
    if [ -f agent.zip ]
    then
        echo "agent.zip already exist"
        echo "Check md5"
        zip_md5=`md5sum agent.zip|awk '{print $1}'`
        if [ ${zip_md5} == '2b3bd1cb0bedb5abe150220cc1895e5d' ]
        then
            return 0
        else
            /bin/mv agent.zip agent.zip.old
        fi
    fi
    
    curl -L -o agent.zip https://ztf.im/dl/ztf/3.2.0/linux/ztf.zip
    ck_ok "download zagent"
    echo "Check md5"
    zip_md5=`md5sum agent.zip|awk '{print $1}'`
    if [ ${zip_md5} == '2b3bd1cb0bedb5abe150220cc1895e5d' ]
    then
        return 0
    fi
    return 1
}
install_zagent()
{
    if [ -f ${HOME}/zagent/zagent/ztf ];then
        if [ ${force} == false ];then
            echo "Already installed zagent"
            return
        fi
    fi
    download_zagent
    ck_ok "download Zagent"
    cd  ${HOME}/zagent
    if [ -f agent.zip ]
    then
        echo "unZip zagent"
        unzip -d ./zagent -o ./agent.zip
        ck_ok "unZip Zagent"
        cd zagent
        nohup ./ztf -P 8085 &
    fi
    /usr/bin/rm ${HOME}/zagent/agent.zip
}

download_novnc()
{
    cd  ${HOME}/zagent
    if [ -f novnc.zip ]
    then
        echo "novnc.zip already exist"
        echo "Check md5"
        zip_md5=`md5sum novnc.zip|awk '{print $1}'`
        if [ ${zip_md5} == '2b3bd1cb0bedb5abe150220cc1895e5d' ]
        then
            return 0
        else
            /bin/mv novnc.zip novnc.zip.old
        fi
    fi
    
    curl -L -o novnc.zip https://ztf.im/dl/ztf/3.2.0/linux/ztf.zip
    ck_ok "download novnc"
    echo "Check md5"
    zip_md5=`md5sum novnc.zip|awk '{print $1}'`
    if [ ${zip_md5} == '2b3bd1cb0bedb5abe150220cc1895e5d' ]
    then
        return 0
    fi
    return 1
}
install_novnc()
{
    if [ -f ${HOME}/zagent/novnc/vnc.html ];then
        if [ ${force} == false ];then
            echo "Already installed novnc"
            return
        fi
    fi
    download_novnc
    ck_ok "download novnc"
    cd  ${HOME}/zagent
    if [ -f novnc.zip ]
    then
        echo "unZip novnc"
        unzip -o -d ./novnc ./novnc.zip
        ck_ok "unZip novnc"
    fi
    /usr/bin/mv ${HOME}/zagent/vnc_lite.html ${HOME}/zagent/index.html
    /usr/bin/rm ${HOME}/zagent/novnc.zip
}
download_nginx()
{
    cd  /usr/local/src
    if [ -f nginx-1.23.0.tar.gz ]
    then
        echo "nginx-1.23.0.tar.gz already exist"
        echo "check md5"
        zip_md5=`md5sum nginx-1.23.0.tar.gz|awk '{print $1}'`
        if [ ${zip_md5} == 'e8768e388f26fb3d56a3c88055345219' ]
        then
            return 0
        else
            sudo /bin/mv nginx-1.23.0.tar.gz nginx-1.23.0.tar.gz.old
        fi
    fi
    
    sudo curl -L -O http://nginx.org/download/nginx-1.23.0.tar.gz
    ck_ok "download Nginx"
    echo "check md5"
    zip_md5=`md5sum nginx-1.23.0.tar.gz|awk '{print $1}'`
    if [ ${zip_md5} == 'e8768e388f26fb3d56a3c88055345219' ]
    then
        return 0
    else
        echo "nginx check md5 fail"
        exit 1
    fi
}
install_nginx()
{
    if command_exist nginx;then
        if [ ${force} == false ];then
            echo "Already installed nginx"
            return
        fi
    fi
    download_nginx
    cd /usr/local/src
    echo "Unzip Nginx"
    sudo tar zxf nginx-1.23.0.tar.gz
    ck_ok "Zip Nginx"
    cd nginx-1.23.0
    
    
    echo "Install depends"
    ##ubuntu
    for pkg in make libpcre++-dev build-essential libssl-dev  zlib1g-dev
    do
        if ! dpkg -l $pkg >/dev/null 2>&1
        then
            sudo apt reinstall -y $pkg
            ck_ok "apt install $pkg"
        else
            echo "$pkg already installed"
        fi
    done
    
    echo "Configure Nginx"
    sudo ./configure --prefix=/usr/local/nginx  --with-http_ssl_module --with-http_stub_status_module --with-stream
    ck_ok "Configure Nginx"
    
    
    echo "Make&install"
    sudo make && sudo make install
    ck_ok "Make&install"
    
    
    echo "Edit systemd config"
    
    
cat > /tmp/nginx.service <<EOF
[Unit]
Description=nginx - high performance web server
Documentation=http://nginx.org/en/docs/
After=network-online.target remote-fs.target nss-lookup.target
Wants=network-online.target

[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/bin/sh -c "/bin/kill -s HUP \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"
ExecStop=/bin/sh -c "/bin/kill -s TERM \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"

[Install]
WantedBy=multi-user.target
EOF
    
    sudo /bin/mv /tmp/nginx.service /lib/systemd/system/nginx.service
    ck_ok "edit nginx.service"
    
    sudo ln -s /usr/local/nginx/sbin/nginx /usr/local/bin/
    echo "Load service"
    sudo systemctl unmask nginx.service
    sudo  systemctl daemon-reload
    sudo systemctl enable nginx
    echo "Start Nginx"
    sudo systemctl start nginx
    ck_ok "Start Nginx"
    sudo /usr/bin/rm /usr/local/src/nginx-1.23.0.tar.gz
}

install_kvm()
{
    if command_exist libvirtd;then
        if [ ${force} == false ];then
            echo "Already installed kvm"
            return
        fi
    fi
    echo "Install kvm"
    sudo apt reinstall -y qemu-kvm libvirt-daemon-system libvirt-clients libvirt-dev qemu virt-manager bridge-utils libosinfo-bin
    ck_ok "Install kvm"
    
    sudo service libvirtd restart
    ck_ok "Start libvirtd"
}

install_websockify()
{
    if [ -f ${HOME}/zagent/websockify/run ];then
        if [ ${force} == false ];then
            echo "Already installed websockify"
            return
        fi
    fi
    echo "Install websockify"
    if ! command_exist git;then
        echo "Install git"
        sudo apt install  -y git
        ck_ok "apt install git"
    fi
    cd ${HOME}/zagent/websockify
    if [ "`ls -A ${HOME}/zagent/websockify`" = "" ]; then
        git clone https://github.com/novnc/websockify.git ./
    else
        git pull
    fi
    ck_ok "Install websockify"
    nohup ./run --token-plugin TokenFile --token-source ../token/ 6080 &
}


install()
{
    sudo apt update
    
    if ! command_exist curl;then
        echo "Install curl"
        sudo apt install  -y curl
        ck_ok "apt install curl"
    fi
    
    if $is_install_zagent;then
        create_dir ${HOME}/zagent
        create_dir ${HOME}/zagent/zagent
        install_zagent
    fi
    if $is_install_nginx;then
        create_dir ${HOME}/zagent/nginx
        create_dir ${HOME}/zagent/nginx/conf.http.d
        create_dir ${HOME}/zagent/nginx/conf.stream.d
        install_nginx
    fi
    if $is_install_kvm;then
        create_dir ${HOME}/zagent/kvm
        create_dir ${HOME}/zagent/token
        install_kvm
    fi
    if $is_install_novnc;then
        install_novnc
    fi
    if $is_install_websockify;then
        create_dir ${HOME}/zagent/websockify
        install_websockify
    fi
    echo -e "\033[32m Install ${soft} success \033[0m"
}

if [ ! -n "$(egrep -o "(vmx|svm)" /proc/cpuinfo)" ];then
    echo -e "\033[31m Not support virtualization \033[0m"
    exit 1
fi

source /etc/os-release
if [ ${ID} != "ubuntu" ];then
    echo -e "\033[31m Not support os \033[0m"
    exit 1
fi

is_install_zagent=true
is_install_nginx=true
is_install_kvm=true
is_install_novnc=true
is_install_websockify=true
soft="zagent,nginx,kvm,novnc,websockify"
force=false
while getopts ":s:r" optname
do
    case "$optname" in
        "s")
            if [ -n $OPTARG ];then
                soft=$OPTARG
                is_install_zagent=false
                is_install_nginx=false
                is_install_kvm=false
                is_install_novnc=false
                is_install_websockify=false
                if [[ $OPTARG =~ zagent ]];then
                    is_install_zagent=true
                fi
                if [[ $OPTARG =~ nginx ]];then
                    is_install_nginx=true
                fi
                if [[ $OPTARG =~ kvm ]];then
                    is_install_kvm=true
                fi
                if [[ $OPTARG =~ novnc ]];then
                    is_install_novnc=true
                fi
                if [[ $OPTARG =~ websockify ]];then
                    is_install_websockify=true
                fi
            fi
        ;;
        "r")
            force=true
        ;;
        *)
            echo "Unknown error while processing options"
        ;;
    esac
done

HOME="`cat /etc/passwd |grep ^${SUDO_USER:-$(id -un)}: | cut -d: -f 6`"
HOME=${HOME:-$HOME}

install
