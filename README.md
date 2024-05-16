# ed2k 链接提取

**ed2k Links Extractor**

这是用来提取某些论坛提供的带有 `ed2k` 文本文件中对应链接的小工具。

## 介绍

### 准备工作

需要进行如下操作：

1. 下载压缩包文件（`rar` / `zip`）
2. 解压得到其中带有链接的文本文件（`txt`）
3. 将该文件移动 / 复制到某一指定目录，例如：`/User/me/ed2k`（MacOS）`D:\ed2k`（Windows）

### 提取链接

请进入对应系统的命令行，并进入 `ed2k` 程序的目录。

命令格式为：

```bash
ed2k <目录>
```
例如：

**MacOS**

```bash
./ed2k /User/me/ed2k
```

**Windows**

```bash
ed2k.exe D:\ed2k
```
执行成功没有错误的话就会生成如下三个文本文件：

- `_ed2k_4k.txt`：提取出的 4K `ed2k` 链接
- `_ed2k_8k.txt`：提取出的 8K `ed2k` 链接
- `_ed2k_all.txt`：提取出的全部 `ed2k` 链接，对于 4K / 8K 并存的文件，保留更高解析度的链接