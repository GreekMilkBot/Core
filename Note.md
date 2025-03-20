# Note

## 加载说明

### 插件加载顺序

```text
system.* -> admin.* -> adapter.* -> handler.*
```

## 消息流向

### 接收消息流向

```text
adapter[async] -> queue -> filter -> handler(loop)
```

### 发送消息流向

```text
any -> queue(1)  -> filter ->queue(2) -> route  ->adapter
```
