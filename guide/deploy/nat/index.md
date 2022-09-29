# 转发NAT虚机端口到宿主机

前面我们说到，虚拟机和宿主机之间建立了一个NAT网络，此种情况下虚拟机没有固定IP地址，他会通过宿主机的代理来访问外部网络。为了访问这些没有IP的虚拟机上的服务（如ssh到Linux虚机的22端口），我们通过Nginx来转发虚拟机端口到宿主机达到目的。比如将一台Ubuntu虚机的22端口转发到宿主机的50022端口，我们就可以通过命令`ssh <宿主机IP> -p 50022`来等登录到虚机。

1. 在Linux宿主机上安装Nginx；

   ```shell
   sudo apt-get install nginx
   ```

2. 执行以下命令，使得非root用户可以运行nginx，其中aaron为普通用户所在的组；

   ```shell
   sudo chown root:aaron /usr/sbin/nginx
   sudo chmod 750 /usr/sbin/nginx
   sudo chmod u+s /usr/sbin/nginx
   ```

3. 新建一个保存nginx配置的文件夹，并修改他的权限。其中conf.http.d用于网站类、conf.stream.d用于TCP/UDP类端口的转发；

   ```
   mkdir -p /home/aaron/zagent/nginx/conf.http.d
   mkdir -p /home/aaron/zagent/nginx/conf.stream.d
   sudo chmod -R 666 /zagent/nginx/conf.*.d
   ```

4. 在conf.http.d目录中创建web端口转发文件，访问宿主机50080端口相当于访问虚机80端口；

   ```
   server{
   	listen      50080;
   	location / {
   		proxy_pass   http://192.168.122.28:80;
   	}
   }
   ```

5. 在conf.stream.d目录中创建ssh端口转发文件，访问宿主机50022端口相当于访问虚机22端口；

   ```
   server {
       listen 50022;            
       proxy_connect_timeout 1h;
       proxy_timeout 1h;
       proxy_pass 192.168.122.28:22;
   }
   ```

6. 执行`nginx -s reload`，重新加载nginx配置；

7. 如果虚拟机上有运行在80端口的Web服务（如Apache），用浏览器打开`http://<宿主机IP>:50080`，确认可访问虚拟机上的网站；

8. 如果虚拟机为Linux且安装了SSH服务，使用命令`ssh <宿主机IP> -p 50022`，确认可登录到虚拟机终端。

以上为端口转发的例子，实际工作中请根据自己的需要来决定转发哪些端口。ZAgent项目提供了映射虚拟机的API接口`api/v1/virtual/addVmPortMap`，具体代码请见[这里](https://github.com/easysoft/zenagent/blob/main/cmd/host/router/handler/virtual.go)。