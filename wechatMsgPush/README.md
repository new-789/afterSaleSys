### 使用说明：
> 本企业微信消息推动程序仅支持企业微信后台自建应用及推送文本消息。  

开发目的：因在实际工作中需要想企业员工推送通知信息，而在找了相关类似程序后使用起来不尽人意，因此在查看学习了企业  
微信API后决定自己造一个仅用来推送文本消息的轮子。  
1. 登录企业微信后台新建应用并获取CorpId(企业Id)、AgentId(应用Id)、Secret。此步操作方法请自行查看企业微信 API 文档
2. 修改 `conf/wxconf.ini` 配置文件中的内容为自己获取到的企业微信内容
3. 在项目中调用 `PushMsgToWechat()` 方法并传参，该函数原型如下：
    ```go
    func PushMsgToWechat(username string, message string) error { ... }
    /*
    参数说明：
        1. username：接收消息内容的成员名
        2. 推送的消息数据
    返回值：
        error: 推送失败返回的错误信息
    */
    ```
4. 示例代码：
    ```go
    // 调用方式示例
    func main() {
        // 定义的消息数据
        msg := "测试内容来了,请让道....\n<a href='https://www.bing.com/'>这是链接</a>\n点击即可前往,恭喜您成功了，几天的功夫没有浪费！"
        // 调用推送方法并传参
        wecahtMsgPush.PushMsgToWechat("测试用户", msg)
    }
    ```
> 本程序支持的功能有限，后续有时间会陆续支持更多消息类型的推送，另谁有比较好的想法亦可通过留言的方式反馈我会逐步更新
> 或将此程序下载自行增加功能后更新到此。最后感谢各位的支持！