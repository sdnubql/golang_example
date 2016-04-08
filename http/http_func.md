从头开始比较pattern是否和path匹配
## func pathMatch(pattern,path string) bool
### 匹配逻辑
* pattern是否为空,为空 false
* pattern不是以/结尾，直接比对 pattern和path 不等 false
* 判断 path的长度大于等于pattern，并且 pattern和path的前pattern个长度的字符比较相等true

### 代码
<pre><code>
    func pathMatch(pattern,path string) bool {
        if len(pattern) == 0 {
            return false
        }

        n := len(pattern)
        if pattern[n-1] != '/' {
            return pattern == path
        } 
        return len(path) >= n && path[0:n] == pattern
    }

</code></pre>


## 示例
<pre><code>
package main

import (
        "fmt"
        "path"
)

func main() {
        paths := map[string]string{
                "":         "/adf",
                "12":       "12sds",
                "56":       "56",
                "567":      "57",
                "as/ff/":   "as/ff",
                "as/fff/":  "as/fff/",
                "as/ffff/": "as/ffff/fff",
        }

        for k, p := range paths {
                fmt.Printf("pattern=%s,path=%s,result=%v \n", k, p, pathMatch(k, p))
        }
}

func cleanPath(p string) string {
        if p == "" {
                return "/"
        }
        if p[0] != '/' {
                p = "/" + p
        }
        np := path.Clean(p)
        if p[len(p)-1] == '/' && np != "/" {
                np += "/"
        }
        return np
}

</code></pre>

## 实例结果
<pre><code>

pattern=as/fff/,path=as/fff/,result=true 
pattern=as/ffff/,path=as/ffff/fff,result=true 
pattern=,path=/adf,result=false 
pattern=12,path=12sds,result=false 
pattern=56,path=56,result=true 
pattern=567,path=57,result=false 
pattern=as/ff/,path=as/ff,result=false
</code></pre>


## func cleanPath(p string) string
### 规范路径,去掉路径中的. 和..,按照shell的目录跳转
*  

### 代码
<pre><code>
    func cleanPath(p string) string {
        if p == "" {
            return "/";    
        }    
        if p[0] != '/' {
            p = "/" + p    
        }
        np := path.Clean(p)
        if p[len(p) - 1] == '/' && np != "/" {
            np += "/"    
        }
        return np
    }
</code></pre>

### 示例代码
<pre><code>
package main

import (
        "fmt"
        "path"
)

func main() {
        paths := []string{
                "a/c",
                "a//c",
                "a/c/.",
                "a/c/b/..",
                "/../a/c",
                "/../a/b/../././/c",
                "/../a/b/../././/c/",
        }

        for _, p := range paths {
                fmt.Printf("Clean(%q) = %q\n", p, cleanPath(p))
        }
}

func cleanPath(p string) string {
        if p == "" {
                return "/"
        }
        if p[0] != '/' {
                p = "/" + p
        }
        np := path.Clean(p)
        if p[len(p)-1] == '/' && np != "/" {
                np += "/"
        }
        return np
}


</code></pre>

### 示例结果
<pre><code>
Clean("a/c") = "/a/c"
Clean("a//c") = "/a/c"
Clean("a/c/.") = "/a/c"
Clean("a/c/b/..") = "/a/c"
Clean("/../a/c") = "/a/c"
Clean("/../a/b/../././/c") = "/a/c"
Clean("/../a/b/../././/c/") = "/a/c/"
</code></pre>

