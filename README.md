我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# cfull
Simple Golang package to set RGB color of fmt.Print output

> It has only one method - SetRGB, but this is enought to do full-color output in Terminal using `fmt.Print()`
>> You can use https://htmlcolorcodes.com/ to pick RGB color

## Using examples

### 1. Set only foreground color 

- foreground = ![#0099cc](https://placehold.it/15/0099cc/000000?text=+) `RGB(0, 153, 204)`
    
```golang
    myColor := cfull.Color{
        ColorFG: cfull.RGB{0, 153, 204},
    }
  
    fmt.Println(cfull.SetRGB("Hello, World!", myColor))
 ```
 
 
 ### 2. Set only background color 

- background = ![#E74C3C](https://placehold.it/15/E74C3C/000000?text=+) `RGB(231, 76, 60)`
    
```golang
    myColor := cfull.Color{
        ColorBG: cfull.RGB{231, 76, 60},
    }
  
    fmt.Println(cfull.SetRGB("Hello, World!", myColor))
 ```
 
### 3. Set background and foreground 
- foreground = ![#000000](https://placehold.it/15/000000/000000?text=+) `RGB(0, 0, 0)`
- background = ![#E74C3C](https://placehold.it/15/E74C3C/000000?text=+) `RGB(231, 76, 60)`
    
```golang
    myColor := cfull.Color{
        ColorFG: cfull.RGB{0, 0, 0},
        ColorBG: cfull.RGB{231, 76, 60},
    }
  
    fmt.Println(cfull.SetRGB("Hello, World!", myColor))
 ```
## Output exaple
```golang
    ColorOne := cfull.Color{
        ColorFG: cfull.RGB{0, 153, 204},
    }

    ColorTwo := cfull.Color{
        ColorBG: cfull.RGB{231, 76, 60},
    }

    ColorThree := cfull.Color{
        ColorFG: cfull.RGB{0, 0, 0},
        ColorBG: cfull.RGB{231, 76, 60},
    }

    fmt.Println(cfull.SetRGB("Hello, World!", ColorOne))
    fmt.Println(cfull.SetRGB("Hello, World!", ColorTwo))
    fmt.Println(cfull.SetRGB("Hello, World!", ColorThree))

    fmt.Println(
        cfull.SetRGB("You can", ColorOne),
        cfull.SetRGB("use different styles", ColorTwo),
        cfull.SetRGB("in one output", ColorThree),
    )
  
    // returns:
```

![example](https://github.com/DERVdice/cfull/blob/master/Example.png)

## Another way

You can make colorfull output without this package (Please, star if it helps).
All you need is use one of construtcions below. Just place it into fmt.Print()

### 1. RGB Foreground:

```golang
    // Set color by changing nums 255
    fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", 255, 255, 255, YourStringText)
```


### 2. RGB Background:

```golang
    // Set color by changing nums 255
    fmt.Sprintf("\033[48;2;%d;%d;%dm%s\033[0m", 255, 255, 255, YourStringText)
```

### 3. RGB Foreground + RGB Background:

```golang
    // Set foreground color by changing first three numbers and background by changing last 3 numbers"
    fmt.Sprintf("\033[38;2;%d;%d;%dm\033[48;2;%d;%d;%dm%s\033[0m", 255, 255, 255, 255, 255, 255, YourStringText)
```
