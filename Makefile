# 定义 Go 编译器命令
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

# 定义目标名称和依赖文件
TARGET=pomment-go

.PHONY: all clean build

all: $(TARGET)

$(TARGET):
	$(GOBUILD) -o $@ ./cli/main

clean:
	$(GOCLEAN)
