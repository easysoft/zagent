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
    if $is_install_zvm && $is_success_zvm;then
        echo -e "\033[32m Install zagent success \033[0m"
    fi
    if $is_install_ztf && $is_success_ztf;then
        echo -e "\033[32m Install ztf success \033[0m"
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

port_is_used(){
    pid=`sudo /usr/bin/lsof -i :$1|grep -v "PID" | awk '{print $2}'`
    if [ "$pid" != "" ];
    then
        return 0
    else
        return 1
    fi
}

service_is_inactive(){
    check_command=`sudo ps -ef | grep $1 | grep -v grep | awk '{print $2}'`
    if [[ -z ${check_command} ]]; then
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
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-host.md5` ]
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
    
    if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-host.md5` ]
    then
        return 0
    fi
    
    return 1
}
restart_zagent()
{
    sudo pkill zagent-host
    if [ -f ${HOME}/zagent/zagent-host.db ]
    then
        rm ${HOME}/zagent/zagent-host.db
    fi
    if service_is_inactive zagent-host;then
        if port_is_used 55001;then
            is_success_zagent=false
            echo -e "\033[31m 端口 55001 已被占用，请清理占用程序后重新执行初始化命令。 \033[0m"
            exit 1
            return
        fi
    fi
    cat > /tmp/zagent.sh <<EOF
#!/bin/sh
### BEGIN INIT INFO
# Provides:          zagent.sh
# Required-start:    $local_fs $remote_fs $network $syslog
# Required-Stop:     $local_fs $remote_fs $network $syslog
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: starts the zagent.sh daemon
# Description:       starts zagent.sh using start-stop-daemon
### END INIT INFO
# zagent.sh

/usr/bin/nohup ${HOME}/zagent/zagent-host -p 55001 -secret ${secret} -s ${zentaoSite} > ${HOME}/zagent/zagent.log 2>&1 &
EOF
    if [ ${ID} = ${ubuntu} ];then
        sudo chmod +x /tmp/zagent.sh
        sudo /bin/mv /tmp/zagent.sh /etc/init.d/
        sudo chmod +x /etc/init.d/zagent.sh
        ck_ok "edit zagent.sh"
        
        echo "Load sh"
        sudo update-rc.d zagent.sh defaults 90
    else
        sudo /bin/mv /tmp/zagent.sh /etc/rc.d/init.d/
        sudo chmod +x /etc/rc.d/init.d/zagent.sh
        ck_ok "edit zagent.sh"
        cd /etc/rc.d/init.d
        sudo chkconfig --add zagent.sh
        sudo chkconfig zagent.sh on
        echo "Load sh"
    fi
    
    echo "Start Zagent"
    /usr/bin/nohup ${HOME}/zagent/zagent-host -p 55001 -secret ${secret} -s ${zentaoSite} > ${HOME}/zagent/zagent.log 2>&1 &
    ck_ok "Start Zagent"
}
install_zagent()
{
    cd  ${HOME}/zagent
    if [ -f ${HOME}/zagent/agent.zip ];then
        zip_md5=`md5sum agent.zip|awk '{print $1}'`
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-host.md5` ]
        then
            if [ ${force} == false ];then
                echo "Already installed zagent"
                if [[ -n $secret ]];then
                    restart_zagent
                fi
                return
            fi
        fi
    fi
    
    download_zagent
    ck_ok "download Zagent"
    cd  ${HOME}/zagent
    
    if [ -f agent.zip ]
    then
        if ! command_exist netstat;then
            install_netTools
        fi
        echo "unZip zagent"
        unzip -o ./agent.zip
        ck_ok "unZip Zagent"
        sudo pkill zagent-host
        sudo chmod +x ./zagent-host
        if [[ -n $secret ]];then
            restart_zagent
        fi
    fi
    
    /usr/bin/rm -rf ${HOME}/zagent/agent.zip
}

install_netTools()
{
    if [[ ${ID} = ${ubuntu} ]];then
        sudo apt install net-tools
    else
        sudo yum install net-tools.x86_64
    fi
}

download_zvm()
{
    cd  ${HOME}/zagent
    
    if [ -f vm.zip ]
    then
        zip_md5=`md5sum vm.zip|awk '{print $1}'`
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-vm.md5` ]
        then
            return 0
        fi
    fi
    
    curl -L -o vm.zip https://pkg.qucheng.com/zenagent/app/zagent-vm.zip
    ck_ok "download zagent-vm"
    echo "Check md5"
    zip_md5=`md5sum vm.zip|awk '{print $1}'`
    
    if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-vm.md5` ]
    then
        return 0
    fi
    
    return 1
}
install_ztf(){
    sleep 3s
    result=$(curl http://127.0.0.1:55201/api/v1/service/setup -X POST -d '{"name":"ztf","secret":"'"$secret"'","server":"'"$zentaoSite"'"}' --header "Content-Type:application/json" | grep success)
    if [[ "$result" != "" ]]
    then
        is_success_ztf=true
    else
        echo -e "\033[31m Install ztf error. \033[0m"
        exit 1
    fi
}
install_zvm()
{
    cd  ${HOME}/zagent
    if [ -f ${HOME}/zagent/vm.zip ];then
        zip_md5=`md5sum vm.zip|awk '{print $1}'`
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/zagent-vm.md5` ]
        then
            if [ ${force} == false ];then
                echo "Already installed zagent-vm"
                if service_is_inactive zagent-vm;then
                    sudo systemctl start zagent-vm
                fi
                install_ztf
                return
            fi
        fi
    fi
    
    download_zvm
    ck_ok "download Zagent-vm"
    cd  ${HOME}/zagent
    
    if [ -f vm.zip ]
    then
        echo "unZip zagent-vm"
        unzip -o ./vm.zip
        ck_ok "unZip Zagent-vm"
        sudo pkill zagent-vm
        sudo chmod +x ./zagent-vm
        cat > /tmp/zagent-vm.service <<EOF
