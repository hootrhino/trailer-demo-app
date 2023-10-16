<!--
 Copyright (C) 2023 wwhai

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU Affero General Public License as
 published by the Free Software Foundation, either version 3 of the
 License, or (at your option) any later version.

 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU Affero General Public License for more details.

 You should have received a copy of the GNU Affero General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
-->

# RPC解码器示例
这是一个RULEX的 RPC 解码器 golang 服务模板，用来开发私有设备接入。
## 构建
```sh
go build
```
## 示例
```lua
    function(data)
        print(rpc:Request('grpcCodec001', "cmd", "arg"))
        return true, data
    end
```
## 社区
- wwhai