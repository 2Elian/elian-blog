#!/bin/bash
# 一键构建脚本：打包前后端
# 用法: bash build.sh

set -e

echo "===== 1/3 构建 Blog 前端 ====="
cd web-blog && pnpm build && cd ..
echo "✓ web-blog/dist/ 已生成"

echo ""
echo "===== 2/3 构建 Admin 前端 ====="
cd web-admin && pnpm build && cd ..
echo "✓ web-admin/dist/ 已生成"

echo ""
echo "===== 3/3 编译 Go 后端 ====="
go build -o server.exe cmd/server/main.go
echo "✓ server.exe 已生成"

echo ""
echo "=========================================="
echo "构建完成！启动方式："
echo "  ./server.exe -f configs/config.yaml"
echo ""
echo "访问地址："
echo "  Blog:   http://localhost:8080/"
echo "  Admin:  http://localhost:8080/admin/"
echo "=========================================="
