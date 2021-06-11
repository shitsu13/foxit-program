#### 請求回傳 403 ERROR (近期新增)



- 發現抓不到資料，請求的回傳狀態是 403，但使用瀏覽器卻可以正常取得資料，而且就算 Headers 帶一樣也沒辦法。

- 將回傳的資料儲存成 html 網頁檔查看：

  ![403錯誤](E:\interview\foxit\403_error.png)

  

  - 主要是因為 Dcard 使用了 Cloudflare 的驗證，會需要經過渲染 JavaScript 才能進入。
  - 有幾種方式可以解決，改使用 Selenium、Pyppeteer 來模擬瀏覽器操作，或者使用 cloudscraper 專門就是要拿來繞過 Cloudflare 頁面的套件，而且它是建立在 Requests 之上，因此幾乎不用修改程式碼。