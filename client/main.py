from ftplib import FTP

def list_callback(line):
    # 在这里对每一行进行处理，例如输出到控制台或存储到列表中
    print(line)

def ftp_client(server, username, password):
    # 连接到服务器
    ftp = FTP()
    server_address, server_port = server.split(":")
    ftp.connect(server_address, int(server_port))

    # 登录
    ftp.login(username, password)

    while True:
        command = input("请输入命令：")
        args = command.split()
        cmd = args[0]

        if cmd == "LIST":
            # 列出文件列表
            if len(args) > 1:
                directory = args[1]
            else:
                directory = "."  # 默认为当前目录
            try:
                ftp.dir(directory, list_callback)
            except Exception as e:
                print("执行 LIST 命令出错:", e)

        elif cmd == "GET":
            # 下载文件
            if len(args) > 1:
                filename = args[1]
                with open(filename, "wb") as file:
                    ftp.retrbinary("RETR " + filename, file.write)
                print(f"文件 {filename} 下载完成。")
            else:
                print("请提供要下载的文件名。")
        elif cmd == "PUT":
            # 上传文件
            if len(args) > 1:
                filename = args[1]
                with open(filename, "rb") as file:
                    ftp.storbinary("STOR " + filename, file)
                print(f"文件 {filename} 上传完成。")
            else:
                print("请提供要上传的文件名。")
        elif cmd == "QUIT":
            # 退出
            ftp.quit()
            break
        else:
            print("未知命令。")

    ftp.quit()


def main():
    # server = input("请输入服务器地址：")
    # username = input("请输入用户名：")
    # password = input("请输入密码：")
    server = "100.76.246.116:8333"
    username = "admin"
    password = "password"

    ftp_client(server, username, password)


if __name__ == "__main__":
    main()