[Unit]
Description=Zagent-vm service
Documentation=https://github.com/easysoft/zenagent
After=network-online.target remote-fs.target nss-lookup.target
Wants=network-online.target

[Service]
User=${USER}
Type=forking
ExecStart=/bin/bash -c "${HOME}/zagent/zagent-vm -p 55201 -s ${zentaoSite} > /dev/null 2>&1 &"

[Install]
WantedBy=multi-user.target
EOF
        
        sudo /bin/mv /tmp/zagent-vm.service /lib/systemd/system/zagent-vm.service
        ck_ok "edit zagent-vm.service"
        
        echo "Load service"
        sudo systemctl unmask zagent-vm.service
        sudo  systemctl daemon-reload
        sudo systemctl enable zagent-vm
        
        echo "Start Zagent-vm"
        sudo systemctl start zagent-vm
        ck_ok "Start Zagent-vm"
    fi
    
    install_ztf
}

download_novnc()
{
    cd  ${HOME}/zagent
    
    if [ -f novnc.zip ]
    then
        echo "novnc.zip already exist"
        echo "Check md5"
        zip_md5=`md5sum novnc.zip|awk '{print $1}'`
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/novnc.md5` ]
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
    
    if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/novnc.md5` ]
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

apt_update()
{
    if [[ ${ID} = ${ubuntu} ]];then
        if ! $is_update_apt;then
            sudo apt update
            is_update_apt=true
        fi
    fi
}

