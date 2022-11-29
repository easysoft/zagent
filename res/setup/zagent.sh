#!/bin/bash
ck_ok(){
    if [ $? -ne 0 ];then
        print_res
        echo -e "\033[31m $1 error. \033[0m"
        exit 1
    fi
}
print_res(){
    if $is_install_zagent && $is_success_zagent;then
        echo -e "\033[32m Install zagent success \033[0m"
        if [ -z $secret ];then
            echo -e "\033[31m Secret is empty, zagent start fail \033[0m"
        fi
    fi
    if $is_install_nginx && $is_success_nginx;then
        echo -e "\033[32m Install nginx success \033[0m"
    fi
    if $is_install_kvm && $is_success_kvm;then
        echo -e "\033[32m Install kvm success \033[0m"
    fi
    if $is_install_novnc && $is_success_novnc;then
        echo -e "\033[32m Install novnc success \033[0m"
    fi
    if $is_install_websockify && $is_success_websockify;then
        echo -e "\033[32m Install websockify success \033[0m"
    fi
}
command_exist(){
    if type $1 >/dev/null 2>&1; then
        return 0
    else
        return 1
    fi
}

service_is_inactive(){
    check_command=`sudo ps -ef | grep $1 | grep -v grep | awk '{print $2}'`
    if test -z $check_command; then
        return 0
    else
        return 1
    fi
}
is_inactive(){
    check_command=`sudo systemctl is-active $1`
    
    if [ ${check_command} == 'active' ]; then
        return 1
    else
        return 0
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
        if [ ${zip_md5} == 'f82ed9ac1a83a9677936e06628b1a338' ]
        then
            return 0
        else
            /bin/mv agent.zip agent.zip.old
        fi
    fi
    
    curl -L -o agent.zip https://pkg.qucheng.com/zenagent/app/zagent-host.zip
    ck_ok "download zagent"
    echo "Check md5"
    zip_md5=`md5sum agent.zip|awk '{print $1}'`
    
    if [ ${zip_md5} == 'f82ed9ac1a83a9677936e06628b1a338' ]
    then
        return 0
    fi
    
    return 1
}
install_zagent()
{
    if [ -f ${HOME}/zagent/zagent-host ];then
        if [ ${force} == false ];then
            echo "Already installed zagent"
            if service_is_inactive zagent-host;then
                if [[ -n $secret ]];then
                    nohup ./zagent-host -p 8086 -secret $secret > /dev/null 2>&1 &
                fi
            fi
            return
        fi
    fi
    
    download_zagent
    ck_ok "download Zagent"
    cd  ${HOME}/zagent
    
    if [ -f agent.zip ]
    then
        echo "unZip zagent"
        unzip -o ./agent.zip
        ck_ok "unZip Zagent"
        sudo chmod +x ./zagent-host
        if [[ -n $secret ]];then
            nohup ./zagent-host -p 8086 -secret $secret > /dev/null 2>&1 &
        fi
    fi
    
    /usr/bin/rm -rf ${HOME}/zagent/zagent
    /usr/bin/rm -rf ${HOME}/zagent/agent.zip
}

download_novnc()
{
    cd  ${HOME}/zagent
    
    if [ -f novnc.zip ]
    then
        echo "novnc.zip already exist"
        echo "Check md5"
        zip_md5=`md5sum novnc.zip|awk '{print $1}'`
        if [ ${zip_md5} == 'ab5127dcac5edfa760814a619bcf9ca2' ]
        then
            return 0
        else
            /bin/mv novnc.zip novnc.zip.old
        fi
    fi
    
    curl -L -o novnc.zip https://pkg.qucheng.com/zenagent/app/novnc.zip
    ck_ok "download novnc"
    echo "Check md5"
    zip_md5=`md5sum novnc.zip|awk '{print $1}'`
    
    if [ ${zip_md5} == 'ab5127dcac5edfa760814a619bcf9ca2' ]
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
    
    /usr/bin/mv ${HOME}/zagent/novnc/vnc_lite.html ${HOME}/zagent/novnc/index.html
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
            if is_inactive nginx;then
                sudo systemctl start nginx
            fi
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
            if ! $is_update_apt;then
                sudo apt update
                is_update_apt=true
            fi
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
            if is_inactive libvirtd;then
                sudo service libvirtd start
            fi
            return
        fi
    fi
    
    echo "Install kvm"
    if ! $is_update_apt;then
        sudo apt update
        is_update_apt=true
    fi
    sudo apt reinstall -y qemu-kvm libvirt-daemon-system libvirt-clients libvirt-dev qemu virt-manager bridge-utils libosinfo-bin
    ck_ok "Install kvm"
    
    sudo service libvirtd restart
    sudo chmod 777 /var/run/libvirt/libvirt-sock
    ck_ok "Start libvirtd"
}

