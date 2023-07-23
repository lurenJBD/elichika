# elichika

Local Server for Love Live! All Stars Japanese / Global

Clone this repository.
```
# clone
git clone https://github.com/YumeMichi/elichika

# or update
git pull
```

Build executable.
```
go build
```

Generate config file.
```
# linux
./elichika

# or windows
elichika.exe
```

Edit cdn server (Usually local ip:port).

Exit `elichika` and run again.

### Assets

Put databases and assets into `static/2d61e7b4e89961c7` (Global) or `static/b66ec2295e9a00aa` (Japanese).

You can download assets from [ll-sifas-cdn-data](https://archive.org/download/ll-sifas-cdn-data).

File list:
| File name                                | Description       |
| :--------------------------------------- | :---------------- |
| sifas-jp-cdn-assets-b66ec2295e9a00aa.tar | assets (Japanese) |
| sifas-gl-cdn-assets-2d61e7b4e89961c7.tar | assets (Global)   |

### Clients
For [3.12.0 clients](https://mega.nz/folder/gwJizZjS#v_fSOadf9yrb_eaOze2r7Q), you need to download [patched databases](https://mega.nz/folder/gwJizZjS#v_fSOadf9yrb_eaOze2r7Q).

### 20230723 测试记录
- 账号系统不完整，只能登入默认的588296696，多台设备登入都共用一个账号信息，故只能作为本地个人服务使用
- 抽卡、CH频道、获取通知详细（通知列表是能查看的）、商店、每日任务、充值页面、交换所 点击会提示错误，返回标题，服务器日志显示404
- 剧情（含语音）、MV、歌曲都能正常使用
- 进行LIVE不会扣LP，且无奖励，集训是不扣AP，不扣道具，也无道具奖励，但能获得技能
