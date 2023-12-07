package autoLoadConfig

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Config struct {
	filename       string
	data           map[string]string
	lastModifyTime int64
	rwLock         sync.RWMutex
	notifyList     []Notifyer
}

func NewConfig(file string) (config *Config, err error) {
	config = &Config{
		filename: file,
		data:     make(map[string]string),
	}

	m, err := config.parse()
	if err != nil {
		fmt.Printf("parse conf error:%v\n", err)
		return
	}

	//将配置文件中的数据解析到结构体中
	config.rwLock.Lock()
	defer config.rwLock.Unlock()
	config.data = m

	//启动一个后台线程去检测配置文件
	go config.reload()
	return
}

func (c *Config) parse() (m map[string]string, err error) {
	m = make(map[string]string, 1024)
	f, err := os.Open(c.filename)
	if err != nil {
		return
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	//声明一个变量存放读取行数据
	var lineNo int
	for {
		line, errRet := reader.ReadString('\n')
		if errRet == io.EOF {
			lineParse(&lineNo, &line, &m)
			break
		}

		if errRet != nil {
			err = errRet
			return
		}
		lineParse(&lineNo, &line, &m)
	}
	return
}

func (c *Config) reload() {
	//定时器
	ticker := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-ticker.C:
			func() {
				f, err := os.Open(c.filename)
				if err != nil {
					fmt.Printf("open file error:%s\n", err)
					return
				}
				defer f.Close()

				fileInfo, err := f.Stat()
				if err != nil {
					fmt.Printf("stat file error:%s\n", err)
					return
				}

				//或者当前文件修改时间
				curModifyTime := fileInfo.ModTime().Unix()
				//fmt.Println(curModifyTime)
				ticker.Reset(time.Second * 1)
				if curModifyTime > c.lastModifyTime {
					//重新解析时，要考虑应用程序正在读取这个配置，因此应该加锁
					m, err := c.parse()
					if err != nil {
						fmt.Printf("parse config error:%v\n", err)
						return
					}

					c.rwLock.Lock()
					//defer c.rwLock.Unlock()
					c.data = m
					c.rwLock.Unlock()
					c.lastModifyTime = curModifyTime

					//配置更新通知所有观察者
					for _, n := range c.notifyList {
						//n.Callback(c)
						n.Callback(c)
					}
				}
			}()
		}
	}
}

func (c *Config) GetString(key string) (value string, err error) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, ok := c.data[key]; ok {
		return value, nil
	}
	return "", fmt.Errorf("key [%s] not found", key)
}

func (c *Config) GetStringDefault(key string, defaultString string) (value string) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if val, ok := c.data[key]; ok {
		return val
	}
	return defaultString
}

func (c *Config) GetIntDefault(key string, defaultInt int) (value int) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if val, ok := c.data[key]; ok {
		v, err := strconv.Atoi(val)
		if err != nil {
			return defaultInt
		}
		return v
	}
	return defaultInt
}

func (c *Config) GetInt(key string) (value int, err error) {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	if value, ok := c.data[key]; ok {
		v, err := strconv.Atoi(value)
		return v, err
	}
	return 0, fmt.Errorf("key [%s] not found", key)
}

func lineParse(lineNo *int, line *string, m *map[string]string) {
	*lineNo++

	l := strings.TrimSpace(*line)
	if len(l) == 0 || l[0] == '\n' {
		return
	}

	itemSlice := strings.Split(l, "=")
	key := strings.TrimSpace(itemSlice[0])
	value := strings.TrimSpace(itemSlice[1])
	(*m)[key] = value
}
