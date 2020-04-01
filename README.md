# cfull
Simple Golang package to set RGB color of fmt.Print output

> It has only one method - SetRGB, but this is enought to do full-color output in Terminal using `fmt.Print()`

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
    // Set color by changing nums "255"
    fmt.Sprintf("\033[38;2;%s;%s;%sm%s\033[0m", "255", "255", "255", YourStringText)
```


### 2. RGB Background:

```golang
    // Set color by changing nums "255"
    fmt.Sprintf("\033[48;2;%s;%s;%sm%s\033[0m", "255", "255", "255", YourStringText)
```

### 3. RGB Foreground + RGB Background:

```golang
// Set foreground color by changing first three numbers and background by changing last 3 numbers"
    fmt.Sprintf("\033[38;2;%s;%s;%sm\033[48;2;%s;%s;%sm%s\033[0m", "255", "255", "255", "255", "255", "255", YourStringText)
```