download_websockify()
{
    cd  ${HOME}/zagent
    
    if [ -f websockify.zip ]
    then
        echo "websockify.zip already exist"
        echo "Check md5"
        zip_md5=`md5sum websockify.zip|awk '{print $1}'`
        if [ ${zip_md5} == '92c8c80ac513d0aba722bf207aab28fc' ]
        then
            return 0
        else
            /bin/mv websockify.zip websockify.zip.old
        fi
    fi
    
    curl -L -o websockify.zip https://pkg.qucheng.com/zenagent/app/websockify.zip
    ck_ok "download websockify"
    echo "Check md5"
    zip_md5=`md5sum websockify.zip|awk '{print $1}'`
    
    if [ ${zip_md5} == '92c8c80ac513d0aba722bf207aab28fc' ]
    then
        return 0
    fi
    
    return 1
}
install_websockify()
{
    if [ -f ${HOME}/zagent/websockify/run ];then
        if [ ${force} == false ];then
            echo "Already installed websockify"
            if service_is_inactive JSONTokenApi;then
                nohup ${HOME}/zagent/websockify/run --token-plugin JSONTokenApi --token-source http://127.0.0.1:8086/api/v1/virtual/getVncAddress?token=%s 6080 > /dev/null 2>&1 &
            fi
            return
        fi
    fi
    
    echo "Install websockify"
    
    cd ${HOME}/zagent
    
    download_websockify
    ck_ok "Download websockify"
    
    cd  ${HOME}/zagent
    
    if [ -f websockify.zip ]
    then
        echo "unZip websockify"
        unzip -o -d ./websockify ./websockify.zip
        ck_ok "unZip websockify"
    fi
    
    /usr/bin/rm -rf ${HOME}/zagent/websockify.zip
    
    sudo chmod +x ${HOME}/zagent/websockify/run
    nohup ${HOME}/zagent/websockify/run --token-plugin JSONTokenApi --token-source http://127.0.0.1:8086/api/v1/virtual/getVncAddress?token=%s 6080 > /dev/null 2>&1 &
}


install()
{
    if ! command_exist curl;then
        echo "Install curl"
        if ! $is_update_apt;then
            sudo apt update
            is_update_apt=true
        fi
        sudo apt install  -y curl
        ck_ok "apt install curl"
    fi
    
    if $is_install_zagent;then
        create_dir ${HOME}/zagent
        create_dir ${HOME}/zagent/zagent
        install_zagent
        is_success_zagent=true
    fi
    
    if $is_install_nginx;then
        create_dir ${HOME}/zagent/nginx
        create_dir ${HOME}/zagent/nginx/conf.http.d
        create_dir ${HOME}/zagent/nginx/conf.stream.d
        install_nginx
        is_success_nginx=true
    fi
    
    if $is_install_kvm;then
        create_dir ${HOME}/zagent/kvm
        create_dir ${HOME}/zagent/token
        install_kvm
        is_success_kvm=true
    fi
    
    if $is_install_novnc;then
        install_novnc
        is_success_novnc=true
    fi
    
    if $is_install_websockify;then
        create_dir ${HOME}/zagent/websockify
        install_websockify
        is_success_websockify=true
    fi
    
    print_res
    # if $is_install_zagent;then
    #     echo -e "\033[32m Install zagent success \033[0m"
    #     if [ -z $secret ];then
    #         echo -e "\033[31m Secret is empty, zagent start fail \033[0m"
    #     fi
    # fi
    # if $is_install_nginx;then
    #     echo -e "\033[32m Install nginx success \033[0m"
    # fi
    # if $is_install_kvm;then
    #     echo -e "\033[32m Install kvm success \033[0m"
    # fi
    # if $is_install_novnc;then
    #     echo -e "\033[32m Install novnc success \033[0m"
    # fi
    # if $is_install_websockify;then
    #     echo -e "\033[32m Install websockify success \033[0m"
    # fi
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

is_success_zagent=false
is_success_nginx=false
is_success_kvm=false
is_success_novnc=false
is_success_websockify=false

is_update_apt=false

soft="zagent,nginx,kvm,novnc,websockify"
force=false
secret=""
isInstall=true

while getopts ":k:s:cr" optname
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
        "k")
            secret=$OPTARG
        ;;
        "c")
            isInstall=false
            if [ ! -f ${HOME}/zagent/zagentinit ];then
                curl -s -L https://pkg.qucheng.com/zenagent/zagentinit -o ${HOME}/zagent/zagentinit
            fi
            sudo chmod +x ${HOME}/zagent/zagentinit
            ${HOME}/zagent//zagentinit -c
        ;;
        *)
            echo "Unknown error while processing options"
        ;;
    esac
done

HOME="`cat /etc/passwd |grep ^${SUDO_USER:-$(id -un)}: | cut -d: -f 6`"
HOME=${HOME:-$HOME}

if [ ${isInstall} == true ];then
    install
fi
