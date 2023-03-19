# job-opportunity-reminder
[![CI](https://github.com/lyf-coder/job-opportunity-reminder/actions/workflows/ci.yml/badge.svg)](https://github.com/lyf-coder/job-opportunity-reminder/actions/workflows/ci.yml)
[![GoDoc](https://godoc.org/github.com/lyf-coder/job-opportunity-reminder?status.svg)](https://pkg.go.dev/github.com/lyf-coder/job-opportunity-reminder)
[![Go Report Card](https://goreportcard.com/badge/github.com/lyf-coder/job-opportunity-reminder)](https://goreportcard.com/report/github.com/lyf-coder/job-opportunity-reminder)

爬取招聘网站上的招聘信息，然后推送到飞书等通讯工具

当前爬取的网站：
* [v2ex 酷工作](https://www.v2ex.com/go/jobs)
* [eleduck 电鸭社区-招聘](https://eleduck.com/categories/5?sort=new)

## 使用


1. `fork` 本项目
2. 然后在 `setting->Environments->New environment->` 配置相关环境变量（也可以在 `secrets and variables` 中进入配置）
3. 飞书推送：需要先在飞书创建一个个人群，然后添加一个[自定义机器人](https://open.feishu.cn/document/ukTMukTMukTM/ucTM5YjL3ETO24yNxkjN)，
   可以获得飞书的自定义机器人 `webhook` 地址,然后设置环境变量 `fei_shu_bot_webhook_url`。

* 可配置环境变量参考：`config.json` 中的配置，当配置了环境变量会覆盖掉该配置文件的配置

备注：假如网络需要代理可设置环境变量： `proxy_url`，格式如：`socks5://127.0.0.1:3128`， `Github action` 执行不需要

## 项目相关文档

* [定时抓取招聘信息推送到飞书机器人-v1.0.0](https://www.liyongfei.com/ding-shi-zhua-qu-zhao-pin-xin-xi-tui-song-dao-fei-shu-ji-qi-ren-v1-0-0/)

