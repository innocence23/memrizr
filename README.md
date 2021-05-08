# memrizr

# 初始化golang的dockerfile及docker-compose环境
- github.com/cespare/reflex 监听文件变化
- traefik反向代理域名  [traefik.http.routers.account.rule=Host(`malcorp.test`) && PathPrefix(`/api/account`)]
- traefik文档：https://doc.traefik.io/traefik/getting-started/quick-start/


# 组织项目结构：call chain (handler -> service -> repository -> data sources)
- handler
- service
- repostitory
- model （entry）