install_depends()
{
    if [[ ${ID} = ${ubuntu} ]];then
        for pkg in make libpcre++-dev build-essential libssl-dev zlib1g-dev
        do
            if ! dpkg -l $pkg >/dev/null 2>&1
            then
                apt_update
                sudo apt reinstall -y $pkg
                ck_ok "install $pkg"
            else
                echo "$pkg already installed"
            fi
        done
    else
        for pkg in gcc-c++ pcre-devel zlib-devel make openssl-devel
        do
            if ! dpkg -l $pkg >/dev/null 2>&1
            then
                apt_update
                if [[ ${ID} = ${ubuntu} ]];then
                    sudo apt reinstall -y $pkg
                else
                    sudo yum install -y $pkg
                fi
                ck_ok "yum install $pkg"
            else
                echo "$pkg already installed"
            fi
        done
    fi
}
install_nginx()
{
    if command_exist nginx;then
        if [ ${force} == false ];then
            echo "Already installed nginx"
            if is_inactive nginx;then
                if port_is_used 80;then
                    echo -e "\033[31m 端口 80 已被占用，请清理占用程序后重新执行初始化命令。 \033[0m"
                    exit 1
                    return
                fi
                sudo systemctl start nginx
                ck_ok "Start Nginx"
            fi
            return
        fi
        sudo systemctl stop nginx
    fi
    
    if port_is_used 80;then
        echo -e "\033[31m 端口 80 已被占用，请清理占用程序后重新执行初始化命令。 \033[0m"
        exit 1
        return
    fi
    
    download_nginx
    cd /usr/local/src
    echo "Unzip Nginx"
    sudo tar zxf nginx-1.23.0.tar.gz
    ck_ok "Zip Nginx"
    cd nginx-1.23.0
    
    echo "Install depends"
    install_depends
    
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
User=${USER}
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/bin/sh -c "/bin/kill -s HUP \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"
ExecStop=/bin/sh -c "/bin/kill -s TERM \$(/bin/cat /usr/local/nginx/logs/nginx.pid)"

[Install]
WantedBy=multi-user.target
EOF
    
sudo bash -c 'cat > /usr/local/nginx/conf/nginx.conf <<EOF
worker_processes  1;

events {
    worker_connections  1024;
}

http {
    include       mime.types;
    default_type  application/octet-stream;

    sendfile        on;
    keepalive_timeout  65;

    server {
        listen       80;
        server_name  localhost;

        location / {
            root   html;
            index  index.html index.htm;
        }

        error_page   500 502 503 504  /50x.html;
        location = /50x.html {
            root   html;
        }
    }

    include '${HOME}'/zagent/nginx/conf.http.d/*.conf;
}

stream{
    upstream tcpssh{
        hash $remote_addr consistent;
        server 	8.8.8.8:389 max_fails=3 fail_timeout=10s;
    }

    include '${HOME}'/zagent/nginx/conf.stream.d/*.conf;
}
EOF'
    
    sudo /bin/mv /tmp/nginx.service /lib/systemd/system/nginx.service
    ck_ok "edit nginx.service"
    
    sudo ln -s /usr/local/nginx/sbin/nginx /usr/local/bin/
    
    sudo chown root.${USER} /usr/local/nginx/sbin/nginx
    sudo chmod 750 /usr/local/nginx/sbin/nginx
    sudo chmod u+s /usr/local/nginx/sbin/nginx
    
    echo "Load service"
    sudo systemctl unmask nginx.service
    sudo  systemctl daemon-reload
    sudo systemctl enable nginx
    
    echo "Start Nginx"
    sudo systemctl start nginx
    ck_ok "Start Nginx"
    
    sudo /usr/bin/rm /usr/local/src/nginx-1.23.0.tar.gz
}

install_libvirt()
{
    if [[ ${ID} = ${ubuntu} ]];then
        sudo apt reinstall -y qemu-kvm libvirt-daemon-system libvirt-clients libvirt-dev qemu virt-manager bridge-utils libosinfo-bin
    else
        sudo yum -y install qemu-kvm python-virtinst libvirt libvirt-python virt-manager libguestfs-tools bridge-utils virt-install
    fi
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
    apt_update
    
    install_libvirt
    ck_ok "Install kvm"
    
    sudo service libvirtd restart
    sudo chmod 777 /var/run/libvirt/libvirt-sock
    sudo setfacl -m u:libvirt-qemu:rx ${HOME}
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
        if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/websockify.md5` ]
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
    
    if [ ${zip_md5} == `curl -s https://pkg.qucheng.com/zenagent/app/websockify.md5` ]
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
                if port_is_used 6080;then
                    echo -e "\033[31m 端口 6080 已被占用，请清理占用程序后重新执行初始化命令。 \033[0m"
                    exit 1
                    return
                fi
                nohup ${HOME}/zagent/websockify/run --token-plugin JSONTokenApi --token-source http://127.0.0.1:55001/api/v1/virtual/getVncAddress?token=%s 6080 > ${HOME}/zagent/websockify/nohup.log 2>&1 &
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
    
    if ! command_exist python3;then
        if [[ ${ID} = ${ubuntu} ]];then
            sudo apt install python3
        else
            sudo yum install python3
        fi
    fi
    
    /usr/bin/rm -rf ${HOME}/zagent/websockify.zip
    
    sudo chmod +x ${HOME}/zagent/websockify/run
    if service_is_inactive JSONTokenApi;then
        if port_is_used 6080;then
            echo -e "\033[31m 端口 6080 已被占用，请清理占用程序后重新执行初始化命令。 \033[0m"
            exit 1
            return
        fi
    fi
    ${HOME}/zagent/websockify/run --token-plugin JSONTokenApi --token-source http://127.0.0.1:55001/api/v1/virtual/getVncAddress?token=%s 6080 -D
}

install_curl()
{
    if [[ ${ID} = ${ubuntu} ]];then
        sudo apt install  -y curl
    else
        sudo yum install curl
    fi
}

install()
{
    if ! command_exist curl;then
        echo "Install curl"
        if ! $is_update_apt;then
            if [[ ${ID} = ${ubuntu} ]];then
                sudo apt update
            fi
            is_update_apt=true
        fi
        install_curl
        ck_ok "install curl"
    fi
    
    create_dir ${HOME}/zagent
    if $is_install_zagent;then
        install_zagent
        is_success_zagent=true
    fi
    
    if $is_install_zvm;then
        install_zvm
        is_success_zvm=true
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
}

if [ ! -n "$(egrep -o "(vmx|svm)" /proc/cpuinfo)" ];then
    echo -e "\033[31m Not support virtualization \033[0m"
    exit 1
fi

if [ '0' = "$(lsmod |grep kvm |grep -v grep | tr -s ' ' | cut -d' ' -f1,3|grep kvm|grep -v grep | tr -s ' ' | cut -d' ' -f2)" ];then
    echo -e "\033[31m Virtualization is disabled in the host BIOS \033[0m"
    exit 1
fi

ubuntu=ubuntu
centos=centos

source /etc/os-release

if [[ ${ID} != ${ubuntu} && ${ID} != ${centos} ]];then
    echo -e "\033[31m Not support os \033[0m"
    exit 1
fi

is_install_zagent=true
is_install_zvm=false
is_install_ztf=false
is_install_nginx=true
is_install_kvm=true
is_install_novnc=true
is_install_websockify=true

is_success_zagent=false
is_success_zvm=false
is_success_ztf=false
is_success_nginx=false
is_success_kvm=false
is_success_novnc=false
is_success_websockify=false

is_update_apt=false

soft="zagent,nginx,kvm,novnc,websockify"
force=false
secret=""
zentaoSite=""
isInstall=true

while getopts ":k:s:z:cr" optname
do
    case "$optname" in
        "s")
            if [ -n $OPTARG ];then
                soft=$OPTARG
                is_install_zagent=false
                is_install_zvm=false
                is_install_ztf=false
                is_install_nginx=false
                is_install_kvm=false
                is_install_novnc=false
                is_install_websockify=false
                if [[ $OPTARG =~ zagent ]];then
                    is_install_zagent=true
                fi
                if [[ $OPTARG =~ zvm ]];then
                    is_install_zvm=true
                    is_install_ztf=true
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
        "z")
            zentaoSite=$OPTARG
        ;;
        "c")
            isInstall=false
            if [ ! -f ${HOME}/zagent/zagentinit ];then
                curl -s -L https://pkg.qucheng.com/zenagent/zagentinit -o ${HOME}/zagent/zagentinit
            fi
            sudo chmod +x ${HOME}/zagent/zagentinit
            ${HOME}/zagent/zagentinit -c
        ;;
        *)
            echo "Unknown error while processing options"
        ;;
    esac
done

if [ -z "$zentaoSite" ]; then
    zentaoSite="192.168.122.1"
fi

HOME="`cat /etc/passwd |grep ^${SUDO_USER:-$(id -un)}: | cut -d: -f 6`"
HOME=${HOME:-$HOME}

if [ ${isInstall} == true ];then
    install
fi
